package runner

import (
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/pkg/errors"
)

type CommandAtom func() error

type EphemeralError struct {
	error
}

func UpgradeToEphemeral(e error) error {
	return EphemeralError{e}
}

type ErrorClassifier struct {
}

func (c ErrorClassifier) Classify(err error) retrier.Action {
	if err == nil {
		return retrier.Succeed
	}
	if _, ok := errors.Cause(err).(*EphemeralError); ok {
		return retrier.Retry
	}
	return retrier.Fail
}

var AttemptInterval = time.Second

func RunAtom(atom CommandAtom, timeout time.Duration) (exit int) {
	r := retrier.New(
		retrier.ConstantBackoff(int(timeout/AttemptInterval), AttemptInterval),
		ErrorClassifier{},
	)
	if e := r.Run(atom); e != nil {
		exit = 1
	}
	return exit
}
