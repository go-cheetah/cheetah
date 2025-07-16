package command

import (
	"flag"
	"fmt"
	"path"

	"github.com/go-cheetah/cheetah/internal/generator"
	"github.com/go-cheetah/cheetah/internal/generator/ansible"
	"github.com/go-cheetah/cheetah/internal/generator/command"
	"github.com/go-cheetah/cheetah/internal/generator/gitbook"
	"github.com/go-cheetah/cheetah/internal/generator/grpc-go"
	"github.com/go-cheetah/cheetah/internal/generator/http"
	"github.com/go-cheetah/cheetah/internal/generator/mdbook"
	"github.com/go-cheetah/cheetah/internal/generator/mvc"
	"github.com/go-cheetah/cheetah/internal/generator/simple"
)

type Command struct {
	Name string
	Path string
	Type string
}

func New() *Command {
	return &Command{}
}

func (s *Command) Run() {
	generator.Init()
	switch {
	case s.Type == "ansible":
		s.gen(ansible.GetFiles(), ansible.Dirs, ansible.Extra)
	case s.Type == "gitbook":
		s.gen(gitbook.Files, gitbook.Dirs, gitbook.GetExtra())
	case s.Type == "mdbook":
		s.gen(mdbook.Files, mdbook.Dirs, mdbook.Extra)
	case s.Type == "simple":
		s.gen(simple.Files, simple.Dirs, simple.Extra)
	case s.Type == "http":
		s.gen(http.Files, http.Dirs, http.Extra)
	case s.Type == "mvc":
		s.gen(mvc.GetFiles(), mvc.GormDirs, mvc.GormExtra)
	case s.Type == "command":
		s.gen(command.GetFiles(), command.Dirs, command.Extra)
	case s.Type == "grpc":
		s.gen(grpc.GetFiles(), grpc.Dirs, grpc.Extra)
	case s.Type == "":
		flag.PrintDefaults()
	default:
		fmt.Printf("还不支持%v生成器\n", s.Type)
	}
}

func (s *Command) gen(files map[string]string, dirs []string, extra map[string]interface{}) {

	opt := generator.Option{
		AbsProjectPath: path.Join(s.Path, s.Name),
		Title:          s.Name,
		Extra:          extra,
	}
	var dirGen generator.DirGenerator
	dirGen.Dirs = dirs
	var fileGen generator.FileGenerator
	fileGen.Files = files
	// 注册，按顺序来进行添加，先创建目录，再创建文件
	generator.Register(&dirGen)
	generator.Register(&fileGen)
	generator.RunGenerator(&opt)
}
