package main

import (
    "path/filepath"
    "fmt"
    "log"
    "os"
)

func printTree(node *SingularTreeNode, indent string, isLeft bool) {
	if node == nil {
		return
	}

	prefix := "└──"
	if isLeft {
		prefix = "├──"
	}
	fmt.Println(indent + prefix + node.Name)

	newIndent := indent
	if isLeft {
		newIndent += "│   "
	} else {
		newIndent += "    "
	}

	printTree(node.Left, newIndent, true)
	printTree(node.Right, newIndent, false)
}

func main() {
	filename := filepath.Join(".", "..", "data", "recipes.json")

	elements, err := readJSON(filename)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	container := buildElementContainer(elements)

	var target string
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		fmt.Print("Target: ")
		fmt.Scanln(&target)
	}

	fullTree := depthFirstSearch(target, &container)

	if fullTree == nil {
		fmt.Println("There's no solution for the %s.", target)
		return
	}

	fmt.Printf("Target %s has %d combinations.\n", fullTree.Name, len(fullTree.Recipes))

	var rootIndex int
	fmt.Printf("Choose recipes(0 - %d): ", len(fullTree.Recipes)-1)
	fmt.Scanln(&rootIndex)

	indexMap := map[string]int{
		target: rootIndex,
	}

	selectedTree := getSingularTree(fullTree, indexMap)

	printTree(selectedTree, "", false)
}
