# Reservoir

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md) [![Build Status](https://travis-ci.org/sbani/reservoir.svg?branch=master)](https://travis-ci.org/sbani/reservoir) [![GoDoc](https://godoc.org/github.com/sbani/reservoir?status.svg)](https://godoc.org/github.com/sbani/reservoir)  


With reservoir you are able to bottle up your function calls in an easy way. Make sure your (or other external services) don't become overloaded by time.

## Installation
```bash
go get github.com/sbani/reservoir
```

## Usage
The example shows us how to print the numbers with a delay.
```go
import (
    "fmt"
    "github.com/sbani/reservoir"
)

func printInt(i int) {
    fmt.Println(i)
}

func main() {
    // Never more than 1 request running at a time.
    // Wait at least 2s between each request.
    limiter := NewReservoir(1, 2 * time.Second)

    for i := 0; i < 5; i++ {
        limiter.add(printInt, i)
    }

    fmt.Println("This is printed first")

    time.Sleep(7 * 2 * time.Second)
}
```

See `reservoir_test.go` for informations about usage as long as there are no docs please.

## Roadmap

- [ ] Add http server (API) functionality

