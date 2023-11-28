
FILENAME=async-api-server

default: help

help:
	@echo "Available 'make' commands"
	@echo "  - make local - compiles the executable for the current architecture"
	@echo "  - make all - compiles for all architectures and copyes results into the necessary folders"
	@echo "  - make config-files - prints-out templates for the configuration files, based on the latest compiled version of the alert"

local:
	CGO_ENABLED=1 go build

all: ensure_dirs .osx .linux .osxm1

ensure_dirs:
	- mkdir -p ../linux_x86_64/bin/
	- mkdir -p ../darwin_x86_64/bin/
	- mkdir -p ../darwin_arm_64/bin/

#Environment settings for cross compilation
#Ref: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

.linux:
	@echo "> Compiling for Linux within ../linux_x86_64/bin/apiserver"
	@echo "  If you are on a Mac, you need 'brew install FiloSottile/musl-cross/musl-cross' before running this"
	CGO_ENABLED=1 CC=x86_64-linux-musl-gcc GOOS=linux GOARCH=amd64 CXX=x86_64-linux-musl-g++ go build -ldflags "-linkmode external -extldflags -static" -o ../linux_x86_64/bin/apiserver
  
.osx:
	@echo "> Compiling for OsX (darwin) w../darwin_x86_64/bin/apiserver"
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ../darwin_x86_64/bin/apiserver

.osxm1:
	@echo "> Compiling for OsX M1 (darwin) ../darwin_arm_64/bin/apiserver"
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ../darwin_arm_64/bin/apiserver