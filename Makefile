build:
	go build -o bin/shellymgr main.go

compile:
	echo "Compiling for multiple platforms"
	GOOS=darwin GOARCH=amd64 go build -o bin/shellymgr-darwin main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/shellymgr-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bin/shellymgr main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/shellymgr-windows-amd64.exe main.go

code_vul_scan:
	time gosec ./...

run:
	go run main.go