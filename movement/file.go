package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"sabey.co/spoofgo/file"
	"sabey.co/spoofgo/log"
	"time"
)

const (
	FILE = `movement.yaml`
)

func init() {
	Load()
	go func() {
		for _ = range time.NewTicker(time.Second * 10).C {
			// save state every 10 seconds
			Save()
		}
		log.Log("./spoofgo/movement.init(): stopped saving movement state???")
	}()
}
func Load() {
	mu.Lock()
	defer mu.Unlock()
	movement = &Movement{}
	movement.Modifier = make(map[int]int)
	movement.state = &State{
		Accelerating: false,
		Speed:        0.0,
		Time:         time.Now(),
	}
	file, err := file.Open(FILE)
	unsafe := false
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/movement.Load(): failed to load \"%s\": \"%s\"\n", FILE, err))
		// movement not found, assuming default install
		unsafe = true
	}
	if !unsafe {
		if err := yaml.Unmarshal(file, movement); err != nil {
			log.Log(fmt.Sprintf("./spoofgo/movement.Load(): failed to unmarshal \"%s\": \"%s\"\n", FILE, err))
			// failed to unmarshal
			unsafe = true
		}
	}
	// fix angle
	a := movement.Angle
	movement.fixAngle()
	if a != movement.Angle {
		// unsafe angle
		movement.Angle = 0
		unsafe = true
	}
	// fix mode
	if !IsMode(movement.Mode) {
		// unsafe mode
		movement.Mode = WALK
		unsafe = true
	}
	// fix modifier
	for mode, modifier := range movement.Modifier {
		if IsMode(mode) {
			// safe mode
			if modifier < 1 {
				// too small
				movement.Modifier[mode] = 1
				unsafe = true
			}
			if modifier > 100 {
				// too big
				movement.Modifier[mode] = 100
				unsafe = true
			}
		} else {
			// unknown mode
			delete(movement.Modifier, mode)
			unsafe = true
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
	bs, err := yaml.Marshal(movement)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/movement.save(): failed to marshal \"%s\": \"%s\"\n", FILE, err))
		return
	}
	err = file.Save(FILE, bs)
	if err != nil {
		log.Log(fmt.Sprintf("./spoofgo/movement.save(): failed to save \"%s\": \"%s\"\n", FILE, err))
		return
	}
	// saved
}
