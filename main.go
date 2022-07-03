package main

import (
	"encoding/json"
	"fmt"
	"os"
  "github.com/schnetzlerjoe/clickscriptcompiler/python"
)

func main() {
  file, err := os.ReadFile("./code.json")
  if err != nil {
    fmt.Println(err)
  }
  script := python.Script{}
  json.Unmarshal(file, &script)
  fmt.Println(script)
}
