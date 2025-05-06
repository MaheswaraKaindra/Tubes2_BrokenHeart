package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Recipe struct {
	Element    string     `json:"element"`
	Components [][]string `json:"components"`
}

type Element struct {
	Name string
	URL  string
}

func MythAndMonstersElements() ([]string, error) {
	url := "https://little-alchemy.fandom.com/wiki/Category:Myths_and_Monsters"
	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var elements []string
	doc.Find(".category-page__member-link").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		if name != "" {
			elements = append(elements, name)
		}
	})
	return elements, nil
}

func BanMythAndMonsters(elements []string) map[string]bool {
	banlist := make(map[string]bool)
	for _, element := range elements {
		banlist[element] = true
	}
	return banlist
}

func AllElements() ([]Element, error) {
	URL := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"
	base := "https://little-alchemy.fandom.com"

	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var elements []Element
	seen := make(map[string]bool)

	myths, _ := MythAndMonstersElements()
	banlist := BanMythAndMonsters(myths)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		title, titleExists := s.Attr("title")

		if exists && titleExists && !strings.Contains(href, ":") {
			if !seen[title] && !banlist[title] {
				seen[title] = true
				elements = append(elements, Element{Name: title, URL: base + href})
			}
		}
	})
	return elements, nil
}

func ElementPage(element Element, validElements map[string]bool) ([][]string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(element.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var combinations [][]string
	found := false

	doc.Find("h3").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if strings.Contains(s.Find(".mw-headline").Text(), "Little Alchemy 2") {
			found = true
			ul := s.NextFiltered("ul").First()
			if ul != nil {
				ul.Find("li").Each(func(i int, li *goquery.Selection) {
					var combo []string
					li.Find("a").Each(func(i int, a *goquery.Selection) {
						name := strings.TrimSpace(a.Text())
						if name != "" {
							combo = append(combo, name)
						}
					})
					if len(combo) >= 2 && validElements[combo[0]] && validElements[combo[1]] {
						combinations = append(combinations, combo[:2])
					}
				})
			}
			return false
		}
		return true
	})

	if !found {
		fmt.Printf("No 'Little Alchemy 2' section found for %s\n", element.Name)
		return nil, nil
	}

	return combinations, nil
}

func GetAllRecipes(elements []Element) ([]Recipe, error) {
	validElements := make(map[string]bool)
	for _, element := range elements {
		validElements[element.Name] = true
	}

	myths, _ := MythAndMonstersElements()
	banlist := BanMythAndMonsters(myths)
	for elem := range banlist {
		validElements[elem] = false
	}

	var allRecipes []Recipe
	for _, element := range elements {
		combinations, err := ElementPage(element, validElements)
		if err != nil {
			fmt.Printf("Error getting recipes for %s: %v\n", element.Name, err)
			continue
		}
		if len(combinations) > 0 {
			allRecipes = append(allRecipes, Recipe{
				Element:    element.Name,
				Components: combinations,
			})
		}
	}
	return allRecipes, nil
}

func main() {
	fmt.Println("Starting scraping...")

	elements, err := AllElements()
	if err != nil {
		log.Fatalf("Failed to fetch elements: %v", err)
	}

	recipes, err := GetAllRecipes(elements)
	if err != nil {
		log.Fatalf("Failed to fetch recipes: %v", err)
	}

	outputPath := filepath.Join("data", "recipes.json")
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(recipes); err != nil {
		log.Fatalf("Failed to encode JSON: %v", err)
	}

	fmt.Printf("Done. Data saved to %s\n", outputPath)
}
