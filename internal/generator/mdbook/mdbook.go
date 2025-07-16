package mdbook

import (
	_ "embed"
)

var (
	//go:embed templates/book.toml.tmpl
	bookTomlContent string
	//go:embed templates/Dockerfile.tmpl
	dockerfileContent string
	//go:embed templates/SUMMARY.md.tmpl
	summaryContent string
	//go:embed templates/README.md.tmpl
	readmeContent string
)

var Files = map[string]string{
	"book.toml":      bookTomlContent,
	"Dockerfile":     dockerfileContent,
	"src/SUMMARY.md": summaryContent,
	"src/README.md":  readmeContent,
}

var Dirs = []string{
	"src",
}

var Extra = map[string]interface{}{}
