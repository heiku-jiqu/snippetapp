package ui

import (
	"embed"
)

// paths using go:embed directive are relative to current source code file
// and cannot contain . or .. or being or end with /
// but it is ok to embed directories or subdirectories
// files that being with . or _ are not embedded unless "all:" prefix is used
// e.g. "all:html"
// Root of embed.FS will be dir that contains the go:embed directive
// i.e. ui directory in this case

//go:embed "html" "static"
var Files embed.FS
