package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kitd3k/benchzribe/mockapi/internal/handlers"
)

func BenchmarkTestHandler(b *testing.B) {
	handler := http.HandlerFunc(handlers.TestHandler)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest("GET", "/api/test", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)
		}
	})
}
