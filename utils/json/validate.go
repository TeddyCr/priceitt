package json

import (
	"fmt"
	"strings"
	"github.com/xeipuuv/gojsonschema"
	
)

type JsonSchemaValidatorResult struct {
	IsValid bool
	Errors []gojsonschema.ResultError
}

func ValidateJsonSchema(schemaPath string, entity interface{}) (JsonSchemaValidatorResult, error) {
	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
	documentLoader := gojsonschema.NewGoLoader(entity)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return JsonSchemaValidatorResult{}, err
	}

	return JsonSchemaValidatorResult{
		IsValid: result.Valid(),
		Errors: result.Errors(),
	}, nil
}

func buildHttpPath(version string, entityType string, extra string) string {
	var path string
	switch {
		case strings.Contains(entityType,"create"):
			path = "createEntities"
		case extra != "":
			path = "entities"
		default:
			path = extra
	}
	root := fmt.Sprintf("https://raw.githubusercontent.com/TeddyCr/priceitt/refs/tags/models/%s/schema/%s/%s.json", version, path, entityType)
	return root
}
