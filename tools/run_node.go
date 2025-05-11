package tools

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

// RunNodeInput defines the input parameters for the run_node tool
type RunNodeInput struct {
	Script string `json:"script" jsonschema_description:"JavaScript code to execute with Node.js."`
}

// RunNodeInputSchema is the schema for the run_node tool's input
var RunNodeInputSchema = GenerateSchema[RunNodeInput]()

// RunNodeDefinition is the complete definition of the run_node tool
var RunNodeDefinition = ToolDefinition{
	Name:        "run_node",
	Description: "Execute JavaScript code using Node.js and return the output. The code should be a complete, valid JavaScript snippet.",
	InputSchema: RunNodeInputSchema,
	Function:    RunNode,
}

// RunNode implements the functionality to run a Node.js script
func RunNode(input json.RawMessage) (string, error) {
	runNodeInput := RunNodeInput{}
	err := json.Unmarshal(input, &runNodeInput)
	if err != nil {
		return "", err
	}

	// Create a temporary JS file
	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, "script.js")

	// Write the script to the temp file
	err = os.WriteFile(tempFile, []byte(runNodeInput.Script), 0644)
	if err != nil {
		return "", err
	}

	// Execute the script with Node.js
	cmd := exec.Command("node", tempFile)
	output, err := cmd.CombinedOutput()

	// Clean up the temp file
	os.Remove(tempFile)

	if err != nil {
		return string(output) + "\nError: " + err.Error(), err
	}

	return string(output), nil
}
