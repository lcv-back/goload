package main

import (
	"fmt"
	"sort"
)

type Item struct {
	a int
	b []int
}

var dataArray = []Item{
	{a: 1, b: []int{1, 2}},
	{a: 1, b: []int{2, 3}},
	{a: 1, b: []int{2, 1}},
	{a: 2, b: []int{4, 5}},
}

func countIdenticalItems() map[string]int {
	counts := make(map[string]int)
	for _, item := range dataArray {
		// Sort the b slice to ensure consistent key generation
		sortedB := make([]int, len(item.b))
		copy(sortedB, item.b)
		sort.Ints(sortedB)

		key := fmt.Sprintf("{a: %d, b: %v}", item.a, sortedB)
		counts[key]++
	}
	return counts
}

func main() {
	fmt.Println("{")
	result := countIdenticalItems()
	for key, count := range result {
		fmt.Printf("  %s: %d,\n", key, count)
	}
	fmt.Println("}")
}
