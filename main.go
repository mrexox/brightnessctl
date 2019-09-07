package main

// fix the step for smaller changes
const brightnessStep   = 250

var Side string

func main() {
	if Side == "Up" {
		HandleBrightness(increase)
	}
	if Side == "Down" {
		HandleBrightness(decrease)
	}
}

func increase(v1 int) int {
	return v1 + brightnessStep
}

func decrease(v1 int) int {
	return v1 - brightnessStep
}
