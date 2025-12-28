package intermediate

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func main() {

	_, err := f(42)
	var ae *argError

	if errors.As(err, &ae) {
		fmt.Println("argError prob:", ae.prob)
		fmt.Println("argError arg:", ae.arg)
	} else {
		fmt.Println("err doesn't match argError")
	}

}
