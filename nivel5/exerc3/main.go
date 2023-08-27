package main

import "fmt"

type pessoa struct {
  nome string
  sobrenome string
  saboresFavoritos []string
}

type veiculo struct {
  portas int
  cor string
}

type caminhote struct {
  veiculo
  tracaoNasQuatro bool
}

type sedan struct {
  veiculo
  modeloLuxo bool
}

func main() {
  caminhote1 := caminhote{
    veiculo: veiculo{
      portas: 2,
      cor: "Azul",
    },
    tracaoNasQuatro: true,
  }

  sedan1 := sedan{
    veiculo: veiculo{
      portas: 4,
      cor: "Vermelho",
    },
    modeloLuxo: false,
  }

  fmt.Println(caminhote1, sedan1)
  fmt.Println(caminhote1.tracaoNasQuatro, sedan1.modeloLuxo)
}

