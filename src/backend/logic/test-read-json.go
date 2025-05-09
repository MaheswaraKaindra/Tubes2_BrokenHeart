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

	container := buildElementContainer(elements)

	var target string
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		fmt.Print("Target: ")
		fmt.Scanln(&target)
	}

	recipes := container.Container[target]
	numRecipes := len(recipes)

	if numRecipes == 0 {
		fmt.Printf("There's no recipes for %s.\n", target)
		return
	}

	var selected int
	for {
		fmt.Printf("Choose combination (0 - %d): ", numRecipes-1)
		_, err := fmt.Scanln(&selected)
		if err == nil && selected >= 0 && selected < numRecipes {
			break
		}
		fmt.Println("Invalid input.")
	}

	var method string
	fmt.Print("Method (DFS/BFS): ")
	fmt.Scanln(&method)
	if method != "DFS" && method != "BFS" {
		fmt.Println("Invalid method. Using DFS.")
		method = "DFS"
	}

	var fullTree *TreeNode
	if method == "BFS" {
		fullTree = firstBreadthFirstSearch(target, &container, selected)
	} else {
		fullTree = firstDepthFirstSearch(target, &container, selected)
	}

	if fullTree == nil {
		fmt.Printf("There's no way to make %s.\n", target)
		return
	}

	printTree(fullTree, "", false)
}

