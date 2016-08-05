package duration

// https://groups.google.com/forum/#!msg/golang-nuts/OWHmTBu16nA/ECOzv1rBCzcJ
// https://play.golang.org/p/QHocTHl8iR

import (
	"time"
)

func Round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}
