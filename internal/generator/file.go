package generator

import (
	"fmt"
	"path"
)

// FileGenerator 文件生成器
type FileGenerator struct {
	Files map[string]string
}

// FileGen 文件生成器实例
var FileGen *FileGenerator

// Run 运行生成器
func (f *FileGenerator) Run(opt *Option) error {
	for fileName, tmpl := range f.Files {
		filePath := path.Join(opt.AbsProjectPath, fileName)
		err := writeFile(tmpl, filePath, opt)
		if err != nil {
			fmt.Printf("create files err：%v", err)
			return err
		}
	}
	return nil
}
