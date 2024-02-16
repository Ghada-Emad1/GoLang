package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0
var m sync.RWMutex

func main() {
	// go HelloWorld();
	// // time.Sleep(1*time.Second) // if i remove this, GoodBye will only executes and exit before helloword execute
	// //GoodBye()
	// go GoodBye()
	// time.Sleep(1*time.Second) //main function will blocked for 1 seconds and all the goroutines would executed successfully

	//waitGroup used to wait for multiple goroutines to finish
	//waitGrooup blocks the execution of a function until its internal counter becomes 0
	// var wg sync.WaitGroup
	// wg.Add(2) //indicates the number of goroutines to wait ,there is two function that will be go routines so the internal counter will be 2
	// go helloWorld(&wg)
	// go goodbye(&wg)
	// wg.Wait() //blocks internal execution of code until the internal counter become 0

	// var msg = "Hello"
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }(msg)
	// msg = "GoodBye"
	// time.Sleep(1 * time.Millisecond)

	//waitGroup deign to synchronous multiple go routines together

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go helloWorld(&wg)
	// 	go Increment(&wg)
	// }
	// wg.Wait()

	//mutex:is basically a lock that the application is going to honor
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go helloWorld(&wg)
		m.RUnlock()
		
		go Increment(&wg)
	}
	wg.Wait()

}

func HelloWorld() {
	fmt.Println("Hello World!")
}
func GoodBye() {
	fmt.Println("GoodBye")
}
func helloWorld(wg *sync.WaitGroup) {
    m.RLock()
	fmt.Printf("Hello World%v\n", counter)
	//defer wg.Done() //reduce internal counter by 1
	m.RUnlock()
	wg.Done()
}
func goodbye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good bye!")
}

func Increment(wg *sync.WaitGroup) {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
