package main

import (
	"context"

	"github.com/Nekroze/ishmael/cmd"
	"github.com/thecodeteam/goodbye"
)

func main() {
	goodbye.Notify(context.Background())
	cmd.Execute()
}
