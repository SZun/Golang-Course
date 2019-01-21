package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		eoa := "even"
		if num%2 != 0 {
			eoa = "odd"
		}
		fmt.Println(num, " is ", eoa)
	}
}
