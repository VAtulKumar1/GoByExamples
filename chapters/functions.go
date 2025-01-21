package chapters

import "fmt"

func Plus(a int, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

func Vals() (int, int) {
	return 2, 3
}

func Sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0

	for _, num := range nums {
		total = total + num
	}

	fmt.Println("sum:", total)
}
