package logic

import (
	"strings"
	"sync"
)

var dfsMutex sync.Mutex

func FirstDepthFirstSearch(target string, container *ElementContainer, index int) *TreeNode {
	target = strings.ToLower(target)
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return nil
		}
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	dfsMutex.Lock()
	if container.IsVisited[target] {
		dfsMutex.Unlock()
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}
	container.IsVisited[target] = true
	dfsMutex.Unlock()

	i := 0
	for _, pair := range container.Container[target] {
		if i < index {
			i++
			continue
		}

		i++
		var left, right *TreeNode
		var wg sync.WaitGroup
		wg.Add(2)

		go func(p ComponentKey) {
			left = depthFirstSearch(p.Component1, container)
			wg.Done()
		}(pair)

		go func(p ComponentKey) {
			right = depthFirstSearch(p.Component2, container)
			wg.Done()
		}(pair)

		wg.Wait()

		dfsMutex.Lock()
		container.IsVisited[target] = false
		dfsMutex.Unlock()

		if left != nil && right != nil {
			return &TreeNode{
				Name:  target,
				Image: container.ElementImage[target],
				Left:  left,
				Right: right,
			}
		}
	}

	return nil
}

func ShortestDepthFirstSearch(target string, container *ElementContainer) *TreeNode {
	target = strings.ToLower(target)
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return nil
		}
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	dfsMutex.Lock()
	if container.IsVisited[target] {
		dfsMutex.Unlock()
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}
	container.IsVisited[target] = true
	dfsMutex.Unlock()

	shortestMap := make(map[int]int)
	i := 0
	for _, pair := range container.Container[target] {
		t1, ok1 := container.ElementTier[pair.Component1]
		t2, ok2 := container.ElementTier[pair.Component2]
		tTarget, okT := container.ElementTier[target]

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
		dfsMutex.Lock()
		container.IsVisited[target] = false
		dfsMutex.Unlock()
		return nil
	}

	i = minKey(shortestMap)
	pair := container.Container[target][i]

	var left, right *TreeNode
	var wg sync.WaitGroup
	wg.Add(2)

	go func(p ComponentKey) {
		left = depthFirstSearch(p.Component1, container)
		wg.Done()
	}(pair)

	go func(p ComponentKey) {
		right = depthFirstSearch(p.Component2, container)
		wg.Done()
	}(pair)

	wg.Wait()

	dfsMutex.Lock()
	container.IsVisited[target] = false
	dfsMutex.Unlock()

	if left != nil && right != nil {
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  left,
			Right: right,
		}
	}

	return nil
}

func depthFirstSearch(target string, container *ElementContainer) *TreeNode {
	target = strings.ToLower(target)
	if _, exists := container.Container[target]; !exists {
		if !isBaseElement(target) {
			return nil
		}
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	if isBaseElement(target) {
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}

	dfsMutex.Lock()
	if container.IsVisited[target] {
		dfsMutex.Unlock()
		return &TreeNode{
			Name:  target,
			Image: container.ElementImage[target],
			Left:  nil,
			Right: nil,
		}
	}
	container.IsVisited[target] = true
	dfsMutex.Unlock()

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

		var left, right *TreeNode
		var wg sync.WaitGroup
		wg.Add(2)

		go func(p ComponentKey) {
			left = depthFirstSearch(p.Component1, container)
			wg.Done()
		}(pair)

		go func(p ComponentKey) {
			right = depthFirstSearch(p.Component2, container)
			wg.Done()
		}(pair)

		wg.Wait()

		dfsMutex.Lock()
		container.IsVisited[target] = false
		dfsMutex.Unlock()

		if left != nil && right != nil {
			return &TreeNode{
				Name:  target,
				Image: container.ElementImage[target],
				Left:  left,
				Right: right,
			}
		}
	}

	dfsMutex.Lock()
	container.IsVisited[target] = false
	dfsMutex.Unlock()

	return nil
}
