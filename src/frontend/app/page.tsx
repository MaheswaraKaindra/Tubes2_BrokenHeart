import Link from "next/link"

export default function Home() {
  return (
    <div className="relative min-h-screen overflow-hidden bg-white px-4 py-12">
      {/* Decorative shapes */}
      <div className="absolute -left-20 top-40 h-96 w-96 rounded-full bg-[#FADADD] opacity-70 blur-3xl"></div>
      <div className="absolute -right-20 top-60 h-96 w-96 rounded-full bg-[#FFDC73] opacity-70 blur-3xl"></div>
      <div className="absolute -left-20 bottom-20 h-64 w-64 rounded-full bg-[#FFDC73] opacity-70 blur-3xl"></div>
      <div className="absolute -right-20 bottom-40 h-64 w-64 rounded-full bg-[#FADADD] opacity-70 blur-3xl"></div>

      <div className="container relative mx-auto flex max-w-5xl flex-col items-center justify-center">
        <h1 className="mb-8 text-center font-poppins text-5xl font-bold text-[#5C1A72] md:text-6xl">
          Find Your Recipe!
        </h1>

        <div className="mb-16 rounded-md bg-[#FFDC73] px-6 py-3 text-center font-poppins text-xl font-medium text-black shadow-md">
          Choose a Search Algorithm
        </div>

        <div className="grid w-full grid-cols-1 gap-6 md:grid-cols-3">
          <Link href="/bfs" className="group">
            <div className="flex h-full flex-col items-center justify-center rounded-xl bg-[#F7E8DC] p-8 text-center shadow-md transition-all duration-300 hover:shadow-lg">
              <h2 className="mb-2 font-poppins text-3xl font-bold text-[#5F3B3B]">BFS</h2>
              <p className="font-poppins text-lg text-[#5F3B3B]">Breadth First Search</p>
            </div>
          </Link>

          <Link href="/dfs" className="group">
            <div className="flex h-full flex-col items-center justify-center rounded-xl bg-[#F7E8DC] p-8 text-center shadow-md transition-all duration-300 hover:shadow-lg">
              <h2 className="mb-2 font-poppins text-3xl font-bold text-[#5F3B3B]">DFS</h2>
              <p className="font-poppins text-lg text-[#5F3B3B]">Depth First Search</p>
            </div>
          </Link>

          <Link href="/bisearch" className="group">
            <div className="flex h-full flex-col items-center justify-center rounded-xl bg-[#F7E8DC] p-8 text-center shadow-md transition-all duration-300 hover:shadow-lg">
              <h2 className="mb-2 font-poppins text-3xl font-bold text-[#5F3B3B]">BiSearch</h2>
              <p className="font-poppins text-lg text-[#5F3B3B]">Bidirectional Search</p>
            </div>
          </Link>
        </div>
      </div>
    </div>
  )
}