package main

import (
	"fmt"
	"unsafe"
)

//import outcaller "utils/outcaller"


type Student struct {
     num int
     num2 int32;
     ok1 bool
	 num3 int32;
}


func main() {

	fmt.Println(unsafe.Sizeof(Student{}))

}