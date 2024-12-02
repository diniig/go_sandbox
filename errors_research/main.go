package main

import (
	"errors"
	"fmt"
	"net"
)

type timeout interface {
	Timeout() bool
}

func main() {

	var err error = getError()
	isTimedOut := isTimeOut(err)
	fmt.Printf("isTimedOut %v", isTimedOut)

}

func isTimeOut(err error) bool {
	t, ok := err.(timeout)
	return ok && t.Timeout()
}

func getError() error {
	return &net.DNSError{
		UnwrapErr: errors.New("error"),
		Err:       "error",
		IsTimeout: true,
	}
}


