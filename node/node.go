package node

type Parser struct {}

func (p Parser) print(content string) (string) {
  return "console.log(" + content + ")"
}