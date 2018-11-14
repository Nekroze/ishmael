package subcommands

import (
	"os"
	"time"

	"github.com/Nekroze/ishmael/pkg/interactions"
	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var Alive = &cobra.Command{
	Use:   "alive [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "check if a container is up and running",
	Long:  `check if a container is up and running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Deadline > 0 {
			s := spinner.New(spinner.CharSets[43], 250*time.Millisecond)
			s.Start()
			defer s.Stop()
		}

		os.Exit(runner.RunAtom(
			func() error {
				return interactions.ContainerIsAlive(args[0])
			},
			time.Duration(Deadline)*time.Second,
		))
	},
}
