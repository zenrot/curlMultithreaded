package hedgedcurl

import (
	"context"
	"hedgedcurl/internal/curl"
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
	res, err := curl.GetURL(strUrl)
	if err != nil {
		return
	}
	if strings.HasPrefix(res, "HTTP/1.0 20") || strings.HasPrefix(res, "HTTP/1.1 20") {
		ch <- res
	}

}
