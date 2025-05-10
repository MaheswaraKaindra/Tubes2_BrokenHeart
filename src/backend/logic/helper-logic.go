package main

func isBaseElement(name string) bool {
	switch name {
	case "air", "water", "fire", "earth", "time":
		return true
	default:
		return false
	}
}

func normalizeKey(a, b string) ComponentKey {
    if a < b {
        return ComponentKey{a, b}
    }
    return ComponentKey{b, a}
}

func minKey(m map[int]int) int {
	minKey := -1
	minVal := int(^uint(0) >> 1)

	for k, v := range m {
		if v < minVal {
			minVal = v
			minKey = k
		}
	}
	return minKey
}

// FUNGSI HELPER BOSS SENGGOL DONG...

func getLength(container *ElementContainer, element string) int {
	return len(container.Container[element])
}

func getRecipe(container *ElementContainer, element string, many int) []ComponentKey {
	returnValue := make([]ComponentKey, 0)
	for i := 0; i < many; i++ {
		returnValue = append(returnValue, container.Container[element][i])
	}

	return returnValue
}
