package social_test

import (
	"strings"
	"testing"

	"github.com/mtlynch/social-go/v2"
)

func TestParseInstagramHandle(t *testing.T) {
	var tests = []struct {
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
	}

	for _, tt := range tests {
		handle, err := social.ParseInstagramHandle(tt.in)
		if err != tt.errExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.in, err, tt.errExpected)
		} else if handle != tt.handleExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.in, handle, tt.handleExpected)
		}
	}
}
