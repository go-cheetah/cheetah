package mvc

// mvc 服务，基于sqlx

import (
	"embed"

	"github.com/go-cheetah/cheetah/internal/generator/gen"
)

//go:embed templates
var fEmbed embed.FS

func GetFiles() (files map[string]string) {
	return gen.GetFiles(fEmbed)
}

var GormDirs = []string{
	"api",
	"build",
	"cmd/app/app/options",
	"config",
	"docs",
	"internal/config",
	"internal/handler/health",
	"internal/model/health",
	"internal/pkg/database",
	"internal/pkg/generator",
	"internal/pkg/middleware/cors",
	"internal/pkg/middleware/log",
	"internal/pkg/response",
	"internal/server",
	// "internal/router",
	"pkg",
}

var GormExtra = map[string]interface{}{}
