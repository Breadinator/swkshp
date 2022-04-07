package main

import (
	"github.com/breadinator/swkshp/cmd"
)

const (
	VERSION = "v1.1.0"
)

func main() {
	autoUpdate()
	cmd.Execute()
}
