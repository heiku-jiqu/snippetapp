package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/heiku-jiqu/snippetapp/internal/assert"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil) // dummy request
	w := httptest.NewRecorder()
	ping(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, 200)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}

func TestPingEndtoEnd(t *testing.T) {
	app := newTestApplication(t)
	svr := newTestServer(t, app.routes())
	defer svr.Close()

	statusCode, _, body := svr.get(t, "/ping")
	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
