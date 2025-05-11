package tools

// GetAllTools returns all available tool definitions
func GetAllTools() []ToolDefinition {
	return []ToolDefinition{
		ReadFileDefinition,
		ListFilesDefinition,
		EditFileDefinition,
		RunNodeDefinition,
	}
}