package main

import (
	"fmt"
  "os"
)

type Code struct {
  
}

type Script struct {
  Script Code `json:"script"`
  Language string `json:"language"`
  Length int64 `json:"length"`
}

func main() {
  file, err := os.ReadFile("./code.json")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(file)
}
