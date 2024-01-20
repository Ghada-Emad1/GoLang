package main
import(
	"fmt"
)
func main(){
	//Boolean Types
	var data bool=true
	fmt.Printf("%v,%T",data,data)

	n:=1==1
	m:=1==2
	fmt.Printf("%v,%T\n",n,n)
	fmt.Printf("%v,%T\n",m,m)

	//Every Variable you initalize in Go have Zero Value
	var i int 
	fmt.Printf("%v,%T\n",i,i)


	//Integer
	num:=24
	fmt.Printf("%v,%T\n",num,num)

	var num2 uint16=42
	fmt.Printf("%v,%T\n",num2,num2)

	//Can't add/sub/mul/div different type of Variables
	var num3 int=3
	var num4 int16=1234
	fmt.Println(num3+int(num4))
	
}