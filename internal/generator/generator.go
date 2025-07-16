package generator

import (
	"fmt"
	"os"
	"text/template"
)

// Option 参数保存
type Option struct {
	// AbsProjectPath 项目路径+项目名称
	AbsProjectPath string
	Title          string
	Dirs           []string
	Extra          map[string]interface{}
}

// GeneratorMgr 生成器管理
type GeneratorMgr struct {
	genMap []Generator
}

var GenMgr GeneratorMgr

// Generator 生成器接口
type Generator interface {
	// Run 生成器run
	Run(opt *Option) error
}

// init 初始化
func Init() {
	GenMgr = GeneratorMgr{
		genMap: []Generator{},
	}
}

// Register 把生成器都注册到map中，然后轮询执行
// 之前是map的方式，现在改成切片，切片就是有序，而map是无序的
func Register(gen Generator) {
	GenMgr.genMap = append(GenMgr.genMap, gen)
}

// writeFile 使用模版文件直接写入文件
func writeFile(tmpl, filePath string, opt *Option) (err error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// t, err := template.ParseFiles(tmpl)
	t, err := template.New(filePath).Parse(tmpl)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	err = t.Execute(file, &opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// RunGenerator 运行所有已经注册的生成器
func RunGenerator(opt *Option) (err error) {
	for _, gen := range GenMgr.genMap {
		err = gen.Run(opt)
		if err != nil {
			fmt.Printf("err: %v", err)
			return err
		}
	}
	return nil
}
