package interactions

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func FindComposeService(project, service string) error {
	ctx, cancel := context.WithDeadline(context.Background(), Deadline)
	defer cancel()

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	if len(containers) < 1 {
		return runner.UpgradeToEphemeral(errors.New("No containers found matching that project and service"))
	}

	for _, container := range containers {
		fmt.Println("found", container.Names)
		if container.Labels["com.docker.compose.oneoff"] == "True" ||
			container.Labels["com.docker.compose.project"] != project ||
			container.Labels["com.docker.compose.service"] != service {
			fmt.Println("unwanted", container.Names)
			continue
		}

		fmt.Println("inspecting", container.Names)
		info, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			return runner.UpgradeToEphemeral(err)
		}

		fmt.Println("checking", container.Names)
		if info.State != nil && info.State.Running && !info.State.Restarting {
			fmt.Println("returning", container.Names)
			fmt.Println(info.ID)
			return nil
		}
	}
	return runner.UpgradeToEphemeral(errors.New("No running containers found matching that project and service"))
}
