package python

import (
  "github.com/schnetzlerjoe/clickscriptcompiler/python/types"
  "os"
  "fmt"
  "encoding/json"
)

type Action struct {
  FunctionName string `json:"function_name"`
  Function func() `json:"function"`
}

type Actions []Action

var actions Actions

func Compile() error {
  file, err := os.ReadFile("./code.json")
  if err != nil {
    return err
  }
  script := types.Script{}
  json.Unmarshal(file, &script)
  fmt.Println(script)
  
  script.Compile()
  return nil
}