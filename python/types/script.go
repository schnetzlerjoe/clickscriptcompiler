package types

type Arg struct {
  Arg string `json:"arg"`
  Type string `json:"type"`
}

type Action struct {
  Line uint64 `json:"line"`
  Action string `json:"action"`
  Args []Arg `json:"args"`
}

type Code struct {
  RawCode string `json:"raw_code"`
}

type Script struct {
  File string `json:"file"`
  Code Code `json:"code"`
  Actions []Action `json:"actions"`
  Length uint64 `json:"length"`
}