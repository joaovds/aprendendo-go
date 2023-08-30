package main

import "fmt"


func main() {
  mostrarUmInt := func (x int) func() {
    return func () {
      fmt.Println("Função retornada de uma função! valor recebido:", x)
    }
  }

  mostrarUmInt(10)()
}

