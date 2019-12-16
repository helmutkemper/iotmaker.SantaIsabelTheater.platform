package tween

import "math"

// en: sinusoidal easing out - decelerating to zero velocity
var KEaseOutSine = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return changeInValue*math.Sin(currentTime/duration*(math.Pi/2)) + startValue
}
