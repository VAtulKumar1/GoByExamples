package chapters

import (
	"fmt"
	"math/rand/v2"
)

func RandomNumbers() {
	p := fmt.Println

	p(rand.IntN(100), ",")
	p(rand.IntN(100))
	p()

	p(rand.Float64())

	p(rand.Float64()*5+5, ",")
	p((rand.Float64() * 5) + 5)
	p()

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	p(r2.IntN(100), ",")
	p(r2.IntN(100))
	p()

}
