package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

type ToolDateTimeNow struct{}

func (t *ToolDateTimeNow) Tool() mcp.Tool {
	return mcp.NewTool(
		"datetime_now",
		mcp.WithDescription("Returns the current date in Europe/Prague timezone"),
	)
}

func (t *ToolDateTimeNow) Call(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	location, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		return nil, fmt.Errorf("error loading Europe/Prague timezone: %w", err)
	}
	return mcp.NewToolResultText(time.Now().In(location).Format(time.DateTime)), nil
}
