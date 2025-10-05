package main

import "fmt"

func main() {
	arr := [...]int{3, 3, 6}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr[i])
	}
}
