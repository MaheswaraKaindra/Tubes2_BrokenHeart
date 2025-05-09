package main

import (
	// Sementara kosong.
	// "fmt"
)

func firstDepthFirstSearch(target string, container *ElementContainer, index int) *TreeNode {
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return nil
		}
		return &TreeNode{
			Name:  target,
			Left:  nil,
			Right: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode {
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

	i := 0
	for _, pair := range container.Container[target] {
		if i < index {
			i++
			continue
		}

		i++
		left := depthFirstSearch(pair.Component1, container)
		right := depthFirstSearch(pair.Component2, container)
		container.IsVisited[target] = false
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

func depthFirstSearch(target string, container *ElementContainer) *TreeNode {
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return nil
		}
		return &TreeNode{
			Name:  target,
			Left:  nil,
			Right: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode {
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
		t1, ok1 := container.ElementTier[pair.Component1]
		t2, ok2 := container.ElementTier[pair.Component2]
		tTarget, okT := container.ElementTier[target]

		if !ok1 || !ok2 || !okT {
			continue
		}
		if t1 > tTarget || t2 > tTarget {
			continue
		}

		left := depthFirstSearch(pair.Component1, container)
		right := depthFirstSearch(pair.Component2, container)
		container.IsVisited[target] = false
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