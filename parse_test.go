package social_test

import (
	"strings"
	"testing"

	"github.com/mtlynch/social-go/v2"
)

func TestParseFacebookUsername(t *testing.T) {
	var tests = []struct {
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
	}

	for _, tt := range tests {
		actualUsername, err := social.ParseFacebookUsername(tt.in)
		if err != tt.expectedErr {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.in, err, tt.expectedErr)
		} else if actualUsername != tt.expectedUsername {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.in, err, tt.expectedErr)
		}
	}
}

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
			social.ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com",
			social.InstagramHandle(""),
			social.ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/",
			social.InstagramHandle(""),
			social.ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/@",
			social.InstagramHandle(""),
			social.ErrInvalidInstagram,
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
			social.ErrInvalidInstagram,
		},
		{
			"trailing space character is invalid",
			"jack ",
			social.InstagramHandle(""),
			social.ErrInvalidInstagram,
		},
		{
			"internal tab character is invalid",
			"jerry\tseinfeld",
			social.InstagramHandle(""),
			social.ErrInvalidInstagram,
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
			social.ErrInvalidInstagram,
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

func TestParseTwitterHandle(t *testing.T) {
	var tests = []struct {
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
			social.ErrInvalidTwitterURL,
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
			social.ErrInvalidTwitterURL,
		},
		{
			"missing handle is invalid",
			"https://twitter.com/",
			social.TwitterHandle(""),
			social.ErrInvalidTwitterURL,
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
	}

	for _, tt := range tests {
		handle, err := social.ParseTwitterHandle(tt.url)
		if err != tt.errExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.url, err, tt.errExpected)
		} else if handle != tt.handleExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.url, handle, tt.handleExpected)
		}
	}
}
