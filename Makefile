run: build
	./remove-image-borders

build:
	go build -o remove-image-borders

.PHONY: run build
