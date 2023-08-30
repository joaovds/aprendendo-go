package main

import "fmt"


func main() {
  func (x int) {
    fmt.Println("Função anônima executada! valor recebido:", x)
  }(10)
}

