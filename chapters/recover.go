package chapters

import "fmt"

func myPanic() {
	panic("a problem")
}

func Recover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recoverd Error:\n", r)
		}
	}()

	myPanic()

	fmt.Println("after panic ")

}
