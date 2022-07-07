package types

import "github.com/pkg/errors"

type Arg struct {
  Arg string `json:"arg"`
  Value string `json:"value"`
}

type Args []Arg

func (a Args) GetArg(argName string) (arg string, err error) {
  for _, arg := range a {
    if arg.Arg == argName {
      return arg, nil
    }
  }
  return arg, errors.Wrap("")
}
