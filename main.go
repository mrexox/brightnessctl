package main

var Side string

func main() {
	if Side == "Up" {
		HandleBrightness(increase)
	}
	if Side == "Down" {
		HandleBrightness(decrease)
	}
}

func increase (v1 int, v2 int) int {
	return v1 + v2
}

func decrease (v1 int, v2 int) int {
	return v1 - v2
}
