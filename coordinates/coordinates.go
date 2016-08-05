package coordinates

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/movement"
	"strings"
	"sync"
)

var (
	coordinates map[string]*Coordinate
	mu          sync.RWMutex
)

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Add(
	key string,
	latitude float64,
	longitude float64,
) *Coordinate {
	if key == "" {
		// empty key
		return nil
	}
	key = strings.ToLower(key)
	mu.Lock()
	defer mu.Unlock()
	coordinates[key] = &Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}
	save()
	// return dereferenced
	return &Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
func Delete(
	key string,
) *Coordinate {
	if key == "" {
		// empty key
		return nil
	}
	key = strings.ToLower(key)
	mu.Lock()
	defer mu.Unlock()
	if c, ok := coordinates[key]; ok {
		// coords exist
		delete(coordinates, key)
		save()
		// return dereferenced
		return &Coordinate{
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		}
	}
	// coords not found
	return nil
}
func Set(
	key string,
	accelerating bool,
) *Coordinate {
	if key == "" {
		// empty key
		return nil
	}
	key = strings.ToLower(key)
	mu.RLock()
	defer mu.RUnlock()
	if c, ok := coordinates[key]; ok {
		// coords exist
		movement.SetLatLong(c.Latitude, c.Longitude, accelerating)
		// return dereferenced
		return &Coordinate{
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		}
	}
	// coords not found
	return nil
}
func Get(
	key string,
) *Coordinate {
	if key == "" {
		// empty key
		return nil
	}
	key = strings.ToLower(key)
	mu.RLock()
	defer mu.RUnlock()
	if c, ok := coordinates[key]; ok {
		// coords exist
		return c
	}
	// coords not found
	return nil
}
func List() map[string]*Coordinate {
	mu.RLock()
	defer mu.RUnlock()
	// dereference
	coords := make(map[string]*Coordinate)
	for k, v := range coordinates {
		coords[k] = v
	}
	return coords
}
