#!/bin/bash

BASH_NAME=$0
SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)
WORK_DIR=$(cd ${SCRIPT_DIR}/..; pwd)

TEMPLATE_DIR=${WORK_DIR}/internal/command/templates

rm -rf ${TEMPLATE_DIR}
mkdir -p ${TEMPLATE_DIR}
git clone -b 0.0.9-develop https://github.com/go-cheetah/ansible-template.git ${TEMPLATE_DIR}/ansible 
git clone -b 0.0.8 https://github.com/go-cheetah/command-template.git ${TEMPLATE_DIR}/command 
git clone -b 0.0.9-develop https://github.com/go-cheetah/gitbook-template.git ${TEMPLATE_DIR}/gitbook 
git clone -b 0.0.8 https://github.com/go-cheetah/grpc-go-template.git ${TEMPLATE_DIR}/grpc-go 
git clone -b 0.0.8 https://github.com/go-cheetah/http-template.git ${TEMPLATE_DIR}/http 
git clone -b 0.0.8 https://github.com/go-cheetah/mdbook-template.git ${TEMPLATE_DIR}/mdbook 
git clone -b 0.0.8 https://github.com/go-cheetah/mvc-template.git ${TEMPLATE_DIR}/mvc 
git clone -b 0.0.8 https://github.com/go-cheetah/simple-template.git ${TEMPLATE_DIR}/simple

cd ${WORK_DIR}
branch=$(git rev-parse --abbrev-ref HEAD)
shot_ID=$(git rev-parse --short HEAD)

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${branch} ${shot_ID}'" -o output/cheetah-linux-amd64 cmd/cheetah/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${branch} ${shot_ID}'" -o output/cheetah-linux-arm64 cmd/cheetah/main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${branch} ${shot_ID}'" -o output/cheetah-windows-amd64 cmd/cheetah/main.go

CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${branch} ${shot_ID}'" -o output/cheetah-darwin-amd64 cmd/cheetah/main.go
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${branch} ${shot_ID}'" -o output/cheetah-darwin-arm64 cmd/cheetah/main.go
