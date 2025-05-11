# Code Editing Agent

A Go-based agent that uses Claude AI to assist with code editing tasks. This agent provides a chat interface to interact with Claude and execute various code-related tools.

## Overview

This repository implements a simple but powerful agent that connects to the Anthropic Claude API and provides several code-related tools:

- **Read File**: Read the contents of files
- **List Files**: List files in a directory
- **Edit File**: Modify file contents
- **Run Node**: Execute JavaScript code using Node.js

The agent acts as a bridge between the user and Claude, allowing Claude to interact with your local filesystem and execute code through a command-line interface.

## Prerequisites

- Go 1.18 or higher
- Node.js (for the `run_node` tool)
- Anthropic API key

## Setup

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/code-editing-agent.git
   cd code-editing-agent
   ```

2. Set your Anthropic API key as an environment variable:

   ```
   export ANTHROPIC_API_KEY=your_api_key_here
   ```

3. Install dependencies:
   ```
   go mod download
   ```

## Usage

Run the agent:

```
go run main.go
```

This will start a chat interface where you can communicate with Claude. The agent will process your messages, send them to Claude, and execute any tools that Claude decides to use.

Example interactions:

- Ask Claude to read a file: "Can you show me the contents of main.go?"
- Ask Claude to modify a file: "Can you add error handling to the Run function in main.go?"
- Ask Claude to run JavaScript code: "Can you run a simple JavaScript function that calculates fibonacci numbers?"

## Tools

### Read File

Reads the contents of a specified file and returns them.

### List Files

Lists files in a specified directory.

### Edit File

Modifies the contents of a specified file.

### Run Node

Executes JavaScript code using Node.js and returns the output.

## Development

To add new tools:

1. Create a new file in the `tools` directory
2. Define the tool's input structure and schema
3. Implement the tool's functionality
4. Add the tool to the list in `tools/tools.go`

## Prior art

Based on https://ampcode.com/how-to-build-an-agent
