package chapters

import "fmt"

type ArgError struct {
	Arg     int
	Message string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Message)
}

func F1(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
