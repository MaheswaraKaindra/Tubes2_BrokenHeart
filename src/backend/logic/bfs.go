package logic

import (
	"strings"
	"sync"
)

func BreadthFirstSearch(target string, container *ElementContainer, index int) *Result {
	target = strings.ToLower(target)
	queue := make(chan *TreeNode, 100)
	var wg sync.WaitGroup

	var root *TreeNode
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return &Result{Node: nil, VisitedCount: 0}
		}
		root = &TreeNode{Name: target, Image: container.ElementImage[target]}
		return &Result{Node: root, VisitedCount: 1}
	} else {
		root = &TreeNode{Name: target, Image: container.ElementImage[target]}
	}

	queue <- root
	first := true

	visitedCount := 0
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		for parentNode := range queue {

			pairs := container.Container[parentNode.Name]
			if len(pairs) == 0 {
				wg.Done()
				continue
			}

			i := index
			if !first {
				shortestMap := make(map[int]int)
				j := 0
				for _, pair := range pairs {
					t1, ok1 := container.ElementTier[pair.Component1]
					t2, ok2 := container.ElementTier[pair.Component2]
					tTarget, okT := container.ElementTier[parentNode.Name]
					if !ok1 || !ok2 || !okT {
						j++
						continue
					}
					if t1 >= tTarget || t2 >= tTarget {
						j++
						continue
					}
					shortestMap[j] = t1 + t2
					j++
				}
				i = minKey(shortestMap)
			}

			pair := pairs[i]
			leftName := pair.Component1
			rightName := pair.Component2

			if leftName == parentNode.Name || rightName == parentNode.Name {
				wg.Done()
				continue
			}

			leftNode := &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
			rightNode := &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}

			if !isBaseElement(leftName) {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Left = leftNode
				wg.Add(1)
				queue <- leftNode
			} else {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Left = &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
			}

			if !isBaseElement(rightName) {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Right = rightNode
				wg.Add(1)
				queue <- rightNode
			} else {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Right = &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}
			}

			first = false
			wg.Done()
		}
	}()

	wg.Wait()
	close(queue)

	return &Result{
		Node:         root,
		VisitedCount: visitedCount,
	}
}

func ShortestBreadthFirstSearch(target string, container *ElementContainer) *Result {
	target = strings.ToLower(target)
	queue := make(chan *TreeNode, 100)
	var wg sync.WaitGroup

	var root *TreeNode
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return &Result{Node: nil, VisitedCount: 0}
		}
		root = &TreeNode{Name: target, Image: container.ElementImage[target]}
		return &Result{Node: root, VisitedCount: 1}
	} else {
		root = &TreeNode{Name: target, Image: container.ElementImage[target]}
	}
	queue <- root

	visitedCount := 0
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		for parentNode := range queue {
			mu.Lock()
			visitedCount++
			mu.Unlock()

			pairs := container.Container[parentNode.Name]
			if len(pairs) == 0 {
				wg.Done()
				continue
			}

			shortestMap := make(map[int]int)
			i := 0
			for _, pair := range pairs {
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
				shortestMap[i] = t1 + t2
				i++
			}

			if len(shortestMap) == 0 {
				wg.Done()
				continue
			}

			index := minKey(shortestMap)
			pair := pairs[index]
			leftName := pair.Component1
			rightName := pair.Component2

			if leftName == parentNode.Name || rightName == parentNode.Name {
				wg.Done()
				continue
			}

			leftNode := &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
			rightNode := &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}

			if !isBaseElement(leftName) {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Left = leftNode
				wg.Add(1)
				queue <- leftNode
			} else {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Left = &TreeNode{Name: leftName, Image: container.ElementImage[leftName]}
			}

			if !isBaseElement(rightName) {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Right = rightNode
				wg.Add(1)
				queue <- rightNode
			} else {
				mu.Lock()
				visitedCount++
				mu.Unlock()
				parentNode.Right = &TreeNode{Name: rightName, Image: container.ElementImage[rightName]}
			}

			wg.Done()
		}
	}()

	wg.Wait()
	close(queue)

	return &Result{
		Node:         root,
		VisitedCount: visitedCount,
	}
}
