package main

import (
	"fmt"
	"os"
)

func main() {
	s := "123"
	cwd, _ := os.Getwd()
	fmt.Printf("%s", cwd)
	func2(&s)
	/*if (err!=nil) {
		fmt.Println("Error")
		panic(nil)
	}*/
	fmt.Printf("%s", s)
}
func func2(s *string) {
	*s += ". its line!"
}
