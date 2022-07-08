package python

import (
	"github.com/schnetzlerjoe/clickscriptcompiler/python/actions"
)

// GetFunction is a lookup function that maps a function name to a function
// and then returns that function. The functions then can be used to compile
// down to string based code (in this case Python) in our Python runner.
// For each newly created Clickscript action, you must add that action as a case
// in the switch statement for Clickscript to be able to use it.
func GetFunction(functionName string) func(args []byte) (code string, err error) {
	switch functionName {
	case "print":
		return actions.Print
	}
	return nil
}
