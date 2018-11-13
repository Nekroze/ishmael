package interactions

import (
	"context"
	"time"

	"github.com/docker/docker/client"
)

var Deadline = time.Now().Add(time.Second)

func ContainerIsAlive(id string) (bool, error) {
	ctx, cancel := context.WithDeadline(context.Background(), Deadline)
	defer cancel()

	cli, err := client.NewEnvClient()
	if err != nil {
		return false, err
	}

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return false, nil
	}

	return info.State != nil &&
			info.State.Running &&
			!info.State.Restarting,
		nil
}
