package main

import (
	"fmt"

	"github.com/VAtulKumar1/GoByExamples/chapters"
)

func main() {

	// fmt.Println("Hello world")
	// fmt.Println("atul" + "kumar")
	// fmt.Println("1+1=", 1+1)
	// fmt.Println("7.0/3.0=", 7.0/3.0)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!false)

	// chapters.Variables()
	// chapters.Constants()
	// chapters.For()
	// chapters.IFElse()
	// chapters.Switch()
	// chapters.Slices()
	// chapters.Maps()
	res := chapters.Plus(2, 3)
	fmt.Println("res:", res)
	res2 := chapters.PlusPlus(1, 2, 3)
	fmt.Println("res2:", res2)
	a, b := chapters.Vals()
	fmt.Println("multi:", a, b)
	_, c := chapters.Vals()
	fmt.Println("subset:", c)
	chapters.Sum(1, 2, 3, 4, 5, 56)
}
