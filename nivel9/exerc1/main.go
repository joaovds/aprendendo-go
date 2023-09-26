package main

import (
  "fmt"
  "sync"
)

func main() {
  wg := sync.WaitGroup{}

  wg.Add(2)
  go func() {
    fmt.Println("Print 1")

    wg.Done()
  }()


  go func() {
    fmt.Println("Print 2")

    wg.Done()
  }()

  wg.Wait()
}

