BIN := /usr/local/bin

all: brightness_up brightness_down

brightness_up: 
	go build -o $@ -ldflags "-X main.Side=Up" 

brightness_down: 
	go build -o $@ -ldflags "-X main.Side=Down" 

# Installing with setuid bit so no root is needed for it
# Change BIN var if you what to have it in anoher place
.PHONY: install
install: brightness_down brightness_up
	install -D -m 4755 brightness_up $(BIN)/brightness_up
	install -D -m 4755 brightness_down $(BIN)/brightness_down

.PHONY: clean
clean:
	rm -f brightness_up brightness_down
