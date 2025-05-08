package logic

import (
	"container/list"
)

func breadthFirstSearch(target string, container *ElementContainer) *TreeNode {
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

	type Node struct {
		Name  string
		Path  *TreeNode
	}

	visited := make(map[string]bool)
	queue := list.New()

	// Enqueue initial target
	queue.PushBack(Node{
		Name: target,
		Path: nil,
	})

	for queue.Len() > 0 {
		// Dequeue
		elem := queue.Remove(queue.Front()).(Node)
		name := elem.Name

		// ketika sudah ada visit, skip
		if visited[name] {
			continue
		}
		visited[name] = true

		for _, pair := range container.Container[name] {
			left := buildSubtreeBFS(pair.Component1, container, visited)
			right := buildSubtreeBFS(pair.Component2, container, visited)

			if left != nil && right != nil {
				return &TreeNode{
					Name:  name,
					Left:  left,
					Right: right,
				}
			}
		}
	}

	return nil
}

func buildSubtreeBFS(name string, container *ElementContainer, visited map[string]bool) *TreeNode {
	if isBaseElement(name) {
		return &TreeNode{Name: name}
	}
	if _, exists := container.Container[name]; !exists {
		return nil
	}

	for _, pair := range container.Container[name] {
		left := buildSubtreeBFS(pair.Component1, container, visited)
		right := buildSubtreeBFS(pair.Component2, container, visited)

		if left != nil && right != nil {
			return &TreeNode{
				Name:  name,
				Left:  left,
				Right: right,
			}
		}
	}

	return nil
}