package docs_swagger

import (
	"embed"
)

//go:embed *.html
var SwaggerTmpl embed.FS
