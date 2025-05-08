package main

import (
	// Sementara kosong.
	// "fmt"
)

func depthFirstSearch(target string, container *ElementContainer) *TreeNode {
	if _, exists := container.Container[target]; !exists {
		return nil
	}
	if isBaseElement(target) {
		return &TreeNode{
			Name:  target,
			Left:  nil,
			Right: nil,
		}
	}

	if container.IsVisited[target] {
		return &TreeNode{
			Name:  target,
			Left:  nil,
			Right: nil,
		}
	}

	container.IsVisited[target] = true


	for _, pair := range container.Container[target] {
		left := depthFirstSearch(pair.Component1, container)
		right := depthFirstSearch(pair.Component2, container)

		if left != nil && right != nil {
			return &TreeNode{
				Name:  target,
				Left:  left,
				Right: right,
			}
		}
	}

	return nil
}



