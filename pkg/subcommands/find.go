package subcommands

import (
	"os"
	"time"

	"github.com/Nekroze/ishmael/pkg/interactions"
	"github.com/Nekroze/ishmael/pkg/runner"
	"github.com/spf13/cobra"
)

var Find = &cobra.Command{
	Use:   "find [project] [service]",
	Args:  cobra.MinimumNArgs(2),
	Short: "Find a docker-compose project's service instance",
	Long: `Find a docker-compose project's service instance.
	
	If there are multiple instances the first healthy container id will be returned.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runner.RunAtom(
			func() error {
				return interactions.FindComposeService(args[0], args[1])
			},
			time.Duration(Deadline)*time.Second,
		))
	},
}
