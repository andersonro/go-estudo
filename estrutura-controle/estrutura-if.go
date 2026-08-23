package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("O numero %d é par.", i)
			fmt.Println()
		} else {
			fmt.Printf("O numero %d é impar.", i)
			fmt.Println()
		}
	}
}