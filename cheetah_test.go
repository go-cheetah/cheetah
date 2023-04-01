package cheetah

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheetah(t *testing.T) {
	c := New()

	c.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET request received"))
	})

	c.Post("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST request received"))
	})

	c.Put("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PUT request received"))
	})

	c.Patch("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PATCH request received"))
	})

	c.Delete("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("DELETE request received"))
	})

	testCases := []struct {
		method   string
		url      string
		expected int
		body     string
	}{
		{http.MethodGet, "/", http.StatusOK, "GET request received"},
		{http.MethodPost, "/", http.StatusOK, "POST request received"},
		{http.MethodPut, "/", http.StatusOK, "PUT request received"},
		{http.MethodPatch, "/", http.StatusOK, "PATCH request received"},
		{http.MethodDelete, "/", http.StatusOK, "DELETE request received"},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		recorder := httptest.NewRecorder()

		c.ServeHTTP(recorder, req)

		if recorder.Code != tc.expected {
			t.Errorf("Expected status code %d, but got %d", tc.expected, recorder.Code)
		}

		if recorder.Body.String() != tc.body {
			t.Errorf("Expected response body %q, but got %q", tc.body, recorder.Body.String())
		}
	}
}

func BenchmarkCheetahGet(b *testing.B) {
	c := New()
	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Cheetah!")
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ServeHTTP(w, req)
	}
}

func BenchmarkCheetahPost(b *testing.B) {
	c := New()
	c.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Cheetah is running!")
	})

	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ServeHTTP(w, req)
	}
}

func BenchmarkCheetahPut(b *testing.B) {
	c := New()
	c.Put("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Cheetah is resting!")
	})

	req, _ := http.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ServeHTTP(w, req)
	}
}

func BenchmarkCheetahPatch(b *testing.B) {
	c := New()
	c.Patch("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Cheetah is napping!")
	})

	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ServeHTTP(w, req)
	}
}

func BenchmarkCheetahDelete(b *testing.B) {
	c := New()
	c.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Cheetah is eating!")
	})

	req, _ := http.NewRequest(http.MethodDelete, "/", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ServeHTTP(w, req)
	}
}
