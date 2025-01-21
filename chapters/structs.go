package chapters

type Person struct {
	Name string
	Age  int
}

func Structs(name string) *Person {
	p := Person{Name: name}
	p.Age = 42
	return &p

}
