package api

import (
	"fmt"

	"ut.practice/golang-testing/api/utils"
)

func Sort(elements []int) []int {
	fmt.Println(elements)
	utils.BubbleSort(elements)
	return elements
}
