package parser

import (
	"net/url"
)

type ParsedURL struct {
	Scheme   string
	Host     string
	Port     string
	Path     string
	RawQuery string
}

func ParseURL(raw string) (ParsedURL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return ParsedURL{}, err
	}

	host := u.Hostname()
	port := u.Port()
	if port == "" {
		if u.Scheme == "http" {
			port = "80"
		} else if u.Scheme == "https" {
			port = "443"
		}
	}

	path := u.EscapedPath()
	if path == "" {
		path = "/"
	}
	fullPath := path
	if u.RawQuery != "" {
		fullPath += "?" + u.RawQuery
	}

	return ParsedURL{
		Scheme:   u.Scheme,
		Host:     host,
		Port:     port,
		Path:     path,
		RawQuery: u.RawQuery,
	}, nil
}
