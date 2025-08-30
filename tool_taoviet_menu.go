package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

type ToolTaoVietMenu struct{}

func (t *ToolTaoVietMenu) Tool() mcp.Tool {
	return mcp.NewTool(
		"restaurant_taoviet_menu",
		mcp.WithDescription("Returns the web page content of the menu of the Tao Viet Vietnamese restaurant in Brno, Czechia"),
	)
}

func (t *ToolTaoVietMenu) Call(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.taorestaurant.cz/tydenni_menu/nabidka/", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response body: %w", err)
	}

	return mcp.NewToolResultText(string(body)), nil
}
