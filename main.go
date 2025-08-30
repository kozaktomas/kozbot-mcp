package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Tomas Kozak personal assistant",
		"1.0.0",
	)

	tools := []Tool{
		&ToolDateTimeNow{},
		&ToolTaoVietMenu{},
	}

	for _, t := range tools {
		s.AddTool(t.Tool(), t.Call)
	}

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
