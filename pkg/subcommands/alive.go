package subcommands

import (
	"os"
	"time"

	"github.com/Nekroze/ishmael/pkg/interactions"
	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/briandowns/spinner"
	isatty "github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var Alive = &cobra.Command{
	Use:   "alive [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "check if a container is up and running",
	Long:  `check if a container is up and running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Deadline > 1 && isatty.IsTerminal(os.Stdout.Fd()) {
			s := spinner.New(spinner.CharSets[12], 250*time.Millisecond)
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
