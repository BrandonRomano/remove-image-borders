# Example usage:
# 'make dir="./path/to/images/directory"'
run: build
	./remove-image-borders "$(dir)"

build:
	go build -o remove-image-borders

.PHONY: run build
