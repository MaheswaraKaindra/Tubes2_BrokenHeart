package logic
import (
	// Sementara kosong.
	// "fmt"
)

// Setelah read sudah tidak digunakan.
type Element struct {
	Name		string     `json:"element"`
	Components	[][]string `json:"components"`
	Tier		int
	Image 		string
}

type ComponentKey struct {
	Component1 string
	Component2 string
}

type ElementContainer struct {
	Container map[string][]ComponentKey
	ElementTier map[string]int
	IsVisited map[string]bool
	ElementImage map[string]string
}

type TreeNode struct {
	Name     string
	Image	 string
	Left     *TreeNode
	Right    *TreeNode
}

type SearchState struct {
	Node   *TreeNode
	Target string
}

type Result struct {
	Node *TreeNode
	VisitedCount int
}