<div align="center">
	<h1>Backero</h1>
	<p>
		<img alt="Codecov" src="https://img.shields.io/codecov/c/gh/dilmorja/backero?color=ff0176&logo=codecov&style=for-the-badge">
	</p>
	<p>
		Backero is the infrastructure that allows you to create and implement multiple targets for the backend of a programming language.
	</p>
</div>

## Example usage

> NOTE: This sample code is currently a long way from being implementable, but it is a goal.

```go
package main

import(
	"fmt"

	"github.com/dilmorja/backero"
	"github.com/dilmorja/backero/targets/ts"
	"github.com/dilmorja/backero/targets/ts/types"
)

func main() {
	vaquero := backero.New()
	vaquero.LoadAndUse(ts.TypeScript())

	mod := vaquero.InitModule("constant")
	mod.Const("x", types.String, "This is a constant", false)

	fmt.Println(mod.String()) // const x: string = "This is a constant"
	fmt.Println(mod.Memory["const"]["x"].Content()) // "This is a constant"
}
```

## Why Go (Golang)?

It's simple: go is easy to learn, relatively modern, scalable, and its standard modular structure is intuitive.

## License

[BSD 3-Clause](LICENSE)