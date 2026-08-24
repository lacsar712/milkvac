package clock

import (
	"time"

	"github.com/lacsar712/milkvac/internal/model"
)

type PulseWindow struct {
	clk      Clock
	duration time.Duration
}

func NewPulseWindow(clk Clock, duration time.Duration) *PulseWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &PulseWindow{clk: clk, duration: duration}
}

func (w *PulseWindow) Active(anchor time.Time) bool {
	// Track the injected process beat, not the wall clock: when the beat-lock
	// freezes the process clock, the closure window must freeze with it instead
	// of advancing on wall time.
	now := anchor
	if w.clk != nil {
		now = w.clk.Now()
	}
	elapsed := now.Sub(anchor)
	if elapsed < 0 {
		elapsed = 0
	}
	return elapsed < w.duration
}

func (w *PulseWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrPulseHold
}
