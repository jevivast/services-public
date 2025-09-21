package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

func TestGetFactura_Success(t *testing.T) {
	// External fake server
	external := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page/insert_factura.php" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("idcontrato"); got != "123" {
			t.Fatalf("expected idcontrato=123, got=%s", got)
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "OK")
	}))
	defer external.Close()

	// Handler under test
	h := &Handler{
		Client:  external.Client(),
		BaseURL: external.URL,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/factura/123", nil)
	// Inject mux vars using gorilla/mux helper
	req = mux.SetURLVars(req, map[string]string{"idcontrato": "123"})
	rec := httptest.NewRecorder()

	h.GetFactura(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct == "" {
		t.Fatalf("expected Content-Type to be set")
	}
	body, _ := io.ReadAll(res.Body)
	if string(body) != "OK" {
		t.Fatalf("unexpected body: %s", string(body))
	}
}

func TestGetFactura_MissingParam(t *testing.T) {
	h := &Handler{}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/factura/", nil)
	// no vars injected -> idcontrato will be empty
	rec := httptest.NewRecorder()

	h.GetFactura(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}
}
