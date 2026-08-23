package main

import "fmt"

func main() {
	mapa := make(map[string]int);
	mapa["a"] = 1
	mapa["b"] = 2

	fmt.Println(mapa)
	fmt.Println(mapa["a"])

	
}

