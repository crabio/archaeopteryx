package security

import (
	// External
	"embed"
)

//go:embed *
var CertFS embed.FS
