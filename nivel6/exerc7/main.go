package main

import "fmt"


func main() {
  mostrarUmInt := func (x int) {
    fmt.Println("Função wm variável! valor recebido:", x)
  }

  mostrarUmInt(10)
}

