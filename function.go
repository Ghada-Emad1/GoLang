package main

import (
	"fmt"
)

type values struct {
	Numbers []int
	id      int
}
type greeter struct {
	greeting string
	name     string
}



func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The index is : ", idx)
}
func sayHello(name, hello string) {
	fmt.Println(name, hello)
}
func sayGreeting(greetin, name *string) {
	fmt.Println(*greetin, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// passing varidaic parameters to send list of same types
// it must be last parameters
// received as a slice
// i can't use return in this function cause i don't specify the type of return function
func sum(values ...int) {
	fmt.Println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	fmt.Println("The sum is:", res)
}

//return statement - here i specify the type of return function by putting the type after declaring the parameters of the function

func mult(values ...int) *int {

	res := 1
	for _, v := range values {
		res *= v
	}

	return &res
}

// i can also write the value i want to return before the type of it in declaration of the function
func sub(values ...int) (res int) {
	for _, v := range values {
		res -= v
	}
	return

}

// multi return value      //typeof a/b ,type of error i passed error the nil value
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cant't divide by 0")
	}
	return a / b, nil // will return two values
}

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	//if the parameters of same type instead of delclare func sayGreeting(greeting string,name string) ,
	//we declare func sayGreeting(greeting,name string)
	sayHello("Ghada", "Hello")

	greeting := "Hello"
	name := "Stacy"
	sayGreeting(&greeting, &name)
	fmt.Println(name) //it will note change because it name in function is just copy
	//but if i passed as pointer to name , it will change the value
	//without pointer ,it will not change the value -[it will be just a copy]

	sum(1, 2, 3, 4, 5)
	m := mult(3, 4, 5, 6)

	fmt.Println("The res is:", *m)

	fmt.Println("The res of sub: ", sub(1, 2, 3, 4))

	d, err := divide(5.6, 6.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//anonymous function:it is just a function but without a name
	func() {
		fmt.Println("Hello Go")
	}()
	//we can assign variable to anonymous function
	f := func() {
		fmt.Println("declare anonymous function as a variable")
	}
	f()

	var divide1 func(float64, float64) (float64, error)
	divide1 = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("Can;t divide by zero")
		} else {
			return a / b, nil
		}

	}
	d1, err1 := divide1(5.6, 1.0)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d1)

	//method:is a function executes in a context of a type
	//receiver can be value or pointer:
	//value receiver gets copy of type
	//pointer receiver gets pointer to type

	g := greeter{
		greeting: "Hello",
		name:     "Ghada",
	}
	g.greet()
	fmt.Println("The New value", g.name)

	val := values{
		Numbers: []int{1, 2, 3, 4},
		id:      1,
	}
	val.PrintSum()

}

// if i pass greeter as pointer ,the value will change otherwise it won't change cause it is a copy
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Sare"
}
func (s *values) PrintSum() {
	sum := 0
	for _, num := range s.Numbers {
		sum += num
	}
	fmt.Printf("Sum of numbers: %d\n", sum)
}
