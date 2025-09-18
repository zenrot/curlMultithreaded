.PHONY: build_go
build_go:
	GOOS=windows GOARCH=amd64 go build -gcflags="-N -l" -ldflags="-s -w" -v -o hedgedcurl.exe ./cmd/app/hedgedcurl/hedgedcurl.go
.PHONY: build_garble
build_garble:
	GOOS=windows GOARCH=amd64 garble build -ldflags="-s -w" -o hedgedcurlG.exe ./cmd/app/hedgedcurl
.PHONY: delete
delete:
	rm -rf ./hedgedcurl
.PHONY: run
run:
	wine hedgedcurl.exe https://example.com 2>/dev/null
.DEFAULT_GOAL := build