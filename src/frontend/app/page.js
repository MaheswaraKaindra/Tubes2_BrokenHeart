import Image from "next/image";
import Link from "next/link";
import TreeVisualizer from "@/components/TreeVisualizer";


const dummyTree = {
  name: 'Root',
  left: {
    name: 'Left Child',
    left: {
      name: 'Left-Left Child',
    },
    right: {
      name: 'Left-Right Child',
    },
  },
  right: {
    name: 'Right Child',
    left: {
      name: 'Right-Left Child',
    },
    right: {
      name: 'Right-Right Child',
    },
  },
};


export default function Home() {
  return (
    <div className="w-full h-full flex flex-col items-center ">
      <main className="flex flex-col">
        <div className="flex flex-col items-center justify-center w-full h-[350px] gap-10">
          <h1 className='text-7xl text-purple-dark font-racing'>Find Your Recipe!</h1>
          <h2 className="bg-orange-bright rounded-full px-8 py-2 shadow-orange font-monts font-bold text-purple-dark " >Choose a Search Algorithm!</h2>
        </div>
        <div className="flex flex-row items-center justify-center gap-8">
          <Link href="/bfs">
            <button className="flex flex-col items-center justify-center gap-2 bg-cream-light text-purple-dark font-monts text-2xl px-8 py-4 rounded-4xl shadow-dark-light hover:scale-105 transition-transform duration-300 ease-in-out h-[150px]">
              <div className="text-5xl font-bold ">BFS</div>
              <div>Breadth First Search</div>
            </button>
          </Link>
          <Link href="/dfs">
            <button className="flex flex-col items-center justify-center gap-2 bg-cream-light text-purple-dark font-monts text-2xl px-8 py-4 rounded-4xl shadow-dark-light hover:scale-105 transition-transform duration-300 ease-in-out h-[150px]">
              <div className="text-5xl font-bold ">DFS</div>
              <div>Depth First Search</div>
            </button>
          </Link>
        </div>

      <TreeVisualizer tree={dummyTree} />


      </main>
    </div>
  );
}
