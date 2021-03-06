# Fly [![Build Status](https://travis-ci.org/frozzare/go-fly.svg?branch=master)](https://travis-ci.org/frozzare/go-fly) [![GoDoc](https://godoc.org/github.com/frozzare/go-fly?status.svg)](https://godoc.org/github.com/frozzare/go-fly) [![Go Report Card](https://goreportcard.com/badge/github.com/frozzare/go-fly)](https://goreportcard.com/report/github.com/frozzare/go-fly)

> Work In Progress

Fly is inspired by [Flysystem](https://flysystem.thephpleague.com/). Fly is a filesystem abstraction which allows you to easily swap out a local filesystem for a remote one.

## Installation

```
$ go get -u github.com/frozzare/go-fly
```

## Adapters

* AWS S3
* Local

## Example

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/frozzare/go-fly"
    "github.com/frozzare/go-fly/adapter/flylocal"
)

func main() {
	fs := fly.NewFly(flylocal.NewAdapter("/tmp/fly"))

	if err := fs.Write("test/file.txt", "Hello, world!"); err != nil {
		log.Fatal(err)
	}

	content, err := fs.Read("test/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
	// Hello, world!
}
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
