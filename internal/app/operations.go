package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/milkvac/internal/model"
)

func (a *App) ValidateVacDrift(ctx context.Context, moistPct float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	limit := a.cfg.TargetMoistPct + a.cfg.MaxGradientDeltaPct
	if moistPct <= limit {
		return nil
	}
	return fmt.Errorf("moisture: %w", model.ErrVacDrift)
}

func (a *App) ConfirmPulseHold(ctx context.Context, anchor time.Time) error {
	if a.avgWindow == nil {
		return model.Wrap("app", "window", model.ErrPulseHold)
	}
	if err := a.avgWindow.Require(anchor); err != nil {
		return fmt.Errorf("gradient hold: %w", err)
	}
	return nil
}
