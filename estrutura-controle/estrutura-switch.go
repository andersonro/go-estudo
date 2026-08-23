package main

import "fmt"
func main() {
	a := 10

	for i := 0; i < a; i++ {
		n := i%2
		switch n {
			case 0: fmt.Println("Valor", i," é par")
			default: fmt.Println("Valor", i," é impar")
		}
	}
}

