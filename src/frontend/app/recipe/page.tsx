"use client"

import { useState } from "react"
import { Navbar } from "@/components/navbar"
import { Switch } from "@/components/ui/switch"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { RecipeTree } from "@/components/recipe-tree"

export default function RecipePage() {
  const [searchMode, setSearchMode] = useState<"shortest" | "multiple">("shortest")
  const [searchQuery, setSearchQuery] = useState("")
  const [isSearched, setIsSearched] = useState(false)

  // Example recipe data - in a real app, this would come from an API call
  const recipeData = {
    targetElement: "Metal",
    timeTaken: 0.23,
    nodesVisited: 73,
    recipe: {
      name: "Metal",
      children: [
        {
          name: "Fire",
          isBaseElement: true,
        },
        {
          name: "Stone",
          children: [
            {
              name: "Earth",
              isBaseElement: true,
            },
            {
              name: "Air",
              isBaseElement: true,
            },
          ],
        },
      ],
    },
  }

  const handleSearch = () => {
    if (searchQuery.trim()) {
      setIsSearched(true)
    }
  }

  return (
    <div className="relative min-h-screen overflow-hidden bg-[#F9F9F9]">
      {/* Decorative shapes */}
      <div className="absolute -left-20 top-40 h-96 w-96 rounded-full bg-[#FADADD] opacity-70 blur-3xl"></div>
      <div className="absolute -right-20 top-60 h-96 w-96 rounded-full bg-[#FFDC73] opacity-70 blur-3xl"></div>
      <div className="absolute -left-20 bottom-20 h-64 w-64 rounded-full bg-[#FFDC73] opacity-70 blur-3xl"></div>
      <div className="absolute -right-20 bottom-40 h-64 w-64 rounded-full bg-[#FADADD] opacity-70 blur-3xl"></div>

      <Navbar />

      <main className="container relative mx-auto max-w-4xl px-4 py-8">
        <h1 className="mb-8 text-center font-poppins text-3xl font-bold text-[#5C1A72] md:text-4xl">
          Enter your Element
        </h1>

        {/* Search bar */}
        <div className="mb-6 flex w-full flex-col items-center justify-center gap-2 sm:flex-row">
          <div className="relative w-full max-w-md">
            <Input
              type="text"
              placeholder="Search..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="h-12 rounded-full border-0 bg-[#FFDC73] px-6 text-black placeholder:text-gray-700 focus-visible:ring-2 focus-visible:ring-[#5C1A72]"
              onKeyDown={(e) => e.key === "Enter" && handleSearch()}
            />
          </div>
          <Button
            className="h-12 rounded-full bg-[#FFDC73] px-6 font-medium text-black hover:bg-[#FFDC73]/90"
            onClick={handleSearch}
          >
            Go!
          </Button>
        </div>

        {/* Toggle control */}
        <div className="mb-6 flex flex-col items-center justify-between gap-4 sm:flex-row">
          <div className="flex items-center gap-4">
            <span className={`font-poppins ${searchMode === "shortest" ? "font-medium" : "text-gray-500"}`}>
              Shortest Route
            </span>
            <Switch
              checked={searchMode === "multiple"}
              onCheckedChange={(checked) => setSearchMode(checked ? "multiple" : "shortest")}
              className="data-[state=checked]:bg-[#5C1A72]"
            />
            <span className={`font-poppins ${searchMode === "multiple" ? "font-medium" : "text-gray-500"}`}>
              Multiple Recipes
            </span>
          </div>

          {searchMode === "multiple" && (
            <div className="w-full sm:w-auto">
              <Input
                type="number"
                placeholder="Enter maximum recipe results..."
                className="h-10 rounded-md border-[#E8DBF2] bg-[#F3EBFA] text-sm placeholder:text-gray-500 focus-visible:ring-2 focus-visible:ring-[#5C1A72]"
              />
            </div>
          )}
        </div>

        {isSearched && (
          <>
            <h2 className="mb-6 text-center font-poppins text-2xl font-semibold text-[#5C1A72]">
              Here is the Shortest Route to Make a {recipeData.targetElement}
            </h2>

            <div className="mb-8 rounded-xl bg-white p-6 shadow-md">
              <RecipeTree recipe={recipeData.recipe} />
            </div>

            <div className="flex flex-col items-center justify-center gap-4 rounded-lg bg-white p-4 text-center shadow-sm sm:flex-row sm:justify-between">
              <div className="flex items-center gap-2 font-poppins text-gray-700">
                <span className="text-xl">⏱️</span>
                <span>Time Taken: {recipeData.timeTaken}s</span>
              </div>
              <div className="flex items-center gap-2 font-poppins text-gray-700">
                <span className="text-xl">🧠</span>
                <span>Nodes Visited: {recipeData.nodesVisited}</span>
              </div>
            </div>
          </>
        )}
      </main>
    </div>
  )
}