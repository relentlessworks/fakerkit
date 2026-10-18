package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/relentlessworks/fakerkit/internal/faker"
)

// Server holds the HTTP server state.
type Server struct {
	faker *faker.Faker
	mu    sync.Mutex
	seed  int64
}

// New creates a new Server.
func New() *Server {
	return &Server{
		faker: faker.New(),
	}
}

// Routes returns the HTTP handler with all routes registered.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/help", s.handleHelp)
	mux.HandleFunc("/.well-known/agent.md", s.handleHelp)
	mux.HandleFunc("/types", s.handleTypes)
	mux.HandleFunc("/generate", s.handleGenerate)
	mux.HandleFunc("/batch", s.handleBatch)
	mux.HandleFunc("/mcp", s.handleMCP)
	mux.HandleFunc("/", s.handleRoot)

	return s.withMiddleware(mux)
}

// withMiddleware wraps the handler with common middleware.
func (s *Server) withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "fakerkit")
		next.ServeHTTP(w, r)
	})
}

// wantsJSON checks if the client wants JSON output.
func wantsJSON(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "application/json") {
		return true
	}
	return r.URL.Query().Get("format") == "json"
}

// writeText writes a plain text response.
func writeText(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	fmt.Fprint(w, body)
}

// writeJSON writes a JSON response.
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// writeError writes an error response with an instructive hint.
func writeError(w http.ResponseWriter, r *http.Request, status int, msg, hint string) {
	if wantsJSON(r) {
		writeJSON(w, status, map[string]string{
			"error": msg,
			"hint":  hint,
		})
		return
	}
	writeText(w, status, fmt.Sprintf("error: %s\nhint: %s\n", msg, hint))
}

// handleRoot handles the root path.
func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeError(w, r, http.StatusNotFound,
			"unknown endpoint: "+r.URL.Path,
			"GET /help for the full list of available endpoints")
		return
	}
	writeText(w, http.StatusOK, "fakerkit — fake data generation service\nGET /help for usage\n")
}

// handleHelp serves the operating manual.
func (s *Server) handleHelp(w http.ResponseWriter, r *http.Request) {
	writeText(w, http.StatusOK, helpText)
}

// handleTypes lists all available generation types.
func (s *Server) handleTypes(w http.ResponseWriter, r *http.Request) {
	types := faker.AvailableTypes()
	if wantsJSON(r) {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"types":  types,
			"count":  len(types),
		})
		return
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("available types (%d):\n", len(types)))
	for _, t := range types {
		sb.WriteString("  " + t + "\n")
	}
	writeText(w, http.StatusOK, sb.String())
}

// generateRequest is the JSON body for /generate.
type generateRequest struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Seed  *int64 `json:"seed"`
}

// handleGenerate handles POST /generate.
func (s *Server) handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, r, http.StatusMethodNotAllowed,
			"method not allowed",
			"POST a JSON body with {\"type\": \"person\", \"count\": 1}")
		return
	}

	var req generateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest,
			"invalid JSON body",
			"send a JSON body like {\"type\": \"person\", \"count\": 1}")
		return
	}

	if req.Type == "" {
		writeError(w, r, http.StatusBadRequest,
			"missing 'type' field",
			"specify a type like \"person\", \"address\", \"company\". GET /types for the full list")
		return
	}

	// Validate type
	valid := false
	for _, t := range faker.AvailableTypes() {
		if t == req.Type {
			valid = true
			break
		}
	}
	if !valid {
		writeError(w, r, http.StatusBadRequest,
			fmt.Sprintf("unknown type: %s", req.Type),
			"GET /types for the full list of available types")
		return
	}

	count := req.Count
	if count <= 0 {
		count = 1
	}
	if count > 1000 {
		count = 1000
	}

	// Use seeded faker if seed provided
	fk := s.faker
	if req.Seed != nil {
		fk = faker.NewSeeded(*req.Seed)
	}

	if wantsJSON(r) {
		results := make([]interface{}, count)
		for i := 0; i < count; i++ {
			results[i] = fk.GenerateJSON(req.Type)
		}
		if count == 1 {
			writeJSON(w, http.StatusOK, results[0])
		} else {
			writeJSON(w, http.StatusOK, map[string]interface{}{
				"type":   req.Type,
				"count":  count,
				"result": results,
			})
		}
		return
	}

	var sb strings.Builder
	for i := 0; i < count; i++ {
		val := fk.Generate(req.Type)
		if count > 1 {
			sb.WriteString(fmt.Sprintf("%s  #%d  %s\n", req.Type, i+1, val))
		} else {
			sb.WriteString(fmt.Sprintf("%s  %s\n", req.Type, val))
		}
	}
	writeText(w, http.StatusOK, sb.String())
}

// batchRequest is the JSON body for /batch.
type batchRequest struct {
	Items []string `json:"items"`
	Count int      `json:"count"`
	Seed  *int64   `json:"seed"`
}

// handleBatch handles POST /batch.
func (s *Server) handleBatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, r, http.StatusMethodNotAllowed,
			"method not allowed",
			"POST a JSON body with {\"items\": [\"person\", \"address\"], \"count\": 1}")
		return
	}

	var req batchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest,
			"invalid JSON body",
			"send a JSON body like {\"items\": [\"person\", \"address\"], \"count\": 1}")
		return
	}

	if len(req.Items) == 0 {
		writeError(w, r, http.StatusBadRequest,
			"missing 'items' field",
			"provide an array of types like [\"person\", \"address\", \"company\"]. GET /types for available types")
		return
	}

	// Validate all types
	for _, t := range req.Items {
		valid := false
		for _, at := range faker.AvailableTypes() {
			if at == t {
				valid = true
				break
			}
		}
		if !valid {
			writeError(w, r, http.StatusBadRequest,
				fmt.Sprintf("unknown type: %s", t),
				"GET /types for the full list of available types")
			return
		}
	}

	count := req.Count
	if count <= 0 {
		count = 1
	}
	if count > 100 {
		count = 100
	}

	fk := s.faker
	if req.Seed != nil {
		fk = faker.NewSeeded(*req.Seed)
	}

	if wantsJSON(r) {
		results := make([]map[string]interface{}, count)
		for i := 0; i < count; i++ {
			entry := make(map[string]interface{})
			for _, t := range req.Items {
				entry[t] = fk.GenerateJSON(t)
			}
			results[i] = entry
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"count":  count,
			"items":  req.Items,
			"result": results,
		})
		return
	}

	var sb strings.Builder
	for i := 0; i < count; i++ {
		for _, t := range req.Items {
			val := fk.Generate(t)
			sb.WriteString(fmt.Sprintf("%s  %s\n", t, val))
		}
		if i < count-1 {
			sb.WriteString("---\n")
		}
	}
	writeText(w, http.StatusOK, sb.String())
}
