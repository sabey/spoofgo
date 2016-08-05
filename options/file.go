package options

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"sabey.co/spoofgo/file"
	"sabey.co/spoofgo/log"
)

const (
	FILE = `options.yaml`
)

func Load() {
	mu.Lock()
	defer mu.Unlock()
	state = make(map[string]bool) // init
	file, err := file.Open(FILE)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/options.Load(): failed to load \"%s\": \"%s\"\n", FILE, err))
		return
	}
	if err := yaml.Unmarshal(file, state); err != nil {
		log.Log(fmt.Sprintf("./spoofgo/options.Load(): failed to unmarshal \"%s\": \"%s\"\n", FILE, err))
		return
	}
	// loaded
}
func Save() {
	mu.Lock()
	defer mu.Unlock()
	bs, err := yaml.Marshal(state)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/options.Save(): failed to marshal \"%s\": \"%s\"\n", FILE, err))
		return
	}
	err = file.Save(FILE, bs)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/options.Save(): failed to save \"%s\": \"%s\"\n", FILE, err))
		return
	}
	// saved
}
