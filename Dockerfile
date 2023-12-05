FROM golang:1.20.5-alpine as builder
RUN apk update

WORKDIR /usr/src/app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o bin/main main.go

### Executable Image
FROM alpine

COPY --from=builder /usr/src/app/bin/main ./main

ENTRYPOINT ["./main"]

