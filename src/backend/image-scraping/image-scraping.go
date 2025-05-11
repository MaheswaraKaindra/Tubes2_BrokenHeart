package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	// "strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ElementImage struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	ImageURL string    `json:"img"` // Ganti sedikit.
}

func getDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Invalid file path.")
	}
	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "data")
}

func saveImages(path string, elements []ElementImage) error {
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
	return enc.Encode(elements)
}

// AWKOAWKO maling dari saveTier.

// func saveTier(path string, tierMap map[string]int) error {
// 	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
// 		return err
// 	}
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	enc := json.NewEncoder(f)
// 	enc.SetIndent("", "  ")
// 	return enc.Encode(tierMap)
// }

func scrapeImage() ([]ElementImage, error) {
	const url = "https://little-alchemy.fandom.com/wiki/Elements_%28Little_Alchemy_2%29"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Parse HTML: %w", err)
	}

	var items []ElementImage
	containers := doc.Find("div.mw-parser-output > h3")

	containers.Each(func(i int, h3 *goquery.Selection) {
		span := h3.Find("span.mw-headline")
		id, _ := span.Attr("id")
		if !strings.HasPrefix(id, "Tier_") {
			return
		}

		sib := h3.Next()
		for sib.Length() > 0 && goquery.NodeName(sib) != "table" {
			sib = sib.Next()
		}
		if sib.Length() == 0 {
			return
		}

		sib.Find("tr").Each(func(j int, row *goquery.Selection) {
			if row.Find("th").Length() > 0 {
				return
			}

			cell := row.Find("td").First()
			// Gajadi ambil name, tapi ambil href nya.
			nameLink := cell.Find("a[title]").First()
			name := strings.TrimSpace(nameLink.Text())
			href, _ := nameLink.Attr("href")

			imageLink := cell.Find("span[typeof='mw:File'] a").First()
			imageHref, _ := imageLink.Attr("href")

			if name == "" || !strings.HasPrefix(href, "/wiki/") || !strings.HasPrefix(imageHref, "https://") {
				return
			}

			items = append(items, ElementImage{
				Name:     name,
				URL:      "https://little-alchemy.fandom.com" + href,
				ImageURL: imageHref,
			})
		})
	})

	if len(items) == 0 {
		return nil, nil
	}
	return items, nil

	// Ngikutin tier-scraping.go diganti dikit awokakawko.
}

func scrapeData() { // Sama aja sih tinggal hapus map.
	items, err := scrapeImage()
	if err != nil {
		log.Fatalf("Failed to scrape: %v", err)
	}

	outputPath := filepath.Join(getDirectory(), "images.json")
	if err := saveImages(outputPath, items); err != nil {
		log.Fatalf("Failed to save image data: %v", err)
	}

	fmt.Printf("%d elements with images saved to %s\n", len(items), outputPath)
}

func main() {
	scrapeData()
	// ScrapeData aja ganti main, tapi gpp aowkakw.
}
