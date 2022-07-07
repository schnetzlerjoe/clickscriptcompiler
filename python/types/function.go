package types

type Function struct {
	FunctionName string          `json:"function_name"`
	Args         Args            `json:"args"`
	Function     func(args Args) `json:"function"`
}

type Command struct {
	Line   uint64   `json:"line"`
	Action Function `json:"action"`
}
