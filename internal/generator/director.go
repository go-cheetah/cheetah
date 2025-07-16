package generator

// 创建目录需要单独运行，不放在GeneratorMgr里面
import (
	"os"
	"path"
)

// DirGenerator 目录生成器
type DirGenerator struct {
	Dirs []string
}

// Run 运行生成器
func (d *DirGenerator) Run(opt *Option) (err error) {

	if len(d.Dirs) == 0 {
		d.Dirs = []string{
			"",
		}
	}

	for _, dir := range d.Dirs {
		fullDir := path.Join(opt.AbsProjectPath, dir)
		err = os.MkdirAll(fullDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
