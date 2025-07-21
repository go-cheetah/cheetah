FROM golang:1.23.4 AS builder
ARG VERSION=0.0.7
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y wget git ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${VERSION}'" -o output/cheetah cmd/cheetah/main.go

FROM debian:bullseye-slim
COPY --from=builder /app/output/cheetah /cheetah
ENTRYPOINT ["/cheetah"]