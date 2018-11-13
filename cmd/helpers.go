package cmd

import (
	"time"

	"github.com/briandowns/spinner"
)

var wait int

func init() {
	rootCmd.PersistentFlags().IntVar(&wait, "wait", 0, "If given a positive integer then it will wait up to this many seconds for the subcommand to succeed.")
}

type cmdFunc func() (bool, error)

func runCmd(cf cmdFunc) int {
	if wait <= 0 {
		if !tickCmd(cf) {
			return 1
		}
		return 0
	}

	s := spinner.New(spinner.CharSets[43], 250*time.Millisecond)
	s.Start()
	defer s.Stop()

	timeout := time.After(time.Duration(wait) * time.Second)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-timeout:
			return 1
		case <-tick:
			if tickCmd(cf) {
				return 0
			}
		}
	}
}

func tickCmd(cf cmdFunc) bool {
	done, err := cf()
	return err == nil && done
}
