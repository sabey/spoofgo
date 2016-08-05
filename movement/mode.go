package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

const (
	WALK = iota
	JOG
	BICYCLE
	CAR
)

func GetMode() int {
	mu.RLock()
	defer mu.RUnlock()
	return movement.Mode
}
func SetMode(
	mode int,
) int {
	mu.Lock()
	defer mu.Unlock()
	if IsMode(mode) {
		movement.Mode = mode
	}
	movement.calculate()
	return movement.Mode
}
func ToggleMode() int {
	mu.Lock()
	defer mu.Unlock()
	if movement.Mode == WALK {
		movement.Mode = JOG
	} else if movement.Mode == JOG {
		movement.Mode = BICYCLE
	} else if movement.Mode == BICYCLE {
		movement.Mode = CAR
	} else {
		movement.Mode = WALK
	}
	movement.calculate()
	return movement.Mode
}
func IsMode(
	mode int,
) bool {
	if mode == WALK ||
		mode == JOG ||
		mode == BICYCLE ||
		mode == CAR {
		return true
	}
	return false
}
func PrintMode(
	mode int,
	full bool,
) string {
	if mode == WALK {
		if full {
			return "Walk"
		}
		return "W"
	}
	if mode == JOG {
		if full {
			return "Jog"
		}
		return "J"
	}
	if mode == BICYCLE {
		if full {
			return "Bicycle"
		}
		return "B"
	}
	if mode == CAR {
		if full {
			return "Car"
		}
		return "C"
	}
	return "?"
}
