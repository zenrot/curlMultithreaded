package curl

import (
	"fmt"
	"hedgedcurl/internal/parser"

	"io"
	"net"
)

func GetURL(URL string) (string, error) {

	parsedUrl, err := parser.ParseURL(URL)
	if err != nil {
		return "", err
	}
	var resp []byte

	resp, err = GetURLHTTP(parsedUrl)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func GetURLHTTP(url parser.ParsedURL) ([]byte, error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(url.Host, url.Port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fullPath := url.Path
	if url.RawQuery != "" {
		fullPath += "?" + url.RawQuery
	}

	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nUser-Agent: hedgedcurl\r\nConnection: close\r\n\r\n",
		fullPath, net.JoinHostPort(url.Host, url.Port))

	_, err = conn.Write([]byte(request))
	if err != nil {

		return nil, err
	}
	res, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
