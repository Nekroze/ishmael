package interactions

import (
	"context"
	"errors"

	"github.com/Nekroze/ishmael/pkg/runner"
)

func ContainerIsHealthy(id string) error {
	info, err := GetClient().ContainerInspect(context.Background(), id)
	if err != nil {
		return err
	}

	if info.State == nil || info.State.Health == nil {
		return errors.New("container does not have health state information")
	}

	if info.State.Health.Status == "unhealthy" {
		return errors.New("container is unhealthy")
	}

	if info.State.Health.Status == "healthy" {
		return nil
	}

	return runner.UpgradeToEphemeral(errors.New("container not yet healthy"))
}
