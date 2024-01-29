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

```go
     fmt.Println("start")
	 panic("something wrong Happend")
	 fmt.Println("end")
     // program will print start ,and then it will start panic and finish program
```
