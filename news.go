package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	n "sabey.co/spoofgo/news"
	"sabey.co/textblock"
	"sync"
)

var (
	textblock_news *textblock.TextBlock
	news           *n.Client
	news_mu        sync.RWMutex
)

func ResetNews() {
	news_mu.Lock()
	textblock_news = news.GetNewsBlock()
	news_mu.Unlock()
}
