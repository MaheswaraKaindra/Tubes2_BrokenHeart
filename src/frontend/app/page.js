import Image from "next/image";
import Link from "next/link";
import ProfileCard from "@/components/ProfileCard";


const Profile = [
  {
    name: "Maheswara Bayu Kaindra",
    nim: "13523015",
    pesan: "''Ini momen pertamaku mengeksplorasi hal yang sangat jauh dari pengetahuanku, menurutku tugas ini workloadnya terlalu berat, jauh melebihi batas normal. Meskipun begitu, tugas masih bisa diselesaikan tepat waktu, karena aku jago.''",
    img: "/fotoindra.jpg",
    bg: "purple-light",
    link:"https://www.linkedin.com/in/maheswarakaindra/"
  },
  {
    name: "Jessica Allen",
    nim: "13523059",
    pesan: "''Ternyata handling API and frontend bersamaan sangat seru.''",
    img: "/fotoallen.jpg",
    bg: "purple-dark-light",
    link:"https://www.linkedin.com/in/jessica-allen-lim/"
  },
  {
    name: "Shanice Feodora Tjahjono",
    nim: "13523097",
    pesan: "''Pergi ke pasar liat orang terbang, puji Tuhan kelar dan ga tumbang.''",
    img: "/fotoshanice.jpg",
    bg: "purple-dark",
    link:"https://www.linkedin.com/in/shanice-feodora-a8368124b/"
  }
]

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

      <div className="p-10 max-w-[1000px] flex flex-col items-center mx-auto">
        <div className="mt-20 text-4xl font-monts font-extrabold text-center mb-4 text-purple-dark p-5">Brought to you by...</div>
        <div className="flex flex-col justify-center items-center ">
          {Profile.map((person, index) => (
            <ProfileCard key={index} data={person} />
          ))}          
        </div>

      </div>
      </main>
    </div>
  );
}
