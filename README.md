# go-methods

[![Build Status](https://travis-ci.org/m90/go-methods.svg?branch=master)](https://travis-ci.org/m90/go-methods)
[![godoc](https://godoc.org/github.com/m90/go-methods?status.svg)](http://godoc.org/github.com/m90/go-methods)

> limit handler access to certain HTTP verbs

Package `methods` provides HTTP middleware for limiting access to a handler to certain HTTP methods.

### Installation using go get

```sh
$ go get github.com/m90/go-methods
```

### Usage

Wrap a handler using `Allow(...string)`, allowing only the given methods:

```go
var getOnlyHandler := Allow(http.MethodGet)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thanks for GETting!"))
}))
```

Wrap a handler using `Disallow(...string)`, disallowing the given methods:

```go
var noDeleteHandler := Disallow(http.MethodPut, http.MethodPatch, http.MethodDelete)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thanks for not changing!"))
}))
```

### License
MIT © [Frederik Ring](http://www.frederikring.com)
