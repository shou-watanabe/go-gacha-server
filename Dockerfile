FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.12

COPY --from=builder /src/main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
