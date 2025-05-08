"use client"

import { useState } from "react"
import { Navbar } from "@/components/navbar"
import { Switch } from "@/components/ui/switch"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

export default function ResultPage() {
  const [searchMode, setSearchMode] = useState<"shortest" | "multiple">("shortest")
  const [searchQuery, setSearchQuery] = useState("")

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

        {/* Result section */}
        <div className="mb-8">
          <h2 className="mb-6 text-center font-poppins text-2xl font-semibold text-[#5C1A72]">
            Here is the Shortest Route to Make a..
          </h2>

          <div className="flex flex-col items-center justify-center">
            <div className="mb-2 flex h-24 w-24 items-center justify-center rounded-lg bg-[#E8DBF2] shadow-sm">
              {/* Element icon would go here */}
            </div>
            <p className="font-poppins text-lg italic text-gray-500">&lt;Element&gt;</p>
          </div>
        </div>

        {/* Recipe visualization container */}
        <div className="mb-8 overflow-hidden rounded-xl bg-[#EDE7F6] p-6 shadow-md">
          <div className="flex h-64 items-center justify-center rounded-lg bg-white/50 p-4">
            <p className="text-center font-poppins text-gray-400">Recipe visualization will appear here</p>
          </div>
        </div>
      </main>
    </div>
  )
}