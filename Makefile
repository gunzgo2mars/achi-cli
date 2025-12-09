build-mac-silicon:
	GOOS=darwin GOARCH=arm64 go build -o bin/mac-os-silicon/achi cmd/main.go

build-mac-intel:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac-os-intel/achi cmd/main.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/achi cmd/main.go

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/achi cmd/main.go
