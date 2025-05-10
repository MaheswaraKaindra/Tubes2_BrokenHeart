package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Element represents a Little Alchemy 2 element
type Element struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

func main() {
	// URL of the wiki page
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to request the page: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Failed to fetch the page, status code: %d", response.StatusCode)
	}

	// Read the HTML
	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Save HTML to file for debugging
	if err := os.WriteFile("wiki_page.html", htmlBytes, 0644); err != nil {
		log.Printf("Warning: Failed to save HTML for debugging: %v", err)
	}

	// Create a new reader with the saved HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlBytes)))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	elements := make([]Element, 0)

	// Find all tier sections that contain element tables
	doc.Find(`[id^="Tier_"], #Starting_elements, #Special_element`).Each(func(i int, section *goquery.Selection) {
		// Move to the element table that follows each tier heading
		sectionID := section.AttrOr("id", "")
		if sectionID == "" {
			return
		}

		// Find the table that follows this section
		elementTable := section.Parent().NextUntil("h2, h3").Find("table").First()
		
		// Process each row in the table
		elementTable.Find("tr").Each(func(j int, row *goquery.Selection) {
			// Skip header rows (usually the first row)
			if j == 0 {
				return
			}
			
			// Get the element name from the link in the first cell
			nameCell := row.Find("td").First()
			link := nameCell.Find("a").First()
			name := strings.TrimSpace(link.Text())
			
			// Get image URL from the image tag
			imgSrc, exists := nameCell.Find("img").Attr("src")
			if !exists {
				imgSrc = "" // No image found
			}
			
			// Skip empty names
			if name == "" {
				return
			}
			
			// Add the element to our list
			elements = append(elements, Element{
				Name:     name,
				ImageURL: imgSrc,
			})
		})
	})

	// If the above approach fails, try a more direct approach to find all element links
	if len(elements) == 0 {
		fmt.Println("Using alternative approach to find elements...")
		
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}
			
			// Check if the link points to an element page
			if strings.HasPrefix(href, "/wiki/") && 
			   !strings.Contains(href, "Category:") &&
			   !strings.Contains(href, "Elements_") {
				
				name := strings.TrimSpace(s.Text())
				
				// Ignore links with no text or very long text (likely not elements)
				if name == "" || len(name) > 50 || 
				   strings.Contains(name, "Tier") || 
				   strings.Contains(name, "element") {
					return
				}
				
				// Get image if available
				imgSrc, _ := s.Find("img").Attr("src")
				if imgSrc == "" {
					imgSrc, _ = s.Parent().Find("img").Attr("src")
				}
				
				// Add the element
				elements = append(elements, Element{
					Name:     name,
					ImageURL: imgSrc,
				})
			}
		})
	}

	// Process results - remove duplicates
	uniqueElements := make(map[string]Element)
	for _, element := range elements {
		// Skip non-element entries
		if strings.Contains(strings.ToLower(element.Name), "tier") ||
		   strings.Contains(strings.ToLower(element.Name), "category:") ||
		   strings.Contains(strings.ToLower(element.Name), "elements") {
			continue
		}
		
		// Store only unique elements by name
		uniqueElements[element.Name] = element
	}

	// Convert map back to slice
	finalElements := make([]Element, 0, len(uniqueElements))
	for _, element := range uniqueElements {
		finalElements = append(finalElements, element)
	}

	fmt.Printf("Found %d unique elements\n", len(finalElements))

	// Create JSON file
	file, err := os.Create("little_alchemy_elements.json")
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()

	// Convert to JSON
	jsonData, err := json.MarshalIndent(finalElements, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Write to file
	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Println("Scraping completed successfully. Results saved to little_alchemy_elements.json")
}
