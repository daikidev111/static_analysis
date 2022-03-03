package a

import (
	"errors"

	"golang.org/x/sync/errgroup"
)

func f() {
	errorCh := make(chan error)
	go func(ch chan error) { // Want "Go test"
		ch <- errors.New("error")
	}(errorCh)

	var eg errgroup.Group
	eg.Go(func() error {
		return errors.New("error")
	})
}
