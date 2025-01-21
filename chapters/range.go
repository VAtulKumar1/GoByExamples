package chapters

import "fmt"

func Range() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[int]int{1: 10, 2: 20}
	for k, v := range kvs {
		fmt.Printf("%d->%d\n", k, v)
	}

	for i, c := range "Atul" {
		fmt.Println(i, c)
	}
}
