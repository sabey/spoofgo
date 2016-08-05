# [Spoof.go](https://www.spoofgo.com)
An Application for Spoofing Movement

* Controllable via Terminal, Browser, or API
  * Browser Controls are extensible with controller plugins, see: api/controllers/controller/standard/
* Persistent Pushing of Coordinates and Movement State
  * Pushing methods are extensible with plugins, see: plugins/plugin/standard/

```
Movement:

Accelerate:	W, Arrow Up
Left:		A, Arrow Left
Decelerate:	S, Arrow Down
Right:		D, Arrow Right

Angles:
NorthWest:	Q
NorthEast:	E
SouthWest:	Z
Flip:		X
SouthEast:	C

Mode:
Toggle:		~
Set:		1-4

Mode Modifier:
Increase:	R, =, +, Page Up
Decrease:	F, -, _, Page Down
Reset:		V, 0
```