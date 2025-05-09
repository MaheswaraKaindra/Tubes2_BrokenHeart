package main

import (
    "path/filepath"
    "fmt"
    "log"
    "os"
)

func printTree(node *TreeNode, indent string, isLeft bool) {
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

func printAllElements(elements []Element) {
	fmt.Println("Element list:")
	for _, el := range elements {
		fmt.Printf("🔹 %s (Tier: %d)\n", el.Name, el.Tier)
		if len(el.Components) == 0 {
			fmt.Println("    -")
		}
		for i, pair := range el.Components {
			if len(pair) == 2 {
				fmt.Printf("    %d. %s + %s\n", i+1, pair[0], pair[1])
			}
		}
	}
}

func main() {
	filename := filepath.Join(".", "..", "data", "recipes.json")
	tiersFile := filepath.Join(".", "..", "data", "tiers.json")

	elements, err := readJSON(filename, tiersFile)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}	

	printAllElements(elements)
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

	fmt.Printf("Target %s has %d combinations.\n", fullTree.Name)

	printTree(fullTree, "", false)
}
