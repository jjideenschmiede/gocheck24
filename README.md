# Library for Check24

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gocheck24.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gocheck24/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gocheck24/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gocheck24)](https://goreportcard.com/report/github.com/jjideenschmiede/gocheck24) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gocheck24?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gocheck24) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gocheck24) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for shopware 6. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.
## Install

```console
go get github.com/jjideenschmiede/gocheck24
```

## How to use?

Currently we have the following functions covered:

- [Orders & Acknowledge](https://github.com/jjideenschmiede/gocheck24#orders-acknowledge)


## Orders & Acknowledge

### Order

If you want to read out an order, you can do this with the following function. The user data is required for this. Only one order can be read out at a time. This must be confirmed afterwards with the acknowledge.

```go
r := gocheck24.Request{
    Username: "partner187",
    Password: "your_password",
}

order, err := gocheck24.Orders(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(order)
}
```

### Acknowledge

With this function you confirm the order. The order number is required for this.

```go
r := gocheck24.Request{
    Username: "partner187",
    Password: "your_password",
}

err := gocheck24.Acknowledge(3, r)
if err != nil {
    log.Fatalln(err)
}
```
