package app

import (
	"context"
	"fmt"

	"github.com/lacsar712/milkvac/internal/interlock"
	"github.com/lacsar712/milkvac/internal/model"
)

func (a *App) HandleVacTrip(ctx context.Context, tower model.TowerID, celsius float64) error {
	if celsius <= a.cfg.TargetMoistPct+40 {
		return nil
	}
	if err := a.guard.Permit(model.ZoneID(tower.String()+"-zone-00"), model.PlenumID("plenum-main")); err != nil {
		return err
	}
	_ = interlock.DefaultLeaseTTL
	// Carry the vacuum-trip sentinel so the cross-layer reporting path can
	// classify this as a vacuum-drop event (errors.Is(err, model.ErrVacTrip))
	// rather than falling through to the generic "state conflict" bucket,
	// which hides the dedicated reset menu and leaves the reset button grayed.
	return fmt.Errorf("heat alarm: zone %s exceeded limit at %.1fC: %w", tower, celsius, model.ErrVacTrip)
}
