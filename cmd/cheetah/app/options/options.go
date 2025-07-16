package options

import (
	"github.com/go-cheetah/cheetah/internal/command"
)

type AppOptions struct {
	Name string
	Path string
	Type string
}

func NewAppOptions() *AppOptions {
	o := &AppOptions{}
	return o
}

func (o *AppOptions) NewServer() (*command.Command, error) {
	s := command.New()
	s.Name = o.Name
	s.Path = o.Path
	s.Type = o.Type
	return s, nil
}
