package main

import "github.com/sigidagi/skystatus/cmd/skystatus/cmd"

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}
