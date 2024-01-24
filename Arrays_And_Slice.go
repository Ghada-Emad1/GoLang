package main

import (
	"fmt"
)

func main() {
	//Create Array
	//Arrays have fixed size that have to be know at compile time
	grades := [3]int{1, 2, 3}
	fmt.Printf("Grades%v \n", grades)

	//Another declare of an array
	grades1 := [...]int{4, 5, 6}
	fmt.Printf("Grades %v\n", grades1)

	//Empty Array
	var students [3]string
	fmt.Printf("Students %v\n", students)
	students[1] = "ghada"
	fmt.Println(students)
	fmt.Printf("Length of Array %v\n", len(students))
	fmt.Println(cap(students)) //capacity of array

	//2D-Array
	var identityMatrix [3][3]int = [3][3]int{[3]int{4, 3, 2}, [3]int{0, 9, 8}, [3]int{7, 5, 6}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][4] int
	identityMatrix2[0]=[4] int{44,55,66,1}
	identityMatrix2[1]=[4] int{9,12,4,13}
	identityMatrix2[2]=[4] int{9,8,7,6}
	fmt.Println(identityMatrix2)
	fmt.Println(identityMatrix2[0][1])

	a:=[...] int{1,2,3}
	b:=a    //b copy the elements of a in it
	//b:=&a // but when we using pointer --> b points to the data that a has , and the original will affect of this change
	b[1]=5  // when we assign new value to copied array the original don't affect of the change
	fmt.Printf("The Values of a %v\n",a) //[1,2,3]
	fmt.Printf("The Values of b %v\n",b) //[1,5,3]


	//Slice :don't have fixed size
	a1:=[]int{12,13,14}
	fmt.Println(a1)
	fmt.Println(len(a1))
	fmt.Println(cap(a1))
	//slice are naturally reference type
	b1:=a1 // with slice --> b1 pointto thesame data that a1 has
	b1[2]=90
	fmt.Printf("Value of a1 %v\n",a1)
	fmt.Printf("Value of b1 %v\n",b1)


	A:=[]int{1,2,3,4,5,6,7,8,9,10}
	B:=A[:] //slice of all elements
	C:=A[2:] // slice from 3th element to the end of slice
	D:=A[:6]  //slice first 6 element
	E:=A[3:6] //slice from 4th element to 6th element
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

	//Another way to create slice
	a3:=make([]int,3)
	a4:=make([]int,3,100) //3 represent size, 100 represent capacity
	fmt.Println(a3)
	fmt.Println(a4)
  
	a5:=[]int{}
	fmt.Println(a5)
	a5=append(a5, 2)
	fmt.Println(a5)
	fmt.Printf("Lenght is %v\n",len(a5))
	fmt.Printf("Capacity is %v\n",cap(a5))
    a5=append(a5, 9,4,5)
	fmt.Println(a5)
    fmt.Println(cap(a5))
	// a5=append(a5, []int{43,11,45}) //Error , the second arg should have the type of slice which is int
	a5=append(a5, []int{12,23,34}...) // it will spread the slice into individual argument
	fmt.Println(a5)
	fmt.Printf("Capcity %v\n",cap(a5))

	a6:=[]int{3,4,5,6,7}
	b6:=a6[1:] //remove first element
	fmt.Println(a6)
	fmt.Println(b6) 
	b7:=a6[:len(a6)-1]  //remove last element
	fmt.Println(b7)

	b8:=append(a6[:2],a6[3:]...) //remove middle element of slice a6
	fmt.Println(b8)
}
