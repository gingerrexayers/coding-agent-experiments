package tools

import (
	"encoding/json"
	"os"
)

// ReadFileInput defines the input parameters for the read_file tool
type ReadFileInput struct {
	Path string `json:"path" jsonschema_description:"The relative path of a file in the working directory."`
}

// ReadFileInputSchema is the schema for the read_file tool's input
var ReadFileInputSchema = GenerateSchema[ReadFileInput]()

// ReadFileDefinition is the complete definition of the read_file tool
var ReadFileDefinition = ToolDefinition{
	Name:        "read_file",
	Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
	InputSchema: ReadFileInputSchema,
	Function:    ReadFile,
}

// ReadFile implements the functionality to read a file's contents
func ReadFile(input json.RawMessage) (string, error) {
	readFileInput := ReadFileInput{}
	err := json.Unmarshal(input, &readFileInput)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(readFileInput.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}