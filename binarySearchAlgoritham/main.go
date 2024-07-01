package main

import (
	"fmt"
	"sort"
)


func main() {

	slice := []int{2, 4, 5, 6, 7, 8, 3, 87, 65, 43, 56, 32}
	sort.Ints(slice)
	fmt.Println(slice)
	value := 5

	high := len(slice) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if value == slice[mid] {
			fmt.Println(mid)
			break
		} else if value > slice[mid] {
			low = mid + 1
		} else if value < slice[mid] {
			high = mid - 1
		}
	}
	fmt.Println("sorry")
}
