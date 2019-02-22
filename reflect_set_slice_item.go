package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{0,0}
	slice := reflect.ValueOf(s)
	fmt.Println(slice)
	slice.Index(0).Set(reflect.ValueOf(1))
	slice.Index(1).Set(reflect.ValueOf(2))
	fmt.Println(s)
}
