package main

import (
	"fmt"
)

func main() {

	s := "Hello World"

	bs := []byte(s)

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Printf("%d  =  %#x\n", i, v)
	}
}
