package cmd

import (
	"github.com/Nekroze/ishmael/pkg/subcommands"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ishmael",
	Short: "watch the waves for whale sign",
	Long:  `Provides various ways to inspect docker objects for scripting.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
func init() {
	rootCmd.PersistentFlags().IntVar(&subcommands.Deadline, "wait", 1, "If given a positive integer then it will wait up to this many seconds for the subcommand to succeed.")

	rootCmd.AddCommand(subcommands.Alive)
	rootCmd.AddCommand(subcommands.Healthy)
	rootCmd.AddCommand(subcommands.Find)
}
