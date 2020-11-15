# tl

[![PkgGoDev](https://pkg.go.dev/badge/github.com/ernado/tl)](https://pkg.go.dev/github.com/ernado/tl)

Package tl implements [TL](https://core.telegram.org/mtproto/TL) (Type Language) schema parser and writer.
Inspired by [grammers](https://github.com/Lonami/grammers) parser.

```console
go get github.com/ernado/tl
```

## Parsing

This program parses schema from stdin and prints all definitions with their
names and types.

```go
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
```

You can use it like that:
```console
$ curl -s "https://raw.githubusercontent.com/tdlib/td/master/td/generate/scheme/td_api.tl" \
    | go run github.com/ernado/tl/cmd/tl-print \
    | less
```

Output:
```tl
double#2210c154 = Double;
string#b5286e24 = String;
int32#5cb934fa = Int32;
int53#6781c7ee = Int53;
int64#5d9ed744 = Int64;
bytes#e937bb82 = Bytes;
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;
error#9bdd8f1a = Error;
ok#d4edbe69 = Ok;
tdlibParameters#d29c1d7b = TdlibParameters;

//...
```

### Generating

You can also generate `.tl` file from `tl.Schema`.
Aby `WriteTo` result is valid input for `Parse`.

```go
package main

import (
	"os"

	"github.com/ernado/tl"
)

func main() {
	def := tl.Definition{
		Name: "error",
		Type: tl.Type{Name: "Error"},
		// Currently you should always pass explicit ID.
		ID: 0x9bdd8f1a,
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
	}
	_, _ = tl.Schema{
		Definitions: []tl.SchemaDefinition{
			{
				Category:   tl.CategoryType,
				Definition: def,
			},
		},
	}.WriteTo(os.Stdout)
}
```

Output
```tl
error#9bdd8f1a code:int32 message:string = Error;
```
