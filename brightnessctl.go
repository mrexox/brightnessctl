package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	backlightDevPath = "/sys/class/backlight"
	brightnessParam  = "brightness"
	brightnessStep   = 500
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getDriver(sysfsPath string) (string, error) {
	return "intel_backlight", nil
}

func concat(strs ...string) string {
	var sumstr string

	for _, s := range strs {
		sumstr += s
	}

	return sumstr
}

func readValue(path string) int {
	val, err := ioutil.ReadFile(path)
	check(err)

	res := strings.TrimSuffix(string(val), "\n")

	i, err := strconv.Atoi(res)
	check(err)

	return i
}

func writeValue(path string, value int) {
	strValue := []byte(strconv.Itoa(value))

	err := ioutil.WriteFile(path, strValue, 0644)
	check(err)
}

func HandleBrightness(f func(int, int) int) {
	// Getting driver from sysfs
	driver, err := getDriver(backlightDevPath)
	check(err)

	syspath := concat(backlightDevPath,
		"/",
		driver,
		"/",
		brightnessParam)

	value := readValue(syspath)

	newvalue := f(value, brightnessStep)

	writeValue(syspath, newvalue)
}
