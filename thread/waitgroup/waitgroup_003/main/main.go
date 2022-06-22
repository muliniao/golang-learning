package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func main() {

	err := Group().Error()
	fmt.Println(err)

}

func Group() error {
	var g multierror.Group

	g.Go(func() error {
		fmt.Println("go routine 1")
		err := errors.New("err from go routine1")
		return err
	})

	g.Go(func() error {
		fmt.Println("go routine 2")
		err := errors.New("err from go routine2")
		return err
	})

	g.Go(func() error {
		fmt.Println("go routine 3")
		err := errors.New("err from go routine3")
		return err
	})

	return g.Wait()
}
