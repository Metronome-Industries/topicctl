package main

import (
	"github.com/Metronome-Industries/topicctl/cmd/topicctl/subcmd"
)

var (
	// Version is the version of this binary. Overridden as part of the build process.
	Version = "metronome-dev"
)

func main() {
	subcmd.Execute(Version)
}
