package grpc_proxy_server

import (
	// External

	"os"
	"path/filepath"
	"regexp"
	// Internal
)

func getOpenAPIFilesPaths() ([]string, error) {
	var filesPaths []string

	libRegEx, err := regexp.Compile(`^.+\.(json)$`)
	if err != nil {
		return nil, err
	}

	swaggerDir := "third_party/OpenAPI"
	err = filepath.Walk(swaggerDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if libRegEx.MatchString(info.Name()) {
			filesPaths = append(filesPaths, "."+path[len(swaggerDir):])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filesPaths, nil
}
