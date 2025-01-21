package chapters

import "fmt"

func Rou(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

}
