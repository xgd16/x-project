#BUILD_SYSTEM=
#BUILD_SYSTEM=GOOS=linux GOARCH=amd64

ARGS=-gcflags='-l -l -l' -ldflags='-s -w' -pgo=auto

#SUFFIX=.exe

build-mac-arm64:
	gf pack x-project packed/data.go -y
	cd ./build && GOOS=darwin GOARCH=arm64 go build $(ARGS) -o ../bin/xInit_darwin_arm64 main.go

build-mac-amd64:
	gf pack x-project packed/data.go -y
	cd ./build && GOOS=darwin GOARCH=amd64 go build $(ARGS) -o ../bin/xInit_darwin_amd64 main.go

build-linux-amd64:
	gf pack x-project packed/data.go -y
	cd ./build && GOOS=linux GOARCH=amd64 go build $(ARGS) -o ../bin/xInit_linux_amd64 main.go

build-windows-amd64:
	gf pack x-project packed/data.go -y
	cd ./build && GOOS=windows GOARCH=amd64 go build $(ARGS) -o ../bin/xInit_windows_amd64.exe main.go

all: build-mac-arm64 build-mac-amd64 build-linux-amd64 build-windows-amd64