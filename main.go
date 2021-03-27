package main

import (
	"os"
	"strconv"
	"strings"
)

// fix the step for smaller changes
const defaultStep = 25

var Side string

func main() {

	var side string
	var step int = defaultStep

	args := len(os.Args[1:])

	switch {
	case args <= 0:
		os.Exit(1)
	case args >= 1:
		side = strings.ToLower(os.Args[1])
	case args >= 2:
		step, err := strconv.Atoi(os.Args[2])
		if step == 0 || err != nil {
			step = defaultStep
		}
	}

	if side == "up" {
		DecreaseBrightness(step)
	} else if side == "down" {
		IncreaseBrightness(step)
	}
}
