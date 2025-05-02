FROM golang:1.24-alpine as Build
COPY . .
RUN GOPATH= go build -o /main cmd/main.go

FROM alpine:latest
COPY --from=Build /main .
EXPOSE 8080
ENTRYPOINT ["./main"]