package options

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

var (
	OPTIONS *Options
)

type Options struct {
	Options                       []*Option
	Option_Speed                  *Option
	Option_Speed_Unit             *Option
	Option_Mode                   *Option
	Option_Mode_Full              *Option
	Option_Modifier               *Option
	Option_Compass                *Option
	Option_Compass_Algorithm      *Option
	Option_Compass_North          *Option
	Option_Angle                  *Option
	Option_Angle_Moving           *Option
	Option_LatLong                *Option
	Option_LatLong_Moving         *Option
	Option_Center_Position        *Option
	Option_Center_Position_Moving *Option
	Option_Alert                  *Option
	Option_Version                *Option
	Option_Plugin                 *Option
	Option_Plugin_Offline         *Option
	Option_Server                 *Option
	Option_Quit                   *Option
}

func init() {
	OPTIONS = &Options{}
	Load()
}
