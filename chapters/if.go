package chapters

import "fmt"

func IFElse() {
	if 7%2 == 0 {
		fmt.Print("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "num is negative")
	} else if num < 10 {
		fmt.Println(num, " has single digit")
	} else {
		fmt.Println(num, " has single digit")
	}
}
