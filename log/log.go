package log

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	file *os.File
)

const (
	FILE = `debug.log`
)

func init() {
	var err error
	file, err = os.OpenFile(FILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("failed to load \"%s\": \"%s\"\n", FILE, err)
		return
	}
}
func Log(
	s string,
) int {
	// we can't properly print to the screen with log because it'll mess up our layout, so a log file is good enough
	// you can follow logs on linux with `tail -f debug.log`
	// we should implement log.Logger in the future!
	// this current logging method should be temporary
	if file != nil {
		t := time.Now()
		n, _ := file.WriteString(fmt.Sprintf("%s %s\n", t.Format(time.RFC850), s))
		return n
	}
	return 0
}
func LogBoth(
	s string,
) {
	// in the current logging state LogBoth shouldn't be used once the CLI starts drawing
	log.Println(s)
	Log(s)
}
func CloseLog() {
	// not concurrent safe
	if file != nil {
		if err := file.Close(); err != nil {
			log.Fatalf("failed to close file: \"%s\"\n", err)
		}
	}
}
