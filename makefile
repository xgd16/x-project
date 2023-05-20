
#BUILD_SYSTEM=GOOS=windows GOARCH=amd64
#BUILD_SYSTEM=GOOS=linux GOARCH=amd64
BUILD_SYSTEM=GOOS=darwin GOARCH=arm64
#SUFFIX=.exe

build-pack:
	gf pack x-project packed/data.go -y
	cd ./build && $(BUILD_SYSTEM) go build -gcflags='-l -l -l' -ldflags='-s -w' -pgo=auto -o ../bin/xInit$(SUFFIX) main.go