# go-colorterm

[![LICENSE](https://img.shields.io/github/license/cions/go-colorterm)](https://github.com/cions/go-colorterm/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/cions/go-colorterm.svg)](https://pkg.go.dev/github.com/cions/go-colorterm)
[![Go Report Card](https://goreportcard.com/badge/github.com/cions/go-colorterm)](https://goreportcard.com/report/github.com/cions/go-colorterm)

Enable ANSI escape sequence support on Windows.

## Usage

```go
package main

import "github.com/cions/go-colorterm"

func main() {
	fmt.Printf("%vred text%v\n", colorterm.FgRed, colorterm.Reset)
}
```

## License

MIT
