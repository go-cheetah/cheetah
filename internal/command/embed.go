package command

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed all:templates\*
var allTemplates embed.FS

func (s *Command) gen(name string) error {
	rootDir := "templates/" + name
	return s.processDir(rootDir, s.AbsProjectPath)
}

// 递归处理目录
func (s *Command) processDir(srcDir, destDir string) error {
	entries, err := allTemplates.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", srcDir, err)
	}

	for _, entry := range entries {
		srcPath := fmt.Sprintf("%s/%s", srcDir, entry.Name())
		destPath := fmt.Sprintf("%s/%s", destDir, entry.Name())

		if entry.IsDir() {
			// 递归处理子目录
			if err := s.processDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// 处理文件
			if strings.HasSuffix(entry.Name(), ".tmpl") {
				// 模板文件：移除.tmpl后缀并渲染
				destPath = strings.TrimSuffix(destPath, ".tmpl")
				if err := s.renderTemplate(srcPath, destPath); err != nil {
					return fmt.Errorf("failed to render template %s: %w", srcPath, err)
				}
			} else {
				// 普通文件：直接复制
				if err := s.copyFile(srcPath, destPath); err != nil {
					return fmt.Errorf("failed to copy file %s: %w", srcPath, err)
				}
			}
		}
	}
	return nil
}

// 渲染模板文件并写入目标路径
func (s *Command) renderTemplate(srcPath, destPath string) error {
	content, err := allTemplates.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", srcPath, err)
	}

	tmpl, err := template.New(filepath.Base(srcPath)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", srcPath, err)
	}

	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(destPath), err)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", destPath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, s); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", srcPath, err)
	}
	return nil
}

// 直接复制文件到目标路径
func (s *Command) copyFile(srcPath, destPath string) error {
	content, err := allTemplates.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", srcPath, err)
	}

	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(destPath), err)
	}

	return os.WriteFile(destPath, content, 0644)
}
