package helpers

import (
	// External
	"crypto/tls"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
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
