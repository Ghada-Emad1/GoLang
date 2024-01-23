package main

import (
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
	a:=10;
	b:=3;
	fmt.Println(a&b);
	fmt.Println(a|b);
	fmt.Println(a^b);
	fmt.Println(a&^b);

	//bit shift
	a1:=8;
	//when we shift to left we multiply
	fmt.Println(a1<<3); //shift to left = 2^3 * 2^3 =2^6
	fmt.Println(a1<<2); // = 2^3 *2^2=2^5
	fmt.Println(a1<<1); // =2^3* 2^1=2^4

	// when we shift to right we divivde
	fmt.Println(a1>>3); // shift right = 2^3/2^3=2^0 =1
	fmt.Println(a1>>2); // =2^3/2^2 = 2^1 =2
	fmt.Println(a1>>1) // =2^3 /2^1 = 2^2 =4

	//Floating number : float is 32 or 64 bit version
	n1:=0.4
	n1=13.1e72
	n1=2.1E14
	fmt.Printf("%v,%T\n",n1,n1);

	v:=10.2
	m1:=3.7
	fmt.Println(v+m1)
	fmt.Println(v-m1)
	//fmt.Println(v%m1) -->error we can't take remainder of two floating numbers

	//Complex Numbers : is 64 or 128 bit version
	var n2 complex64=1+2i 
	fmt.Printf("%v,%T\n",n2,n2);
	//To separate real and imaginary numbers we can use built-in function (real-imag)
	fmt.Printf("%v,%T\n",real(n2),real(n2));
	fmt.Printf("%v,%T\n",imag(n2),imag(n2))

	//another way to decalre complex number
	var complx complex128=complex(5,12);
	fmt.Printf("%v,%T\n",complx,complx);

	//String  : represent Utf8 -it is immutable
	s:="This is a string variable"
	
	fmt.Printf("%v,%T\n",s,s)
	fmt.Printf("%v,%T\n",s[1],s[1]) //string in go alises for bytes
	fmt.Printf("%v,%T\n",string(s[1]),s[1]) //type - uint8
	//concating
	s2:=" another string"
	fmt.Println(s+s2);
   
	b1:=[]byte(s) //convert string into a collection of byte
	fmt.Printf("%v,%T\n",b1,b1);
	

	//Rune : represent  utf32 character- Rune is an integer 32
	r:='r'
	fmt.Printf("%v,%T\n",r,r)
	fmt.Printf("%v,%T\n",string(r),r)
}