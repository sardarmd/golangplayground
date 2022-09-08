package api

import "testing"

func TestSortService(t *testing.T) {
	elements := []int{9, 8, 6, 5, 2}

	Sort(elements)

	if elements[0] != 2 {
		t.Error("First element should not be 9")
	}
	if elements[len(elements)-1] != 9 {
		t.Error(("last element should not be 2"))
	}
}
