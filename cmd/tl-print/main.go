package main

import (
	"fmt"
	"os"

	"github.com/ernado/tl"
)

func main() {
	schema, err := tl.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, d := range schema.Definitions {
		fmt.Printf("%s#%x = %s;\n", d.Definition.Name, d.Definition.ID, d.Definition.Type)
	}
}
