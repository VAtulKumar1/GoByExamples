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
	//

	nextInt := chapters.IntSeq()

	fmt.Println("c1:", nextInt())
	fmt.Println("c2:", nextInt())

	newInts := chapters.IntSeq()
	fmt.Println("c3:", newInts())
}
