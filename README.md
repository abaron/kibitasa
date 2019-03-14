# Kibitasa Project
Golang REST API using Gin.

## Requirements
| Require | Minimum Version | Recomended Version |
| ------ | ------ | ------ |
| Go | v1.7 | v1.11 |
| Dep | - | devel |
| Docker | 18.0 | 18.09 |

## Project Installation
- Install Kibitasa `go get github.com/abaron/kibitasa` or clone `$ git clone https://github.com/abaron/kibitasa.git`
`
- Install Gin `go get -u github.com/gin-gonic/gin`
- Go to project dir `$ cd $GOPATH/src/github.com/abaron/kibitasa`

## Start Project
### With Docker
```
$ sh run.sh
```
OR
```
$ sudo docker build . -t go-gin
$ sudo docker run -i -t -p 8080:8080 go-gin
```
### Without Docker
```
$ go run src/main.go
```

## Endpoints
- Addition `GET` `http://localhost:8080/sum/3/5`
- Multiplication `GET` `http://localhost:8080/multiple/-3/5`
- Prime `GET` `http://localhost:8080/prime/5`
- Fibonacci `GET` `http://localhost:8080/fibonacci/5`

## TODO
- [x] Unit test
- [ ] Swagger docs
- [ ] GUI version

Thank you
