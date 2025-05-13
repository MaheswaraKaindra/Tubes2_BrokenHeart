package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"log"
	"os"
	"strings"
	"time"
	"runtime"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

import "github.com/MaheswaraKaindra/Tubes2_BrokenHeart/src/backend/logic"

type SearchRequest struct {
	Target string `json:"target"`
	Index  int    `json:"index"`
}

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

type ElementTier struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Tier int    `json:"tier"`
}

func getDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Invalid file path.")
	}
	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "data")
}

func saveTier(path string, tierMap map[string]int) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(tierMap)
}

func scrapeTier() ([]ElementTier, error) {
	const url = "https://little-alchemy.fandom.com/wiki/Elements_%28Little_Alchemy_2%29"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Parse HTML: %w", err)
	}

	var items []ElementTier

	tempContainer := doc.Find("div.mw-parser-output > h3")

	tempContainer.Each(func(i int, h3 *goquery.Selection) {
		span := h3.Find("span.mw-headline")
		hdr := strings.TrimSpace(span.Text())

		id, _ := span.Attr("id")
		if !strings.HasPrefix(id, "Tier_") {
			return
		}

		parts := strings.Fields(hdr)
		if len(parts) < 2 {
			return
		}
		tierNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			return
		}

		sib := h3.Next()
		for sib.Length() > 0 && goquery.NodeName(sib) != "table" {
			sib = sib.Next()
		}
		if sib.Length() == 0 {
			return
		}

		// TableMAXXING.
		rows := sib.Find("tr")
		rows.Each(func(j int, row *goquery.Selection) {
			if row.Find("th").Length() > 0 {
				return
			}
			cell := row.Find("td").First()
			a := cell.Find("a[title]").First()
			name := strings.TrimSpace(a.Text())
			href, _ := a.Attr("href")

			if name == "" || !strings.HasPrefix(href, "/wiki/") {
				return
			}

			items = append(items, ElementTier{
				Name: name,
				URL:  "https://little-alchemy.fandom.com" + href,
				Tier: tierNumber,
			})
		})
	})

	if len(items) == 0 {
		return nil, nil
	}
	return items, nil
}

func scrapeData() {
	items, err := scrapeTier()
	if err != nil {
		log.Fatalf("Failed to scrape: %v", err)
	}

	tierMap := make(map[string]int)
	for _, item := range items {
		tierMap[item.Name] = item.Tier
	}

	outputPath := filepath.Join(getDirectory(), "tiers.json")
	if err := saveTier(outputPath, tierMap); err != nil {
		log.Fatalf("Failed to save: %v", err)
	}

	fmt.Printf("%d elements was saved to %s.\n", len(tierMap), outputPath)
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

	doc.Find("h3").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if strings.Contains(selection.Find(".mw-headline").Text(), "Little Alchemy 2") {
			found = true
			ul := selection.NextFiltered("ul").First()
			if ul != nil {
				// Web structure T_T.
				ul.Find("li").Each(func(i int, li *goquery.Selection) {
					var combo []string
					li.Find("a").Each(func(i int, a *goquery.Selection) {
						name := strings.TrimSpace(a.Text())
						if name != "" {
							combo = append(combo, name)
						}
					})
					if len(combo) >= 2 && validElements[combo[0]] && validElements[combo[1]] {
						// Intinya nambahin di sini.
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

func scrapeAllData() {
	fmt.Println("Starting scraping...")
	scrapeData()

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

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func loadElementContainerFromFiles() (*logic.ElementContainer, error) {
	recipesPath := filepath.Join(".", "data", "recipes.json")
	tiersPath := filepath.Join(".", "data", "tiers.json")
	imagesPath := filepath.Join(".", "data", "images.json")
	
	elements, err := logic.ReadJSON(recipesPath, tiersPath, imagesPath)
	if err != nil {
		return nil, err
	}
	container := logic.BuildElementContainer(elements)
	return &container, nil
}

func shortestbfsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	root := logic.ShortestBreadthFirstSearch(req.Target, container)
	duration := time.Since(start)
	fmt.Printf("Target: %s, Result: %+v, Duration: %s\n", req.Target, root, duration)


	response := map[string]interface{}{
		"data":          root,
		"executionTime": duration.String(),
 	} 	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func shortestdfsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var visitedCount = new(int)
	*visitedCount = 0
	root := logic.ShortestDepthFirstSearch(req.Target, container, visitedCount)
	duration := time.Since(start)

	response := map[string]interface{}{
		"data":          root,
		"executionTime": duration.String(),
 	} 	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func multiplebfsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	loopCount := logic.GetLength(container, req.Target)
	recipes := logic.GetRecipe(container, req.Target, loopCount)
	var trees []interface{}

	for i := 0; i < loopCount; i++ {
		tree := logic.BreadthFirstSearch(req.Target, container, i)
		trees = append(trees, tree)
	}
	duration := time.Since(start)

	response := map[string]interface{}{
		"trees":   trees,
		"recipes": recipes,
		"executionTime": duration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func multipledfsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	container, err := loadElementContainerFromFiles()
	if err != nil {
		http.Error(w, "Error loading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	loopCount := logic.GetLength(container, req.Target)
	recipes := logic.GetRecipe(container, req.Target, loopCount)
	var trees []interface{}

	for i := 0; i < loopCount; i++ {
		var visitedCount = new(int)
		*visitedCount = 0
		tree := logic.FirstDepthFirstSearch(req.Target, container, i, visitedCount)
		trees = append(trees, tree)
	}
	duration := time.Since(start)

	response := map[string]interface{}{
		"trees":   trees,
		"recipes": recipes,
		"executionTime": duration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// scrapeAllData()
	http.HandleFunc("/api/bfs", shortestbfsHandler)
	http.HandleFunc("/api/dfs", shortestdfsHandler)
	http.HandleFunc("/api/bfsmultiple", multiplebfsHandler)
	http.HandleFunc("/api/dfsmultiple", multipledfsHandler)

	fmt.Println("Server running on :8080")
	logErr := http.ListenAndServe(":8080", nil)
	if logErr != nil {
		fmt.Println("Server error:", logErr)
	}
}