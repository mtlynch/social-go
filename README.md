# social-go

[![CircleCI](https://circleci.com/gh/mtlynch/social-go.svg?style=svg)](https://circleci.com/gh/mtlynch/social-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/mtlynch/social-go.svg)](https://pkg.go.dev/github.com/mtlynch/social-go)
[![License](https://img.shields.io/badge/license-Unlicense-blue)](LICENSE)

A Go parser for various social media handles and URLs.

## Overview

social-go parses usernames from various social media platforms, verifying that the supplied username matches validity rules for that platform.

social-go's checks are local - they don't verify whether the account actually exists on the platform, just that it conforms to the platform's published schemas for usernames.

## Twitter

```golang
raw := "jack"
t1, _ := social.ParseTwitterHandle(raw)
fmt.Printf("%s -> %s\n", raw, t1)
// jack -> jack

raw = "@jack"
t2, _ := social.ParseTwitterHandle(raw)
fmt.Printf("%s -> %s\n", raw, t2)
// @jack -> jack

raw = "twitter.com/@jack"
t3, _ := social.ParseTwitterHandle(raw)
fmt.Printf("%s -> %s\n", raw, t3)
// twitter.com/@jack -> jack
```

## Facebook

```golang
raw := "zuck"
fb, _ := social.ParseFacebookUsername(raw)
fmt.Printf("%s -> %s\n", raw, fb)
// zuck -> zuck

raw = "https://facebook.com/zuck"
fb2, _ := social.ParseFacebookUsername(raw)
fmt.Printf("%s -> %s\n", raw, fb2)
// https://facebook.com/zuck -> zuck
```

## Instagram

```golang
raw := "chelseahandler"
h1, _ := social.ParseInstagramHandle(raw)
fmt.Printf("%s -> %s\n", raw, h1)
// chelseahandler -> chelseahandler

raw = "@chelseahandler"
h2, _ := social.ParseInstagramHandle(raw)
fmt.Printf("%s -> %s\n", raw, h2)
// @chelseahandler -> chelseahandler

raw = "instagram.com/@chelseahandler"
h3, _ := social.ParseInstagramHandle(raw)
fmt.Printf("%s -> %s\n", raw, h3)
// instagram.com/@chelseahandler -> chelseahandler
```
