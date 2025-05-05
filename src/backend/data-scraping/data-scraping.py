import requests
from bs4 import BeautifulSoup
import json

URL = "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"
HEADERS = {"User-Agent": "Mozilla/5.0"}

response = requests.get(URL, headers=HEADERS)
soup = BeautifulSoup(response.content, "html.parser")

elements = {}

element_links = soup.select("table a[href^='/wiki/']")

visited = set()

for link in element_links:
    href = link.get("href")
    name = link.get_text().strip()

    if not href.startswith("/wiki/") or name in visited or ":" in href:
        continue

    visited.add(name)
    page_url = f"https://little-alchemy.fandom.com{href}"
    page = requests.get(page_url, headers=HEADERS)
    psoup = BeautifulSoup(page.content, "html.parser")

    recipe_section = psoup.find("span", id="Combinations")
    if recipe_section:
        ul = recipe_section.find_next("ul")
        if ul:
            recipes = []
            for li in ul.find_all("li"):
                combo = [x.strip() for x in li.get_text(separator="+").split("+")]
                if len(combo) == 2:
                    recipes.append(combo)
            if recipes:
                elements[name] = recipes

with open("recipes.json", "w", encoding="utf-8") as f:
    json.dump(elements, f, indent=2, ensure_ascii=False)

print("Done.")
