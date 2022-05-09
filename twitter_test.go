package social_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mtlynch/social-go/v2"
)

func TestParseTwitterHandle(t *testing.T) {
	for _, tt := range []struct {
		explanation    string
		url            string
		handleExpected social.TwitterHandle
		errExpected    error
	}{
		{
			"regular handle on its own is valid",
			"jerry",
			social.TwitterHandle("jerry"),
			nil,
		},
		{
			"regular handle in URL is valid",
			"https://twitter.com/jerry",
			social.TwitterHandle("jerry"),
			nil,
		},
		{
			"leading @ character is valid",
			"https://twitter.com/@jack",
			social.TwitterHandle("jack"),
			nil,
		},
		{
			"missing scheme is valid",
			"twitter.com/jerry",
			social.TwitterHandle("jerry"),
			nil,
		},
		{
			"query string is valid",
			"http://twitter.com/jerry?ref=somejunk",
			social.TwitterHandle("jerry"),
			nil,
		},
		{
			"invalid scheme is invalid",
			"ftp://twitter.com/jerry",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with underscore is valid",
			"https://twitter.com/jerry_seinfeld",
			social.TwitterHandle("jerry_seinfeld"),
			nil,
		},
		{
			"handle with numbers is valid",
			"https://twitter.com/jerry123",
			social.TwitterHandle("jerry123"),
			nil,
		},
		{
			"empty string is invalid",
			"",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"missing handle is invalid",
			"https://twitter.com/",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"single character handle is invalid",
			"https://twitter.com/q",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with internal whitespace is invalid",
			"jerry seinfeld",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with internal tab is invalid",
			"jerry\tseinfeld",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with leading whitespace is invalid",
			" jerry",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with trailing whitespace is invalid",
			"jerry ",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with dots is invalid",
			"https://twitter.com/jerry.seinfeld",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
		{
			"handle with exactly 15 characters is valid",
			"https://twitter.com/" + strings.Repeat("A", 15),
			social.TwitterHandle(strings.Repeat("A", 15)),
			nil,
		},
		{
			"handle with more than 15 characters is invalid",
			"https://twitter.com/" + strings.Repeat("A", 16),
			social.TwitterHandle(""),
			social.ErrInvalidTwitterHandle,
		},
	} {
		t.Run(fmt.Sprintf("%s [%s]", tt.explanation, tt.url), func(t *testing.T) {
			handle, err := social.ParseTwitterHandle(tt.url)

			if got, want := err, tt.errExpected; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := handle, tt.handleExpected; got != want {
				t.Errorf("handle=%v, want=%v", got, want)
			}
		})
	}
}
