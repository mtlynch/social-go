package social

import (
	"fmt"
	"net/url"
	"strings"
)

func parseSocialMediaUsername(s string) (string, error) {
	cleaned := strings.TrimSpace(s)
	if len(cleaned) == 0 {
		return "", nil
	}
	// Short-circuit the logic if there's no URL, just a username.
	if !strings.Contains(cleaned, "/") {
		return stripLeadingAtSymbol(s), nil
	}

	u, err := parseURL(cleaned)
	if err != nil {
		return "", err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("invalid URL scheme: %v", u.Scheme)
	}

	splitFn := func(c rune) bool {
		return c == '/'
	}
	pathParts := strings.FieldsFunc(u.Path, splitFn)
	if len(pathParts) == 0 || len(pathParts[0]) == 0 {
		return "", fmt.Errorf("invalid social media URL: %s - bad path: %s", s, u.Path)
	}
	return stripLeadingAtSymbol(pathParts[0]), nil
}

func stripLeadingAtSymbol(s string) string {
	if strings.HasPrefix(s, "@") {
		return s[1:]
	}
	return s
}

func parseURL(s string) (*url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "" {
		return u, nil
	}
	return url.Parse("https://" + s)
}
