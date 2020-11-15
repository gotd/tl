package main

import (
	"os"

	"github.com/ernado/tl"
)

func main() {
	// error#9bdd8f1a code:int32 message:string = Error;
	_, _ = tl.Schema{
		Definitions: []tl.SchemaDefinition{
			{
				Category: tl.CategoryType,
				Definition: tl.Definition{
					Name: "error",
					Type: tl.Type{Name: "Error"},
					ID:   0x9bdd8f1a,
					Params: []tl.Parameter{
						{
							Name: "code",
							Type: tl.Type{Name: "int32", Bare: true},
						},
						{
							Name: "message",
							Type: tl.Type{Name: "string", Bare: true},
						},
					},
				},
			},
		},
	}.WriteTo(os.Stdout)
}
