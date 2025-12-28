package intermediate

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrOutofTea = errors.New("out of tea")
var ErrPower = errors.New("can't boil water")


func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutofTea
	} else if arg == 4 {
		return fmt.Errorf("makeTea failed: %w", ErrPower)
}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		
		if r, e := f(i); e != nil {
			fmt.Println("f failed")
		} else  {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutofTea) {
				fmt.Println("We need to buy more tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("We need to fix the power")
			} else {
				fmt.Println("Unknown error:", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}

}
