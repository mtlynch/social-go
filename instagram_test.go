package social_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mtlynch/social-go/v2"
)

func TestParseInstagramHandle(t *testing.T) {
	for _, tt := range []struct {
		explanation    string
		in             string
		handleExpected social.InstagramHandle
		errExpected    error
	}{
		{
			"regular handle is valid",
			"jerry",
			social.InstagramHandle("jerry"),
			nil,
		},
		{
			"handle with dots is valid",
			"jerry.seinfeld",
			social.InstagramHandle("jerry.seinfeld"),
			nil,
		},
		{
			"handle with underscore is valid",
			"jerry_seinfeld",
			social.InstagramHandle("jerry_seinfeld"),
			nil,
		},
		{
			"handle with numbers is valid",
			"jerry123",
			social.InstagramHandle("jerry123"),
			nil,
		},
		{
			"empty string is invalid",
			"",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"missing handle is invalid",
			"https://instagram.com",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/@",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"single character handle is valid",
			"q",
			social.InstagramHandle("q"),
			nil,
		},
		{
			"leading @ character is valid",
			"@jack",
			social.InstagramHandle("jack"),
			nil,
		},
		{
			"leading @ character in URL is valid",
			"http://instagram.com/@jack",
			social.InstagramHandle("jack"),
			nil,
		},
		{
			"leading space character is invalid",
			" jack",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"trailing space character is invalid",
			"jack ",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"internal tab character is invalid",
			"jerry\tseinfeld",
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
		{
			"handle with exactly 30 characters is valid",
			"https://instagram.com/" + strings.Repeat("A", 30),
			social.InstagramHandle(strings.Repeat("A", 30)),
			nil,
		},
		{
			"handle with more than 30 characters is invalid",
			"https://instagram.com/" + strings.Repeat("A", 31),
			social.InstagramHandle(""),
			social.ErrInvalidInstagramHandle,
		},
	} {
		t.Run(fmt.Sprintf("%s [%s]", tt.explanation, tt.in), func(t *testing.T) {
			handle, err := social.ParseInstagramHandle(tt.in)

			if got, want := err, tt.errExpected; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := handle, tt.handleExpected; got != want {
				t.Errorf("handle=%v, want=%v", got, want)
			}
		})
	}
}
