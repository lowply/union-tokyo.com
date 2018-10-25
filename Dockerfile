FROM golang:latest AS golang
WORKDIR /go/src/app
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main .

FROM alpine:latest
WORKDIR /app
COPY --from=golang /go/src/app/bin/main .
COPY ./templates ./templates
COPY ./public ./public
CMD ["./main"]

MAINTAINER Sho Mizutani <lowply@gmail.com>
