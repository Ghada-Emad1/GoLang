package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Switch Case
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not one of those")
	}
	fmt.Print("Go runs On ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Os X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)

	}

	switch 3 {
	case 1, 5, 3: // if value 1 or 5 or 3
		fmt.Println("one,five and three")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one , five and ten")
	case 4, 6, 8:
		fmt.Println("four,six and eight")
	}

	//in switch without fallthrough the program exits after first match "case"
	//with fallthrough  transfer control to the next case clause, even if the expression associated with that case evaluates to false
	i := 45
	switch {
	//it will execute
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
		//even i not greater 100 it will execute because of fallthrough
	case i > 100:
		fmt.Println("i is less than 100")
	}

	num := 2
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}

	num2 := 10
	switch {
	case num2 <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough
	case num2 > 20:
		fmt.Println("less than or equal to 20")
	}
    //fallthrough can't be in the final case

	

}
