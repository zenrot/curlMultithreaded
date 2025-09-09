package parser

import (
	"hedgedcurl/internal/secure/secfmt"
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
		if u.Scheme == secfmt.Sprintf("HRscFg==") {
			port = secfmt.Sprintf("TV8=")
		} else if u.Scheme == secfmt.Sprintf("HRscFhw=") {
			port = secfmt.Sprintf("QVtb")
		}
	}

	path := u.EscapedPath()
	if path == "" {
		path = secfmt.Sprintf("Wg==")
	}
	fullPath := path
	if u.RawQuery != "" {
		fullPath += secfmt.Sprintf("Sg==") + u.RawQuery
	}

	return ParsedURL{
		Scheme:   u.Scheme,
		Host:     host,
		Port:     port,
		Path:     path,
		RawQuery: u.RawQuery,
	}, nil
}
