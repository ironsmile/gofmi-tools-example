package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ironsmile/moons"
)

var arguments = struct {
	planet string
}{}

func init() {
	flag.Usage = func() {
		fmt.Fprintln(
			flag.CommandLine.Output(),
			"Command line utility for greeting all the planets.\n\nOPTIONS",
		)
		flag.PrintDefaults()
	}

	flag.StringVar(&arguments.planet, "planet", "World", "Configure which planet to greet")
}

func main() {
	flag.Parse()

	fmt.Printf("Hello, %s!\n", strings.Title(arguments.planet))

	count, err := moons.Count(moons.BodyFromString(arguments.planet))
	if err == nil {
		fmt.Printf("How are your %d moons today?\n", count)
	}
}
