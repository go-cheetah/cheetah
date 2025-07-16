FROM golang:1.23.4 as builder
ARG VERSION=0.0.7
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/go-cheetah/cheetah/cmd/cheetah/app.Version=${VERSION}'" -o output/cheetah

FROM debian:bullseye-slim
COPY --from=builder /app/output/cheetah /cheetah
ENTRYPOINT ["/cheetah"]