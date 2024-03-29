FROM golang:1.22-alpine AS builder

COPY . /github.com/FreylGit/auth/sourse/
WORKDIR /github.com/FreylGit/auth/sourse/

RUN go mod download
RUN go build -o ./bin/auth_server ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/FreylGit/auth/sourse/bin/auth_server .
COPY .env .


CMD ["./auth_server"]