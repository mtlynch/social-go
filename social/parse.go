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

// InstagramHandle parses an InstagramHandle from either a raw string or a URL
// that includes the user's handle.
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

// TwitterHandle parses a TwitterHandle from a raw URL or handle string,
// validating that the handle matches Twitter's restrictions on handles.
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
