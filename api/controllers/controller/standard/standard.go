package standard

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/api"
	"sabey.co/spoofgo/api/controllers/controller"
	"sabey.co/spoofgo/duration"
	"sabey.co/spoofgo/movement"
	"time"
)

func no_cache_query() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func Controller(
	control *controller.Controller,
	api *api.API,
	buffer *bytes.Buffer,
	r *http.Request,
) {
	// write
	// stop?
	reload := true
	if r.URL != nil && r.URL.Query().Get("reload") == "no" {
		reload = false
	}
	// top
	buffer.WriteString(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
<meta http-equiv="Pragma" content="no-cache" />
<meta http-equiv="Expires" content="0" />
`)
	if reload {
		buffer.WriteString(`<meta http-equiv="refresh" content="1; url=/controller?`)
		buffer.WriteString(no_cache_query())
		buffer.WriteString(`">
`)
	}
	buffer.WriteString(`<style type="text/css">
td a {display: block; height:100%; width:100%;}
td a:hover {background-color: rgb(239, 239, 239);}
</style>
<title>Spoof.go | Standard Controller</title>
</head>
<body>
`)
	// controller
	buffer.WriteString(`<table style="width: 100%;">
<tbody>
<tr>
<td style="width: 33%;">Angle: `)
	buffer.WriteString(fmt.Sprintf("%d", api.Angle))
	buffer.WriteString(`</td>
<td style="width: 35%;">Lat: `)
	buffer.WriteString(fmt.Sprintf("%.6f", api.Latitude))
	buffer.WriteString(`</td>
<td style="width: 33%;">Long: `)
	buffer.WriteString(fmt.Sprintf("%.6f", api.Longitude))
	buffer.WriteString(`</td>
</tr>
</tbody>
</table>
<table style="width: 100%;">
<tbody>
<tr>
<td style="width: 20%;">Speed: `)
	buffer.WriteString(fmt.Sprintf("%.2f", api.Speed))
	buffer.WriteString(` km/h</td>
<td style="width: 20%;">Modifier: `)
	buffer.WriteString(fmt.Sprintf("%d", api.Modifier))
	buffer.WriteString(`</td>
<td style="width: 20%;"><a href="/modifier/increase?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">Increase</a></td>
<td style="width: 20%;"><a href="/modifier/decrease?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">Decrease</a></td>
<td style="width: 20%;"><a href="/modifier/reset?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">Reset</a></td>
</tr>
</tbody>
</table>
<table style="width: 100%;">
<tbody>
<tr>
<td style="width: 20%; text-align: center;"><a href="/mode/toggle?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">TOGGLE</a></td>
<td style="width: 20%; text-align: center; `)
	if api.Mode == movement.WALK {
		buffer.WriteString("background-color: rgb(84, 172, 210);")
	}
	buffer.WriteString(`"><a href="/mode/walk?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">WALK</a></td>
<td style="width: 20%; text-align: center; `)
	if api.Mode == movement.JOG {
		buffer.WriteString("background-color: rgb(84, 172, 210);")
	}
	buffer.WriteString(`"><a href="/mode/jog?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">JOG</a></td>
<td style="width: 20%; text-align: center; `)
	if api.Mode == movement.BICYCLE {
		buffer.WriteString("background-color: rgb(84, 172, 210);")
	}
	buffer.WriteString(`"><a href="/mode/bicycle?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">BICYCLE</a></td>
<td style="width: 20%; text-align: center; `)
	if api.Mode == movement.CAR {
		buffer.WriteString("background-color: rgb(84, 172, 210);")
	}
	buffer.WriteString(`"><a href="/mode/car?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">CAR</a></td>
</tr>
</tbody>
</table>
<table style="width: 100%;">
<tbody>
<tr>
<td style="width: 33%; text-align: center;"><a href="/northwest?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">NW</a></td>
<td style="width: 34%; text-align: center; `)
	if api.Accelerating {
		buffer.WriteString("background-color: rgb(97, 189, 109);")
	}
	buffer.WriteString(`"><a href="/accelerate?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">ACCELERATE</a></td>
<td style="width: 33%; text-align: center;"><a href="/northeast?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">NE</a></td>
</tr>
<tr>
<td style="width: 33%; text-align: center;"><a href="/left?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">LEFT</a></td>
<td style="width: 34%; text-align: center; `)
	if !api.Accelerating {
		buffer.WriteString("background-color: rgb(226, 80, 65);")
	}
	buffer.WriteString(`"><a href="/decelerate?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">DECELERATE</a></td>
<td style="width: 33%; text-align: center;"><a href="/right?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">RIGHT</a></td>
</tr>
<tr>
<td style="width: 33%; text-align: center;"><a href="/southwest?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">SW</a></td>
<td style="width: 34%; text-align: center;"><a href="/flip?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">FLIP</a></td>
<td style="width: 33%; text-align: center;"><a href="/southeast?controller=index&`)
	if !reload {
		buffer.WriteString(`reload=no&`)
	}
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">SE</a></td>
</tr>
</tbody>
</table>
<hr />
`)
	if control.GetState().Plugin == nil {
		buffer.WriteString("<h3>No Plugin Selected!</h3>")
	} else {
		buffer.WriteString(fmt.Sprintf("<h3>Running Plugin \"%s\" every %s!</h3>\n", control.GetState().Plugin.GetUsing(), control.GetState().Plugin.GetEvery().String()))
		info := control.GetState().Plugin.GetInfo()
		if info.LastSuccess.IsZero() {
			buffer.WriteString("<h1 style=\"color: red;\">Plugin Not Online Yet!</h1>")
		} else {
			if info.Offline {
				buffer.WriteString(fmt.Sprintf("<h1 style=\"color: red;\">Plugin Offline For %s</h1>", duration.Round(info.SinceLast, time.Second).String()))
			}
		}
	}
	buffer.WriteString("<hr />\n")
	if reload {
		buffer.WriteString(`<h1><a href="/controller?reload=no&`)
		buffer.WriteString(no_cache_query())
		buffer.WriteString(`">Stop Reloading</a></h1>
`)
	} else {
		buffer.WriteString(`<h1><a href="/controller?&`)
		buffer.WriteString(no_cache_query())
		buffer.WriteString(`">Start Reloading</a></h1>
`)
	}
	buffer.WriteString(`<h3><a href="/?`)
	buffer.WriteString(no_cache_query())
	buffer.WriteString(`">Home</a></h3>
`)
	// bottom
	buffer.WriteString(`</body>
</html>
`)
}
