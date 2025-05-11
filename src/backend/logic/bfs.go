package logic

import (
	"container/list"
	"strings"
	// "fmt"
)

// Untuk enqueue : queue.PushBack(node)
// Untuk dequeue : queue.Remove(queue.Front()), bisa simpan ke value jika butuh.

func BreadthFirstSearch(target string, container *ElementContainer, index int) *TreeNode {
	target = strings.ToLower(target)
	queue := list.New()

	root := &TreeNode{Name: target, Image: container.ElementImage[target]}
	queue.PushBack(root)

	first := true

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		parentNode := element.Value.(*TreeNode)

		pairs := container.Container[parentNode.Name]
		if len(pairs) == 0 {
			continue
		}

		i := index

		if !first {
			shortestMap := make(map[int]int)
			i = 0
			for _, pair := range container.Container[parentNode.Name] {
				// fmt.Printf("[EXPAND] %s → %s + %s\n", parentNode.Name, pair.Component1, pair.Component2)
				t1, ok1 := container.ElementTier[pair.Component1]
				t2, ok2 := container.ElementTier[pair.Component2]
				tTarget, okT := container.ElementTier[parentNode.Name]
		
				if !ok1 || !ok2 || !okT {
					i++
					continue
				}
				if t1 >= tTarget || t2 >= tTarget {
					i++
					continue
				}

				shortestMap[i] = container.ElementTier[pair.Component1] + container.ElementTier[pair.Component2]
				i++
			}
			i = minKey(shortestMap)
		}

		pair := pairs[i]
		leftName := pair.Component1
		rightName := pair.Component2
		
		if (leftName == parentNode.Name || rightName == parentNode.Name) {
			continue
		}
		
		leftNode := &TreeNode{ Name: leftName, Image: container.ElementImage[leftName] }
		rightNode := &TreeNode{ Name: rightName, Image: container.ElementImage[rightName] }

		if !isBaseElement(leftName) {
			parentNode.Left = leftNode
			queue.PushBack(leftNode)
		} else {
			parentNode.Left = &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
		}
		
		if !isBaseElement(rightName) {
			parentNode.Right = rightNode
			queue.PushBack(rightNode)
		} else {
			parentNode.Right = &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}
		}

		first = false
	}

	return root
}

func ShortestBreadthFirstSearch(target string, container *ElementContainer) *TreeNode {
	target = strings.ToLower(target)
	queue := list.New()

	root := &TreeNode{Name: target, Image: container.ElementImage[target]}
	queue.PushBack(root)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		parentNode := element.Value.(*TreeNode)

		pairs := container.Container[parentNode.Name]
		if len(pairs) == 0 {
			continue
		}

		shortestMap := make(map[int]int)

		i := 0
		for _, pair := range container.Container[parentNode.Name] {
			// fmt.Printf("[EXPAND] %s → %s + %s\n", parentNode.Name, pair.Component1, pair.Component2)
			t1, ok1 := container.ElementTier[pair.Component1]
			t2, ok2 := container.ElementTier[pair.Component2]
			tTarget, okT := container.ElementTier[parentNode.Name]
	
			if !ok1 || !ok2 || !okT {
				i++
				continue
			}
			if t1 >= tTarget || t2 >= tTarget {
				i++
				continue
			}

			shortestMap[i] = container.ElementTier[pair.Component1] + container.ElementTier[pair.Component2]
			i++
		}

		index := minKey(shortestMap)

		pair := pairs[index]
		leftName := pair.Component1
		rightName := pair.Component2
		
		if (leftName == parentNode.Name || rightName == parentNode.Name) {
			continue
		}
		
		leftNode := &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
		rightNode := &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}

		if !isBaseElement(leftName) {
			parentNode.Left = leftNode
			queue.PushBack(leftNode)
		} else {
			parentNode.Left = &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
		}
		
		if !isBaseElement(rightName) {
			parentNode.Right = rightNode
			queue.PushBack(rightNode)
		} else {
			parentNode.Right = &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}
		}
	}

	return root
}
