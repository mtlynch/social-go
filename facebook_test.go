package social_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mtlynch/social-go/v2"
)

func TestParseFacebookUsername(t *testing.T) {
	for _, tt := range []struct {
		explanation      string
		in               string
		expectedUsername social.FacebookUsername
		expectedErr      error
	}{
		{
			"regular username is valid",
			"jerry",
			social.FacebookUsername("jerry"),
			nil,
		},
		{
			"regular URL is valid",
			"https://facebook.com/jerry",
			social.FacebookUsername("jerry"),
			nil,
		},
		{
			"username with dots is valid",
			"jerry.seinfeld",
			social.FacebookUsername("jerry.seinfeld"),
			nil,
		},
		{
			"long page name is valid",
			"Joe-Smith-Magician-129875824511529",
			social.FacebookUsername("Joe-Smith-Magician-129875824511529"),
			nil,
		},
		{
			"long page URL is valid",
			"https://facebook.com/Joe-Smith-Magician-129875824511529",
			social.FacebookUsername("Joe-Smith-Magician-129875824511529"),
			nil,
		},
		{
			"username with numbers is valid",
			"jerry123",
			social.FacebookUsername("jerry123"),
			nil,
		},
		{
			"empty string is invalid",
			"",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"single character username is valid",
			"a",
			social.FacebookUsername("a"),
			nil,
		},
		{
			"leading @ character is invalid",
			"@mark",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"leading space character is invalid",
			" jerry",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"trailing space character is invalid",
			"jerry ",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"internal tab character is invalid",
			"jerry\tseinfeld",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"username with underscore is invalid",
			"jerry_seinfeld",
			"",
			social.ErrInvalidFacebookUsername,
		},
		{
			"username with exactly 60 characters is valid",
			strings.Repeat("A", 60),
			social.FacebookUsername(strings.Repeat("A", 60)),
			nil,
		},
		{
			"username with more than 60 characters is invalid",
			strings.Repeat("A", 61),
			"",
			social.ErrInvalidFacebookUsername,
		},
	} {
		t.Run(fmt.Sprintf("%s [%s]", tt.explanation, tt.in), func(t *testing.T) {
			username, err := social.ParseFacebookUsername(tt.in)

			if got, want := err, tt.expectedErr; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := username, tt.expectedUsername; got != want {
				t.Errorf("username=%v, want=%v", got, want)
			}
		})
	}
}
