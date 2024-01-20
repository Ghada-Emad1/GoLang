package main
import(
	"fmt"
	"strconv"
) 
 var i int=23;
func main(){
	// var num int =42;
	// Num2:=2.0
	// fmt.Printf("%v,%T",Num2,Num2)
	// fmt.Printf("%v,%T",num,num)
	var theHTTP string="https://google.com";
	fmt.Println(theHTTP);
	var i int=42;
	fmt.Printf("%v,%T\n",i,i);

	var j float32
	j=float32(i)
	fmt.Printf("%v,%T\n",j,j)

	var II int=54
	fmt.Printf("%v,%T\n",II,II);

	var jj string
	jj=strconv.Itoa(II)
	fmt.Printf("%v,%T\n",jj,jj)
	
}