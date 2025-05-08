import { ChevronRight } from "lucide-react"

interface RecipesListProps {
  recipes: {
    id: number
    steps: string[]
  }[]
}

export function RecipesList({ recipes }: RecipesListProps) {
  return (
    <div className="flex max-h-80 flex-col gap-4 overflow-y-auto pr-2">
      {recipes.map((recipe) => (
        <div key={recipe.id} className="flex flex-wrap items-center gap-2 rounded-lg bg-white p-3 shadow-sm">
          {recipe.steps.map((step, index) => (
            <div key={index} className="flex items-center">
              <div className="rounded-full bg-[#E8DBF2] px-4 py-1 font-poppins text-sm font-medium">{step}</div>
              {index < recipe.steps.length - 1 && <ChevronRight className="mx-1 h-4 w-4 text-gray-400" />}
            </div>
          ))}
        </div>
      ))}
    </div>
  )
}