package social_test

import (
	"fmt"

	"github.com/mtlynch/social-go/v2"
)

func Example_facebook() {
	raw := "zuck"
	fb, _ := social.ParseFacebookUsername(raw)
	fmt.Printf("%s -> %s\n", raw, fb)

	raw = "https://facebook.com/zuck"
	fb2, _ := social.ParseFacebookUsername(raw)
	fmt.Printf("%s -> %s\n", raw, fb2)
}

func Example_twitter() {
	raw := "jack"
	t1, _ := social.ParseTwitterHandle(raw)
	fmt.Printf("%s -> %s\n", raw, t1)

	raw = "@jack"
	t2, _ := social.ParseTwitterHandle(raw)
	fmt.Printf("%s -> %s\n", raw, t2)

	raw = "twitter.com/@jack"
	t3, _ := social.ParseTwitterHandle(raw)
	fmt.Printf("%s -> %s\n", raw, t3)
}

func Example_instagram() {
	raw := "chelseahandler"
	h1, _ := social.ParseInstagramHandle(raw)
	fmt.Printf("%s -> %s\n", raw, h1)

	raw = "@chelseahandler"
	h2, _ := social.ParseInstagramHandle(raw)
	fmt.Printf("%s -> %s\n", raw, h2)

	raw = "instagram.com/@chelseahandler"
	h3, _ := social.ParseInstagramHandle(raw)
	fmt.Printf("%s -> %s\n", raw, h3)
}
