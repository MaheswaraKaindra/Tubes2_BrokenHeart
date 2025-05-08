import Link from "next/link"

export default function Home() {
  return (
    <div className="relative min-h-screen bg-white py-16 px-4 overflow-hidden">
      {/* Decorative background */}
      <div className="absolute -left-40 top-40 h-[30rem] w-[30rem] rounded-full bg-[#FFDC73] opacity-60 blur-[100px] z-0" />
      <div className="absolute -right-40 bottom-20 h-[30rem] w-[30rem] rounded-full bg-[#FADADD] opacity-60 blur-[100px] z-0" />

      <div className="relative z-10 mx-auto max-w-6xl text-center">
        <h1 className="mb-6 text-5xl font-bold text-primary font-racing">
          Find Your Recipe!
        </h1>

        <div className="mb-12 inline-block rounded-full bg-yellow-400 px-6 py-3 text-lg font-semibold text-primary shadow-lg">
          Choose a Search Algorithm
        </div>

        {/* Grid container */}
        <div className="flex flex-wrap justify-center gap-8">
          {/* BFS */}
          <Link href="/bfs">
            <div className="w-[250px] cursor-pointer rounded-2xl bg-[#f5efe7] p-8 text-center shadow-md transition hover:scale-105 hover:shadow-xl">
              <h2 className="mb-2 text-3xl font-bold text-[#5C1A72]">BFS</h2>
              <p className="text-[#5C1A72]">Breadth First Search</p>
            </div>
          </Link>

          {/* DFS */}
          <Link href="/dfs">
            <div className="w-[250px] cursor-pointer rounded-2xl bg-[#f5efe7] p-8 text-center shadow-md transition hover:scale-105 hover:shadow-xl">
              <h2 className="mb-2 text-3xl font-bold text-[#5C1A72]">DFS</h2>
              <p className="text-[#5C1A72]">Depth First Search</p>
            </div>
          </Link>

          {/* BiSearch */}
          <Link href="/bisearch">
            <div className="w-[250px] cursor-pointer rounded-2xl bg-[#f5efe7] p-8 text-center shadow-md transition hover:scale-105 hover:shadow-xl">
              <h2 className="mb-2 text-3xl font-bold text-[#5C1A72]">BiSearch</h2>
              <p className="text-[#5C1A72]">Bidirectional Search</p>
            </div>
          </Link>
        </div>
      </div>
    </div>
  )
}