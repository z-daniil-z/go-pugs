package search

import "errors"

var (
	ErrWrongParse = errors.New("url searching error")
	ErrBlock      = errors.New("you've blocked")
)
