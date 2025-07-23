#!/bin/bash

BASH_NAME=$0
SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)

cd ${SCRIPT_DIR}/../internal/command/templates 
git clone -b 0.0.8 https://github.com/go-cheetah/ansible-template.git ansible 
git clone -b 0.0.8 https://github.com/go-cheetah/command-template.git command 
git clone -b 0.0.8 https://github.com/go-cheetah/gitbook-template.git gitbook 
git clone -b 0.0.8 https://github.com/go-cheetah/grpc-go-template.git grpc-go 
git clone -b 0.0.8 https://github.com/go-cheetah/http-template.git http 
git clone -b 0.0.8 https://github.com/go-cheetah/mdbook-template.git mdbook 
git clone -b 0.0.8 https://github.com/go-cheetah/mvc-template.git mvc 
git clone -b 0.0.8 https://github.com/go-cheetah/simple-template.git simple
cd - 

shot_ID=$(git rev-parse --short HEAD)

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${shot_ID}'" -o output/cheetah-linux-amd64 cmd/cheetah/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${shot_ID}'" -o output/cheetah-linux-arm64 cmd/cheetah/main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${shot_ID}'" -o output/cheetah-windows-amd64 cmd/cheetah/main.go

CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${shot_ID}'" -o output/cheetah-darwin-amd64 cmd/cheetah/main.go
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${shot_ID}'" -o output/cheetah-darwin-arm64 cmd/cheetah/main.go