package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/relentlessworks/fakerkit/internal/faker"
)

// mcpRequest represents a Model Context Protocol request.
type mcpRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

// mcpResponse represents a Model Context Protocol response.
type mcpResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  *mcpError   `json:"error,omitempty"`
}

type mcpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// handleMCP handles POST /mcp for Model Context Protocol integration.
func (s *Server) handleMCP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, r, http.StatusMethodNotAllowed,
			"method not allowed",
			"POST a JSON-RPC style request with {\"method\": \"generate\", \"params\": {\"type\": \"person\"}}")
		return
	}

	var req mcpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, mcpResponse{
			Error: &mcpError{Code: -32700, Message: "parse error"},
		})
		return
	}

	switch req.Method {
	case "initialize":
		writeJSON(w, http.StatusOK, mcpResponse{
			Result: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"serverInfo": map[string]string{
					"name":    "fakerkit",
					"version": "1.0.0",
				},
				"capabilities": map[string]interface{}{
					"tools": map[string]bool{"listChanged": false},
				},
			},
		})

	case "tools/list":
		tools := []map[string]interface{}{
			{
				"name":        "generate",
				"description": "Generate fake data of a given type",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"type": map[string]interface{}{
							"type":        "string",
							"description": "The type of fake data to generate",
							"enum":        faker.AvailableTypes(),
						},
						"count": map[string]interface{}{
							"type":        "integer",
							"description": "Number of items to generate (default 1, max 1000)",
						},
					},
					"required": []string{"type"},
				},
			},
			{
				"name":        "batch",
				"description": "Generate multiple types of fake data at once",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"items": map[string]interface{}{
							"type":        "array",
							"items":       map[string]interface{}{"type": "string"},
							"description": "Array of types to generate",
						},
						"count": map[string]interface{}{
							"type":        "integer",
							"description": "Number of iterations (default 1, max 100)",
						},
					},
					"required": []string{"items"},
				},
			},
		}
		writeJSON(w, http.StatusOK, mcpResponse{
			Result: map[string]interface{}{"tools": tools},
		})

	case "tools/call":
		var params struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			writeJSON(w, http.StatusBadRequest, mcpResponse{
				Error: &mcpError{Code: -32602, Message: "invalid params"},
			})
			return
		}

		switch params.Name {
		case "generate":
			var args struct {
				Type  string `json:"type"`
				Count int    `json:"count"`
				Seed  *int64 `json:"seed"`
			}
			if err := json.Unmarshal(params.Arguments, &args); err != nil {
				writeJSON(w, http.StatusBadRequest, mcpResponse{
					Error: &mcpError{Code: -32602, Message: "invalid arguments"},
				})
				return
			}
			if args.Type == "" {
				writeJSON(w, http.StatusBadRequest, mcpResponse{
					Error: &mcpError{Code: -32602, Message: "missing 'type' argument"},
				})
				return
			}
			count := args.Count
			if count <= 0 {
				count = 1
			}
			if count > 1000 {
				count = 1000
			}
			fk := s.faker
			if args.Seed != nil {
				fk = faker.NewSeeded(*args.Seed)
			}
			results := make([]interface{}, count)
			for i := 0; i < count; i++ {
				results[i] = fk.GenerateJSON(args.Type)
			}
			writeJSON(w, http.StatusOK, mcpResponse{
				Result: map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": fmt.Sprintf("Generated %d %s: %+v", count, args.Type, results),
						},
					},
				},
			})

		case "batch":
			var args struct {
				Items []string `json:"items"`
				Count int      `json:"count"`
				Seed  *int64   `json:"seed"`
			}
			if err := json.Unmarshal(params.Arguments, &args); err != nil {
				writeJSON(w, http.StatusBadRequest, mcpResponse{
					Error: &mcpError{Code: -32602, Message: "invalid arguments"},
				})
				return
			}
			if len(args.Items) == 0 {
				writeJSON(w, http.StatusBadRequest, mcpResponse{
					Error: &mcpError{Code: -32602, Message: "missing 'items' argument"},
				})
				return
			}
			count := args.Count
			if count <= 0 {
				count = 1
			}
			if count > 100 {
				count = 100
			}
			fk := s.faker
			if args.Seed != nil {
				fk = faker.NewSeeded(*args.Seed)
			}
			results := make([]map[string]interface{}, count)
			for i := 0; i < count; i++ {
				entry := make(map[string]interface{})
				for _, t := range args.Items {
					entry[t] = fk.GenerateJSON(t)
				}
				results[i] = entry
			}
			writeJSON(w, http.StatusOK, mcpResponse{
				Result: map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": fmt.Sprintf("Generated batch: %+v", results),
						},
					},
				},
			})

		default:
			writeJSON(w, http.StatusBadRequest, mcpResponse{
				Error: &mcpError{Code: -32601, Message: "unknown tool: " + params.Name},
			})
		}

	default:
		writeJSON(w, http.StatusBadRequest, mcpResponse{
			Error: &mcpError{Code: -32601, Message: "unknown method: " + req.Method},
		})
	}
}
