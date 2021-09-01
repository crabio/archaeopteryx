package swagger

import (
	// External
	"embed"
	"io/fs"
	"regexp"
	// Internal
)

func GetOpenAPIFilesPaths(fileSystem embed.FS, dirName string, prefix string) ([]string, error) {
	var filesPaths []string

	libRegEx, err := regexp.Compile(`^.+\.(json)$`)
	if err != nil {
		return nil, err
	}

	if err := fs.WalkDir(fileSystem, dirName, func(path string, d fs.DirEntry, err error) error {
		// Check input error
		if err != nil {
			return err
		}
		// Check that we walk onto file
		if d.IsDir() {
			return nil
		}
		// Check regexp
		if libRegEx.MatchString(d.Name()) {
			// Remove root dir from path
			pathWithoutRoot := path[len(dirName):]
			// Add file path prefix
			pathWithPrefix := prefix + pathWithoutRoot

			filesPaths = append(filesPaths, pathWithPrefix)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// Check if no available files
	if len(filesPaths) == 0 {
		// Append empty to have at lease 1 element on html template for default
		filesPaths = append(filesPaths, "")
	}

	return filesPaths, nil
}
