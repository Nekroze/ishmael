package cmd

import (
	"os"

	"github.com/Nekroze/ishmael/interactions"
	"github.com/spf13/cobra"
)

var aliveCmd = &cobra.Command{
	Use:   "alive [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "check if a container is up and running",
	Long:  `check if a container is up and running.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runCmd(func() (bool, error) {
			return interactions.ContainerIsAlive(args[0])
		}))
	},
}

func init() {
	rootCmd.AddCommand(aliveCmd)
}
