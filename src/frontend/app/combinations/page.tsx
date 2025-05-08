"use client"

import { useState } from "react"
import { Navbar } from "@/components/navbar"
import { Switch } from "@/components/ui/switch"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { CombinationCard } from "@/components/combination-card"
import { RecipesList } from "@/components/recipes-list"

// Sample data
const combinations = [
  { id: 1, element1: "Fire", element2: "Water" },
  { id: 2, element1: "Earth", element2: "Air" },
  { id: 3, element1: "Water", element2: "Earth" },
]

const recipes = [
  { id: 1, steps: ["Fire", "Steam", "Cloud"] },
  { id: 2, steps: ["Fire", "Energy", "Steam", "Cloud"] },
  { id: 3, steps: ["Water", "Pressure", "Steam", "Cloud"] },
]

export default function CombinationsPage() {
  const [searchMode, setSearchMode] = useState<"shortest" | "multiple">("multiple")
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCombination, setSelectedCombination] = useState<null | { element1: string; element2: string }>(null)

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
            />
          </div>
          <Button className="h-12 rounded-full bg-[#FFDC73] px-6 font-medium text-black hover:bg-[#FFDC73]/90">
            Go!
          </Button>
        </div>

        {/* Toggle control */}
        <div className="mb-8 flex flex-col items-center justify-between gap-4 sm:flex-row">
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
        </div>

        {/* Content based on selection */}
        {selectedCombination ? (
          <div>
            <h2 className="mb-6 text-center font-poppins text-2xl font-bold text-[#5C1A72]">
              Here are All the Recipes for..
            </h2>

            {/* Selected combination card */}
            <div className="mb-6 flex justify-center">
              <CombinationCard
                element1={selectedCombination.element1}
                element2={selectedCombination.element2}
                onClick={() => {}}
                isHighlighted
              />
            </div>

            {/* Recipes list */}
            <div className="mb-8 overflow-hidden rounded-xl bg-[#EDE7F6] p-6 shadow-md">
              <RecipesList recipes={recipes} />
            </div>

            <div className="flex justify-center">
              <Button variant="outline" className="font-poppins" onClick={() => setSelectedCombination(null)}>
                Back to Combinations
              </Button>
            </div>
          </div>
        ) : (
          <div>
            <h2 className="mb-6 text-center font-poppins text-2xl font-bold text-[#5C1A72]">
              Here are All the Combinations to Make a..
            </h2>

            <div className="mb-4 flex justify-center">
              <div className="flex h-24 w-24 items-center justify-center rounded-lg bg-[#E8DBF2] shadow-sm">
                {/* Element icon would go here */}
              </div>
            </div>
            <p className="mb-8 text-center font-poppins text-lg italic text-gray-500">&lt;Element&gt;</p>

            <p className="mb-6 text-center font-poppins text-sm text-[#5C1A72]">
              Click one of the combinations to see their multiple recipes!
            </p>

            <div className="flex flex-col items-center gap-4">
              {combinations.map((combo) => (
                <CombinationCard
                  key={combo.id}
                  element1={combo.element1}
                  element2={combo.element2}
                  onClick={() => setSelectedCombination({ element1: combo.element1, element2: combo.element2 })}
                />
              ))}
            </div>
          </div>
        )}
      </main>
    </div>
  )
}