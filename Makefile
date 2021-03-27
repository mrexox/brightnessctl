#
# Change BIN_DIR var if you what to have brightnessctl
# installed into another place
#
BINDIR ?= /usr/local/bin

.PHONY: build
build:
	go build -o brightnessctl

# Installing with setuid bit so no root is needed for it
.PHONY: install
install: build
	install -D -m 4755 brightnessctl $(DESTDIR)/$(BINDIR)/brightnessctl

.PHONY: clean
clean:
	rm -f brightnessctl
