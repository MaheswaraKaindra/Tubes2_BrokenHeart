package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
)

import "github.com/MaheswaraKaindra/Tubes2_BrokenHeart/src/backend/logic"

type SearchRequest struct {
	Target string `json:"target"`
	Index  int    `json:"index"`
}

// CORS Middleware
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Utility: Load container from JSON files
func loadElementContainerFromFiles() (*logic.ElementContainer, error) {
	recipesPath := filepath.Join(".", "data", "recipes.json")
	tiersPath := filepath.Join(".", "data", "tiers.json")
	imagesPath := filepath.Join(".", "data", "images.json")
	
	elements, err := logic.ReadJSON(recipesPath, tiersPath, imagesPath)
	if err != nil {
		return nil, err
	}
	container := logic.BuildElementContainer(elements)
	return &container, nil
}

func bfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	root := logic.BreadthFirstSearch(req.Target, container, req.Index)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

func dfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	root := logic.FirstDepthFirstSearch(req.Target, container, req.Index)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

func shortestbfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	root := logic.ShortestBreadthFirstSearch(req.Target, container)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

func shortestdfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	root := logic.ShortestDepthFirstSearch(req.Target, container)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

func multiplebfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	loopCount := logic.GetLength(container, req.Target)
	recipes := logic.GetRecipe(container, req.Target, loopCount)
	var trees []interface{}

	for i := 0; i < loopCount; i++ {
		tree := logic.BreadthFirstSearch(req.Target, container, i)
		trees = append(trees, tree)
	}

	response := map[string]interface{}{
		"trees":   trees,
		"recipes": recipes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func multipledfsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	loopCount := logic.GetLength(container, req.Target)
	recipes := logic.GetRecipe(container, req.Target, loopCount)
	var trees []interface{}

	for i := 0; i < loopCount; i++ {
		tree := logic.FirstDepthFirstSearch(req.Target, container, i)
		trees = append(trees, tree)
	}

	response := map[string]interface{}{
		"trees":   trees,
		"recipes": recipes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/bfs", shortestbfsHandler)
	http.HandleFunc("/api/dfs", shortestdfsHandler)
	http.HandleFunc("/api/bfsmultiple", multiplebfsHandler)
	http.HandleFunc("/api/dfsmultiple", multipledfsHandler)

	fmt.Println("Server running on :8080")
	logErr := http.ListenAndServe(":8080", nil)
	if logErr != nil {
		fmt.Println("Server error:", logErr)
	}
}