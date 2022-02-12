package social_test

import (
	"strings"
	"testing"

	"github.com/mtlynch/social-go"
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
		/*{
			"regular URL is valid",
			"https://facebook.com/jerry",
			true,
		},
		{
			"username with dots is valid",
			"jerry.seinfeld",
			true,
		},
		{
			"long page name is valid",
			"Matt-Woodland-Comedian-119671814711529",
			true,
		},
		{
			"long page URL is valid",
			"https://facebook.com/Matt-Woodland-Comedian-119671814711529",
			true,
		},
		{
			"username with numbers is valid",
			"jerry123",
			true,
		},
		{
			"empty string is invalid",
			"",
			false,
		},
		{
			"single character username is valid",
			"q",
			true,
		},
		{
			"leading @ character is invalid",
			"@mark",
			false,
		},
		{
			"leading space character is invalid",
			" jerry",
			false,
		},
		{
			"trailing space character is invalid",
			"jerry ",
			false,
		},
		{
			"internal tab character is invalid",
			"jerry\tseinfeld",
			false,
		},
		{
			"username with underscore is invalid",
			"jerry_seinfeld",
			false,
		},
		{
			"username with exactly 60 characters is valid",
			strings.Repeat("A", 60),
			true,
		},
		{
			"username with more than 60 characters is invalid",
			strings.Repeat("A", 61),
			false,
		},
		{
			"'undefined' as a username is invalid",
			"undefined",
			false,
		},*/
	}

	for _, tt := range tests {
		_, err := social.ParseFacebookUsername(tt.username)
		if (err == nil) != tt.validExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.username, err, tt.validExpected)
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
			ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
		},
		{
			"missing handle is invalid",
			"https://instagram.com/@",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
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
			ErrInvalidInstagram,
		},
		{
			"trailing space character is invalid",
			"jack ",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
		},
		{
			"internal tab character is invalid",
			"jerry\tseinfeld",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
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
			ErrInvalidInstagram,
		},
		{
			"'undefined' as a URL is invalid",
			"undefined",
			social.InstagramHandle(""),
			ErrInvalidInstagram,
		},
	}

	for _, tt := range tests {
		handle, err := ParseInstagramHandle(tt.in)
		if err != tt.errExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.in, err, tt.errExpected)
		} else if err == nil && handle != tt.handleExpected {
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
			ErrInvalidTwitterURL,
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
			ErrInvalidTwitterURL,
		},
		{
			"missing handle is invalid",
			"https://twitter.com/",
			social.TwitterHandle(""),
			ErrInvalidTwitterURL,
		},
		{
			"single character handle is invalid",
			"https://twitter.com/q",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
		},
		{
			"handle with internal whitespace is invalid",
			"jerry seinfeld",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
		},
		{
			"handle with internal tab is invalid",
			"jerry\tseinfeld",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
		},
		{
			"handle with leading whitespace is invalid",
			" jerry",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
		},
		{
			"handle with trailing whitespace is invalid",
			"jerry ",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
		},
		{
			"handle with dots is invalid",
			"https://twitter.com/jerry.seinfeld",
			social.TwitterHandle(""),
			ErrInvalidTwitterHandle,
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
			ErrInvalidTwitterHandle,
		},
		{
			"'undefined' as a URL is invalid",
			"undefined",
			social.TwitterHandle(""),
			ErrInvalidTwitterURL,
		},
	}

	for _, tt := range tests {
		handle, err := TwitterHandle(tt.url)
		if err != tt.errExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.url, err, tt.errExpected)
		} else if handle != tt.handleExpected {
			t.Errorf("%s: input [%s], got %v, want %v", tt.explanation, tt.url, handle, tt.handleExpected)
		}
	}
}
