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

func isBaseElement(name string) bool {
	switch name {
	case "Air", "Water", "Fire", "Earth", "Time":
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
