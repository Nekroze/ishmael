package interactions

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/docker/docker/client"
)

func ContainerIsHealthy(id string) (bool, error) {
	ctx, cancel := context.WithDeadline(context.Background(), Deadline)
	defer cancel()

	cli, err := client.NewEnvClient()
	if err != nil {
		return false, nil
	}

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return false, nil
	}

	spew.Dump(info.State)

	return info.State != nil &&
			info.State.Health != nil &&
			info.State.Health.Status == "healthy",
		nil
}
