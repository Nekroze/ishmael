package interactions

import (
	"context"
	"errors"
	"time"

	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/docker/docker/client"
)

var Deadline = time.Now().Add(time.Second)

func ContainerIsAlive(id string) error {
	ctx, cancel := context.WithDeadline(context.Background(), Deadline)
	defer cancel()

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	if info.State == nil || !info.State.Running || info.State.Restarting {
		return runner.UpgradeToEphemeral(errors.New("Container not running"))
	}

	return nil
}
