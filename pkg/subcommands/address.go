package subcommands

import (
	"os"
	"time"

	"github.com/Nekroze/ishmael/pkg/interactions"
	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/spf13/cobra"
)

var Address = &cobra.Command{
	Use:   "address [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get ip and port addresses this container exposes",
	Long:  `Get ip and port addresses this container exposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runner.RunAtom(
			func() error {
				return interactions.ContainerAddresses(args[0])
			},
			time.Duration(Deadline)*time.Second,
			false,
		))
	},
}
