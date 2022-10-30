FROM golang:1.19.1-alpine
RUN apk update && apk add git

RUN mkdir /go/src/hello_docker
WORKDIR /go/src/hello_docker
ADD . /go/src/hello_docker

RUN go install golang.org/x/tools/gopls@latest