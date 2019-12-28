package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
func NewEaseInOutSineInfiniteLoop(duration time.Duration, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), onStartFunc func(value float64), onEndFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		OnStart:     onStartFunc,
		OnEnd:       onEndFunc,
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutSine,
		Interaction: interactionFunc,
		Repeat:      -1,
	}
	t.Start()

	return t
}