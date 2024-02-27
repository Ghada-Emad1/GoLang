# Defer-Panic-Recover

- Defer : defer statement pushes a function onto a list ,the list of saved calls is executed after the surrounding function returns
- defer will execute after the main function  executes

```go
    fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
    // output:start-end-middle
``` 
- defer work lifo(last in first out)
```go
    fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
    //program will store defer statments  in stack [second,third] so after print first ,third will print and then second
```


# panic :
- program can't continue to execute because of a certian statement that gets thrown
- panic is built-in function that stops the ordinary flow of control and begins panicking
- the program ends after the panic statement is run.

```go
     fmt.Println("start")
	 panic("something wrong Happend")
	 fmt.Println("end")
     // program will print start ,and then it will start panic and finish program,end statemnt will not execute
```

# recover
- recover function is used to recover function from panic. Recover is only useful inside deferred functions. In a normal execution, a call to recover will return nil and have no other effect. A recover function should always be called inside a defer function because the deferred function does not stop its execution if the program panics, so the recover function stops the panicking.
```go
func main(){
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
//if we remove defer statment the program won't be able to recover , and reaches the top of the goroutine's call stack, terminating the program
```

# interfaces
- is kind of like a definition ,that defines and describes methods that some other type must have

```go
     type stringer interface{
		String() string
	 }
```
- we say that something satisifies this interface if it has a method with the exact signature String() string
```go
    type Book struct{
		title string
		Author string
	}
	func(b Book) String() string{
		return fmt.Println(b.title)
	}
```
- It's not really important what this Book type is or does. The only thing that matters is that is has a method called String() which returns a string value.
- Wherever you see declaration in Go (such as a variable, function parameter or struct field) which has an interface type, you can use an object of any type so long as it satisfies the interface.


# Channels :
- channels are designed to synchronize data transmission between muliple go routines
 ## There is Two type of Channels:
- buffer channel and unbuffer channel 
### unbuffer Channel:
- doesn't have any capacity to store values
- when a sender sends a value on unbuffered channel ,it blocks until a receiver  receives that value
- unbuffered channels are typically used for synchronized communication between goroutines
```go
     ch:=make(chan int)

	 //this attemp to send value 12 into the channel,however since the channel is unbuffered and there is no gorountines to receive the value ,it will cause deadlock 
	ch<-12

	//the deadlock accour before meet this line
	fmt.Println(<-ch)
    
``` 
so how can we fix it ??
- you should either create buffered channel or create a goroutine to send the data to the channel while having the main goroutine ready to receive
```go
    ch:=make(chan int,1)
	ch<-12
	fmt.Println(<-ch)
```


## what is buffered channel
- have defined capacity to store values
- when a sender sends a value on a buffered channel it blocks only when channel is full
- when a reciver tries to recevie from channel ,it blocks only if the channel is empty
- buffered channela are useful when you want to decouple the timing of sender and receiver , allowing them to operate more independently
- they can provide a degree of asynchrony, as they sender can continue sending even if the receiver is not immediately ready
```go
    
	//buffered channel with capacity 2
	ch1:=make(chan int,2)

	//sender can send without blocking
	ch1<-12
	ch1<-23
	res1:=<-ch1
	res2:=<-ch1
	fmt.Println(res1,res2)
```
##  Channel Directions
When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

  <-chan datatype — A channel that only receives data

chan<- datatype — A channel that only sends data
```go
    
var waitgroup sync.WaitGroup

func main() {
	//make channel -unbuffered channel
	ch := make(chan int)
	//unbuffered channel work inside goroutines
	waitgroup.Add(2)
	go func(ch <-chan int) {
		//receive value from channel
		i := <-ch

		fmt.Println(i)
		waitgroup.Done()

	}(ch)
	
	go func(ch chan<- int) {
		//send value to channel
		ch <- 42

		waitgroup.Done()
	}(ch) //Note:declare that to call function

	waitgroup.Wait()

}

```
### Another example

```go
//function to send only channel
func pings(pings chan<-string,msg string){
	//msg will send to pings channel
	pings<-msg
}
//function to receive only channel and send only channel
func pongs(pings <-chan string,pongs chan<- string){
	//pings will recive value from msg
	msg:=<-pings

	//msg will send to channel pongs
	pongs<-msg
}
  
func main(){
	 ping:=make(chan string,1)
     pong:=make(chan string,1)
     pings(ping,"Passed Message")
     pongs(ping,pong)
     fmt.Println(<-pong)
}
```

### when you want to pass multi value into channel to handle that is by using some kind of a looping construct
```go
var waitgroup sync.WaitGroup
func main(){
	 ch := make(chan int, 40)
	waitgroup.Add(2)
	go func(ch <-chan int) {
		for i := range ch {

			fmt.Println(i)

		}

		waitgroup.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 24
		ch <- 12
		//you must close the channel[tell the channel that we're done working with you] or it will happened deadlock
		close(ch)
		//you are not allowed to pass anything after closing channel
		waitgroup.Done()
	}(ch)
	waitgroup.Wait()

}
  
```
## Select Statment
- allows goroutines to monitor several channels at once
 
  - blocks if all channels block
  - if multiple channels receive balue simultaneously,behaviour is undefined
```go
func main(){
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

```


