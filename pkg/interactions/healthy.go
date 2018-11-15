package interactions

import (
	"context"
	"errors"

	"github.com/Nekroze/ishmael/pkg/runner"
)

func ContainerIsHealthy(id string) error {
	ctx, cancel := context.WithDeadline(context.Background(), Deadline)
	defer cancel()

	info, err := GetClient().ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	if info.State == nil || info.State.Health == nil {
		return errors.New("Container does not have health state information")
	}

	if info.State.Health.Status == "unhealthy" {
		return errors.New("Container is unhealthy")
	}

	if info.State.Health.Status == "healthy" {
		return nil
	}

	return runner.UpgradeToEphemeral(errors.New("Container not yet healthy"))
}
