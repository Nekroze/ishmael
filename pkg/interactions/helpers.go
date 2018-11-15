package interactions

import (
	"fmt"
	"os"
	"sync"

	"github.com/docker/docker/client"
)

type Client = *client.Client

var dcli Client
var dcliOnce sync.Once

func GetClient() Client {
	dcliOnce.Do(func() {
		ncli, err := client.NewEnvClient()
		possibleFatality(err)
		dcli = ncli
	})
	return dcli
}

func possibleFatality(e error) {
	if e != nil {
		fatality(e)
	}
}

func fatality(e error) {
	fmt.Fprintln(os.Stderr, e)
	os.Exit(1)
}
