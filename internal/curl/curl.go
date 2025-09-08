package curl

import (
	"crypto/tls"
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

	if parsedUrl.Scheme == "https" {
		resp, err = GetURLHTTPS(parsedUrl)
		if err != nil {
			return "", err
		}
		return string(resp), nil
	}

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
	if fullPath == "" {
		fullPath = "/"
	}
	if url.RawQuery != "" {
		fullPath += "?" + url.RawQuery
	}
	hostHeader := url.Host
	if url.Port != "80" && url.Port != "" {
		hostHeader = net.JoinHostPort(url.Host, url.Port)
	}

	request := fmt.Sprintf(
		"GET %s HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"User-Agent: hedgedcurl\r\n"+
			"Accept: */*\r\n"+
			"Accept-Encoding: identity\r\n"+
			"Connection: close\r\n\r\n",
		fullPath, hostHeader)

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

func GetURLHTTPS(url parser.ParsedURL) ([]byte, error) {
	conn, err := tls.Dial("tcp", net.JoinHostPort(url.Host, url.Port), &tls.Config{
		ServerName: url.Host, // для проверки сертификата
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fullPath := url.Path
	if fullPath == "" {
		fullPath = "/"
	}
	if url.RawQuery != "" {
		fullPath += "?" + url.RawQuery
	}

	hostHeader := url.Host
	if url.Port != "443" && url.Port != "" {
		hostHeader = net.JoinHostPort(url.Host, url.Port)
	}

	request := fmt.Sprintf(
		"GET %s HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"User-Agent: hedgedcurl\r\n"+
			"Accept: */*\r\n"+
			"Accept-Encoding: identity\r\n"+
			"Connection: close\r\n\r\n",
		fullPath, hostHeader)

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
