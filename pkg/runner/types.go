package runner

import (
	"context"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/pkg/errors"
	"github.com/thecodeteam/goodbye"
)

type CommandAtom func() error

type EphemeralError struct {
	error
}

func UpgradeToEphemeral(e error) error {
	return &EphemeralError{e}
}

type ErrorClassifier struct {
	shouldSpin bool
}

func Classifier(nospin bool) *ErrorClassifier {
	return &ErrorClassifier{nospin}
}

func (c *ErrorClassifier) Classify(err error) retrier.Action {
	if err == nil {
		return retrier.Succeed
	}
	if _, ok := errors.Cause(err).(*EphemeralError); ok {
		c.maybeStartSpinning()
		return retrier.Retry
	}
	return retrier.Fail
}

var AttemptInterval = time.Second

func RunAtom(atom CommandAtom, timeout time.Duration, spinOnRetry bool) (exit int) {
	r := retrier.New(
		retrier.ConstantBackoff(int(timeout/AttemptInterval), AttemptInterval),
		Classifier(!spinOnRetry),
	)
	if e := r.Run(atom); e != nil {
		exit = 1
	}
	return exit
}

func (c *ErrorClassifier) maybeStartSpinning() {
	if c.shouldSpin {
		return
	}

	s := spinner.New(spinner.CharSets[12], 250*time.Millisecond)
	goodbye.Register(func(_ context.Context, _ os.Signal) {
		defer s.Stop()
	})
	s.Start()
	c.shouldSpin = true
}
