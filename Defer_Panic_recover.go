package main

import (
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer src.Close()
	return io.Copy(dst, src)
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func main() {
	//defer statement pushes a function onto a list ,the list of saved calls is executed after the surrounding function returns
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	//defer will execute after the main function  executes
	//defer work lifo(last in first out)
	fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	//first ,third,second

	//1.deferred function's arguments are evaluted when the defer statment is evaluted

	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}

	res, err := http.Get("https://www.google.com/robots.txt")

	if err != nil {
		return
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Printf("%s", robots)
	a := "start"
	defer fmt.Println(a) //the value will be start because it already stored to start
	a = "end"

	// fmt.Println("start")
	// panic("something wrong Happend")
	// fmt.Println("end")

	// Recover is a built-in function that regains control of a panicking goroutine.
	// Recover is only useful inside deferred functions.
	// During normal execution, a call to recover will return nil and have no other effect.
	//If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
	f()
	fmt.Println("Returned Normaly from f")

}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normaly from g.")
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprint("%v ---", i))

	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
