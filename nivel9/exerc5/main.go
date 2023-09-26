package main

import (
  "fmt"
  "sync"
  "runtime"
  "sync/atomic"
)

func main() {
  var contagem int64
  contagem = 0
  numerosDeGoRoutines := 100

  wg := sync.WaitGroup{}
  wg.Add(numerosDeGoRoutines)


  for i := 0; i < numerosDeGoRoutines; i++ {
    go func() {
      atomic.AddInt64(&contagem, 1)

      runtime.Gosched()
      fmt.Println(contagem)

      wg.Done()
    }()
  }

  wg.Wait()

  fmt.Println("Contador:", contagem)
}

