package util

import "net/url"

// URLUserPassword is a replacement/wrapper around url.UserPassword that treats empty string arguments as not specified
// If no user or password is specified, it returns nil (which serializes in url.URL to "")
func URLUserPassword(user, password string) *url.Userinfo { _ = "STUB: not implemented"; return nil }
