### Bulder
FROM golang:1.15.11-alpine3.13 as builder
RUN apk update
RUN apk add git ca-certificates upx tzdata

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod tidy
# install dependencies

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main main.go;
# compile & pack

### Executable Image
FROM scratch

COPY --from=builder /usr/src/app/bin/main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]