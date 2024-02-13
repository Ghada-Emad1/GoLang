package main

import (
	"bytes"
	"fmt"

	"io"
)

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface{
	Close() error
}
type WriterCloser interface{
	Writer
	Closer
}
type BufferedWriterCloser struct{
	buffer *bytes.Buffer
}
type Counter interface {
	Increment() int
}
type IntCounter int


func (bwc *BufferedWriterCloser) Write(data []byte) (int,error){
  n,err:=bwc.buffer.Write(data)
  if err!=nil{
	return 0,err
  }
  v:=make([]byte,8)
  for bwc.buffer.Len()>8{
	_,err=bwc.buffer.Read(v)
  }
  _,err=fmt.Println(string(v))
  if err!=nil{
	return 0,err
  }
  return n,nil
}
func(bwc*BufferedWriterCloser)Close() error{
	for bwc.buffer.Len()>0{
		data:=bwc.buffer.Next(8)
		_,err:=fmt.Println(string(data))
		if err!=nil{
			return err
		}
	}
	return nil
}
func NewBufferedWriterCloser()*BufferedWriterCloser{
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
func main() {
	// var w Writer = ConsoleWriter{}
	// w.Write([]byte("Hello World!"))

	// Mycount:=IntCounter(0)
	// var inc Counter=&Mycount
	// for i:=0;i<10;i++{
	// 	fmt.Println(inc.Increment())
	// }
    
	var wc WriterCloser=NewBufferedWriterCloser()
	wc.Write([]byte("Hello from my Terminal"))
	wc.Close()

	r,ok:=wc.(*BufferedWriterCloser)
	if ok{
		fmt.Println(r)
	}else{
		fmt.Println("Coveration is failed")
	}
 

	var myobj interface{}=NewBufferedWriterCloser()
	if wc1,ok1:=myobj.(WriterCloser);ok1{
		wc1.Write([]byte("Hello From My Termainal"))
		wc1.Close()
	}
	r1,ok1:=myobj.(io.Reader)
	if ok1{
		fmt.Println(r1)
	}else{
		fmt.Println("Conversion failed")
	}

    var i interface{}=true
	switch i.(type){
	case int:
		fmt.Println("i is integer")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i don't know what i is")
	}
}
