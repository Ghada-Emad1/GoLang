package main

import "fmt"
type myStruct struct{
	foo int
}

func main() {
	a := 42
	b := a // is a copy of a if i change a later b will not change
	fmt.Println(a)
	fmt.Println(b)
	a = 45
	fmt.Println(a)
	fmt.Println(b)

	var a1 int = 54
	var b1 *int = &a1 //b1 is point to the address of a1
	fmt.Println(&a1) //the address of a1
	fmt.Println(a1) //value of a1
	fmt.Println(*b1) //value of b1
	fmt.Println(b1) //the address of a1

	//in c,c++ we can do arthmtic operation with pointer but not with go ,we can't do that
	//c:=&a1+1 we can't do that

	a2:=[3]int{1,2,3}
	b2:=&a2[1]
	c2:=&a2[2]
	fmt.Printf("%v %p %p\n",a2,b2,c2)

	var ms* myStruct
	ms=&myStruct{foo: 32}//ms is holding the address of an object that has a field with value 42 in it
	fmt.Println(ms)

	//another declaration mys is point to myStruct
	mys:=new(myStruct)
	mys.foo=97
	fmt.Println(mys.foo)

	var mystrkt *myStruct //point to myStruct with nil value
	
	mystrkt.foo=90
	fmt.Println((*mystrkt).foo)

	//all assignment perations in Go are copy operations
	//instead of slices and maps containe interna pointers , do copies point to same underlying data
	
}
