package http

// 简单的http服务，例如做个metrics的任务
import (
	"embed"
)

//go:embed templates/*
var embedContent embed.FS

var Files = map[string]string{
	"cmd/app/app/options/options.go": readFileGetString("templates/cmd/app/app/options/options.go.tmpl"),
	"cmd/app/app/server.go":          readFileGetString("templates/cmd/app/app/server.go.tmpl"),
	"cmd/app/main.go":                readFileGetString("templates/cmd/app/main.go.tmpl"),

	"internal/cors/config.go": readFileGetString("templates/internal/cors/config.go.tmpl"),
	"internal/cors/cors.go":   readFileGetString("templates/internal/cors/cors.go.tmpl"),
	"internal/cors/utils.go":  readFileGetString("templates/internal/cors/utils.go.tmpl"),

	"internal/handler/health/handler.go": readFileGetString("templates/internal/handler/health/handler.go.tmpl"),
	"internal/handler/health/service.go": readFileGetString("templates/internal/handler/health/service.go.tmpl"),
	"internal/response/common.go":        readFileGetString("templates/internal/response/common.go.tmpl"),
	"internal/response/response.go":      readFileGetString("templates/internal/response/response.go.tmpl"),
	"internal/server/router.go":          readFileGetString("templates/internal/server/router.go.tmpl"),
	"internal/server/server.go":          readFileGetString("templates/internal/server/server.go.tmpl"),

	"README.md": readFileGetString("templates/README.md.tmpl"),
	"go.mod":    readFileGetString("templates/go.mod.tmpl"),
}

var Dirs = []string{
	"cmd/app/app/options",
	"internal/handler/health",
	"internal/cors",
	"internal/response",
	"internal/server",
}

var Extra = map[string]interface{}{}

func readFileGetString(fileName string) string {
	data, err := embedContent.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}
