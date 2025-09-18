.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -gcflags="-N -l" -v -o hedgedcurl.exe ./cmd/app/hedgedcurl/hedgedcurl.go
.PHONY: delete
delete:
	rm -rf ./hedgedcurl
.PHONY: run
run:
	wine hedgedcurl.exe https://example.com 2>/dev/null
.DEFAULT_GOAL := build