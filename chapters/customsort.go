package chapters

import (
	"cmp"
	"fmt"
	"slices"
)

func CustomSort() {

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "jax", age: 27},
		Person{name: "max", age: 29},
		Person{name: "sax", age: 10},
		Person{name: "ajax", age: 87},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)

}
