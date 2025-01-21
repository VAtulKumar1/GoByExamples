package chapters

import "fmt"

type Base struct {
	Num int
}

func (b Base) Describe() string {
	return fmt.Sprintf("base with num=%v", b.Num)
}

type Container struct {
	Base
	Str string
}
