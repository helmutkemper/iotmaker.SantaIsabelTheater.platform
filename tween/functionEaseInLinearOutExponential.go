package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInLinearOutExponential = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta*interactionCurrent/interactionTotal + startValue
	}
	interactionCurrentToCalc--
	return delta/2*(-1*math.Pow(2, -10*interactionCurrentToCalc)+2) + startValue
}
