package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"backend/logic" // import the logic package
)

// Your existing struct types
// type ElementPair struct {
// 	Component1 string `json:"component1"`
// 	Component2 string `json:"component2"`
// }

// type ElementContainer struct {
// 	Container     map[string][]ElementPair `json:"container"`
// 	ElementTier   map[string]int           `json:"element_tier"`
// }

// type TreeNode struct {
// 	Name  string    `json:"name"`
// 	Left  *TreeNode `json:"left,omitempty"`
// 	Right *TreeNode `json:"right,omitempty"`
// }

// Helper: minimal mockup
func isBaseElement(name string) bool {
	// Adjust to your logic
	base := []string{"fire", "water", "earth", "air"}
	name = strings.ToLower(name)
	for _, b := range base {
		if b == name {
			return true
		}
	}
	return false
}

// Helper: find the index of minimum value in map
func minKey(m map[int]int) int {
	min := int(^uint(0) >> 1) // Max int
	minKey := -1
	for k, v := range m {
		if v < min {
			min = v
			minKey = k
		}
	}
	return minKey
}

// Import your BFS algorithm (already defined)

// Request struct
type BFSRequest struct {
	Target    string           `json:"target"`
	Container logic.ElementContainer `json:"container"`
	Index     int              `json:"index"`
}

// Handler function
func BFSHandler(w http.ResponseWriter, r *http.Request) {
    // CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight request
    if r.Method == http.MethodOptions {
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var req BFSRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
        return
    }

    root := logic.BreadthFirstSearch(req.Target, &req.Container, req.Index)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(root)
}

// Server entry point
func main() {
	http.HandleFunc("/bfs", BFSHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
