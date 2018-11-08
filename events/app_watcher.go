package events

import (
	"crypto/tls"
	"sync"
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry/noaa/consumer"
	sonde_events "github.com/cloudfoundry/sonde-go/events"
	"github.com/prometheus/client_golang/prometheus"
)

//go:generate counterfeiter -o mocks/appWatcher_process.go . AppWatcherProcess
type AppWatcherProcess interface {
	Run() error
}

type AppWatcher struct {
	config             *cfclient.Config
	cfClient           *cfclient.Client
	metricsForInstance []InstanceMetrics
  app                cfclient.App
	appUpdateChan      chan cfclient.App
	sync.RWMutex       // TODO: what's this?
}

type InstanceMetrics struct {
	cpu prometheus.Gauge
}

func NewInstanceMetrics(instanceIndex int) InstanceMetrics {
	im := InstanceMetrics{
		cpu: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cpu",
				Help: " ",
				ConstLabels: prometheus.Labels{
					"instance": fmt.Sprintf("%d", instanceIndex),
				},
			},
		),
	}
	prometheus.MustRegister(im.cpu)
	return im
}

func NewAppWatcher(
	config        *cfclient.Config,
	app           cfclient.App,
	appUpdateChan chan cfclient.App,
) *AppWatcher {
	appWatcher := &AppWatcher{
		metricsForInstance: make([]InstanceMetrics, 0),
		config:             config,
		app:                app,
		appUpdateChan:      appUpdateChan,
	}
	appWatcher.scaleTo(app.Instances)
	return appWatcher
}

// RefreshAuthToken satisfies the `consumer.TokenRefresher` interface.
func (m *AppWatcher) RefreshAuthToken() (token string, authError error) {
	token, err := m.cfClient.GetToken()
	if err != nil {
		err := m.authenticate()

		if err != nil {
			return "", err
		}

		return m.cfClient.GetToken()
	}

	return token, nil
}

func (m *AppWatcher) authenticate() (err error) {
	client, err := cfclient.NewClient(m.config)
	if err != nil {
		return err
	}

	m.cfClient = client
	return nil
}

func (m *AppWatcher) Run() error {
	err := m.authenticate()
	if err != nil {
		return err
	}
	tlsConfig := tls.Config{InsecureSkipVerify: false} // TODO: is this needed?
	conn := consumer.New(m.cfClient.Endpoint.DopplerEndpoint, &tlsConfig, nil)
	conn.RefreshTokenFrom(m)

	authToken, err := m.cfClient.GetToken()
	if err != nil {
		return err
	}

	msgs, errs := conn.Stream(m.app.Guid, authToken)

	// log.Printf("Started reading %s events\n", app.Name)
	for {
		select {
		case message, ok := <-msgs:
			if !ok {
				// delete all instances

				return nil
			}
			switch message.GetEventType() {
			case sonde_events.Envelope_ContainerMetric:
				metric := message.GetContainerMetric()
				index := metric.GetInstanceIndex()
				if int(index) < len(m.metricsForInstance) {
					instance := m.metricsForInstance[index]
				  instance.cpu.Set(metric.GetCpuPercentage())
				}
			}
		case err, ok := <-errs:
			if !ok {
				errs = nil
				continue
			}
			if err == nil {
				continue
			}
			// TODO: do something with errors
			break
			// m.errorChan <- err
		case updatedApp, ok := <-m.appUpdateChan:
			if !ok {
				m.appUpdateChan = nil
				conn.Close()
				break
			}

			if updatedApp.Instances != m.app.Instances {
				m.scaleTo(updatedApp.Instances)
			}
			m.app = updatedApp
		}
	}
}

func (m *AppWatcher) scaleTo(newInstanceCount int) {
	currentInstanceCount := len(m.metricsForInstance)

	if currentInstanceCount < newInstanceCount {
		for i := currentInstanceCount; i < newInstanceCount; i++ {
			m.metricsForInstance = append(m.metricsForInstance, NewInstanceMetrics(i))
		}
	}	else {
		for i := currentInstanceCount; i > newInstanceCount; i-- {
			m.unregisterInstanceMetrics(i-1)
		}
		m.metricsForInstance = m.metricsForInstance[0:len(m.metricsForInstance)-1]
	}
}

func (m *AppWatcher) unregisterInstanceMetrics(instanceIndex int) {
	prometheus.Unregister(m.metricsForInstance[instanceIndex].cpu)
}
