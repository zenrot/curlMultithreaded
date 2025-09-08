.PHONY: build
build:
	go build -v -o hedgedcurlMac ./cmd/app/hedgedcurl/hedgedcurl.go
	GOOS=windows GOARCH=amd64 go build -v -o hedgedcurl.exe ./cmd/app/hedgedcurl/hedgedcurl.go
.PHONY: delete
delete:
	rm -rf ./hedgedcurl
.DEFAULT_GOAL := build