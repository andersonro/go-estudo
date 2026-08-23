package main

var arr [3]int;

func main() {
	arr[0] = 8
	arr[1] = 2
	arr[2] = 10

	var total int = 0;

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	println("TOTAL: ", total);
	println("MEDIA: ", total / len(arr))

}
