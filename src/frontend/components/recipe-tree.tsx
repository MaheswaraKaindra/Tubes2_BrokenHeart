"use client"

type RecipeNode = {
  name: string
  children?: RecipeNode[]
  isBaseElement?: boolean
}

interface RecipeTreeProps {
  recipe: RecipeNode
}

export function RecipeTree({ recipe }: RecipeTreeProps) {
  return (
    <div className="flex flex-col items-center">
      <RecipeNode node={recipe} isRoot={true} />
    </div>
  )
}

interface RecipeNodeProps {
  node: RecipeNode
  isRoot?: boolean
}

function RecipeNode({ node, isRoot = false }: RecipeNodeProps) {
  const hasChildren = node.children && node.children.length > 0

  return (
    <div className="flex flex-col items-center">
      {/* Element card */}
      <div
        className={`mb-2 flex h-16 w-32 items-center justify-center rounded-lg ${
          node.isBaseElement
            ? "bg-gradient-to-r from-blue-100 to-green-100 shadow-sm"
            : isRoot
              ? "bg-gradient-to-r from-purple-100 to-pink-100 shadow-md"
              : "bg-[#E8DBF2] shadow-sm"
        } p-2 text-center font-poppins font-medium transition-all hover:shadow-md`}
      >
        {node.name}
      </div>

      {/* Connection lines and children */}
      {hasChildren && (
        <>
          {/* Vertical connection line */}
          <div className="h-6 w-0.5 bg-gray-300"></div>

          {/* Horizontal line connecting children */}
          <div className="relative flex w-full justify-center">
            <div className="absolute top-0 h-0.5 w-40 bg-gray-300"></div>

            {/* Children nodes */}
            <div className="mt-6 flex w-full justify-center gap-16">
              {node.children!.map((child, index) => (
                <RecipeNode key={`${child.name}-${index}`} node={child} />
              ))}
            </div>
          </div>
        </>
      )}
    </div>
  )
}