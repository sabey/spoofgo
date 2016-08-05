package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

func GetModeModifier() (
	int, // mode
	int, // modifier
) {
	mu.RLock()
	defer mu.RUnlock()
	return movement.Mode, movement.Modifier[movement.Mode]
}
func SetModifier(
	mode int,
	modifier int,
) (
	int, // mode
	int, // modifier
) {
	mu.Lock()
	defer mu.Unlock()
	if modifier < 1 {
		modifier = 1
	}
	if modifier > 100 {
		modifier = 100
	}
	if IsMode(mode) {
		// real mode
		movement.Mode = mode                        // set mode
		movement.Modifier[movement.Mode] = modifier // set modifier
	}
	// not a mode
	movement.calculate()
	return movement.Mode, movement.Modifier[movement.Mode]
}
func ResetModifier() (
	int, // mode
	int, // modifier
) {
	mu.Lock()
	defer mu.Unlock()
	movement.Modifier[movement.Mode] = 50
	movement.calculate()
	return movement.Mode, movement.Modifier[movement.Mode]
}
func IncreaseModifier() (
	int, // mode
	int, // modifier
) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := movement.Modifier[movement.Mode]; ok {
		// found
		movement.Modifier[movement.Mode] += 10
	} else {
		// not found
		movement.Modifier[movement.Mode] = 50
	}
	// fix modifier
	if movement.Modifier[movement.Mode] < 1 {
		movement.Modifier[movement.Mode] = 1
	}
	if movement.Modifier[movement.Mode] > 100 {
		movement.Modifier[movement.Mode] = 100
	}
	movement.calculate()
	return movement.Mode, movement.Modifier[movement.Mode]
}
func DecreaseModifier() (
	int, // mode
	int, // modifier
) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := movement.Modifier[movement.Mode]; ok {
		// found
		movement.Modifier[movement.Mode] -= 10
	} else {
		// not found
		movement.Modifier[movement.Mode] = 50
	}
	// fix modifier
	if movement.Modifier[movement.Mode] < 1 {
		movement.Modifier[movement.Mode] = 1
	}
	if movement.Modifier[movement.Mode] > 100 {
		movement.Modifier[movement.Mode] = 100
	}
	movement.calculate()
	return movement.Mode, movement.Modifier[movement.Mode]
}
