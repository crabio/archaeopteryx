package helpers

import (
	// External
	"crypto/tls"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
)

func CreateTlsConfig(c *config.Config) (*tls.Config, error) {
	serverCert, err := tls.X509KeyPair(c.Secutiry.Cert, c.Secutiry.Key)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
		// TODO Valid only for local test
		InsecureSkipVerify: true,
	}

	return tlsConfig, nil
}
