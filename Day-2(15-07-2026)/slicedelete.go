package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	index := 2 // Remove value 3

	nums = append(nums[:index], nums[index+1:]...)//[:2],[3:]...

	fmt.Println(nums)
}