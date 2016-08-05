package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"math/rand"
	"time"
)

type State struct {
	Accelerating bool
	Speed        float64
	Time         time.Time
}

func GetState() *State {
	mu.Lock()
	defer mu.Unlock()
	movement.calculate()
	// copy state
	return &State{
		Accelerating: movement.state.Accelerating,
		Speed:        movement.state.Speed,
		Time:         movement.state.Time,
	}
}
func IsAccelerating() bool {
	mu.Lock()
	defer mu.Unlock()
	movement.calculate()
	// copy state
	return movement.state.Accelerating
}
func GetLastState() *State {
	mu.RLock()
	defer mu.RUnlock()
	// copy the current state
	return &State{
		Accelerating: movement.state.Accelerating,
		Speed:        movement.state.Speed,
		Time:         movement.state.Time,
	}
}
func GetLastTime() int64 {
	mu.RLock()
	defer mu.RUnlock()
	// get the current state
	return movement.state.Time.UnixNano()
}
func Calculate() {
	mu.Lock()
	defer mu.Unlock()
	movement.calculate()
}
func (self *Movement) calculate() {
	// should we do our movement calculations?
	if !self.state.Accelerating {
		// no need to do calculations
		// fix our angle
		self.fixAngle()
		self.state.Speed = 0
		return
	}
	if rand.Intn(10)%10 == 0 {
		// add a little wobble to our angle 1/10 times that we calculate
		// angle isn't a float so this is a good alternative
		self.Angle += deviate(1)
	}
	// fix our angle
	self.fixAngle()
	// get current time
	now := time.Now()
	// get duration between last/current state
	duration := now.UnixNano() - movement.state.Time.UnixNano()
	// update time
	movement.state.Time = now
	// Calculate Speed
	movement.state.Speed = GetModeModifierMaxSpeed(movement.Mode, movement.Modifier[movement.Mode])
	// Calculate Distance
	// Distance = Speed × Time
	distance := (movement.state.Speed * float64(duration)) / float64(time.Hour)
	// Update Lat/Long
	movement.Latitude, movement.Longitude = PointAtDistanceAndBearing(
		movement.Latitude,
		movement.Longitude,
		distance,
		float64(movement.Angle),
	)
}
func deviate(
	d int,
) int {
	d++ // bit start at 0
	if rand.Intn(2)%2 == 0 {
		// deviate right
		return rand.Intn(d)
	}
	// deviate left
	return -rand.Intn(d)
}
