"use client"

import { Plus } from "lucide-react"
import { cn } from "@/lib/utils"

interface CombinationCardProps {
  element1: string
  element2: string
  onClick: () => void
  isHighlighted?: boolean
}

export function CombinationCard({ element1, element2, onClick, isHighlighted = false }: CombinationCardProps) {
  return (
    <div
      className={cn(
        "flex items-center gap-2 rounded-xl p-2 transition-all hover:shadow-md cursor-pointer",
        isHighlighted ? "bg-[#F7E8DC]/80 shadow-md" : "bg-[#F7E8DC] shadow-sm",
      )}
      onClick={onClick}
    >
      <div className="flex h-16 w-24 items-center justify-center rounded-lg bg-white p-2 text-center font-poppins font-medium">
        {element1}
      </div>
      <div className="flex h-8 w-8 items-center justify-center rounded-full bg-[#5C1A72] text-white">
        <Plus className="h-5 w-5" />
      </div>
      <div className="flex h-16 w-24 items-center justify-center rounded-lg bg-white p-2 text-center font-poppins font-medium">
        {element2}
      </div>
    </div>
  )
}