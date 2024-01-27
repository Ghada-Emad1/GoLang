package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = i*2 + 1
		}
	}
	// for i:=0,j:=0;i<5;i++,j++ that is not allowed to do
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {

	}
	//we can do that also
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 3 {
			break
		}
	}

	//Nested loop
	fmt.Println("Nest Loop ------")

Loop:
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 3 {
				//break the whole loop not only the inner loop
				break Loop

			}

		}

	}
	//To loop in range
	s:=[]int{1,2,3,4}
	for k,v:=range s{
		fmt.Println(k,v) // k is the index and v is value
		
	}
	names:=[]string{"omar","ahmed","ali"}
	//To print only index
	for k:=range names{
		
		fmt.Println(k)
	}
	//To print only values
	for _,v:=range names{
		fmt.Println(v)
	}

}
