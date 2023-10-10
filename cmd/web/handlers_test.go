package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)
	svr := newTestServer(t, app.routes())
	defer svr.Close()

	t.Run("Successful snippetview", func(t *testing.T) {
		statusCode, _, body := svr.get(t, "/snippet/view/1")
		assert.Equal(t, statusCode, http.StatusOK)
		assert.StringContains(t, body, "An old silent pond")
	})
	t.Run("Not found snippetview", func(t *testing.T) {
		statusCode, _, _ := svr.get(t, "/snippet/view/2")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
	t.Run("Negative ID", func(t *testing.T) {
		statusCode, _, _ := svr.get(t, "/snippet/view/-1")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
	t.Run("Decimal ID", func(t *testing.T) {
		statusCode, _, _ := svr.get(t, "/snippet/view/2.2")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
	t.Run("String ID", func(t *testing.T) {
		statusCode, _, _ := svr.get(t, "/snippet/view/foo")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
	t.Run("Empty ID", func(t *testing.T) {
		statusCode, _, _ := svr.get(t, "/snippet/view/")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
}

func TestUserSignupPost(t *testing.T) {
	app := newTestApplication(t)
	svr := newTestServer(t, app.routes())
	defer svr.Close()

	_, _, body := svr.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Bob"
		validPassword = "pa$$word"
		validEmail    = "bob@example.com"
		formTag       = "<form action='/user/signup' method='POST' novalidate>"
	)

	testCases := []struct {
		name            string
		userName        string
		userPassword    string
		userEmail       string
		csrfToken       string
		expectedCode    int
		expectedFormTag string
	}{
		{
			name:         "Valid Signup",
			userName:     validName,
			userPassword: validPassword,
			userEmail:    validEmail,
			csrfToken:    validCSRFToken,
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF",
			userName:     validName,
			userPassword: validPassword,
			userEmail:    validEmail,
			csrfToken:    "InvalidCSRF",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:            "Empty Name",
			userName:        "",
			userPassword:    validPassword,
			userEmail:       validEmail,
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Empty Password",
			userName:        validName,
			userPassword:    "",
			userEmail:       validEmail,
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Empty Email",
			userName:        validName,
			userPassword:    validPassword,
			userEmail:       "",
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Invalid Email Format",
			userName:        validName,
			userPassword:    validPassword,
			userEmail:       "bob@example.",
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Password less than 8 Chars Long",
			userName:        validName,
			userPassword:    "abcd",
			userEmail:       validEmail,
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Email already in use",
			userName:        validName,
			userPassword:    validPassword,
			userEmail:       "dupe@example.com",
			csrfToken:       validCSRFToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			vals := url.Values{}
			vals.Set("name", c.userName)
			vals.Set("email", c.userEmail)
			vals.Set("password", c.userPassword)
			vals.Set("csrf_token", c.csrfToken)

			code, _, body := svr.postForm(t, "/user/signup", vals)
			assert.Equal(t, code, c.expectedCode)

			if c.expectedFormTag != "" {
				assert.StringContains(t, body, c.expectedFormTag)
			}
		})
	}
}
