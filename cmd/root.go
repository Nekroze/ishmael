package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
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

var wait int

func init() {
	rootCmd.PersistentFlags().IntVar(&wait, "wait", 0, "If given a positive integer then it will wait up to this many seconds for the subcommand to succeed.")
}

type cmdFunc func() (bool, error)

func runCmd(cf cmdFunc) {
	if wait <= 0 {
		if !tickCmd(cf) {
			os.Exit(1)
		}
		return
	}

	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	timeout := time.After(time.Duration(wait) * time.Second)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-timeout:
			os.Exit(1)
		case <-tick:
			if tickCmd(cf) {
				os.Exit(0)
			}
		}
	}
}

func tickCmd(cf cmdFunc) bool {
	done, err := cf()
	if err != nil {
		panic(err)
	}
	fmt.Println("TICK OUTPUT", done)
	return done
}
