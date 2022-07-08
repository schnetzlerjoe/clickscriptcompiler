package actions

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobuffalo/plush"
)

type PrintArgs struct {
	Values []any `json:"values"`
}

func Print(args []byte) (code string, err error) {
	var rawVals PrintArgs
	var currentVal string
	json.Unmarshal(args, &rawVals)
	for _, val := range rawVals.Values {
		currentVal = fmt.Sprintf("%q, %q", currentVal, val)
	}

	template := `print("<%= value %>")`
	ctx := plush.NewContext()
	ctx.Set("value", currentVal)

	code, err = plush.Render(template, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return code, nil
}
