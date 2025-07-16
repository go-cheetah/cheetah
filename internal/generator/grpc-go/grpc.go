package grpc

import (
	"embed"

	"github.com/go-cheetah/cheetah/internal/generator/gen"
)

//go:embed templates
var fEmbed embed.FS

func GetFiles() (files map[string]string) {
	return gen.GetFiles(fEmbed)
}

var Dirs = []string{
	"api",
	"cmd/server/options",
	"conf",
	"docs",
	"internal/conf",
	"internal/data",
	"internal/service",
	"pb/helloworld",
}

var Extra = map[string]interface{}{}
