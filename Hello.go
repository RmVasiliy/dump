package main

import "fmt"

func main() {
	a := 14545
	p := &a
	fmt.Println("%x - %d", p, *p)
}
