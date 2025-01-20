package chapters

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4}
	fmt.Println("dec1:", b)

	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("dc1:", c)

	d := [...]int{1, 3: 300, 500}
	fmt.Println("idx:", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("twoD:", twoD)

}
