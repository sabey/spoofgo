package options

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

type Option struct {
	label string
	key   string
	on    rune
	on_f  func()
	off   rune
	off_f func()
	noop  bool
}

func Create(
	label string,
	key string,
	on rune,
	on_f func(),
	off rune,
	off_f func(),
	noop bool,
) *Option {
	return &Option{
		label: label,
		key:   key,
		on:    on,
		on_f:  on_f,
		off:   off,
		off_f: off_f,
		noop:  noop,
	}
}
func (self *Option) GetLabel() string {
	return self.label
}
func (self *Option) GetKey() string {
	return self.key
}
func (self *Option) Exists() bool {
	return Exists(self.key)
}
func (self *Option) Get() (
	rune,
	bool, // on
) {
	if Get(self.key) {
		// on
		return self.on, true
	}
	// off
	return self.off, false
}
func (self *Option) Set(
	v bool,
) rune {
	if !self.noop {
		// only set key if we're not noop
		Set(self.key, v)
	}
	if v {
		// on
		if self.on_f != nil {
			go self.on_f()
		}
		return self.on
	}
	// off
	if self.off_f != nil {
		go self.off_f()
	}
	return self.off
}
func (self *Option) Toggle() (
	rune,
	bool, // on
) {
	if self.noop {
		// noop
		c, v := self.Get()
		// we still want our toggle functions to be called
		if v {
			// on
			if self.on_f != nil {
				go self.on_f()
			}
		} else {
			// off
			if self.off_f != nil {
				go self.off_f()
			}
		}
		return c, v
	}
	if Toggle(self.key) {
		// on
		if self.on_f != nil {
			go self.on_f()
		}
		return self.on, true
	}
	// off
	if self.off_f != nil {
		go self.off_f()
	}
	return self.off, false
}
