package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/menublock"
	o "sabey.co/spoofgo/options"
	plug "sabey.co/spoofgo/plugins/plugin"
	v "sabey.co/spoofgo/version"
	"sabey.co/textblock"
)

type Frame struct {
	Location    int
	Options     *o.Options
	MenuOptions *menublock.MenuBlock
	Help        *textblock.TextBlock
	News        *textblock.TextBlock
	Version     *v.Client
	Plugin      *plug.Plugin
}
