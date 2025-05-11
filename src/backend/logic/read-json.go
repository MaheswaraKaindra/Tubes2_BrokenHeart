package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadJSON(recipes string, tiersFile string, imagesFile string) ([]Element, error) {
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

	for i := range elements {
		elements[i].Name = strings.ToLower(elements[i].Name)
		for j := range elements[i].Components {
			for k := range elements[i].Components[j] {
				elements[i].Components[j][k] = strings.ToLower(elements[i].Components[j][k])
			}
		}
	}

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

	normalizedTiers := make(map[string]int)
	for k, v := range tiers {
		normalizedTiers[strings.ToLower(k)] = v
	}

	for i := range elements {
		if tier, ok := normalizedTiers[elements[i].Name]; ok {
			elements[i].Tier = tier
		} else {
			elements[i].Tier = 0
		}
	}

	type imageEntry struct {
		Name string `json:"name"`
		Img  string `json:"img"`
	}

	file3, err := os.Open(imagesFile)
	if err != nil {
		return nil, fmt.Errorf("Error opening images file: %w", err)
	}
	defer file3.Close()

	byteImage, err := ioutil.ReadAll(file3)
	if err != nil {
		return nil, fmt.Errorf("Error reading images file: %w", err)
	}

	var images []imageEntry
	if err := json.Unmarshal(byteImage, &images); err != nil {
		return nil, fmt.Errorf("Error parsing images JSON: %w", err)
	}

	normalizedImages := make(map[string]string)
	for _, img := range images {
		normalizedImages[strings.ToLower(img.Name)] = img.Img
	}

	for i := range elements {
		if img, ok := normalizedImages[elements[i].Name]; ok {
			elements[i].Image = img
		}
	}

	required := []string{"fire", "time"} // HARDCODED BOZO HAHAHAHA
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

func BuildElementContainer(elements []Element) ElementContainer {
	container := make(map[string][]ComponentKey)
	isVisited := make(map[string]bool)
	elementTier := make(map[string]int)
	elementImage := make(map[string]string)

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
		elementImage[el.Name] = el.Image
	}

	return ElementContainer{
		Container:    container,
		IsVisited:    isVisited,
		ElementTier:  elementTier,
		ElementImage: elementImage,
	}
}
