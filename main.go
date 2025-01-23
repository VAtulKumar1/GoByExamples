package main

import "github.com/VAtulKumar1/GoByExamples/chapters"

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

	// i := 1
	// fmt.Println("iVal", i)

	// chapters.ZeroVal(i)
	// fmt.Println("pass by value:", i)

	// chapters.ZeroPtr(&i)
	// fmt.Println("pass by reference:", i)

	// chapters.Strings()
	// p := chapters.Person{Name: "bob", Age: 28}
	// fmt.Println(p.Name)

	// fmt.Println(chapters.Person{Name: "atul"})
	// fmt.Println(&chapters.Person{Name: "Atul"})
	// fmt.Println(chapters.Structs("Sharad"))
	// sp := &p
	// fmt.Println(sp.Name)

	// r := chapters.Rect{Width: 10, Height: 30}

	// fmt.Println("area:", r.Area())
	// fmt.Println("Perim:", r.Perim())

	// rp := &r
	// fmt.Println("area:", rp.Area())
	// fmt.Println("Perim:", rp.Perim())

	// r := chapters.Rectangle{Width: 3, Height: 4}
	// c := chapters.Circle{R: 5}

	// chapters.Measure(r)
	// chapters.Measure(c)

	// chapters.DetectCicle(r)
	// chapters.DetectCicle(c)

	// ns := chapters.Transition(chapters.StateIdle)
	// fmt.Println(ns)

	// ns2 := chapters.Transition(ns)
	// fmt.Print(ns2)

	// co := chapters.Container{
	// 	Base: chapters.Base{
	// 		Num: 1,
	// 	},

	// 	Str: "some time",
	// }

	// fmt.Printf("co={num: %v,str:%v}\n", co.Num, co.Str)
	// fmt.Println("also num:", co.Describe())

	// type describer interface {
	// 	Describe() string
	// }

	// var d describer = co
	// fmt.Println("describer:", d.Describe())

	// for _, i := range []int{7, 42} {
	// 	if r, e := chapters.F(i); e != nil {
	// 		fmt.Println("f failed:", e)
	// 	} else {
	// 		fmt.Println("f worked:", r)
	// 	}
	// }

	// for i := range 5 {
	// 	if err := chapters.MakeTea(i); err != nil {
	// 		if errors.Is(err, chapters.ErrOutOfTea) {
	// 			fmt.Println("we should by tea!")
	// 		} else if errors.Is(err, chapters.ErrPower) {
	// 			fmt.Println("now it is dark")
	// 		} else {
	// 			fmt.Printf("unknown error: %s\n", err)
	// 		}
	// 		continue
	// 	}

	// 	fmt.Println("tea is ready")
	// }

	// _, err := chapters.F1(42)
	// var ae *chapters.ArgError

	// if errors.As(err, &ae) {
	// 	fmt.Println(ae.Arg)
	// 	fmt.Println(ae.Message)
	// } else {
	// 	fmt.Println("err does'nt match argError")
	// }

	//

	// chapters.Channel()

	// done := make(chan bool, 1)

	// go chapters.Worker(done)

	// // <-done

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)

	// chapters.Ping(pings, "passed message")
	// chapters.Pong(pings, pongs)

	// fmt.Println(<-pongs)

	// chapters.Timeouts()
	// chapters.NonBlockingChannelOperations()

	chapters.ClosingChannel()
}
