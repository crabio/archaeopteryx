package helpers

import (
	// External
	"crypto/tls"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
)

// Constants for testing ONLY!
// Don't use these certificates in your apps
var (
	MockCertBytes = []byte(`-----BEGIN CERTIFICATE-----
MIIB4TCCAYugAwIBAgIUA7V9TBkILumPXFXBXpvl2N3/w2kwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA5MDgyMTExMThaFw0yMTEw
MDgyMTExMThaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwXDANBgkqhkiG9w0BAQEF
AANLADBIAkEAq0tuIM5bDAmiI+cwpom70XGovlfoPAYf0/xzPNaE7vrajzeqp557
Gr3tS//5D7mv9URdJJLmQS8Bi6IE+MiwNQIDAQABo1MwUTAdBgNVHQ4EFgQUZxRI
pEalJuPBYdU0xOl3Qg3GNpswHwYDVR0jBBgwFoAUZxRIpEalJuPBYdU0xOl3Qg3G
NpswDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAANBAJzvfb9lvkWlkRMY
NTJKdeyQKRBXzrfkkHjJRCQLaWDTF12KkIZaVjGjuZ4RUZHfy5Cjday1bjH7nt3O
uptwg8o=
-----END CERTIFICATE-----`)

	MockKeyBytes = []byte(`-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAq0tuIM5bDAmiI+cw
pom70XGovlfoPAYf0/xzPNaE7vrajzeqp557Gr3tS//5D7mv9URdJJLmQS8Bi6IE
+MiwNQIDAQABAkEAnfT23wWra9ROUjFU6Z3FNoRLCQtjOkajfwYi9g0TlJL4Vbt9
Ea11KLOiV5wwGESwYlJOlokDMHT7NuGkvpA0QQIhANJYLQx+E9VL6TWTb1ESeZcZ
ANEjLJIcX1S5Ujnu3JCdAiEA0HlyOBpdrKtsQLnIZSSQ28F0kiCPxm8cQ4o4KF2y
znkCIBE9DrwWXRO+++bbJWVUiUh70RhStKVo09tCsN10mPj1AiBV5nTF4TdP+qJ0
WRjVdCesJR5fR8N2RDolKkLRfyo6IQIgBq9N9MGfLD6ku2VN6jJGOarPYfgksb1m
RR0T6ErpcV4=
-----END PRIVATE KEY-----`)
)

func CreateTlsConfig(c *config.Config) (*tls.Config, error) {
	certBytes, err := c.Secutiry.CertFS.ReadFile(*c.Secutiry.CertName)
	if err != nil {
		return nil, err
	}
	keyBytes, err := c.Secutiry.CertFS.ReadFile(*c.Secutiry.KeyName)
	if err != nil {
		return nil, err
	}
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
		// TODO Valid only for local test
		InsecureSkipVerify: true,
	}

	return tlsConfig, nil
}
