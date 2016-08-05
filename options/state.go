package options

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sync"
)

var (
	state map[string]bool
	mu    sync.RWMutex
)

func Exists(
	key string,
) bool {
	mu.RLock()
	defer mu.RUnlock()
	if _, ok := state[key]; !ok {
		// does not exist
		return false
	}
	// key exists
	return true
}
func Get(
	key string,
) bool {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := state[key]
	if !ok {
		// not found
		return false
	}
	// key found
	return v
}
func Set(
	key string,
	v bool,
) {
	mu.Lock()
	defer mu.Unlock()
	state[key] = v
	go Save()
}
func Toggle(
	key string,
) bool {
	mu.Lock()
	defer mu.Unlock()
	v, ok := state[key]
	if ok {
		// found
		if v {
			// toggle off
			state[key] = false
		} else {
			// toggle on
			state[key] = true
		}
	} else {
		// not found, toggle on
		state[key] = true
	}
	go Save()
	return state[key]
}
