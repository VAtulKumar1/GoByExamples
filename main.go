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

	//

	// res := chapters.Fact(10)
	// fmt.Println("fact:", res)

	// chapters.Range()

	i := 1
	fmt.Println("iVal", i)

	chapters.ZeroVal(i)
	fmt.Println("pass by value:", i)

	chapters.ZeroPtr(&i)
	fmt.Println("pass by reference:", i)

}
