package constant

import "errors"

var (
	ErrLinkNotFound = errors.New("link not found")
	ErrLinkExpired  = errors.New("link has been expired")
)
