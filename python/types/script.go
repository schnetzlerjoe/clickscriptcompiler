package types

type Command struct {
	LineNumber int64    `json:"line"`
	Action     Function `json:"action"`
}

type Script struct {
	FileName string    `json:"file_name"`
	Script   []Command `json:"script"`
	Length   uint64    `json:"length"`
}
