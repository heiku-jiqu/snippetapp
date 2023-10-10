package main

import (
	"bytes"
	"io"
	"log"
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
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
	svr := httptest.NewTLSServer(app.routes())
	defer svr.Close()

	rs, err := svr.Client().Get(svr.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
