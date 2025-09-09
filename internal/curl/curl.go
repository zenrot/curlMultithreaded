package curl

import (
	"crypto/tls"
	"fmt"
	"hedgedcurl/internal/parser"
	"hedgedcurl/internal/secure/secfmt"

	"io"
	"net"
)

func GetURL(URL string) (string, error) {

	parsedUrl, err := parser.ParseURL(URL)
	if err != nil {
		return "", err
	}
	var resp []byte

	if parsedUrl.Scheme == secfmt.Sprintf("HRscFhw=") {
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
	conn, err := net.Dial(secfmt.Sprintf("AQwY"), net.JoinHostPort(url.Host, url.Port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fullPath := url.Path
	if fullPath == "" {
		fullPath = secfmt.Sprintf("Wg==")
	}
	if url.RawQuery != "" {
		fullPath += secfmt.Sprintf("Sg==") + url.RawQuery
	}
	hostHeader := url.Host
	if url.Port != secfmt.Sprintf("TV8=") && url.Port != "" {
		hostHeader = net.JoinHostPort(url.Host, url.Port)
	}

	request := fmt.Sprintf(
		secfmt.Sprintf("Mio8RkoBRTo7MSZdVFxTOR0qASoaHBxcT1cWLh05GCcWFxBILhEKDAFVSA4KFgIXCwYDAAkuEDkBNwwBEB8cXE9YSlgzFyocJBEBAB8CQicbDAcCBhwCSE8MEhcLBgsRFiodPhssBwgBFwYGBgoYSEURDgocEzMQKQE0FDMc"),
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
	conn, err := tls.Dial(secfmt.Sprintf("AQwY"), net.JoinHostPort(url.Host, url.Port), &tls.Config{
		ServerName: url.Host, // для проверки сертификата
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fullPath := url.Path
	if fullPath == "" {
		fullPath = secfmt.Sprintf("Wg==")
	}
	if url.RawQuery != "" {
		fullPath += secfmt.Sprintf("Sg==") + url.RawQuery
	}
	hostHeader := url.Host
	if url.Port != secfmt.Sprintf("TV8=") && url.Port != "" {
		hostHeader = net.JoinHostPort(url.Host, url.Port)
	}

	request := fmt.Sprintf(
		secfmt.Sprintf("Mio8RkoBRTo7MSZdVFxTOR0qASoaHBxcT1cWLh05GCcWFxBILhEKDAFVSA4KFgIXCwYDAAkuEDkBNwwBEB8cXE9YSlgzFyocJBEBAB8CQicbDAcCBhwCSE8MEhcLBgsRFiodPhssBwgBFwYGBgoYSEURDgocEzMQKQE0FDMc"),
		fullPath, hostHeader)

	fmt.Println(request)

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
