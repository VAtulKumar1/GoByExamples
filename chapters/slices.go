package chapters

import (
	"fmt"
	"slices"
)

func Slices() {
	var s []string
	fmt.Println("uninti:", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Set:", s)
	fmt.Println("get:", s[1])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("Sl3:", l)

	t := []string{"g", "f", "h"}
	fmt.Println("dc1:", t)

	t1 := []string{"g", "f", "h"}

	if slices.Equal(t, t1) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD:", twoD)

}
