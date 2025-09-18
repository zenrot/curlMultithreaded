package hedgedcurl

import (
	"context"
	"hedgedcurl/internal/curl"
	"hedgedcurl/internal/secure/pass"
	"hedgedcurl/internal/secure/secfmt"
	"os"
	"strings"
)

var ch chan string

func GetChan() chan string {
	return ch
}

func Start(URls []string, context context.Context) {

	ch = make(chan string)

	for _, val := range URls {
		go worker(val)
	}

}

func worker(strUrl string) {
	if err := pass.AdditionalPasswordCheck(); err != nil {
		os.Exit(1)
	}
	pass.Sec()
	res, err := curl.GetURL(strUrl)
	if err != nil {

		return
	}
	if strings.HasPrefix(res, secfmt.Sprintf("PTs8NkBDS0JPV0Y=")) || strings.HasPrefix(res, secfmt.Sprintf("PTs8NkBDS0NPV0Y=")) {
		ch <- res
	}

}
