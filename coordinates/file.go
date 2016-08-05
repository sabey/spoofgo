package coordinates

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"sabey.co/spoofgo/file"
	"sabey.co/spoofgo/log"
	"strings"
)

const (
	FILE = `coordinates.yaml`
)

func init() {
	Load()
}
func Load() {
	mu.Lock()
	defer mu.Unlock()
	coordinates = make(map[string]*Coordinate)
	file, err := file.Open(FILE)
	unsafe := false
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/coordinates.Load(): failed to load \"%s\": \"%s\"\n", FILE, err))
		// coordinates not found, assuming default install
		unsafe = true
	}
	if !unsafe {
		if err := yaml.Unmarshal(file, coordinates); err != nil {
			log.Log(fmt.Sprintf("./spoofgo/coordinates.Load(): failed to unmarshal \"%s\": \"%s\"\n", FILE, err))
			// failed to unmarshal
			unsafe = true
		}
	}
	// fix coordinates
	for key, value := range coordinates {
		if key == "" {
			// empty
			unsafe = true
			delete(coordinates, key)
		} else {
			if value == nil {
				// nil
				unsafe = true
				delete(coordinates, key)
			} else {
				key2 := strings.ToLower(key)
				if key != key2 {
					// make lowercase
					unsafe = true
					coordinates[key2] = coordinates[key]
					delete(coordinates, key)
				}
			}
		}
	}
	if unsafe {
		save()
	}
	// loaded
}
func Save() {
	mu.Lock()
	defer mu.Unlock()
	save()
}
func save() {
	bs, err := yaml.Marshal(coordinates)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/coordinates.save(): failed to marshal \"%s\": \"%s\"\n", FILE, err))
		return
	}
	err = file.Save(FILE, bs)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/coordinates.save(): failed to save \"%s\": \"%s\"\n", FILE, err))
		return
	}
	// saved
}
