
# Cheetah - Go脚手架工具

![logo](logo/logo.png)

---
[English](README_en.md)
---

## 猎豹的故事

猎豹（Cheetah）是地球上速度最快的动物，凭借其敏捷的动作和极速的奔跑能力，它能够迅速追上猎物，完成捕猎任务。猎豹的快速和高效不仅是它在自然界中的生存优势，也成为了现代技术领域中许多工具和框架追求的目标——快速、高效和简洁。

正如猎豹凭借速度和敏捷征服大自然，Cheetah 脚手架工具的设计理念就是为了让开发者能够像猎豹一样，快速、敏捷地启动和构建 Go 项目，省去繁琐的配置和重复性工作，让你专注于业务逻辑的实现。

## 项目介绍

`Cheetah` 是一个 Go 语言的脚手架工具，旨在帮助开发者快速构建和启动常见的应用程序架构和代码模板。它通过封装和自动化一系列常见的技术栈和开发流程，使开发者能够高效地启动项目，减少手动配置和重复劳动。

使用规范 [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)

## 工具使用

使用`./cheetah --help`查看帮助

|功能|示例|介绍|demo|
|---|---|---|---|
|【golang】基于gin的mvc框架|`cheetah -i mvc`|[项目结构](./docs/mvc.md)|[demo地址](https://github.com/go-cheetah/mvc-demo)|
|【golang】简单函数|`cheetah -i simple`|[项目结构](./docs/simple.md)|暂未创建|
|【golang】简单http|`cheetah -i http`|[项目结构](./docs/http.md)|暂未创建|
|【golang】简单grpc-go|`cheetah -i grpc`|[项目结构](./docs/grpc-go.md)|暂未创建|
|【golang】命令行方式|`cheetah -i command`|[项目结构](./docs/command.md)|暂未创建|
|【ansible】基于shell|`cheetah -i ansible`|[项目结构](./docs/ansible.md)|暂未创建|
|【文档】gitbook|`cheetah -i gitbook`|[项目结构](./docs/gitbook.md)|暂未创建|
|【文档】mdbook|`cheetah -i mdbook`|[项目结构](./docs/mdbook.md)|暂未创建|

### 安装

二进制方式

```bash
# linux
curl -sSL https://raw.githubusercontent.com/go-cheetah/cheetah/refs/heads/main/install.sh | bash
```

docker

```bash
docker pull gocheetah/cheetah

your_pro_path=/tmp/demo
docker run -it --rm -v ${your_pro_path}:/demo gocheetah/cheetah -i mvc

your_pro_path=/tmp/test
your_project_name=test
docker run -it --rm -v ${your_pro_path}:/${your_project_name} gocheetah/cheetah -i mvc -n $your_project_name
```

## gitee和github仓库地址

> github: https://github.com/go-cheetah/cheetah  
> gitee: https://gitee.com/go-cheetah/cheetah  

## 欢迎提 Issue

感谢您使用 Cheetah 生成代码脚手架！如果在使用过程中遇到任何问题，或者有改进建议，欢迎在 [Issues](https://github.com/go-cheetah/cheetah/issues) 中提交问题。我们会尽快处理并回复。

请提供详细的错误信息或功能需求，以便我们更好地帮助您。

感谢您的支持与反馈！

## 许可证

本项目采用 MIT 许可证，详情请查看 [LICENSE](LICENSE) 文件。
