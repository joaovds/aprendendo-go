package main

import "fmt"

func main() {
  const (
    a = 2023 + iota
    b
    c
    d
  )

  fmt.Println(a, b, c, d)
}

