package security

import (
	_ "embed"
)

//go:embed cert.pem
var Cert []byte

//go:embed key.pem
var Key []byte
