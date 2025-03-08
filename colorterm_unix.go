// Copyright (c) 2024-2025 cions
// Licensed under the MIT License. See LICENSE for details.

//go:build !windows

package colorterm

import (
	"os"

	"golang.org/x/term"
)

var Enabled = false

func init() {
	if os.Getenv("NO_COLOR") != "" {
		return
	}

	if os.Getenv("TERM") == "dumb" {
		return
	}

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return
	}

	Enabled = true
}
