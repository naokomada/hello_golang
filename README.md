# hello_golang

## build and run
```
$ docker-compose build
$ docker-compose up -d
$ docker exec -it hello-golang /bin/sh
(attach hello-golang container)
$ go run src/cmd/hello/main.go
(hello world)
```
