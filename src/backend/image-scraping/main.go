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

// fetchRealImageURL tries to find a real image URL for an element by visiting its wiki page
func fetchRealImageURL(elementURL string) string {
	// Make HTTP request to the element's wiki page
	response, err := http.Get(elementURL)
	if err != nil {
		log.Printf("Warning: Failed to fetch image for %s: %v", elementURL, err)
		return ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Warning: Failed to fetch image, status code: %d", response.StatusCode)
		return ""
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Warning: Failed to parse HTML: %v", err)
		return ""
	}

	// Try to find the element image in the infobox or main content
	// First look for the infobox image
	imgURL, exists := doc.Find(".pi-image img").First().Attr("src")
	if !exists || strings.Contains(imgURL, "data:image/gif;base64") {
		// Try to find an image in the main content
		imgURL, exists = doc.Find(".mw-parser-output img").First().Attr("src")
		if !exists || strings.Contains(imgURL, "data:image/gif;base64") {
			return ""
		}
	}

	return imgURL
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
			
			// Skip empty names
			if name == "" {
				return
			}
			
			// Get the URL for the wiki page of this element
			href, exists := link.Attr("href")
			if !exists {
				href = ""
			}
			
			// Construct image URL based on the element name
			// Get image URL from the image tag, but filter out placeholder GIFs
			imgSrc, exists := nameCell.Find("img").Attr("src")
			if !exists || strings.Contains(imgSrc, "data:image/gif;base64") {
				// If no image or it's a placeholder, create a URL for the element's wiki page to get it later
				if href != "" {
					imgSrc = "https://little-alchemy.fandom.com" + href
				} else {
					imgSrc = ""
				}
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
				
				// Construct a URL for the element's wiki page to get the image later
				elementURL := "https://little-alchemy.fandom.com" + href
				
				// Add the element
				elements = append(elements, Element{
					Name:     name,
					ImageURL: elementURL, // Store the wiki page URL to get a real image
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
	
	// Try to fetch real image URLs for elements (with a limit to avoid too many requests)
	maxImagesToFetch := 10
	fetchCount := 0
	
	for i := range finalElements {
		// Skip elements that already have valid image URLs
		if finalElements[i].ImageURL != "" && !strings.HasPrefix(finalElements[i].ImageURL, "https://little-alchemy.fandom.com/wiki/") {
			continue
		}
		
		// Limit the number of requests to avoid overloading the server
		if fetchCount >= maxImagesToFetch {
			break
		}
		
		// Try to get a real image URL
		if strings.HasPrefix(finalElements[i].ImageURL, "https://little-alchemy.fandom.com/wiki/") {
			realImageURL := fetchRealImageURL(finalElements[i].ImageURL)
			if realImageURL != "" {
				finalElements[i].ImageURL = realImageURL
				fmt.Printf("Found image for %s: %s\n", finalElements[i].Name, realImageURL)
			} else {
				// If we can't get a real image, use a generic one or clear it
				finalElements[i].ImageURL = ""
			}
			fetchCount++
		}
	}

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
