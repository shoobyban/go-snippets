package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"a":0,"b":0,"c":0}
	fmt.Println(m)
	v := reflect.ValueOf(m)
	v.SetMapIndex(reflect.ValueOf("b"),reflect.ValueOf(1))
	fmt.Println(m)
}
