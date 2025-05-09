package main

import (
	"container/list"
)

func firstBreadthFirstSearch(target string, container *ElementContainer, index int) *TreeNode {
	if isBaseElement(target) {
		return &TreeNode{Name: target}
	}
	if _, exists := container.Container[target]; !exists {
		return nil
	}

	visited := make(map[string]bool)
	queue := list.New()

	i := 0
	for _, pair := range container.Container[target] {
		
		if i < index {
			i++
			continue
		}

		i++

		queue.PushBack(SearchState {
			Node: &TreeNode{
				Name:  target,
				Left:  &TreeNode{Name: pair.Component1},
				Right: &TreeNode{Name: pair.Component2},
			},
			Target: target,
		})
	}

	for queue.Len() > 0 {
		elem := queue.Remove(queue.Front()).(SearchState)
		node := elem.Node

		if !isBaseElement(node.Left.Name) && !visited[node.Left.Name] {
			visited[node.Left.Name] = true
			node.Left = depthFirstSearch(node.Left.Name, container)
		}
		if !isBaseElement(node.Right.Name) && !visited[node.Right.Name] {
			visited[node.Right.Name] = true
			node.Right = depthFirstSearch(node.Right.Name, container)
		}

		if node.Left != nil && node.Right != nil {
			return node
		}
	}

	return nil
}

func breadthFirstSearch(target string, container *ElementContainer) *TreeNode {
	if isBaseElement(target) {
		return &TreeNode{Name: target}
	}
	if _, exists := container.Container[target]; !exists {
		return nil
	}

	visited := make(map[string]bool)
	queue := list.New()

	for _, pair := range container.Container[target] {
		t1, ok1 := container.ElementTier[pair.Component1]
		t2, ok2 := container.ElementTier[pair.Component2]
		tTarget := container.ElementTier[target]

		if !ok1 || !ok2 || t1 > tTarget || t2 > tTarget {
			continue
		}

		queue.PushBack(SearchState{
			Node: &TreeNode{
				Name:  target,
				Left:  &TreeNode{Name: pair.Component1},
				Right: &TreeNode{Name: pair.Component2},
			},
			Target: target,
		})
	}

	for queue.Len() > 0 {
		elem := queue.Remove(queue.Front()).(SearchState)
		node := elem.Node

		if !isBaseElement(node.Left.Name) && !visited[node.Left.Name] {
			visited[node.Left.Name] = true
			node.Left = depthFirstSearch(node.Left.Name, container)
		}
		if !isBaseElement(node.Right.Name) && !visited[node.Right.Name] {
			visited[node.Right.Name] = true
			node.Right = depthFirstSearch(node.Right.Name, container)
		}

		if node.Left != nil && node.Right != nil {
			return node
		}
	}

	return nil
}
