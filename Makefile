all: brightness_up brightness_down

brightness_up: 
	go build -o $@ -ldflags "-X main.Side=Up" 

brightness_down: 
	go build -o $@ -ldflags "-X main.Side=Down" 

clean:
	rm -f brightness_up brightness_down
