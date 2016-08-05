package movement

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

const (
	NORTH_WEST = 315
	NORTH_EAST = 45
	SOUTH_EAST = 125
	SOUTH_WEST = 225
)

func GetAngle() int {
	mu.RLock()
	defer mu.RUnlock()
	return movement.Angle
}
func SetAngle(
	angle int,
) int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle = angle
	movement.calculate()
	return movement.Angle
}
func Right() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle += 10
	movement.calculate()
	return movement.Angle
}
func Left() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle -= 10
	movement.calculate()
	return movement.Angle
}
func NorthWest() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle = NORTH_WEST
	movement.calculate()
	return movement.Angle
}
func NorthEast() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle = NORTH_EAST
	movement.calculate()
	return movement.Angle
}
func SouthWest() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle = SOUTH_WEST
	movement.calculate()
	return movement.Angle
}
func SouthEast() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle = SOUTH_EAST
	movement.calculate()
	return movement.Angle
}
func Flip() int {
	mu.Lock()
	defer mu.Unlock()
	movement.Angle += 180
	movement.calculate()
	return movement.Angle
}
func (self *Movement) fixAngle() {
	if self.Angle < 0 {
		self.Angle += 360 // back around
	}
	self.Angle %= 360 // fix with modulus
}
