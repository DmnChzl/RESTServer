# REST Server

## Intro

This is a RESTful BookStore API made with **Go**.

### Mandatory

_Install **MongoDB** and run it before lauching this program._

```
mkdir db
mongo --dbpath db
```

## Process

Repository :

```
git clone https://github.com/mrdoomy/restserver.git
```

Dependencies :

```
go get -v github.com/gorilla/mux
go get -v go.mongodb.org/mongo-driver/mongo
```

Build :

```
go build main.go
```

Enjoy =)

## Docs

- The **Go** Programming Language : [https://golang.org](https://golang.org)
- Gorillax / Mux : [https://github.com/gorilla/mux](https://github.com/gorilla/mux)
- **MongoDB** Go Driver : [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

## License

```
"THE BEER-WARE LICENSE" (Revision 42):
<phk@FreeBSD.ORG> wrote this file. As long as you retain this notice you
can do whatever you want with this stuff. If we meet some day, and you think
this stuff is worth it, you can buy me a beer in return Poul-Henning Kamp
```
