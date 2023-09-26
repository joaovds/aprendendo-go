package main

import (
  "fmt"
  "sync"
  "runtime"
)

func main() {
  contagem := 0
  numerosDeGoRoutines := 1000

  wg := sync.WaitGroup{}
  wg.Add(numerosDeGoRoutines)

  mutex := sync.Mutex{}

  for i := 0; i < numerosDeGoRoutines; i++ {
    go func() {
      mutex.Lock()
      atualContagem := contagem

      runtime.Gosched()

      atualContagem++
      contagem = atualContagem
      mutex.Unlock()

      wg.Done()
    }()
  }

  wg.Wait()

  fmt.Println("Contador:", contagem)
}

