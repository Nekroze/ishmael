package subcommands

import (
	"os"
	"time"

	"github.com/Nekroze/ishmael/pkg/interactions"
	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/spf13/cobra"
)

var Healthy = &cobra.Command{
	Use:   "healthy [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "check if a container is up and healthy",
	Long:  `check if a container is up and healthy.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runner.RunAtom(
			func() error {
				return interactions.ContainerIsHealthy(args[0])
			},
			time.Duration(Deadline)*time.Second,
			true,
		))
	},
}
