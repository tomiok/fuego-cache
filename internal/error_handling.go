package internal

import "github.com/tomiok/fuego-cache/logs"

//OnCloseError will - not handle - the error but just logging when an I/O operations is closed.
//This only should be used when a close function is called
func OnCloseError(fn func() error) {
	if err := fn(); err != nil {
		logs.LogError(err)
	}
}
