package json

import (
	"github.com/xeipuuv/gojsonschema"
)

type JsonSchemaValidatorResult struct {
	IsValid bool
	Errors  []gojsonschema.ResultError
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
		Errors:  result.Errors(),
	}, nil
}
