package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"time"
)

const (
	SPEED_KM_WALK    = 5.0   // the average human walking speed is about 5.0 kilometres per hour (km/h), or about 3.1 miles per hour
	SPEED_KM_JOG     = 10.0  // jogging as running slower than 6 miles per hour (10 km/h)
	SPEED_KM_BICYCLE = 15.5  // For cyclists in Copenhagen, the average cycling speed is 15.5 km/h (9.6 mph).
	SPEED_KM_CAR     = 120.0 // the max legal highway speed in canada is like 120km/h, people shouldn't travel that in the cities but we will allow people to accelerate up to it
)
const (
	SPEED_KMH_TO_MPH = 0.621371192237
	SPEED_MPH_TO_KMH = 1.60934400061
)
const (
	// we're going to assume we're always going the max speed while walking or jogging
	TIME_TO_MAX_WALK    = time.Second * 3
	TIME_TO_MAX_JOG     = time.Second * 5
	TIME_TO_MAX_BICYCLE = time.Second * 6
	TIME_TO_MAX_CAR     = time.Second * 20
)

func GetSpeed() float64 {
	mu.Lock()
	defer mu.Unlock()
	movement.calculate()
	// copy state
	return movement.state.Speed
}
func GetModeModifierMaxSpeed(
	mode int,
	modifier int,
) float64 {
	if modifier < 1 {
		modifier = 1
	}
	if modifier > 100 {
		modifier = 100
	}
	p := float64(modifier) / 100
	if mode == WALK {
		return p * SPEED_KM_WALK
	}
	if mode == JOG {
		return p * SPEED_KM_JOG
	}
	if mode == BICYCLE {
		return p * SPEED_KM_BICYCLE
	}
	if mode == CAR {
		return p * SPEED_KM_CAR
	}
	return 0
}
