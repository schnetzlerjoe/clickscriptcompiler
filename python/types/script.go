package types

import (
	"log"
	"os"
  "github.com/schnetzlerjoe/clickscriptcompiler/python/actions"
)

type Script struct {
  FileName string `json:"file_name"`
  Commands []Command `json:"script"`
  Length uint64 `json:"length"`
}

func (s Script) Compile() {
  var code string
  var err error
  for _, command := range s.Commands {
    run := command.Action.Function
    run(args)
  }
  
  f, err := os.Create("script.py")
  if err != nil {
      log.Fatal(err)
  }
  defer f.Close()

  _, err2 := f.WriteString(code)
  if err2 != nil {
      log.Fatal(err2)
  }
}
