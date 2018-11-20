package interactions

import (
	"context"
	"errors"
	"time"

	"github.com/Nekroze/ishmael/pkg/runner"
)

var Deadline = time.Now().Add(time.Second)

func ContainerIsAlive(id string) error {
	info, err := GetClient().ContainerInspect(context.Background(), id)
	if err != nil {
		return err
	}

	if info.State == nil || !info.State.Running || info.State.Restarting {
		return runner.UpgradeToEphemeral(errors.New("Container not running"))
	}

	return nil
}
