// Code generated by go-bindata.
// sources:
// fixtures/CA.crt
// fixtures/CA.key
// fixtures/client.crt
// fixtures/client.key
// fixtures/invalid-ca.crt
// fixtures/server.crt
// fixtures/server.key
// DO NOT EDIT!

package loggregator_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _caCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIE1zCCAr+gAwIBAgIBATANBgkqhkiG9w0BAQsFADANMQswCQYDVQQDEwJDQTAe
Fw0xNzAzMDIyMTM4NTZaFw0yNzAzMDIyMTM5MDJaMA0xCzAJBgNVBAMTAkNBMIIC
IjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA7EJgvUeZxE1QdgXTwjiDuk1B
dXNSLjXS6STSMMEbFTk098Fa2CqIkAHSAhM0iYqIeEwRLGuBzQSvWZPOGu7b0WeU
7i1siir4xdgXQ0OK0xlsWIqS4w15eFpVZmA1apEKncEfjSrj/F639rISJD0hHGho
tSSZ4kWunNqMWbVpd7agp4yFL/nwVGOoa49KRByQ5S90YqVEbz+hhlyYxUmYh0bi
nt+Zto/oKW4yfa70tZU0b53Ax0uiJZizhmCEkICn221YofvhA1RELj7VT5g+8ahZ
s6J8uqtDjPREpiOrYUddU4IkqtZDDwQySdzYPvxfeDINd3Nhuwk4AwD0fsdbGFKL
0ZlbpMLCWj927J5Te80fiKrX9EAom2i3WfViNebXHPdQ5FaLkeflOudjgpvMbYjE
4o/UgiEELO7i5UXl89SMir1bpve1CZELcfYcSkqFDlxc6Q/PgS4y1CtQ9LAX4DdA
jXti2S/q1Larhl74iY6EdKIqfdlx9lWH5O3eBAFI9TEtQoljxKXkZG+J7lSpm1pw
XqzaDyRBWG4pTGwcV8dcU+6S1Z95waOgHo8KVAi3DBUAOOpfgTACop8D7dv2NNWD
ZHRcxSDXthanLC0fhNXzlkI0SzO3BxacNi0e23QCF8l4s0U3960FJf33IKycG8yR
YPEhKR5tPEjiCk7+jG0CAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFIZ668SYdlEncMeDV85tR7R6hwHIMA0GCSqGSIb3
DQEBCwUAA4ICAQCPpur8XpJHWlSJIAl6V4+7ZEXiV2b6Z4dZWtFJFx/gJjASpC+b
MP1+x6Oarh9H+NyxChAqQK54Y6q30S/IL2wZVT+5HhTub/GFwWFRWqCm8zabe8pQ
6w8iFOZ3gNg14p7gRsqNbv7Ci/67eE9S1qnux4H85Eeo/qTRYtlVsZPOxO0PdkEB
tuyVu6PA6ifGwBq3+z8tpgiMGhpVkoYxFgcgJx0qlozzOEBR5akwNcMKl2V2GOSN
IbfT8R7XrRjS5cytNyhjHVD2ZejoQ6j9XZ7/MjsEEN33CS3+PEsh9bpFVwotrZnq
e09DJjYvsmhMV7H9vVo3fP6HKNEzLyxVYwwpPJrYLfU3vfrvf/VskNu0gBukHED9
FErITDiOyfx/Y6LaxJgQR+iiFFDor1u3WUPd070/h4ByCKsTsZkhtHoIl6m93Xoo
7Nxouili21GAkSWNQVxKK2nOIrdxlkOMCJwgK9qLe6OPfhfxxsl5+soiezq4qpQ1
kDuIxk+xVVofy6AWz/TN9Fu2J+MHztjYtOaPe131gQ+G5w3E5oYwZqdpv+HFdS05
IajPpSvKSBuredA2nhPUmsnwXV6XvGQEkmg6qHvXCjiH4IFf0EnezAX/j8xWa60l
yDO02zqeN43dCTYzT2dnIM1e/KDCwRNrhuCbJqGFQVHxNK4pUzCcpplfpg==
-----END CERTIFICATE-----
`)

func caCrtBytes() ([]byte, error) {
	return _caCrt, nil
}

func caCrt() (*asset, error) {
	bytes, err := caCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "CA.crt", size: 1740, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func caKeyBytes() ([]byte, error) {
	return _caKey, nil
}

func caKey() (*asset, error) {
	bytes, err := caKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "CA.key", size: 3243, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIEGjCCAgKgAwIBAgIRANbJG+A8oSBVz9DB7G1//1IwDQYJKoZIhvcNAQELBQAw
DTELMAkGA1UEAxMCQ0EwHhcNMTcwMzAyMjE0MDQzWhcNMTkwMzAyMjE0MDQzWjAR
MQ8wDQYDVQQDEwZjbGllbnQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDQk09TaDP+LKwx7qFhmODiTO7u3VdUFiIXQZWhTXB5bSzrH6kUWLZGiZSL4SBx
ulDkZCeUrSSC6v5/r0pTjBWlprYP6qc0yL8iC5G0CrptoZgtVT6B71zhBzzzA8bF
FY/9CU7N4sLffNvNdhw0hwoWhRbDaK0RiR2n5K6OskaMCAMCOptQbzB97NkdkUJO
c//DPhsFPevPdclzQhFsTkvKrdtrAATeFk4stBVQFbgH8jAT9ZZbrLW9/Q9lRH2n
f7In/KMHfU6KSM7dTJrcYm4U8PBCpfgNrZJGOY1sbZUL2lUmiZhjsWfrzPPUyKx1
eV7vnjtITF4dHhB57biZvJg5AgMBAAGjcTBvMA4GA1UdDwEB/wQEAwIDuDAdBgNV
HSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFKXg0FSOMQPujhYr
Xu6aKHtLVUXVMB8GA1UdIwQYMBaAFIZ668SYdlEncMeDV85tR7R6hwHIMA0GCSqG
SIb3DQEBCwUAA4ICAQAIDEoIKdKNeqbN/CB23oS8eStQGJWskng4fY319JAiVph6
fcDYmarwYUQAHL3tC39L2aeb9//FWnwF5NSZp4jqWQpUxtHXnd8vmCMDHe5+BeaV
cDMDttTqklINy/wkUNC1NJuNGIyTXLOsM1VQVMrAQw2/RmH1Hs8eXCgJwBmjyQHZ
LwhbS9+CvoV09bleMPBzEAmH0zITjuBxvGaxQBT/DyEfD5XK5tIovByzlfKGHJdf
GVV2Vjmodw9EHJW4atkJphbwhg/G//WWnF4tMhhKEVaOV04J9kLHmvmFxcVqRVAF
Lh6MXio8keHkuypCNsiA9jodR4A4TLvIPnaAg0rrYAck6su+ZbgYliVodHNHarGd
pIBlmd9uv0J8G+0Wd6XRXqfr3rNdjGnyjcGa/9xfYuaryLyU2wJ9XSeJDccr5MBB
lo1AbTLIb4GeXlTzn80x2tR/j6JRF2sEpZjEpVv+0+VT2+Am1V+gvLx2JGRX1dBO
beFV+F8h1YqhmeK2dCiQ7vrjMHcP3QNjGpJr4lX9ON5Pon2yPYPwXpSHIhdALXcV
SpFgM6lHofMjxmiOtikqDs5s6zV9yGBwNU+KeTxLGdtWLVpFC1EBeq/AXs3XDbRL
qri1ums73u9nf9fdq/LTL9jUGoBy3pwYpZ/fc9TqxM9rpOqY2V1toT1qsv5d1Q==
-----END CERTIFICATE-----
`)

func clientCrtBytes() ([]byte, error) {
	return _clientCrt, nil
}

func clientCrt() (*asset, error) {
	bytes, err := clientCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client.crt", size: 1484, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0JNPU2gz/iysMe6hYZjg4kzu7t1XVBYiF0GVoU1weW0s6x+p
FFi2RomUi+EgcbpQ5GQnlK0kgur+f69KU4wVpaa2D+qnNMi/IguRtAq6baGYLVU+
ge9c4Qc88wPGxRWP/QlOzeLC33zbzXYcNIcKFoUWw2itEYkdp+SujrJGjAgDAjqb
UG8wfezZHZFCTnP/wz4bBT3rz3XJc0IRbE5Lyq3bawAE3hZOLLQVUBW4B/IwE/WW
W6y1vf0PZUR9p3+yJ/yjB31OikjO3Uya3GJuFPDwQqX4Da2SRjmNbG2VC9pVJomY
Y7Fn68zz1MisdXle7547SExeHR4Qee24mbyYOQIDAQABAoIBAQCDQvG9L+ronvyP
P/pYDXoZcQk9UhbPWE6EDEY6iR2rjUWjgv1JTLsJkTPnA/sDF5oo4c73Bi/gJDud
55cDFTnaUhxpz+cClQOUTWuYvtesfvHiNa9s3lK5W3DEDzBcf8FqIR1y/K5fL5i0
qQv/Eq0Klx9IQuwICR6ctpIUQ6WaIZzqfFjt4parVemNRExFlEoVSCq+3tsykLYv
A0ik/KT9uxV4hnT8y3ChUWmR5vsyZywhPcNkkmyZ++NT/dSiEfm2sQHxBHrbCvrR
IA/tqueW0dvqNn/5i7a80cKvF3yYaP7diUDmaF07AknklgP4pDql+a57fZUyU1OP
dRU/f5KxAoGBAPm29QMyAniqF4UyHLF1lT0e8eF5mgmxHbolNydP6hRI+tt5zooO
Hm98Emc8vvUmwjDeBf2Q/uzMlybb7l/jEWMdGpSWF4fHbMW6p/O337AsexRs4CYW
fGlxefSgwn2eLohrF75JUoIkOBWPTghOb1rriFDYPPfIQ2798qRqL0eXAoGBANXT
RV/YVrE8dEuIdIgL405QRC4uJ0JqvWcF+6wdM+oUIii2QJ2Gv40s6DkKkj3xVfFn
YZo3nLiN1dlXssUbj3rowK8RDOsWWB0fABy3cv5AStcEm5/DbYoBA238sJrxMv3s
4FIfejXgz0H8rzLFo2G6OH+Fj4gyfTEe+L2bf5ivAoGBAJ3iC2+FwwI/qLdSQ8n/
qegHFK+Wjsnp7f2vYfu9DybcLZ9umCO5YYzjSU90nNerNcRn3CnB7ywQ/tBW5ZRC
oM+lqxSfASmb8S3jIiV58DgXJ4p2fBIYzXhtgsokilOaG5ms/EoXWc9mzv7yqEeg
L8+wc3OKBKI1QE2AylUhvdrpAoGBAJKiSs36nyD2M05lKnztohjfhJIixnuOlYqk
GL4kXKim2/INYexlLtFSvYH4lOdzrBtbaRyIDKqnYcKlmIwBZXjFSvYDfFdvQEfg
LYSD7ZEcZ/ZiS5LZQzJDlaWEh1wKEoz+N6s5uOZf3+IBtefcjzmCqiIRZqfg7+eS
tMPacjszAoGAIMYV/ODqdpRjzwWVpJnBmWkX/sWaAMwS5kZ7FrP5Gskn4WOFjuwX
L9XuEqBl5to9vn6SAE50dRqbnBK3HLx7v0Y0QVaq/JRkGAS2n8N0ZEk6J7QctPtu
lmEvvZIj8yFantPLxX9j/PNmm9r+8pWQDFmWwLnSGTSHuBqV9PpTT24=
-----END RSA PRIVATE KEY-----
`)

func clientKeyBytes() ([]byte, error) {
	return _clientKey, nil
}

func clientKey() (*asset, error) {
	bytes, err := clientKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client.key", size: 1679, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _invalidCaCrt = []byte(`foobar
`)

func invalidCaCrtBytes() ([]byte, error) {
	return _invalidCaCrt, nil
}

func invalidCaCrt() (*asset, error) {
	bytes, err := invalidCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "invalid-ca.crt", size: 7, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serverCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIEGTCCAgGgAwIBAgIQBmJzvh3KxC4lVxa7yaSc5jANBgkqhkiG9w0BAQsFADAN
MQswCQYDVQQDEwJDQTAeFw0xNzAzMDIyMTUxMzJaFw0xOTAzMDIyMTUxMzJaMBEx
DzANBgNVBAMTBm1ldHJvbjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJxDv+O8Qp6xkNwCGU5AoY4uRiXUxvPA3TWjX2thtSy/pcKBVI1GsFtLGDQC3fsA
qCTAu7OrrSfWeBkyLlPhvJEeAznH51AKR6zR2qIMzVndtG0oJUIBEUPHFlP3pMYB
jEX3tNmF3LX9WG3mQOec7aIQ53Xb/ck2+p2gllIkXmmULhizxi74RWdLxjzQXI30
2fkVgjzpbDAttvBQjsCNz8yIcBIfPMoayoXXI8FhLc8MfSjlOjUCQ9uUaQmUV10y
zy5Es3KUQVqyYEcsvpZBWFwFV2C83OhKGURpbkYIsB1NSFI805RdIwch5XyX2Bih
bfOFFWI264Jo8oxNiyH9OOECAwEAAaNxMG8wDgYDVR0PAQH/BAQDAgO4MB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUkHbRKLFt2NIH9fim
grpBPRTu140wHwYDVR0jBBgwFoAUhnrrxJh2USdwx4NXzm1HtHqHAcgwDQYJKoZI
hvcNAQELBQADggIBANhWs616uSc8blBuMF3Ubzja6OW95farvbGwW0Xqp/lenn/C
6H19yXoNfeVFW3+05MMSQl3UZW5FEsnvgbKi5yC7MuZV6ligpIbr19UTuBsguoV/
6y6ilaBlx88V9FB3eKUlUG69j7FYqNv2uLmKdOcfCBDRPOhDv8fykoxo9hz1tn2Q
B44Y47PlyKF8bajpiw7kUdW0Bj2VZRvDrivX6Vlrgj5DLc7QYYFKcGSuNGkMwTFP
XHmq96kdO4Qpswv83n5CZeIbHWxe4LX2iQlPkcLOpi8wNBtvwrY8riNYHaqg2Ad0
kSCYiZMXQT/nBc9yz6hD3Sf5JyOAtJEJN75QgJcguJHFXNs3WwulSGD30UBwi9SC
icCeRsoFdwUY96FotaFu/s8kTtKjhS3iKUoboxa0TVf5JvIznhILuOWvRPilHnc4
XHGCUHU6yaxAbpWf8RtXoRaZgrsHpNDrE2Vuo35DHlWtQuefjotdF2Vtn8wds7cR
RPCG1K4goMj7604IVano8FR9f1IB2l9h+M3GZRatt9Va7CQ4dzf5CI8lV96m/JO+
aW2lBbIIOc5KWHBpCsrYiT6z8xqmCr03jkERpQWW7a8wz0UGQDd0GHeYnWz7vq79
dVRUfeUdSa7cPBy5xJwBxusW7ExOqjhCByRYxSQWIrT6p5VVFFeLHhDrbFlb
-----END CERTIFICATE-----
`)

func serverCrtBytes() ([]byte, error) {
	return _serverCrt, nil
}

func serverCrt() (*asset, error) {
	bytes, err := serverCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "server.crt", size: 1480, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAnEO/47xCnrGQ3AIZTkChji5GJdTG88DdNaNfa2G1LL+lwoFU
jUawW0sYNALd+wCoJMC7s6utJ9Z4GTIuU+G8kR4DOcfnUApHrNHaogzNWd20bSgl
QgERQ8cWU/ekxgGMRfe02YXctf1YbeZA55ztohDnddv9yTb6naCWUiReaZQuGLPG
LvhFZ0vGPNBcjfTZ+RWCPOlsMC228FCOwI3PzIhwEh88yhrKhdcjwWEtzwx9KOU6
NQJD25RpCZRXXTLPLkSzcpRBWrJgRyy+lkFYXAVXYLzc6EoZRGluRgiwHU1IUjzT
lF0jByHlfJfYGKFt84UVYjbrgmjyjE2LIf044QIDAQABAoIBAEoQ1qgDuHPtZ+LP
39y2R5zTlA4kXzRzyqW8zOJhynTuVYvurvvTcI2k91DXtw1cj1X0QEN18Lm8LuqK
XWkGytDhr9e4R4wpgzWHtiAu/zpEGlvOtnIlfWBpwcyEyYXMgAc91Gtxxl5AOaTF
iSf+uqehaYP0hMLVm04RKo6+jD1v8Biw56UpHYXtzH34YqVR+Wlbj/hS31sn/HBE
xcRhNDZyW7uhKEF5iUQ0Nt8HneLmMs1ZD9W4yqOw+yNBwwXkRaTkzkmxSYSIrB5B
NVaBZsNcxeIL0tufy5kL7W5Asr08jHnA08wKC3vh1x9LtgquWEEuyzpUaAnrn6so
JUyjKU0CgYEAygy59F1a3C3CDY+KR6JPvvDlYCsPPqPtWYPmOr7eYeW2M3xxlRAU
mhpmc80JFlVUrcBMnWIjm9N2S/035h5ZuvnvHPTZC/k5qFqlMo0tjV0pg1G421+Z
68Bb52ignbMGjX/fYmNqdzLn1VY3q1bI4ltL7rMzSyPOylwaMaGGs2MCgYEAxf1Y
9nvfzCUYy6AjWyF0o//62auz01JaddUJgCmMezKg3qRbZkkP6oCkIQX34AHUtOqe
2YoNdLFMaiK2Xopiq+jB/ADpbJDqTZpzlwPnZyuhHX2CZ2J537ozwzkITLb50OGZ
x9EYMIoeQ9zwTe76+mDXWRhCP8lv6DqvYK6JT+sCgYEAk74uk6NX9zx2c1gMs0ja
qzKXZ3TVFubjfdtXFInRLyC71HYgz/EMP4sic3bwOCZ5XvwIieTjSOSd0f79SwXS
c0ijMjdQJtM8+AAgYBCfyIjg/Cdb1v3XHW5xRYtaNzZbikiA3f74MOVGZNdAqccx
6WL86TScQ0uRwKMMfFIynyUCgYBP/OutIRyoUDB1S6wdDCYgf7FrsVfEP1g2lKvy
8FZGC408hbA7YbchaXllshCTmonB0av9fS20gFC77Gw7Q7Nbenysf/3OE7nlQkoC
r9wULWc5D9TYScutHH//FhKJU78XqZ1EQfCA47wYdLFM+wjAkNFPuGJDfPFOJC9o
YvssowKBgCsvXF7S0TPfWqa3JkJuNH8+39GF6jB1rJTBZgpTw3U2Jo3Y8Z5Up1+G
0vRWMWjtiav3gZYjm0+3n4gGVUGxDuc94E+cWoHfWciouKl2/TvDbORtjZ4hMY3z
KseA2F/yzDOYHLTsX3JRKh0hrTr2Ar4f3F81ho85HWcKl34uX8+A
-----END RSA PRIVATE KEY-----
`)

func serverKeyBytes() ([]byte, error) {
	return _serverKey, nil
}

func serverKey() (*asset, error) {
	bytes, err := serverKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "server.key", size: 1675, mode: os.FileMode(436), modTime: time.Unix(1502835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"CA.crt":         caCrt,
	"CA.key":         caKey,
	"client.crt":     clientCrt,
	"client.key":     clientKey,
	"invalid-ca.crt": invalidCaCrt,
	"server.crt":     serverCrt,
	"server.key":     serverKey,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"CA.crt":         &bintree{caCrt, map[string]*bintree{}},
	"CA.key":         &bintree{caKey, map[string]*bintree{}},
	"client.crt":     &bintree{clientCrt, map[string]*bintree{}},
	"client.key":     &bintree{clientKey, map[string]*bintree{}},
	"invalid-ca.crt": &bintree{invalidCaCrt, map[string]*bintree{}},
	"server.crt":     &bintree{serverCrt, map[string]*bintree{}},
	"server.key":     &bintree{serverKey, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
