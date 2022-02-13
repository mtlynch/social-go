package social

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidTwitterHandle = errors.New("invalid twitter handle")

	twitterHandlePattern = regexp.MustCompile("^[A-Za-z0-9_]{4,15}$")
)

// ParseTwitterHandle parses an untrusted string into a TwitterHandle if the
// string contained a well-formed Twitter handle. If the string does not include
// a handle that conforms to Instagram's rules, this function returns an error.
// This function does not verify whether the handle exists, only that it adheres
// to Twitter's public rules for a handle.
func ParseTwitterHandle(twitterURL string) (TwitterHandle, error) {
	t, err := parseSocialMediaUsername(twitterURL)
	if err != nil {
		return TwitterHandle(""), ErrInvalidTwitterHandle
	}
	if t == "" {
		return TwitterHandle(""), ErrInvalidTwitterHandle
	}

	if !twitterHandlePattern.MatchString(t) {
		return TwitterHandle(""), ErrInvalidTwitterHandle
	}
	return TwitterHandle(t), nil
}
