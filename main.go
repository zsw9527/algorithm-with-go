package main

import (
	"algorithm-with-go/algorithm"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	res := algorithm.Permute(arr)
	fmt.Println(res)
}
