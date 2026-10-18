package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoot(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "fakerkit") {
		t.Errorf("root should mention fakerkit: %s", w.Body.String())
	}
}

func TestRootNotFound(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "hint:") {
		t.Errorf("404 should include a hint: %s", w.Body.String())
	}
}

func TestHelp(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/help", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "fakerkit") {
		t.Errorf("help should mention fakerkit")
	}
	if !strings.Contains(body, "ENDPOINTS") {
		t.Errorf("help should list endpoints")
	}
	if !strings.Contains(body, "GENERATION TYPES") {
		t.Errorf("help should list generation types")
	}
}

func TestAgentMD(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/.well-known/agent.md", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestTypes(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/types", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "person") {
		t.Errorf("types should include person: %s", body)
	}
	if !strings.Contains(body, "address") {
		t.Errorf("types should include address: %s", body)
	}
}

func TestTypesJSON(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/types", nil)
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if result["types"] == nil {
		t.Error("JSON response should include types")
	}
}

func TestGeneratePerson(t *testing.T) {
	srv := New()
	body := `{"type":"person"}`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	resp := w.Body.String()
	if !strings.Contains(resp, "name=") {
		t.Errorf("person response should contain name=: %s", resp)
	}
	if !strings.Contains(resp, "email=") {
		t.Errorf("person response should contain email=: %s", resp)
	}
}

func TestGeneratePersonJSON(t *testing.T) {
	srv := New()
	body := `{"type":"person"}`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if result["first_name"] == nil {
		t.Error("JSON person should have first_name")
	}
}

func TestGenerateMultiple(t *testing.T) {
	srv := New()
	body := `{"type":"person.name","count":5}`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	lines := strings.Split(strings.TrimSpace(w.Body.String()), "\n")
	if len(lines) != 5 {
		t.Errorf("expected 5 lines, got %d: %s", len(lines), w.Body.String())
	}
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			t.Errorf("multi-generate lines should contain #: %s", line)
		}
	}
}

func TestGenerateSeeded(t *testing.T) {
	srv := New()
	body1 := `{"type":"person.name","seed":42}`
	body2 := `{"type":"person.name","seed":42}`

	req1 := httptest.NewRequest("POST", "/generate", strings.NewReader(body1))
	w1 := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w1, req1)

	req2 := httptest.NewRequest("POST", "/generate", strings.NewReader(body2))
	w2 := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w2, req2)

	if w1.Body.String() != w2.Body.String() {
		t.Errorf("seeded generation should be deterministic:\n%s\n!=\n%s", w1.Body.String(), w2.Body.String())
	}
}

func TestGenerateMissingType(t *testing.T) {
	srv := New()
	body := `{}`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "hint:") {
		t.Errorf("error should include hint: %s", w.Body.String())
	}
}

func TestGenerateInvalidType(t *testing.T) {
	srv := New()
	body := `{"type":"nonexistent"}`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "hint:") {
		t.Errorf("error should include hint: %s", w.Body.String())
	}
}

func TestGenerateInvalidJSON(t *testing.T) {
	srv := New()
	body := `not json`
	req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestGenerateWrongMethod(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/generate", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}

func TestBatch(t *testing.T) {
	srv := New()
	body := `{"items":["person","address","company"],"count":2}`
	req := httptest.NewRequest("POST", "/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	resp := w.Body.String()
	if !strings.Contains(resp, "person") {
		t.Errorf("batch should contain person: %s", resp)
	}
	if !strings.Contains(resp, "address") {
		t.Errorf("batch should contain address: %s", resp)
	}
	if !strings.Contains(resp, "company") {
		t.Errorf("batch should contain company: %s", resp)
	}
	if !strings.Contains(resp, "---") {
		t.Errorf("batch with count>1 should have separator: %s", resp)
	}
}

func TestBatchJSON(t *testing.T) {
	srv := New()
	body := `{"items":["person","address"]}`
	req := httptest.NewRequest("POST", "/batch", strings.NewReader(body))
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if result["result"] == nil {
		t.Error("JSON batch should have result")
	}
}

func TestBatchMissingItems(t *testing.T) {
	srv := New()
	body := `{}`
	req := httptest.NewRequest("POST", "/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestBatchInvalidType(t *testing.T) {
	srv := New()
	body := `{"items":["nonexistent"]}`
	req := httptest.NewRequest("POST", "/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestMCPInitialize(t *testing.T) {
	srv := New()
	body := `{"method":"initialize"}`
	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if result["result"] == nil {
		t.Error("MCP initialize should return result")
	}
}

func TestMCPToolsList(t *testing.T) {
	srv := New()
	body := `{"method":"tools/list"}`
	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	res := result["result"].(map[string]interface{})
	tools := res["tools"].([]interface{})
	if len(tools) < 2 {
		t.Errorf("expected at least 2 tools, got %d", len(tools))
	}
}

func TestMCPToolsCallGenerate(t *testing.T) {
	srv := New()
	body := `{"method":"tools/call","params":{"name":"generate","arguments":{"type":"person.name"}}}`
	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(body))
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if result["result"] == nil {
		t.Error("MCP tools/call should return result")
	}
}

func TestMCPWrongMethod(t *testing.T) {
	srv := New()
	req := httptest.NewRequest("GET", "/mcp", nil)
	w := httptest.NewRecorder()
	srv.Routes().ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}

func TestGenerateAllTypes(t *testing.T) {
	srv := New()
	// Test a representative subset of types
	types := []string{
		"person", "person.name", "person.first", "person.last",
		"address", "address.street", "address.city", "address.state", "address.zip", "address.country",
		"company", "company.name", "company.catchphrase", "company.buzzwords",
		"internet", "internet.username", "internet.domain", "internet.url", "internet.ip", "internet.useragent",
		"commerce", "commerce.product", "commerce.price", "commerce.category",
		"finance", "finance.creditcard", "finance.iban", "finance.btc",
		"date", "date.past", "date.future", "date.recent", "date.birthday",
		"system", "system.filepath", "system.mimetype", "system.extension",
	}
	for _, typ := range types {
		body := `{"type":"` + typ + `"}`
		req := httptest.NewRequest("POST", "/generate", strings.NewReader(body))
		w := httptest.NewRecorder()
		srv.Routes().ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Generate(%q) expected 200, got %d: %s", typ, w.Code, w.Body.String())
		}
		if strings.TrimSpace(w.Body.String()) == "" {
			t.Errorf("Generate(%q) returned empty body", typ)
		}
	}
}
