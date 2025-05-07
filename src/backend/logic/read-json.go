package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Element struct {
	Name	string     `json:"element"`
	Components	[][]string `json:"components"`
}

type ComponentKey struct {
	Component1 string
	Component2 string
}

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

func buildCacheMap (elements []Element) map[ComponentKey]string {
	cacheMap := make(map[ComponentKey]string)
    for _, el := range elements {
        for _, pair := range el.Components {
            if len(pair) == 2 {
                a, b := pair[0], pair[1]
                cacheMap[ComponentKey{a, b}] = el.Name
                cacheMap[ComponentKey{b, a}] = el.Name
            }
        }
    }
    return cacheMap
}