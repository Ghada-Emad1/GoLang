package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {
	//make channel -unbuffered channel
	// 	ch := make(chan int)
	// 	//unbuffered channel work inside goroutines
	// 	waitgroup.Add(2)
	// 	go func(ch <-chan int) {
	// 		//receive value from channel
	// 		i := <-ch

	// 		fmt.Println(i)
	// 		waitgroup.Done()

	// 	}(ch)
	// 	//
	// 	go func(ch chan<- int) {
	// 		//send value to channel
	// 		ch <- 42

	// 		waitgroup.Done()
	// 	}(ch) //Note:declare that to call function

	// 	waitgroup.Wait()

	// 	ch2:=make(chan int)
	// 	for j:=0;j<5;j++{
	// 		waitgroup.Add(2)
	// 		go func(){
	// 			i:=<-ch2
	// 			fmt.Println(i)
	// 			waitgroup.Done()
	// 		}()
	// 		go func(){
	// 			ch2<-123
	// 			waitgroup.Done()
	// 		}()
	// 		waitgroup.Wait()
	// 	}

	//    ping:=make(chan string,1)
	//    pong:=make(chan string,1)
	//    pings(ping,"Passed Message")
	//    pongs(ping,pong)
	//    fmt.Println(<-pong)

	// ch := make(chan int, 40)
	// waitgroup.Add(2)
	// go func(ch <-chan int) {
	// 	for i := range ch {

	// 		fmt.Println(i)

	// 	}

	// 	waitgroup.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 24
	// 	ch <- 12
	// 	close(ch)
	// 	waitgroup.Done()
	// }(ch)
	// waitgroup.Wait()

	c1 := make(chan string)
	c2 := make(chan string)
	waitgroup.Add(2)
	go func() {
		// time.Sleep(1*time.Second)
		c1 <- "one"
		waitgroup.Done()
	}()
	go func() {
		// time.Sleep(2*time.Second)
		c2 <- "two"
		waitgroup.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	waitgroup.Wait()
}

// function to send only channel
func pings(pings chan<- string, msg string) {
	//msg will send to pings channel
	pings <- msg
}

// function to receive only channel and send only channel
func pongs(pings <-chan string, pongs chan<- string) {
	//pings will recive value from msg
	msg := <-pings

	//msg will send to channel pongs
	pongs <- msg
}
