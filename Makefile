# Get current directory
current_dir := $(shell pwd)

# Get current commit hash
# commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
.PHONY: build

all: clean build assets

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/VisionaryQuery $(current_dir)/cmd/visionaryquery.go

assets:
	@echo "Copy assets"
	cp example.yml $(current_dir)/build/config.yml
	cp README.md $(current_dir)/build/README.md

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build
