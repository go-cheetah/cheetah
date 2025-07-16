package command

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
	"build",
	"cmd/app/app/options",
	"config",
	"docs",
	"internal/config",
	"internal/command",
	"internal/log",
}

func readFileGetString(fileName string) string {
	data, err := fEmbed.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

var Extra = map[string]interface{}{}
