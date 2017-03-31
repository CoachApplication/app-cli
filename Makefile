.PHONY: all build local deps fmt test binary install clean

# This script is the switchboard for the make targets
MAKE_SCRIPT="./make/make.sh"

# BY default do the all build
default: all

# Typical make all target
all: build

# Typical build routine
build: deps binary install

# Quick developer build
local: fmt binary install
# Full developer build, best to do before Pull Requests
local-full: fmt test deps local binary install


# Run the formatting on all source
fmt:
	${MAKE_SCRIPT} fmt

# Run all go tests
test:
	${MAKE_SCRIPT} test

# Build a binary executable
binary:
	${MAKE_SCRIPT} binary

# Make sure that all the vendor dependencies are available and properly versioned
deps:
	${MAKE_SCRIPT} deps

# Install a compiled binary
install:
	${MAKE_SCRIPT} install

# Remove the compiled binary
clean:
	${MAKE_SCRIPT} clean
