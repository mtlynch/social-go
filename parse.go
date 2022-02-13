package social

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	ErrInvalidFacebookUsername = errors.New("invalid facebook username")
	ErrInvalidInstagram        = errors.New("invalid instagram handle")
	ErrInvalidTwitterURL       = errors.New("invalid twitter URL")
	ErrInvalidTwitterHandle    = errors.New("invalid twitter handle")

	facebookUsernamePattern = regexp.MustCompile(`^[a-zA-Z0-9\.\-]{1,60}$`)
	instagramHandlePattern  = regexp.MustCompile(`^[a-zA-Z0-9_\.]{1,30}$`)
	twitterHandlePattern    = regexp.MustCompile("^[A-Za-z0-9_]{4,15}$")
)

// ParseFacebookUsername parses an untrusted string into a FacebookUsername if
// the string contained a well-formed Facebook username. If the string does not
// include a username that conforms to Facebook's username rules, this function
// returns an error. This function does not verify whether the username exists,
// only that it adheres to Facebook's public rules for a username.
func ParseFacebookUsername(username string) (FacebookUsername, error) {
	if strings.ContainsAny(username, "@") {
		return FacebookUsername(""), ErrInvalidFacebookUsername
	}

	fb, err := parseSocialMediaUsername(username)
	if err != nil {
		return FacebookUsername(""), ErrInvalidFacebookUsername
	}

	if !facebookUsernamePattern.MatchString(fb) {
		return FacebookUsername(""), ErrInvalidFacebookUsername
	}
	return FacebookUsername(fb), nil
}

// ParseInstagramHandle parses an untrusted string into a InstagramHandle if the
// string contained a well-formed Instagram handle. If the string does not
// include a handle that conforms to Instagram's rules, this function returns an
// error. This function does not verify whether the handle exists, only that it
// adheres to Instagram's public rules for a handle.
func ParseInstagramHandle(s string) (InstagramHandle, error) {
	insta, err := parseSocialMediaUsername(s)
	if err != nil {
		return InstagramHandle(""), ErrInvalidInstagram
	}

	if !instagramHandlePattern.MatchString(insta) {
		return InstagramHandle(""), ErrInvalidInstagram
	}

	return InstagramHandle(insta), nil
}

// ParseTwitterHandle parses an untrusted string into a TwitterHandle if the
// string contained a well-formed Twitter handle. If the string does not include
// a handle that conforms to Instagram's rules, this function returns an error.
// This function does not verify whether the handle exists, only that it adheres
// to Twitter's public rules for a handle.
func ParseTwitterHandle(twitterURL string) (TwitterHandle, error) {
	t, err := parseSocialMediaUsername(twitterURL)
	if err != nil {
		return TwitterHandle(""), ErrInvalidTwitterURL
	}
	if t == "" {
		return TwitterHandle(""), ErrInvalidTwitterURL
	}

	if !twitterHandlePattern.MatchString(t) {
		return TwitterHandle(""), ErrInvalidTwitterHandle
	}
	return TwitterHandle(t), nil
}

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
