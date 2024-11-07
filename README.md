# go-colorterm

[![LICENSE](https://img.shields.io/github/license/cions/go-colorterm)](https://github.com/cions/go-colorterm/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/cions/go-colorterm.svg)](https://pkg.go.dev/github.com/cions/go-colorterm)
[![Go Report Card](https://goreportcard.com/badge/github.com/cions/go-colorterm)](https://goreportcard.com/report/github.com/cions/go-colorterm)

Enable ANSI escape sequence handling on Windows.

## Usage

```go
package main

import "github.com/cions/go-colorterm"

func main() {
	if colorterm.Enabled {
		fmt.Printf("\x1b[31mred text\x1b[0m")
	} else {
		fmt.Printf("plain text")
	}
}
```

## License

MIT
