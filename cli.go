package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/cli"
	o "sabey.co/spoofgo/options"
)

func DrawCLI() {
	frame := &cli.Frame{
		Options:     o.OPTIONS,
		MenuOptions: menu_option,
		Help:        textblock_help,
		Version:     version,
		Plugin:      plugin,
	}
	// location is unsafe!
	location_mu.RLock()
	frame.Location = location
	location_mu.RUnlock()
	// news is unsafe!
	news_mu.RLock()
	frame.News = textblock_news
	news_mu.RUnlock()
	cli.Draw(frame)
}
