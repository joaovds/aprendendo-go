package main

import "fmt"


func main() {
  chamada1 := funcaoClosure()
  chamada2 := funcaoClosure()

  fmt.Println(chamada1())
  fmt.Println(chamada1())
  fmt.Println(chamada1())
  fmt.Println(chamada2())
  fmt.Println(chamada2())
  fmt.Println(chamada2())
}

func funcaoClosure() func() int {
  escopoSuperior := 0

  return func() int {
    escopoSuperior++
    return escopoSuperior
  }
}

