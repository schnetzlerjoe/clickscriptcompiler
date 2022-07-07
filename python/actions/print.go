package actions

import (
	"log"

	"github.com/gobuffalo/plush"
)

func Print(value string) (code string, err error) {
	template := `print("<%= value %>")`
	ctx := plush.NewContext()
	ctx.Set("value", value)

	code, err = plush.Render(template, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return code, nil
}
