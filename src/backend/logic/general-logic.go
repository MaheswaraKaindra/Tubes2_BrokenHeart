package logic

import (
	// Sementara kosong.
	"fmt"
)


// Setelah read sudah tidak digunakan.
type Element struct {
	Name	string     `json:"element"`
	Components	[][]string `json:"components"`
}

type ComponentKey struct {
	Component1 string
	Component2 string
}

type ElementContainer struct {
	Container map[string][]ComponentKey
	IsVisited map[string]bool
}

type TreeNode struct {
	Name     string
	Left     *TreeNode
	Right    *TreeNode
}


// Tidak tahu akan digunakan atau tidak, tapi aman jika ada.
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
