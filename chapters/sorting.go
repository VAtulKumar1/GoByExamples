package chapters

import (
	"fmt"
	"slices"
)

func Soring() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strings:", strs)

	s := slices.IsSorted(strs)
	fmt.Println("sorted:", s)

}
