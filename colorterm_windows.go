// Copyright (c) 2024-2025 cions
// Licensed under the MIT License. See LICENSE for details.

//go:build windows

package colorterm

import (
	"os"

	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

var Enabled = false

func init() {
	if os.Getenv("NO_COLOR") != "" {
		return
	}

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return
	}

	handle := windows.Handle(os.Stdout.Fd())
	var mode uint32

	if err := windows.GetConsoleMode(handle, &mode); err != nil {
		return
	}

	if err := windows.SetConsoleMode(handle, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err != nil {
		return
	}

	Enabled = true
}
