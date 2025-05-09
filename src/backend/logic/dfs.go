package main

import (
	// Sementara kosong.
	// "fmt"
)

func depthFirstSearch(target string, container *ElementContainer) *TreeNode {
	if _, exists := container.Container[target]; !exists {
		return &TreeNode{
			Name:    target,
			Recipes: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode{
			Name:    target,
			Recipes: nil,
		}
	}

	if container.IsVisited[target] {
		return &TreeNode{
			Name:    target,
			Recipes: nil,
		}
	}
	
	container.IsVisited[target] = true
	var recipes []*RecipeNode

	for _, raw := range container.Container[target] {
		left := depthFirstSearch(raw.Component1, container)
		right := depthFirstSearch(raw.Component2, container)
		
		if left != nil && right != nil {
			recipes = append(recipes, &RecipeNode{
				Left:  left,
				Right: right,
			})
		}
	}
	
	return &TreeNode{
		Name:    target,
		Recipes: recipes,
	}
	container.IsVisited[target] = false

	return &TreeNode{
		Name:    target,
		Recipes: nil,
	}
}