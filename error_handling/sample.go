package error_handling

import (
	"errors"

	"golang.org/x/sync/errgroup"
)

func main() {
	errorCh := make(chan error)
	go func(ch chan error) {
		ch <- errors.New("error")
	}(errorCh)

	var eg errgroup.Group
	eg.Go(func() error {
		return errors.New("error")
	})
}

// ast.GoStmt
// call calling func
// func lit has type and literal
// list len(1) means thehre is on ly one param
// if go func has the error channel then it can be replaceable
// ch <- chan for sending & chan -> error for receiving
// arrow type <- , -> , -

// if func lit of the gostmt
