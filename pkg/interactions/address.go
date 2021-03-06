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
		return runner.UpgradeToEphemeral(errors.New("container has no network settings"))
	}

	for _, mappings := range info.NetworkSettings.Ports {
		for _, mapping := range mappings {
			fmt.Printf("%s:%s\n", mapping.HostIP, mapping.HostPort)
		}
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
		return runner.UpgradeToEphemeral(errors.New("container has no config settings"))
	}

	if len(info.Config.ExposedPorts) == 0 {
		return errors.New("container has no exposed ports")
	}
	seen := map[string]bool{}
	for port := range info.Config.ExposedPorts {
		pnum := strings.Split(string(port), "/")[0]
		if !seen[pnum] {
			fmt.Printf("%s:%s\n", addr, pnum)
			seen[pnum] = true
		}
	}
	return nil
}
