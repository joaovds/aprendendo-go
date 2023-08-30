package main

import "fmt"


func main() {
  funcaoCallback(func(x int) {
    fmt.Println(x)
  })
}

func funcaoCallback(funcaoRecebida func(int)) {
  fmt.Println("Estou passando uma função como argumento")
  funcaoRecebida(10)
}

