package main

import (
	// Sementara kosong.
	// "fmt"
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
	Name    string
	Recipes []*RecipeNode
}

type RecipeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

type SingularTreeNode struct {
	Name    string
	Left  	*SingularTreeNode
	Right 	*SingularTreeNode
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

func getSingularTree(node *TreeNode, indexMap map[string]int) *SingularTreeNode {
	if node == nil {
		return nil
	}

	if len(node.Recipes) == 0 {
		return &SingularTreeNode{
			Name:  node.Name,
			Left:  nil,
			Right: nil,
		}
	}

	idx := indexMap[node.Name]
	if idx >= len(node.Recipes) {
		idx = 0
	}

	selected := node.Recipes[idx]

	return &SingularTreeNode{
		Name:  node.Name,
		Left:  getSingularTree(selected.Left, indexMap),
		Right: getSingularTree(selected.Right, indexMap),
	}
}
