package main

import (
  "fmt"
  "encoding/json"
)

type user struct {
  First string `json:"first"`
  Age int `json:"age"`
}

func main() {
  u1 := user{
    First: "James",
    Age: 32,
  }

  u2 := user{
    First: "Moneypenny",
    Age: 27,
  }

  u3 := user{
    First: "M",
    Age: 54,
  }

  users := []user{u1, u2, u3}

  fmt.Println(users)

  usersJson, err := json.Marshal(users)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println("JSON: ", string(usersJson))
}

