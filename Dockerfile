FROM golang:1.23.4 AS builder
ARG VERSION=0.0.7
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y wget git ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
RUN cd /app/internal/command/templates && \
    wget https://github.com/go-cheetah/ansible-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv ansible-template-0.0.8 ansible && \
    wget https://github.com/go-cheetah/command-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv command-template-0.0.8 command && \
    wget https://github.com/go-cheetah/gitbook-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv gitbook-template-0.0.8 gitbook && \
    wget https://github.com/go-cheetah/grps-go-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv grps-go-template-0.0.8 grps-go && \
    wget https://github.com/go-cheetah/http-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv http-template-0.0.8 http && \
    wget https://github.com/go-cheetah/mdbook-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv mdbook-template-0.0.8 mdbook && \
    wget https://github.com/go-cheetah/mvc-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv mvc-template-0.0.8 mvc && \
    wget https://github.com/go-cheetah/simple-template/archive/refs/tags/0.0.8.zip && \
    unzip 0.0.8.zip && \
    rm -f 0.0.8.zip && \
    mv simple-template-0.0.8 simple
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${VERSION}'" -o output/cheetah cmd/cheetah/main.go

FROM debian:bullseye-slim
COPY --from=builder /app/output/cheetah /cheetah
ENTRYPOINT ["/cheetah"]