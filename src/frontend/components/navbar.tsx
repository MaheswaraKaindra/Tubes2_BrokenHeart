"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import { cn } from "@/lib/utils"
import { Home, Search } from "lucide-react"

type NavItem = {
  label: string
  href: string
  icon?: React.ReactNode
}

const navItems: NavItem[] = [
  { label: "Home", href: "/", icon: <Home className="mr-1 h-4 w-4" /> },
  {
    label: "BFS", href: "/bfs", icon: <Search className="mr-1 h-4 w-4 text-purple-900" />
  },
  {
    label: "DFS", href: "/dfs", icon: <Search className="mr-1 h-4 w-4 text-purple-900" />
  },
]

export function Navbar() {
  const [activeItem, setActiveItem] = useState("Home")

  return (
    <header className="sticky top-0 z-50 w-full bg-[#FDF5E6] shadow-sm border border-red-500">
      <div className="flex h-24 items-center justify-between px-8">
        <Link href="/" className="flex items-center gap-2 text-black font-bold">
          <span className="text-2xl">💔</span>
          <span className="text-xl font-[var(--font-racing-sans)]">BrokenHeart</span>
        </Link>
        <nav className="flex items-center gap-4">
          <Link href="/" className="flex items-center text-fuchsia-900 text-lg font-medium gap-1">
            <Home className="w-5 h-5" />
            Home
          </Link>
          <Link href="/bfs" className="flex items-center text-fuchsia-900 text-lg font-medium gap-1">
            <Search className="w-5 h-5" />
            BFS
          </Link>
          <Link href="/dfs" className="flex items-center text-fuchsia-900 text-lg font-medium gap-1">
            <Search className="w-5 h-5" />
            DFS
          </Link>
        </nav>
      </div>
    </header>
  )
}