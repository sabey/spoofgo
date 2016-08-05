package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sync"
)

var (
	movement *Movement
	mu       sync.RWMutex
)

type Movement struct {
	Angle     int
	Mode      int
	Modifier  map[int]int
	Latitude  float64
	Longitude float64
	state     *State
}

func Accelerate() {
	mu.Lock()
	defer mu.Unlock()
	movement.state.Accelerating = true
	movement.calculate()
}
func Decelerate() {
	mu.Lock()
	defer mu.Unlock()
	movement.state.Accelerating = false
	movement.calculate()
}
