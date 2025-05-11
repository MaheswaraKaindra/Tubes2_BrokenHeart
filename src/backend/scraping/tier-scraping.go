package scraping

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ElementTier struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Tier int    `json:"tier"`
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

func scrapeData(outputPath string) {
	items, err := scrapeTier()
	if err != nil {
		log.Fatalf("Failed to scrape: %v", err)
	}

	tierMap := make(map[string]int)
	for _, item := range items {
		tierMap[item.Name] = item.Tier
	}

	if err := saveTier(outputPath, tierMap); err != nil {
		log.Fatalf("Failed to save: %v", err)
	}

	fmt.Printf("%d elements was saved to %s.\n", len(tierMap), outputPath)
}
