package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	//if won't work if there is no {} after it
	// --if statement
	if true {
		fmt.Println("Hello world!")
	}
	statePop := map[string]int{
		"calforina": 12345,
		"egypt":     4321495833,
		"taxes":     98531,
	}
	// if value of key egypt exist print the value,ok value is true or false
	if pop, ok := statePop["egypt"]; ok {
		fmt.Print(pop)
	}

	num := 500
	guess := 300
	if num > guess {
		fmt.Println("Too High!")
	}
	if num < guess {
		fmt.Println("Too Low")
	}
	if num == guess {
		fmt.Println("Equal To each Other")
	}

	//working with logical operator
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must between 1 and 100")
	}
	if guess >= 1 || guess <= 100 {
		if guess < num {
			fmt.Println("Too Low")
		}
	}
	if guess > 100 && ReturnTrue() {
		fmt.Println("The guess must be betweeen 1 and 100")
	}

	// --if-else statement
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must between 1 and 100")
	} else if guess > 100 {
		fmt.Println("the guess must less than 100")
	} else {
		if guess < num {
			fmt.Println("Too Low")
		}
	}
	mynum := 0.1
	if mynum == math.Pow(math.Sqrt(mynum), 2) {
		fmt.Println("These are the same")

	} else {
		fmt.Println("They are different")
	}

	

}

func ReturnTrue() bool {
	fmt.Println("Return true")
	return true
}
