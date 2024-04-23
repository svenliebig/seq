# seq

A simple sequence package for go.

## Installation

```bash
go get github.com/svenliebig/seq
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/svenliebig/seq"
)

func main() {
    // Create a sequence of integers from 1 to 10
    s := seq.Ints(1, 10)
    for v := range s.Iterator() {
        fmt.Println(s)
    }

    // Filter the sequence by even numbers
    f := seq.Filter(s, func(v int) bool {
        return v % 2 == 0
    })
    for v := range f.Iterator() {
        fmt.Println(f)
    }

    // Map the sequence to quoted strings
    m := seq.Map(f, func(v int) string {
        return fmt.Sprintf("%q", v)
    })
    for v := range m.Iterator() {
        fmt.Println(m)
    }
}
```

## Implement your own sequence

```go
// todo
```

