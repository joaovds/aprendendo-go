package main

import "fmt"

func main() {
  sliceCom10Elementos := []int{ 2, 4, 6, 8, 10, 12, 14, 16, 18, 20 }

  fmt.Printf("slice do 1° ao 3°: %v\n", sliceCom10Elementos[:3])
  fmt.Printf("slice do 5° ao último: %v\n", sliceCom10Elementos[4:])
  fmt.Printf("slice do 2° ao 7°: %v\n", sliceCom10Elementos[1:7])
  fmt.Printf("slice do 3° ao penúltimo°: %v\n", sliceCom10Elementos[2:9])
  fmt.Printf("slice do 3° ao penúltimo°: %v\n", sliceCom10Elementos[2:len(sliceCom10Elementos) - 1])
}

