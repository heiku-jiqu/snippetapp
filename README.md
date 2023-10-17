# snippetapp
Building a Golang webserver to store / retrieve text snippets.

## Motivation

Learning about backend development concepts while familiarizing with Golang:

- Routing and http handlers
- Middleware for logging, auth, etc.
- Interacting with Postgres
- HTML Templating
- Processing and validating forms
- Sessions with cookies
- Security (server headers, TLS, CSRF prevention, password storage with bcrypt)
- User authorization and authentication
- Golang Request Context to pass data to downstream handlers
- Flash messages
- Golang embedded files using embed.FS
- Testing http handlers
- Mocking data models for testing
- Integrating testing with test database

## Using

1. `mkdir tls`
1. `cd tls`
1. Using `crypto/tls` `generate_cert.go` tool, run `go run /path/to/generate_cert.go --rsa-bites=2048 --host=localhost`
1. Install postgres and run `sql/snippets.sql` setup script to create necessary tables and user
1. `go run ./cmd/web`
