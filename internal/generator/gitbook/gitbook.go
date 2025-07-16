package gitbook

import (
	_ "embed"
)

var (
	//go:embed templates/book.json.tmpl
	bookJsonContent string
	//go:embed templates/Dockerfile.tmpl
	dockerfileContent string
	//go:embed templates/SUMMARY.md.tmpl
	summaryContent string
	//go:embed templates/README.md.tmpl
	readmeContent string
	//go:embed templates/website.css.tmpl
	websiteContent string
)

var Files = map[string]string{
	"book.json":          bookJsonContent,
	"Dockerfile":         dockerfileContent,
	"SUMMARY.md":         summaryContent,
	"README.md":          readmeContent,
	"styles/website.css": websiteContent,
}

var Dirs = []string{
	"styles",
	"assets",
}

func GetExtra() map[string]interface{} {

	return map[string]interface{}{}
}
