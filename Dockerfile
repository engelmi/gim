FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gim -ldflags="-w -s" cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/gim /usr/bin/gim

ENTRYPOINT ["/usr/bin/gim"]
