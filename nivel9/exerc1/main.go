package main

import (
  "fmt"
  "sync"
)

func main() {
  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    fmt.Println("Print 1")

    wg.Done()
  }()


  wg.Add(1)
  go func() {
    fmt.Println("Print 2")

    wg.Done()
  }()

  wg.Wait()
}

