package tools

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/invopop/jsonschema"
	"encoding/json"
)

// ToolDefinition represents a callable tool with its metadata and implementation
type ToolDefinition struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	InputSchema anthropic.ToolInputSchemaParam `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}

// GenerateSchema generates an Anthropic-compatible schema for a given type
func GenerateSchema[T any]() anthropic.ToolInputSchemaParam {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}

	var v T

	schema := reflector.Reflect(v)

	return anthropic.ToolInputSchemaParam{
		Properties: schema.Properties,
	}
}