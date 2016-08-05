package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"sabey.co/menublock"
	o "sabey.co/spoofgo/options"
)

var (
	menu_option *menublock.MenuBlock
)

func OptionToggle() {
	o.OPTIONS.Options[menu_option.GetItem()].Toggle()
}
func BuildOptions() {
	o.OPTIONS.Option_Speed = o.Create(
		"Show Speed",
		"speed",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Speed.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Speed.Set(true)
	}
	o.OPTIONS.Option_Speed_Unit = o.Create(
		"Speed Format: mi/h / km/h",
		"speed_unit",
		'M', // on
		nil,
		'K', // off
		nil,
		false, // noop
	)
	o.OPTIONS.Option_Mode = o.Create(
		"Show Movement Mode",
		"mode",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Mode.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Mode.Set(true)
	}
	o.OPTIONS.Option_Mode_Full = o.Create(
		"Show Movement Mode Full",
		"mode_full",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Mode_Full.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Mode_Full.Set(true)
	}
	o.OPTIONS.Option_Modifier = o.Create(
		"Show Movement Modifier",
		"modifier",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Modifier.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Modifier.Set(true)
	}
	o.OPTIONS.Option_Compass = o.Create(
		"Show Compass",
		"compass",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Compass.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Compass.Set(true)
	}
	o.OPTIONS.Option_Compass_Algorithm = o.Create(
		"Compass Algorithm: Fixed / Dynamic",
		"compass_algorithm",
		'D', // on
		nil,
		'F', // off
		nil,
		false, // noop
	)
	o.OPTIONS.Option_Compass_North = o.Create(
		"Show Compass North",
		"compass_north",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Compass_North.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Compass_North.Set(true)
	}
	o.OPTIONS.Option_Angle = o.Create(
		"Show Angle",
		"angle",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Angle.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Angle.Set(true)
	}
	o.OPTIONS.Option_Angle_Moving = o.Create(
		"Show Angle Moving",
		"angle_moving",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	o.OPTIONS.Option_LatLong = o.Create(
		"Show Latitude/Longitude",
		"latlong",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_LatLong.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_LatLong.Set(true)
	}
	o.OPTIONS.Option_LatLong_Moving = o.Create(
		"Show Latitude/Longitude Moving",
		"latlong_moving",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	o.OPTIONS.Option_Center_Position = o.Create(
		"Show Center Position",
		"center_position",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Center_Position.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Center_Position.Set(true)
	}
	o.OPTIONS.Option_Center_Position_Moving = o.Create(
		"Show Center Position Moving",
		"center_position_moving",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	if !o.OPTIONS.Option_Center_Position_Moving.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Center_Position_Moving.Set(true)
	}
	o.OPTIONS.Option_Alert = o.Create(
		"Menu Alert",
		"alert",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	o.OPTIONS.Option_Version = o.Create(
		"Check Version",
		"version",
		'*', // on
		func() {
			version.Start()
		},
		' ', // off
		func() {
			version.Stop()
		},
		false, // noop
	)
	if !o.OPTIONS.Option_Version.Exists() {
		// option not found, assuming default install
		o.OPTIONS.Option_Version.Set(true)
	}
	o.OPTIONS.Option_Quit = o.Create(
		"Quit",
		"quit",
		' ', // on
		quit_func,
		' ', // off
		quit_func,
		true, // noop
	)
	// plugin show offline
	o.OPTIONS.Option_Plugin_Offline = o.Create(
		"Show Plugin Offline Status",
		"plugin_offline",
		'*', // on
		nil,
		' ', // off
		nil,
		false, // noop
	)
	// plugin noop indicator
	plugin_running := "Plugin Not Running!"
	if *flag_plugin != "" {
		plugin_running = fmt.Sprintf("Plugin Running: \"%s\" Addr: \"%s\"", *flag_plugin, *flag_plugin_addr)
	}
	plugin_running += " (toggled on restart only)"
	o.OPTIONS.Option_Plugin = o.Create(
		plugin_running,
		"plugin",
		'*', // on
		nil,
		' ', // off
		nil,
		true, // noop
	)
	if *flag_plugin != "" {
		o.Set("plugin", true)
	}
	// server noop indicator
	api_running := "API/Controller Not Running!"
	if *flag_api {
		scheme := "http"
		if *flag_api_tls {
			scheme += "s"
		}
		// we can't access server.ServerAddr() yet
		authority := *flag_api_addr
		if len(authority) > 0 && authority[0] == ':' {
			// prepend localhost?
			// it would be nice to get the actual proper server address, if we were listening on ":0" this couldn't be correct
			// we would have to implement our own listener to figure out what we're bound to
			// we would also not even be listening yet!
			authority = "localhost" + authority
		}
		api_running = fmt.Sprintf("API/Controller Running: \"%s://%s/\" Controller: \"%s\"", scheme, authority, *flag_controller)
	}
	api_running += " (toggled on restart only)"
	o.OPTIONS.Option_Server = o.Create(
		api_running,
		"api",
		'*', // on
		nil,
		' ', // off
		nil,
		true, // noop
	)
	// set the state of our -api flag
	o.Set("api", *flag_api)
	// build slice
	o.OPTIONS.Options = []*o.Option{
		o.OPTIONS.Option_Mode,
		o.OPTIONS.Option_Mode_Full,
		o.OPTIONS.Option_Speed,
		o.OPTIONS.Option_Speed_Unit,
		o.OPTIONS.Option_Modifier,
		o.OPTIONS.Option_Compass,
		o.OPTIONS.Option_Compass_Algorithm,
		o.OPTIONS.Option_Compass_North,
		o.OPTIONS.Option_Angle,
		o.OPTIONS.Option_Angle_Moving,
		o.OPTIONS.Option_LatLong,
		o.OPTIONS.Option_LatLong_Moving,
		o.OPTIONS.Option_Center_Position,
		o.OPTIONS.Option_Center_Position_Moving,
		o.OPTIONS.Option_Alert,
		o.OPTIONS.Option_Version,
		o.OPTIONS.Option_Plugin_Offline,
		o.OPTIONS.Option_Plugin,
		o.OPTIONS.Option_Server,
		o.OPTIONS.Option_Quit,
	}
	menu_option = menublock.Create(len(o.OPTIONS.Options))
}
