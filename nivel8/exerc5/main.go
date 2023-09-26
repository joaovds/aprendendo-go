package main

import (
  "fmt"
  "sort"
)

type user struct {
  First   string
  Last    string
  Age     int
  Sayings []string
}

type byLast []user

func (bL byLast) Len() int { return len(bL) }
func (bL byLast) Less(i, j int) bool { return bL[i].Last < bL[j].Last }
func (bL byLast) Swap(i, j int) { bL[i], bL[j] = bL[j], bL[i] }

type byAge []user

func (bA byAge) Len() int { return len(bA) }
func (bA byAge) Less(i, j int) bool { return bA[i].Age < bA[j].Age }
func (bA byAge) Swap(i, j int) { bA[i], bA[j] = bA[j], bA[i] }

func main() {
  user1 := user{
    First: "João",
    Last: "Victor1",
    Age: 15,
    Sayings: []string{
      "bla1", "bla2", "bla3",
    },
  }

  user2 := user{
    First: "João2",
    Last: "Victor0",
    Age: 19,
    Sayings: []string{
      "bla89", "bla0", "bla2",
    },
  }

  users := []user{user1, user2}

  for _, usr := range users {
    sort.Strings(usr.Sayings)
  }

  sort.Sort(byAge(users))
  fmt.Println("Usuários por idade: ")
  fmt.Println("---------------------------")
  prettier(users)

  sort.Sort(byLast(users))
  fmt.Println("Usuários por sobrenome: ")
  fmt.Println("---------------------------")
  prettier(users)
}

func prettier(users []user) {
  for index, usr := range users {
    fmt.Println("User", index, "\tNome Completo:", usr.First, usr.Last, "\tIdade:", usr.Age)
    fmt.Println("Sayings:")

    for _, saying := range usr.Sayings {
      fmt.Println("\t", saying)
    }

    fmt.Println("---------------------------")
  }
}

