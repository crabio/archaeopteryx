package helpers

import (
	// External
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// Must marshal data struct into the JSON string.
// On any error function will raise Fatal error.
func MustMarshal(data interface{}) string {
	dataJson, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't encode JSON into string")
	}
	return string(dataJson)
}
