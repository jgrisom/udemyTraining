package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d, %b, %o, %x, %X, %c, %U, %q, \n", i, i, i, i, i, i, i, i)
		fmt.Println()
	}
	fmt.Printf("%d, %b, %o, %x, %X, %c, %U, \n", 42, 42, 42, 42, 42, 42, 42)

}
