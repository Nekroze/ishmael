package interactions

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Nekroze/ishmael/pkg/runner"
)

func ContainerAddresses(id string) error {
	info, err := GetClient().ContainerInspect(context.Background(), id)
	if err != nil {
		return err
	}

	if info.NetworkSettings == nil {
		return runner.UpgradeToEphemeral(errors.New("Container has no network settings"))
	}

	addr := info.NetworkSettings.IPAddress
	for _, net := range info.NetworkSettings.Networks {
		if net.IPAddress != "" {
			addr = net.IPAddress
			break
		}
	}
	if addr == "" {
		addr = "127.0.0.1"
	}

	if info.Config == nil {
		return runner.UpgradeToEphemeral(errors.New("Container has no config settings"))
	}

	if len(info.Config.ExposedPorts) == 0 {
		return errors.New("Container has no exposed ports")
	}
	for port := range info.Config.ExposedPorts {
		fmt.Printf("%s:%s\n",
			addr,
			strings.Split(string(port), "/")[0],
		)
	}
	return nil
}
