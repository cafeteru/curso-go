package main

import (
	"fmt"
	"reflect"
)

var global = "Hello world"

func main() {
	local := 2
	localFloat := 2.0
	fmt.Println(local)
	fmt.Println(global)
	fmt.Printf("The value of the \"local\" variable is %d and its type is %T\n", local, local)
	fmt.Printf("The value of the \"global\" variable is \"%s\" and its type is %T\n", global, global)
	s1 := fmt.Sprint("The value of the \"localFloat\" variable is ", localFloat, " and its type is ",
		reflect.TypeOf(localFloat))
	fmt.Println(s1)
}
