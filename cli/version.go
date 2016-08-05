package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"github.com/nsf/termbox-go"
	v "sabey.co/spoofgo/version"
)

func GetVersion(
	frame *Frame,
) string {
	// don't print build
	ver := frame.Version.GetVersion()
	return fmt.Sprintf("%d.%d", ver.Major, ver.Minor)
}
func GetVersionColour(
	frame *Frame,
) termbox.Attribute {
	state := frame.Version.GetVersionState()
	if state == v.SAFE {
		// safe
		return termbox.ColorGreen
	} else if state == v.OLD_MINOR {
		// new minor
		return termbox.ColorCyan // BLUE
	} else if state == v.OLD_MAJOR {
		// new major
		return termbox.ColorYellow
	} else if state == v.OLD_BOTH {
		// new major and minor!!!
		return termbox.ColorMagenta // PINK
	}
	// dangerous
	return termbox.ColorRed
}
