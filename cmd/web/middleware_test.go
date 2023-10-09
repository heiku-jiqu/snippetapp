package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/heiku-jiqu/snippetapp/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	secureHeaders(handler).ServeHTTP(w, req)
	res := w.Result()

	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, res.Header.Get("Content-Security-Policy"), expectedValue)

	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, res.Header.Get("Referrer-Policy"), expectedValue)
	expectedValue = "nosniff"
	assert.Equal(t, res.Header.Get("X-Content-Type-Options"), expectedValue)
	expectedValue = "deny"
	assert.Equal(t, res.Header.Get("X-Frame-Options"), expectedValue)
	expectedValue = "0"
	assert.Equal(t, res.Header.Get("X-XSS-Protection"), expectedValue)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
