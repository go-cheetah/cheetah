package ansible

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
	"config",
	"playbooks",
	"core",
	"playbooks/roles/base/tasks",
}

var Extra = map[string]interface{}{}
