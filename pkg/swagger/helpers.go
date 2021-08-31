package swagger

import (
	// External
	"os"
	"path/filepath"
	"regexp"
	// Internal
)

func getOpenAPIFilesPaths(dir string) ([]string, error) {
	var filesPaths []string

	libRegEx, err := regexp.Compile(`^.+\.(json)$`)
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if libRegEx.MatchString(info.Name()) {
			filesPaths = append(filesPaths, path[len(dir):])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Check if no available files
	if len(filesPaths) == 0 {
		// Append empty to have at lease 1 element on html template for default
		filesPaths = append(filesPaths, "")
	}

	return filesPaths, nil
}
