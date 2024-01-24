package main

import (
	"fmt"
)

// Iota is a counter that we can use when we are creating what is called enumerated constants
// iota represents successive untyped integer constants
// it will start from zero
const (
	a3 = iota //0
	b3 = iota //1
	c3 = iota //2
)

// if we don't assign the value of const after the first one it will act that you assign it, so it will behave as a counter
// it will start from zero
const (
	a4 = iota
	b4
	c4
)

// if you want to ignore the first value =0
const (
	_ = iota
	a5
	b5
	c5
)

// Examples
const (
	errorSpecialist = iota + 6 //6
	catSpecialist              //7
	dogSpecialist              //8
	snakeSpecialist            //9
)
const (
	_  = iota
	KB = 1 << (10 * iota) // = 1<<(10*1)= 2^0  * 2^10 =2^10=1024
	GB                    // = 1<<(10*2) = 1<<(20)
	TB
	PB
)
const (
	isAdmin          = 1 << iota // =(1<<0)=2^0 * 2^0 = 2^0 =1
	isHeadquarters               // =(1<<1)=2^0 * 2^1 = 2^1=2
	canSeeFinancials             // =(1<<2)=2^0 * 2^2 =2^2= 4
	canSeeAfrica                 // =(1<<3)=2^0 * 2^3 = 2^3 =8
	canSeeAsia                   // (1<<4)
	canSeeEurope                 // (1<<5)
)

func main() {
	//Declare Constants
	const myConst int = 42
	fmt.Printf("%v,%T\n", myConst, myConst)

	const a int = 2
	const b string = "boo"
	const c bool = false

	//fmt.Println(a+b)// Error can't make operation with different types

	const a1 = 12
	var b1 int = 8
	fmt.Printf("%v,%T", a1+b1, a1+b1) // we can that because i don't specify the type of a1

	const a2 int16 = 23
	//var b2 int32;
	//can't add a2 and b2 ,as we know they are different types

	fmt.Println("Working with Iota-----------")
	fmt.Printf("%v,%T\n", a3, a3)
	fmt.Printf("%v,%T\n", b3, b3)
	fmt.Printf("%v,%T\n", c3, c3)

	fmt.Printf("%v,%T\n", a4, a4)
	fmt.Printf("%v,%T\n", b4, b4)
	fmt.Printf("%v,%T\n", c4, c4)

	fmt.Printf("%v,%T\n", a5, a5)
	fmt.Printf("%v,%T\n", b5, b5)
	fmt.Printf("%v,%T\n", c5, c5)

	var specialType int // it's value =0
	fmt.Printf("%v\n", specialType == catSpecialist)

	filesize := 40000000
	fmt.Printf("%.2fGB\n", filesize/GB)

	//iota don't work with main func it work with const

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope // 37
	fmt.Printf("%v\n", roles)                                  // represent roles as integer number
	fmt.Printf("%b\n", roles)                                  //represent roles as binary number

	fmt.Printf("isAdmin?%v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ?%v", isHeadquarters&roles == isHeadquarters)

}
