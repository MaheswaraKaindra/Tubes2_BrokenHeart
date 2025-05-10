package main
import (
	// Sementara kosong.
	// "fmt"
)

// Setelah read sudah tidak digunakan.
type Element struct {
	Name		string     `json:"element"`
	Components	[][]string `json:"components"`
	Tier		int
}

type ComponentKey struct {
	Component1 string
	Component2 string
}

type ElementContainer struct {
	Container map[string][]ComponentKey
	ElementTier map[string]int
	IsVisited map[string]bool
}

type TreeNode struct {
	Name     string
	Left     *TreeNode
	Right    *TreeNode
}

type SearchState struct {
	Node   *TreeNode
	Target string
}

func isBaseElement(name string) bool {
	switch name {
	case "air", "water", "fire", "earth", "time":
		return true
	default:
		return false
	}
}

func normalizeKey(a, b string) ComponentKey {
    if a < b {
        return ComponentKey{a, b}
    }
    return ComponentKey{b, a}
}

func minKey(m map[int]int) int {
	minKey := -1
	minVal := int(^uint(0) >> 1) // nilai maksimum int

	for k, v := range m {
		if v < minVal {
			minVal = v
			minKey = k
		}
	}
	return minKey
}
