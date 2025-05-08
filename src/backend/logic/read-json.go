package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSON (filename string) ([]Element, error) {
	file, err := os.Open(filename)
	if err != nil {
		// Hanya untuk debugging.
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("Error reading file: %w", err)
    }

	var elements []Element
    if err := json.Unmarshal(byteValue, &elements); err != nil {
        return nil, fmt.Errorf("Error converting to data struct from JSON: %w", err)
    }

	return elements, nil
}

func BuildElementContainer(elements []Element) ElementContainer {
	container := make(map[string][]ComponentKey)
	isVisited := make(map[string]bool)

	for _, el := range elements {
		for _, pair := range el.Components {
			if len(pair) == 2 {
				key := ComponentKey{
					Component1: pair[0],
					Component2: pair[1],
				}
				container[el.Name] = append(container[el.Name], key)
			}
		}
		isVisited[el.Name] = false
	}

	return ElementContainer{
		Container: container,
		IsVisited: isVisited,
	}
}