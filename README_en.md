# Cheetah - Go Scaffold Tool

![logo](logo/logo.png)

---
[中文简体](README.md)
---

## The Story of the Cheetah

The cheetah (Cheetah) is the fastest animal on Earth, capable of catching prey with its agile movements and incredible running speed. The cheetah's speed and efficiency are not only its survival advantages in the wild, but also a goal pursued by many tools and frameworks in modern technology—fast, efficient, and simple.

Just as the cheetah conquers nature with speed and agility, the design concept of the Cheetah scaffold tool is to enable developers to quickly and agilely start and build Go projects, eliminating tedious configuration and repetitive tasks, and allowing them to focus on implementing business logic.

## Project Overview

`Cheetah` is a Go language scaffold tool designed to help developers quickly build and launch common application architectures and code templates. By encapsulating and automating a series of common technology stacks and development workflows, it enables developers to start projects efficiently and reduces manual configuration and repetitive work.

using standards [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## Tool Usage

Use `./cheetah --help` to view help

| Feature | Command Example | Description | Demo |
| --- | --- | --- | --- |
| **Golang** Gin-based MVC Framework | `cheetah create mvc` | [Project Structure](./docs/mvc.md) | [Demo Link](https://github.com/go-cheetah/mvc-demo) |
| **Golang** Simple Functions | `cheetah create simple` | [Project Structure](./docs/simple.md) | Not Created Yet |
| **Golang** Simple HTTP | `cheetah create http` | [Project Structure](./docs/http.md) | Not Created Yet |
| **Golang** Simple grpc-go | `cheetah create grpc` | [Project Structure](./docs/grpc.md) | Not Created Yet |
| **Golang** Command Line Approach | `cheetah create command` | [Project Structure](./docs/command.md) | Not Created Yet |
| **Ansible** Based on Shell | `cheetah create ansible` | [Project Structure](./docs/ansible.md) | Not Created Yet |
| **Documentation** gitbook | `cheetah create gitbook` | [Project Structure](./docs/gitbook.md) | Not Created Yet |
| **Documentation** mdbook | `cheetah create mdbook` | [Project Structure](./docs/mdbook.md) | Not Created Yet |

### guide

binary

```bash
# linux
curl -sSL https://raw.githubusercontent.com/go-cheetah/cheetah/refs/heads/main/install.sh | bash
```

docker

```bash
docker pull gocheetah/cheetah

your_pro_path=/tmp/demo
docker run -it --rm -v ${your_pro_path}:/demo gocheetah/cheetah create mvc 

your_pro_path=/tmp/test
your_project_name=test
docker run -it --rm -v ${your_pro_path}:/${your_project_name} gocheetah/cheetah create mvc  -n $your_project_name
```

## Gitee and GitHub Repository Links

> github: https://github.com/go-cheetah/cheetah  
> gitee: https://gitee.com/go-cheetah/cheetah

## Welcome to Open an Issue

Thank you for using Cheetah, the code generation scaffold! If you encounter any issues or have suggestions for improvements, feel free to open an issue in the [Issues](https://github.com/go-cheetah/cheetah/issues) section. We’ll review and respond as soon as possible.

Please provide detailed error messages or feature requests to help us assist you better.

Thanks for your support and feedback!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
