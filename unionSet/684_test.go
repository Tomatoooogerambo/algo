package leet

import (
	"fmt"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	// edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	res := findRedundantConnection(edges)
	fmt.Println(res)
}

func findRedundantConnection(edges [][]int) []int {
	unionSet := make([]int, len(edges)+1)

	for i := 0; i < len(edges)+1; i++ {
		unionSet[i] = i
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if find(a, &unionSet) != find(b, &unionSet) {
			conn(a, b, &unionSet)
		} else {
			return edge
		}
	}

	return nil
}

func find(a int, unionSet *[]int) int {
	if (*unionSet)[a] == a {
		return a
	}

	(*unionSet)[a] = find((*unionSet)[a], unionSet)

	return (*unionSet)[a]
}

func conn(a, b int, unionSet *[]int) {
	(*unionSet)[find(a, unionSet)] = find(b, unionSet)
}
