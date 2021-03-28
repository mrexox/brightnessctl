# brightnessctl

[![Go](https://github.com/mrexox/brightnessctl/actions/workflows/go.yml/badge.svg)](https://github.com/mrexox/brightnessctl/actions/workflows/go.yml)

A small utility to UP and DOWN the screen brightness (for laptops)

:warning: :penguin: Linux only :warning:

## Installation

```
make
make install BINDIR=/usr/bin
```

## Usage

Increase brightness: `brightnessctl up` or `brightnessctl up 55`

Decrease brightness: `brightnessctl down` or `brightnessctl down 35`

You can bind it in your lovely wm. For spectrwm it is like:

## Binding keys

**~/.spectrwm.conf**
```
program[br_up]          = brightnessctl up
program[br_down]        = brightnessctl down

bind[br_up]             = MOD+F6
bind[br_down]           = MOD+F5
```

or despite the WM you use:

**~/.xbindkeysrc**
```
"brightnessctl up 10"
  XF86MonBrightnessUp

"brightnessctl down 10"
  XF86MonBrightnessDown
```

## Description

Simply sets a new value into a sysfs parameter for backlight device. Tested on Intel Core i3 (ASUS X200M) and AMD Ryzen 7 (HP Probook 445). Installed with setuid bit, so no root is needed to run the binaries.
