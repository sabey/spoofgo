package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sabey.co/spoofgo/api"
	controls "sabey.co/spoofgo/api/controllers"
	control "sabey.co/spoofgo/api/controllers/controller"
	"sabey.co/spoofgo/coordinates"
	"sabey.co/spoofgo/file"
	clog "sabey.co/spoofgo/log"
	"sabey.co/spoofgo/movement"
	n "sabey.co/spoofgo/news"
	o "sabey.co/spoofgo/options"
	plugs "sabey.co/spoofgo/plugins"
	plug "sabey.co/spoofgo/plugins/plugin"
	v "sabey.co/spoofgo/version"
	"sync"
	"time"
)

const (
	// our current version info
	// we could control these with build parameters but we will probably just hardcode them
	VERSION_MAJOR = 0
	VERSION_MINOR = 1
	// possibly only build will be set as a build parameter?
	VERSION_BUILD = 0
)

var (
	quit       chan struct{}
	quit_once  sync.Once
	server     *api.Server
	version    *v.Client
	plugin     *plug.Plugin
	controller *control.Controller
	// general
	flag_read  = flag.Bool("read", false, fmt.Sprintf("Run even if \"%s\" exists", LOCK_FILE))
	flag_nogui = flag.Bool("nogui", false, "Don't start the GUI")
	// plugin
	flag_plugin_list  = flag.Bool("plugin-list", false, "List all Plugins and exit")
	flag_plugin       = flag.String("plugin", "", "Which Plugin to use?")
	flag_plugin_addr  = flag.String("plugin-addr", "", "Plugin Call Address")
	flag_plugin_every = flag.String("plugin-every", "500ms", `Plugin Call Every ("ns", "us" (or "µs"), "ms", "s", "m", "h")?`)
	flag_plugin_kv    = flag.String("plugin-kv", `{"key":"value"}`, "Plugin KeyValue Json Object map[string]interface{} (*Optional, Plugin Dependent!*)")
	// api
	flag_api      = flag.Bool("api", false, "Start Local Browser Controller/API Server")
	flag_api_addr = flag.String("api-addr", ":8844", "Server Local Browser Controller/API Address \"localhost:8844\"")
	flag_api_tls  = flag.Bool("api-tls", false, "Start Local TLS Browser Controller/API Server")
	flag_api_cert = flag.String("api-cert", "", "Local TLS Browser Controller/API Server Cert")
	flag_api_key  = flag.String("api-key", "", "Local TLS Browser Controller/API Server Key")
	// controller
	flag_controller_list = flag.Bool("controller-list", false, "List all Controllers and exit")
	flag_controller      = flag.String("controller", "standard", "Which Controller to use?")
	flag_controller_kv   = flag.String("controller-kv", `{"key":"value"}`, "Controller KeyValue Json Object map[string]interface{} (*Optional, Controller Dependent!*)")
	// coordinates
	// if we have multiple instances of spoofgo running, adding and deleting coordinates won't work currently if the running app decides to save and overwrite
	// we can fix this by not storing coordinates in memory but we will have to reload from the file on every call
	// coordinates shouldn't be accessed that often so this is a good idea
	// we need to look into a cross platform file lock that will work
	// jumping coordinates should also not be encouraged if the goal of the project is to simulate movement data, but has to be allowed
	flag_coordinates_list     = flag.Bool("coordinates-list", false, "List all Saved Coordinates and exit")
	flag_coordinates_set      = flag.String("coordinates-set", "", "Which Cordinate Key to set and use? To use the last known location DO NOT SET!")
	flag_coordinates_add_key  = flag.String("coordinates-add-key", "", "Add Cordinate Key and exit")
	flag_coordinates_add_lat  = flag.Float64("coordinates-add-lat", 0.0, "Add Coordinate Latitude and exit")
	flag_coordinates_add_long = flag.Float64("coordinates-add-long", 0.0, "Add Coordinate Longitude and exit")
	flag_coordinates_delete   = flag.String("coordinates-delete", "", "Which Cordinate Key to delete and exit?")
)

const (
	LOCK_FILE = `read-then-delete-this.txt`
)

func init() {
	rand.Seed(time.Now().UnixNano())
	quit = make(chan struct{})
	// print hello
	clog.LogBoth("hello")
	clog.LogBoth(fmt.Sprintf("Version: %d.%d Build: %d", VERSION_MAJOR, VERSION_MINOR, VERSION_BUILD))
	clog.LogBoth("Website: http://spoofgo.com - for the latest news, guides and downloads")
	clog.LogBoth("Email: spoofgo@gmail.com - feel free to send feeback")
	clog.LogBoth("Donations: http://spoofgo.com/donate")
	clog.LogBoth("we are an open source project, contributions are appreciated!")
	// parse flags
	if !flag.Parsed() {
		flag.Parse()
	}
	// display plugins list and quit?
	if *flag_plugin_list {
		// list plugins and exist
		bs, _ := json.MarshalIndent(plugs.GetPlugins(), "", "\t")
		log.Fatalf("Plugins: \"%s\"\n", bs)
		return
	}
	// display controllers list and quit?
	if *flag_controller_list {
		// list controllers and exist
		bs, _ := json.MarshalIndent(controls.GetControllers(), "", "\t")
		log.Fatalf("Controllers: \"%s\"\n", bs)
		return
	}
	// display coordinates list and quit?
	if *flag_coordinates_list {
		// list coordinates and exist
		bs, _ := json.MarshalIndent(coordinates.List(), "", "\t")
		log.Fatalf("Coordinates: \"%s\"\n", bs)
		return
	}
	if *flag_coordinates_add_key != "" {
		if c := coordinates.Add(*flag_coordinates_add_key, *flag_coordinates_add_lat, *flag_coordinates_add_long); c != nil {
			log.Fatalf("Add Coordinate: \"%s\" %.5f %.5f Added!\n", *flag_coordinates_add_key, c.Latitude, c.Longitude)
		} else {
			log.Fatalf("Add Coordinate: \"%s\" %.5f %.5f Failed to Add!\n", *flag_coordinates_add_key, *flag_coordinates_add_lat, *flag_coordinates_add_long)
		}
		return
	}
	if *flag_coordinates_delete != "" {
		if c := coordinates.Delete(*flag_coordinates_delete); c != nil {
			log.Fatalf("Delete Coordinate: \"%s\" %.5f %.5f Deleted!\n", *flag_coordinates_delete, c.Latitude, c.Longitude)
		} else {
			log.Fatalf("Delete Coordinate: \"%s\" DOES NOT EXIST!\n", *flag_coordinates_delete)
		}
		return
	}
	if *flag_coordinates_set != "" {
		// don't accelerate by default
		c := coordinates.Set(*flag_coordinates_set, false)
		if c == nil {
			log.Fatalf("Set Coordinate: \"%s\" DOES NOT EXIST!\n", *flag_coordinates_set)
			return
		}
		log.Printf("Set Coordinate: \"%s\" %.5f %.5f SET!\n", *flag_coordinates_set, c.Latitude, c.Longitude)
	} else {
		lat, long := movement.GetLatLong()
		log.Printf("USING COORDINATES FROM LAST USE! %.5f %.5f\n", lat, long)
	}
	// create version
	version = v.Create(
		http.DefaultClient,
		VERSION_MAJOR,
		VERSION_MINOR,
		VERSION_BUILD,
	)
	// create plugin?
	if *flag_plugin == "" {
		clog.LogBoth("Plugin Not Selected!")
	} else {
		// parse plugin
		plugin_call, ok := plugs.GetPlugin(*flag_plugin)
		if !ok {
			clog.LogBoth(fmt.Sprintf("Plugin \"%s\" was not found!", *flag_plugin))
			bs, _ := json.MarshalIndent(plugs.GetPlugins(), "", "\t")
			log.Fatalf("Plugins: \"%s\"\n", bs)
			return
		}
		// parse duration
		durf := (*flag_plugin_every == "")
		var every time.Duration
		if durf {
			clog.LogBoth("Plugin Duration was empty!")
		} else {
			var err error
			every, err = time.ParseDuration(*flag_plugin_every)
			if err != nil {
				clog.LogBoth(fmt.Sprintf("Plugin Duration was invalid!: \"%s\"", err))
				durf = true
			}
		}
		if durf {
			log.Fatalln(`A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`)
			return
		}
		clog.LogBoth(fmt.Sprintf("Running Plugin \"%s\" every %s!", *flag_plugin, every.String()))
		// create plugin
		plugin_state := &plug.State{
			Addr:       *flag_plugin_addr,
			HTTPClient: http.DefaultClient,
			KeyValue:   make(map[string]interface{}),
		}
		if len(*flag_plugin_kv) > 0 {
			d := json.NewDecoder(bytes.NewReader([]byte(*flag_plugin_kv)))
			d.UseNumber()
			if err := d.Decode(&plugin_state.KeyValue); err != nil {
				log.Fatalf("FAILED TO UNMARSHAL -plugin-kv: \"%s\"\n", err)
				return
			}
		}
		plugin = plug.Create(
			version.GetVersion(), // version
			plugin_state,         // state
			*flag_plugin,         // using
			plugs.GetPlugins(),   // list of plugins, since we can't include plugins package
			func() bool {
				// plugin and server objects aren't thread safe!
				// call plugin.Start() after creating this plugin object AND the server object and we won't have to worry about that!
				return plugin_call(plugin, server.GetAPI())
			}, // func
			every, // every
		)
	}
	// create controller?
	if *flag_controller == "" {
		clog.LogBoth("Controller Not Selected!")
	} else {
		// parse controller
		controller_call, ok := controls.GetController(*flag_controller)
		if !ok {
			clog.LogBoth(fmt.Sprintf("Controller \"%s\" was not found!", *flag_controller))
			bs, _ := json.MarshalIndent(controls.GetControllers(), "", "\t")
			log.Fatalf("Controllers: \"%s\"\n", bs)
			return
		}
		clog.LogBoth(fmt.Sprintf("Using API Controller \"%s\"!", *flag_controller))
		// create controller
		controller_state := &control.State{
			Plugin:   plugin,
			KeyValue: make(map[string]interface{}),
		}
		if len(*flag_controller_kv) > 0 {
			d := json.NewDecoder(bytes.NewReader([]byte(*flag_controller_kv)))
			d.UseNumber()
			if err := d.Decode(&controller_state.KeyValue); err != nil {
				log.Fatalf("FAILED TO UNMARSHAL -controller-kv: \"%s\"\n", err)
				return
			}
		}
		controller = control.Create(
			version.GetVersion(),      // version
			controller_state,          // state
			*flag_controller,          // using
			controls.GetControllers(), // list of controllers, since we can't include controllers package
			func(
				w *bytes.Buffer,
				r *http.Request,
			) {
				// controller and server objects aren't thread safe!
				// starting our api server after creating this object and we won't have to worry about that!
				controller_call(
					controller,
					server.GetAPI(),
					w,
					r,
				)
			}, // func
		)
	}
	// create news
	news = n.Create(
		http.DefaultClient,
		version.GetVersion(),
	)
	// init api, does not start an http server!!!
	// we need a api object for plugin/controller use
	// api also needs the plugin/controller object so it would be bad to start the server!
	server = api.Create(
		&http.Server{
			Addr:           *flag_api_addr,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		version,
		news,
		plugin,                    // plugin can be nil
		plugs.GetPlugins(),        // list of plugins, since we can't include plugins package
		controller,                // controller can be nil
		controls.GetControllers(), // list of controllers, since we can't include controllers package
	)
	// build options
	BuildOptions()
	// start tickers
	if _, on := o.OPTIONS.Option_Version.Get(); on {
		go version.Start()
	}
}
func quit_func() {
	movement.Save()
	quit_once.Do(func() {
		close(quit)
	})
}
func main() {
	if !*flag_read {
		if bs, err := file.Open(LOCK_FILE); err == nil && len(bs) > 0 {
			clog.LogBoth("")
			clog.LogBoth(fmt.Sprintf("can't start spoof.go yet, you must read and then delete: -> \"%s\" <-  or use the flag -read", LOCK_FILE))
			clog.LogBoth("")
			return
		}
	}
	// briefly show init log
	<-time.After(time.Second / 2)
	// start controller/api server?
	if *flag_api {
		// start controller/api server
		clog.LogBoth("")
		if *flag_api_tls {
			clog.LogBoth(fmt.Sprintf("Starting Local API TLS Server: \"%s\" Cert File: \"%s\" Key File: \"%s\"", *flag_api_addr, *flag_api_cert, *flag_api_key))
			go func() {
				err := server.ListenAndServeTLS(*flag_api_cert, *flag_api_key)
				// wait
				if err != nil {
					clog.LogBoth(fmt.Sprintf("Local API Server ListenAndServeTLS error: \"%s\"", err))
				}
				quit_func()
			}()
		} else {
			clog.LogBoth(fmt.Sprintf("Starting Local API Server: \"%s\"", *flag_api_addr))
			go func() {
				err := server.ListenAndServe()
				// wait
				if err != nil {
					clog.LogBoth(fmt.Sprintf("Local API Server ListenAndServe error: \"%s\"", err))
				}
				quit_func()
			}()
		}
		clog.LogBoth("")
	}
	// start plugin?
	if plugin != nil {
		// start plugin!
		plugin.Start()
	}
	// show help menu briefly
	fmt.Println("")
	fmt.Println(help)
	fmt.Println("")
	if *flag_api && *flag_controller != "" {
		scheme := "http"
		if *flag_api_tls {
			scheme += "s"
		}
		authority := server.ServerAddr()
		if len(authority) > 0 && authority[0] == ':' {
			// prepend localhost?
			// it would be nice to get the actual proper server address, if we were listening on ":0" this couldn't be correct
			// we would have to implement our own listener to figure out what we're bound to
			authority = "localhost" + authority
		}
		fmt.Printf("You can use the controller \"%s\" in your browser: -> %s://%s/controller <- instead of the command line!\n", *flag_controller, scheme, authority)
		fmt.Println("")
	}
	// briefly show help/controller message
	<-time.After(time.Second * 2)
	if *flag_nogui {
		clog.LogBoth("-nogui=\"true\" was set! not spawning CLI GUI! CTRL+C or Close this terminal when finished.")
	} else {
		// launch cli
		Loop()
	}
	// block
	<-quit
	clog.LogBoth("bye")
}
