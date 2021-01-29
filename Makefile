# Get current directory
current_dir := $(shell pwd)

# Get current commit hash
commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
.PHONY: build

all: clean build assets

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/VisionaryQuery $(current_dir)/cmd/visionaryquery.go

assets:
	@echo "Copy assets"
	cp example.yml $(current_dir)/build/
	cp README.md $(current_dir)/build/

bundle: clean build assets
	tar czf visionaryquery-build_$(commit_hash).tar.gz -C $(current_dir)/build .

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build
