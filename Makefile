#
# Change BIN_DIR var if you what to have brightnessctl
# installed into another place
#
BINDIR ?= /usr/local/bin

MANDIR ?= /usr/share/man

.PHONY: build
build:
	go build -o brightnessctl
	gzip --keep --force brightnessctl.1

# Installing with setuid bit so no root is needed for it
.PHONY: install
install:
	install -D -m 4755 brightnessctl $(DESTDIR)/$(BINDIR)/brightnessctl
	install -D -m 0644 brightnessctl.1.gz $(DESTRID)/$(MANDIR)/man1/brightnessctl.1.gz

.PHONY: clean
clean:
	rm -f brightnessctl brightnessctl.1.gz
