package uno

import (
	"fmt"
	"github.com/louis-trevino/go-uno/util"
)

func AcceptLnDemo() {
	fmt.Printf("\n* Demo of uno.AcceptLn (accept multiple space-separated values in a line.) \n")
	fmt.Print("Enter a string: ")
	if inStr, err := util.AcceptLn(); err == nil {
		fmt.Printf("In Str: %s", inStr)
	} else {
		fmt.Printf("Error: %V \n", err)
	}
}
