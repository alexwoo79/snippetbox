package ui

import (
	"embed"
)

//go:embed "static" "html"
var Files embed.FS

//go:embed "public"
var PublicFiles embed.FS
