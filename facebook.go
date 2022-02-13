package social

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidFacebookUsername = errors.New("invalid facebook username")

	facebookUsernamePattern = regexp.MustCompile(`^[a-zA-Z0-9\.\-]{1,60}$`)
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
