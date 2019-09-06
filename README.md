# brightnessctl

A small utility to UP and DOWN the brightness

## Installation

```
make
make install
```

## Usage

Increase brightness: `brightness_up`

Decrease brightness: `brightness_down`

You can bind it in your lovely wm. For spectrwm it is like:

*~/.spectrwm.conf*
```
program[br_up]          = brightness_up
program[br_down]        = brightness_down

bind[br_up]             = MOD+F6
bind[br_down]           = MOD+F5
```

## Description

Simply sets a new value into a sysfs parameter for backlight device. Tested on intel ASUS X200M. Installed with as setuid bit, so no root is needed to run the binaries. 
