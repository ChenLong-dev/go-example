.PHONY:  clean build_x86 build_arm64 build_win64 build_mac64

all: build_x86 build_arm64 build_win64 build_mac64

clean:
	rm -rf build*

build_x86:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o build_x86 main.go

build_x64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build_x64 main.go

build_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build_arm64 main.go

build_win64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build_win64 main.go

build_mac64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build_mac64 main.go
