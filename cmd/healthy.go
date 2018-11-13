package cmd

import (
	"os"

	"github.com/Nekroze/ishmael/interactions"
	"github.com/spf13/cobra"
)

var healthyCmd = &cobra.Command{
	Use:   "healthy [container_id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "check if a container is up and healthy",
	Long:  `check if a container is up and healthy.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(runCmd(func() (bool, error) {
			return interactions.ContainerIsHealthy(args[0])
		}))
	},
}

func init() {
	rootCmd.AddCommand(healthyCmd)
}
