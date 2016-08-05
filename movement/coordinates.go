package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

func SetLatLong(
	lat float64,
	long float64,
	accelerating bool,
) {
	mu.Lock()
	defer mu.Unlock()
	movement.Latitude = lat
	movement.Longitude = long
	if accelerating {
		// calculate movement
		movement.calculate()
		// are we accelerating?
		if !movement.state.Accelerating {
			// start accelerating!
			movement.state.Accelerating = true
		}
	} else {
		// don't calculate movement
		// we're at a stand still
		movement.state.Accelerating = false
		movement.state.Speed = 0.0
	}
	save()
}
func GetLatLong() (
	float64, // lat
	float64, // long
) {
	mu.Lock()
	defer mu.Unlock()
	movement.calculate()
	return movement.Latitude, movement.Longitude
}
