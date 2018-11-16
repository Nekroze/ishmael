package interactions

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/docker/docker/api/types"
)

func FindComposeService(project, service string) error {
	ctx := context.Background()

	cli := GetClient()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	for _, container := range containers {
		if container.Labels["com.docker.compose.oneoff"] == "True" ||
			container.Labels["com.docker.compose.project"] != project ||
			container.Labels["com.docker.compose.service"] != service {
			continue
		}

		info, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			return runner.UpgradeToEphemeral(err)
		}

		if info.State != nil && info.State.Running && !info.State.Restarting {
			fmt.Println(info.ID)
			return nil
		}
	}
	return runner.UpgradeToEphemeral(errors.New("No running containers found matching that project and service"))
}
