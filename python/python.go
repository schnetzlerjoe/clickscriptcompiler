package python

import (
	"encoding/json"
	"log"
	"os"

	"github.com/schnetzlerjoe/clickscriptcompiler/python/types"
)

type Action struct {
	FunctionName string `json:"function_name"`
	Function     func() `json:"function"`
}

type Actions []Action

func Compile() error {
	// create initial base variables needed later
	var code string
	var err error

	// read in the json file that will be converted to code
	file, err := os.ReadFile("./code.json")
	if err != nil {
		return err
	}
	// unmarshal the raw json bytes into golang struct
	script := types.Script{}
	json.Unmarshal(file, &script)

	// loop through each command in the json file (golang struct)
	// and use handler to lookup function and convert to python string code.
	// Each compiled string is concated into the current code string
	for _, command := range script.Script {
		run := GetFunction(command.Action.FunctionName)
		code, err = run(command.Action.Args)
		if err != nil {
			return err
		}
	}

	// create the base python script file
	f, err := os.Create("script.py")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write the generated python script (in the code variable) to the file
	_, err2 := f.WriteString(code)
	if err2 != nil {
		log.Fatal(err2)
	}

	return nil
}
