package command

import (
	"fmt"
	"path"
)

type Command struct {
	Name           string
	Path           string
	Type           string
	Title          string
	AbsProjectPath string
}

func New() *Command {
	return &Command{}
}

func (s *Command) Run() {

	s.Title = s.Name
	s.AbsProjectPath = path.Join(s.Path, s.Name)
	err := s.gen(s.Type)
	if err != nil {
		fmt.Printf("还不支持%v生成器\n", s.Type)
	}
}
