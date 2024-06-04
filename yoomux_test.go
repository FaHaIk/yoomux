package yoomux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fahaik/yoomux"
)

func TestYoomux_Use(t *testing.T) {
	y := yoomux.New()

	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	y.Use(middleware)

}

func TestYoomux_Get(t *testing.T) {
	y := yoomux.New()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET / HTTPMethod"))
	}

	y.Get("/path", handler)

	req := httptest.NewRequest(http.MethodGet, "/path", nil)

	res := httptest.NewRecorder()

	y.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}

	expectedBody := "GET / HTTPMethod"
	if res.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, but got %q", expectedBody, res.Body.String())
	}

}

var result string

func BenchmarkYoomuxGet(b *testing.B) {
	var s string
	y := yoomux.New()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET / HTTPMethod"))
	}

	y.Get("/path", handler)

	var req *http.Request
	for i := 0; i < 1000000; i++ {
		req = httptest.NewRequest(http.MethodGet, "/path", nil)
	}

	res := httptest.NewRecorder()

	y.ServeHTTP(res, req)
	result = s
}
func BenchmarkYoomuxPost(b *testing.B) {
	var s string
	y := yoomux.New()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST / HTTPMethod"))
	}

	y.Get("/path", handler)

	req := httptest.NewRequest(http.MethodPost, "/path", nil)

	res := httptest.NewRecorder()

	y.ServeHTTP(res, req)
	result = s
}
