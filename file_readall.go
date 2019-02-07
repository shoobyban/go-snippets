package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: file_readall {filename}")
		os.Exit(0)
	}
	filename := os.Args[1]
	fh, err := os.Open(filename)
	defer fh.Close()
	if err != nil {
		log.Panicf("can't open file %s", filename)
	}
	filebytes, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Panicf("can't read file %s", filename)
	}
	fmt.Printf("File content as string: %s", string(filebytes))
}
