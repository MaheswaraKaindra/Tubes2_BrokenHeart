package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
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

	// Save the HTML for debugging
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

	// Try various approaches to find elements

	// First approach: Look for the list elements directly
	fmt.Println("Attempting to scrape elements from lists...")
	doc.Find("#mw-content-text .mw-parser-output ul li").Each(func(i int, s *goquery.Selection) {
		linkText := strings.TrimSpace(s.Find("a").First().Text())
		if linkText != "" {
			imgSrc, _ := s.Find("img").Attr("src")
			elements = append(elements, Element{Name: linkText, ImageURL: imgSrc})
		}
	})

	// Second approach: Try to find tables with element data
	if len(elements) == 0 {
		fmt.Println("Attempting to scrape from tables...")
		doc.Find("table").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(j int, row *goquery.Selection) {
				// Skip header rows
				if j > 0 {
					cells := row.Find("td")
					if cells.Length() >= 2 {
						nameCell := cells.Eq(1)
						name := strings.TrimSpace(nameCell.Text())

						// If cell text is empty, try finding link text
						if name == "" {
							name = strings.TrimSpace(nameCell.Find("a").Text())
						}

						// Find image in first cell
						imgSrc, _ := cells.Eq(0).Find("img").Attr("src")

						if name != "" {
							elements = append(elements, Element{
								Name:     name,
								ImageURL: imgSrc,
							})
						}
					}
				}
			})
		})
	}

	// Third approach: Try to find by HTML patterns in the text
	if len(elements) == 0 {
		fmt.Println("Attempting to parse HTML content directly...")

		// Look for patterns in the HTML that might indicate elements
		htmlContent := string(htmlBytes)

		// Define a regular expression to find element entries
		// This is a simplified pattern and might need adjustment
		elementPattern := regexp.MustCompile(`<a[^>]*?href="[^"]*?/wiki/([^"]+)"[^>]*?>([^<]+)</a>`)

		matches := elementPattern.FindAllStringSubmatch(htmlContent, -1)
		for _, match := range matches {
			if len(match) >= 3 {
				slug := match[1]
				name := strings.TrimSpace(match[2])

				// Skip navigation links and other non-element links
				if strings.Contains(slug, "Category:") ||
					strings.Contains(slug, "Help:") ||
					strings.Contains(slug, "Special:") {
					continue
				}

				// Only add if it looks like an element name
				if name != "" && len(name) < 50 {
					elements = append(elements, Element{
						Name:     name,
						ImageURL: "", // We don't have images with this fallback method
					})
				}
			}
		}
	}

	// Fourth approach: Look for the specific section that contains the list of elements
	if len(elements) == 0 {
		fmt.Println("Looking for sections with element lists...")

		// Find headings that might contain element lists
		doc.Find("h2, h3").Each(func(i int, s *goquery.Selection) {
			heading := strings.TrimSpace(s.Text())
			if strings.Contains(strings.ToLower(heading), "list of element") ||
				strings.Contains(strings.ToLower(heading), "elements") {

				// Find lists after this heading
				elementsList := s.NextUntil("h2, h3").Find("li")

				elementsList.Each(func(j int, li *goquery.Selection) {
					link := li.Find("a").First()
					name := strings.TrimSpace(link.Text())
					imgSrc, _ := li.Find("img").Attr("src")

					if name != "" {
						elements = append(elements, Element{
							Name:     name,
							ImageURL: imgSrc,
						})
					}
				})
			}
		})
	}

	// Last approach: Try to find any links that look like elements
	if len(elements) == 0 {
		fmt.Println("Scanning all links on the page...")

		visited := make(map[string]bool)

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			name := strings.TrimSpace(s.Text())
			href, _ := s.Attr("href")

			// Only consider links that look like wiki pages for elements
			if href != "" &&
				strings.Contains(href, "/wiki/") &&
				!strings.Contains(href, "Category:") &&
				!strings.Contains(href, "Special:") &&
				!strings.Contains(href, "Help:") &&
				name != "" &&
				len(name) < 50 {

				// Skip if we've already seen this name
				if visited[name] {
					return
				}
				visited[name] = true

				// Try to find an image near this link
				imgSrc, _ := s.Find("img").Attr("src")
				if imgSrc == "" {
					imgSrc, _ = s.Parent().Find("img").Attr("src")
				}

				elements = append(elements, Element{
					Name:     name,
					ImageURL: imgSrc,
				})
			}
		})
	}

	// Print the number of elements found
	fmt.Printf("Found %d elements\n", len(elements))

	// Process results - remove duplicates and filter out non-elements
	uniqueElements := make(map[string]Element)
	for _, element := range elements {
		// Skip very short or very long names (likely not elements)
		if len(element.Name) < 2 || len(element.Name) > 50 {
			continue
		}

		// Skip navigation and UI elements
		if strings.Contains(strings.ToLower(element.Name), "category:") ||
			strings.Contains(strings.ToLower(element.Name), "help:") ||
			strings.Contains(strings.ToLower(element.Name), "special:") ||
			strings.Contains(strings.ToLower(element.Name), "edit") ||
			strings.Contains(strings.ToLower(element.Name), "jump to") {
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

	fmt.Printf("After filtering, found %d unique elements\n", len(finalElements))

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
