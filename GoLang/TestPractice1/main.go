package main

import (
	"fmt"

	api "ut.practice/golang-testing/api/services"
)

func main() {
	elements := []int{1, 2, 4, 6}
	fmt.Println(api.Sort(elements))
}
