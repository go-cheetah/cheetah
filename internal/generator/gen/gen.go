package gen

import (
	"embed"
	"io/fs"
	"strings"
)

func GetFiles(fEmbed embed.FS) (files map[string]string) {
	entries, err := fEmbed.ReadDir("templates")
	if err != nil {
		panic(err)
	}
	files = make(map[string]string)
	processDirectory(fEmbed, entries, "templates", &files)
	return files
}

func processDirectory(fEmbed embed.FS, entries []fs.DirEntry, basePath string, files *map[string]string) {
	for _, entry := range entries {
		if !entry.IsDir() {
			// 获取文件路径
			filePath := basePath + "/" + entry.Name()
			fileNameSuffix := strings.TrimPrefix(filePath, "templates/") // 去除模板目录前缀
			fileName := strings.TrimSuffix(fileNameSuffix, ".tmpl")      // 去除模板目录前缀

			// 读取文件内容
			content := readFileGetString(fEmbed, filePath)

			// 将文件内容保存到 Files 中
			(*files)[fileName] = content
		} else {
			subDirEntries, err := fEmbed.ReadDir(basePath + "/" + entry.Name())
			if err != nil {
				panic(err)
			}
			processDirectory(fEmbed, subDirEntries, basePath+"/"+entry.Name(), files)
		}
	}
}

func readFileGetString(fEmbed embed.FS, fileName string) string {
	data, err := fEmbed.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}
