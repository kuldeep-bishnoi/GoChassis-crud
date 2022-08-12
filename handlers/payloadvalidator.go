package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/xeipuuv/gojsonschema"
)

func Validate(schemaPath string, payload map[string]interface{}) (interface{}, error) {
	payload_bytes, _ := json.Marshal(payload)
	schema_bytes, _ := ioutil.ReadFile(schemaPath)
	schemaLoader := gojsonschema.NewBytesLoader(schema_bytes)
	documentLoader := gojsonschema.NewBytesLoader(payload_bytes)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err, errors.New("internal server error")
	}
	if result.Valid() {
		return nil, nil
	} else {

		validationErrors := make([]string, 0)
		for _, desc := range result.Errors() {
			validationErrors = append(validationErrors, desc.String())
		}
		return validationErrors, errors.New("invalid payload")
	}
}
