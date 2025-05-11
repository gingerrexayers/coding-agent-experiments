package tools

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ListFilesInput defines the input parameters for the list_file tool
type ListFilesInput struct {
	Path string `json:"path,omitempty" jsonschema_description:"Optional relative path to list files from. Defaults to current directory if not provided."`
}

// ListFilesInputSchema is the schema for the list_file tool's input
var ListFilesInputSchema = GenerateSchema[ListFilesInput]()

// ListFilesDefinition is the complete definition of the list_file tool
var ListFilesDefinition = ToolDefinition{
	Name:        "list_file",
	Description: "List files and directories at a given path. If no path is provided, lists files in the current directory. Do not use this with file names.",
	InputSchema: ListFilesInputSchema,
	Function:    ListFiles,
}

// ListFiles implements the functionality to list files in a directory
func ListFiles(input json.RawMessage) (string, error) {
	listFilesInput := ListFilesInput{}
	err := json.Unmarshal(input, &listFilesInput)
	if err != nil {
		return "", err
	}

	dir := "."
	if listFilesInput.Path != "" {
		dir = listFilesInput.Path
	}

	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if relPath != "." {
			if info.IsDir() {
				files = append(files, relPath+"/")
			} else {
				files = append(files, relPath)
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(files)
	if err != nil {
		return "", err
	}

	return string(result), nil
}