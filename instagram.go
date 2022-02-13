package social

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidInstagramHandle = errors.New("invalid instagram handle")

	instagramHandlePattern = regexp.MustCompile(`^[a-zA-Z0-9_\.]{1,30}$`)
)

// ParseInstagramHandle parses an untrusted string into a InstagramHandle if the
// string contained a well-formed Instagram handle. If the string does not
// include a handle that conforms to Instagram's rules, this function returns an
// error. This function does not verify whether the handle exists, only that it
// adheres to Instagram's public rules for a handle.
func ParseInstagramHandle(s string) (InstagramHandle, error) {
	insta, err := parseSocialMediaUsername(s)
	if err != nil {
		return InstagramHandle(""), ErrInvalidInstagramHandle
	}

	if !instagramHandlePattern.MatchString(insta) {
		return InstagramHandle(""), ErrInvalidInstagramHandle
	}

	return InstagramHandle(insta), nil
}
