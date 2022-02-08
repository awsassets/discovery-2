# Discovery
[![Go Report Card](https://goreportcard.com/badge/github.com/nothinux/discovery)](https://goreportcard.com/report/github.com/nothinux/discovery)  ![test status](https://github.com/nothinux/discovery/actions/workflows/test.yml/badge.svg?branch=master)

DNS service discovery library for Go


## Documentation
see [pkg.go.dev](https://pkg.go.dev/github.com/nothinux/discovery)


## Installation

```
$ go get github.com/nothinux/discovery
```

### Getting Started
``` go
package main

import (
    "github.com/nothinux/discovery"
    "log"
)

func main() {
    d := discovery.NewResolver("127.0.0.1:8600")
    services, err := d.Discover(
        context.Background(),
        "",
        "",
        "redis.service.dc1.consul",
    )
    if err != nil {
        log.Fatal(err)
    }
    
    for _, s := range services {
        log.Println(s.Target, s.Address, s.Port)
    }
}
```

