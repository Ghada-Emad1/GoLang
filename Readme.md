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
//if we remove defer statment the progrma won't be able to recover , and reaches the top of the goroutine’s call stack, terminating the program
```