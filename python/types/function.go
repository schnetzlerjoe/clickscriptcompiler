package types

type Function struct {
	FunctionName string `json:"function_name"`
	Args         []byte `json:"args"`
}

func (f Function) Compile() {
}
