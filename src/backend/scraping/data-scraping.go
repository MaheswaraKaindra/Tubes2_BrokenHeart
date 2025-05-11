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