FROM golang:1.23.4 AS builder
ARG VERSION=0.0.7
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y git ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
RUN cd /app/internal/command/templates && \
    git clone -b 0.0.8 https://github.com/go-cheetah/ansible-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/command-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/gitbook-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/grps-go-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/http-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/mdbook-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/mvc-template.git && \
    git clone -b 0.0.8 https://github.com/go-cheetah/simple-template.git 
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${VERSION}'" -o output/cheetah cmd/cheetah/main.go

FROM debian:bullseye-slim
COPY --from=builder /app/output/cheetah /cheetah
ENTRYPOINT ["/cheetah"]