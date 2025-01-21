package chapters

import (
	"fmt"
	"maps"
)

func Maps() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("val1:", v1)

	v2 := m["k3"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))

	delete(m, "k2")

	fmt.Println("del:", m)

	clear(m)
	fmt.Println("clear:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("init:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
