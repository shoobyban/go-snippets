package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := reflect.MakeSlice(reflect.TypeOf([]int{}), 2, 2)
	fmt.Println(slice)
	slice.Index(0).Set(reflect.ValueOf(1))
	slice.Index(1).Set(reflect.ValueOf(2))
	fmt.Println(slice)
}
