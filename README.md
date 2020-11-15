# tl

Package tl implements TL (Type Language) schema parser and writer.

## Example

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
		fmt.Printf("%s -> %s\n", d.Definition.Name, d.Definition.Type)
	}
}
```

You can use like that:
```console
$ curl -s "https://raw.githubusercontent.com/tdlib/td/master/td/generate/scheme/td_api.tl" \
    | go run github.com/ernado/tl/cmd/tl-print \
    | less
```
