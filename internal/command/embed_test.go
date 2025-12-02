package command

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommand_gen(t *testing.T) {
	type fields struct {
		Name           string
		Path           string
		Type           string
		Title          string
		AbsProjectPath string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Command{
				Name:           tt.fields.Name,
				Path:           tt.fields.Path,
				Type:           tt.fields.Type,
				Title:          tt.fields.Title,
				AbsProjectPath: tt.fields.AbsProjectPath,
			}
			if err := s.gen(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Command.gen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGen(t *testing.T) {
	// 注意：因为使用了 all:templates/*，所以根目录是 "templates"
	err := fs.WalkDir(allTemplates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 计算缩进
		depth := strings.Count(path, string(filepath.Separator))
		indent := strings.Repeat("  ", depth)

		if d.IsDir() {
			if path == "." {
				fmt.Println("templates/")
			} else {
				fmt.Printf("%s📁 %s/\n", indent, filepath.Base(path))
			}
		} else {
			info, _ := d.Info()
			fmt.Printf("%s📄 %s (%d bytes)\n", indent, d.Name(), info.Size())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历错误: %v\n", err)
	}
}
