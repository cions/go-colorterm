// Copyright (c) 2024 cions
// Licensed under the MIT License. See LICENSE for details.

// Package colorterm enables ANSI escape sequence support on Windows.
package colorterm

import "fmt"

type EscapeCode string

func (ec EscapeCode) String() string {
	if !Enabled {
		return ""
	}
	return string(ec)
}

var (
	Escape = '\x1b'
	CSI    = "\x1b["
)

var (
	Reset = EscapeCode("\x1b[0m")

	FgBlack   = EscapeCode("\x1b[30m")
	FgRed     = EscapeCode("\x1b[31m")
	FgGreen   = EscapeCode("\x1b[32m")
	FgYellow  = EscapeCode("\x1b[33m")
	FgBlue    = EscapeCode("\x1b[34m")
	FgMagenta = EscapeCode("\x1b[35m")
	FgCyan    = EscapeCode("\x1b[36m")
	FgWhite   = EscapeCode("\x1b[37m")
	FgReset   = EscapeCode("\x1b[39m")

	BgBlack   = EscapeCode("\x1b[40m")
	BgRed     = EscapeCode("\x1b[41m")
	BgGreen   = EscapeCode("\x1b[42m")
	BgYellow  = EscapeCode("\x1b[43m")
	BgBlue    = EscapeCode("\x1b[44m")
	BgMagenta = EscapeCode("\x1b[45m")
	BgCyan    = EscapeCode("\x1b[46m")
	BgWhite   = EscapeCode("\x1b[47m")
	BgReset   = EscapeCode("\x1b[49m")
)

func Fg256Color(c uint8) EscapeCode {
	return EscapeCode(fmt.Sprintf("\x1b[38;5;%dm", c))
}

func FgRGB(r, g, b uint8) EscapeCode {
	return EscapeCode(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b))
}

func Bg256Color(c uint8) EscapeCode {
	return EscapeCode(fmt.Sprintf("\x1b[48;5;%dm", c))
}

func BgRGB(r, g, b uint8) EscapeCode {
	return EscapeCode(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b))
}
