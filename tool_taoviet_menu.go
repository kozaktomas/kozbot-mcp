package main

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

type ToolTaoVietMenu struct{}

func (t *ToolTaoVietMenu) Tool() mcp.Tool {
	desc := `
Returns the web page content of the lunch menu of the Tao Viet Vietnamese restaurant in Brno, Czechia
`

	return mcp.NewTool(
		"restaurant_tao_viet_menu",
		mcp.WithDescription(strings.TrimSpace(desc)),
	)
}

func (t *ToolTaoVietMenu) Call(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	const URL = "https://www.taorestaurant.cz/tydenni_menu/nabidka/"
	return ReturnPage(ctx, URL)
}
