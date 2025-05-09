package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSON(recipes string, tiersFile string) ([]Element, error) {
	file, err := os.Open(recipes)
	if err != nil {
		return nil, fmt.Errorf("Error opening recipes file: %w", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading recipes file: %w", err)
	}

	var elements []Element
	if err := json.Unmarshal(byteValue, &elements); err != nil {
		return nil, fmt.Errorf("Error parsing recipes JSON: %w", err)
	}

	// Baca tiers
	file2, err := os.Open(tiersFile)
	if err != nil {
		return nil, fmt.Errorf("Error opening tiers file: %w", err)
	}
	defer file2.Close()

	byteTier, err := ioutil.ReadAll(file2)
	if err != nil {
		return nil, fmt.Errorf("Error reading tiers file: %w", err)
	}

	var tiers map[string]int
	if err := json.Unmarshal(byteTier, &tiers); err != nil {
		return nil, fmt.Errorf("Error parsing tiers JSON: %w", err)
	}

	// Assign tier ke masing-masing elemen
	for i, el := range elements {
		if tier, ok := tiers[el.Name]; ok {
			elements[i].Tier = tier
		} else {
			elements[i].Tier = 0 // atau -1 sebagai tanda tidak diketahui
		}
	}

	required := []string{"Fire", "Time"}
	existing := make(map[string]bool)

	for _, el := range elements {
		existing[el.Name] = true
	}

	for _, name := range required {
		if !existing[name] {
			elements = append(elements, Element{
				Name:       name,
				Components: [][]string{},
				Tier:       0,
			})
		}
	}

	return elements, nil
}

func buildElementContainer(elements []Element) ElementContainer {
	container := make(map[string][]ComponentKey)
	isVisited := make(map[string]bool)
	elementTier := make(map[string]int)

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
		elementTier[el.Name] = el.Tier
	}

	return ElementContainer{
		Container:    container,
		IsVisited:    isVisited,
		ElementTier:  elementTier,
	}
}
