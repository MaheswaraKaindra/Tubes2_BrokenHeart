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
  { label: "BFS", href: "/bfs", icon: <Search className="mr-1 h-4 w-4" /> },
  { label: "DFS", href: "/dfs", icon: <Search className="mr-1 h-4 w-4" /> },
]

export function Navbar() {
  const [activeItem, setActiveItem] = useState("Home")

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-[#FDF5E6] shadow-sm">
      <div className="container flex h-16 items-center justify-between">
        <Link href="/" className="flex items-center gap-2 font-bold text-black">
          <span className="text-2xl">💔</span>
          <span className="font-poppins text-xl">BrokenHeart</span>
        </Link>
        <nav className="flex items-center gap-2">
          {navItems.map((item) => (
            <Link
              key={item.label}
              href={item.href}
              onClick={() => setActiveItem(item.label)}
              className={cn(
                "flex items-center font-poppins rounded-full py-2 px-4 font-medium text-purple-900 transition-all hover:shadow-md",
                activeItem === item.label ? "bg-[#E3C7F3]" : "bg-purple-300",
              )}
            >
              {item.icon}
              {item.label}
            </Link>
          ))}
        </nav>
      </div>
    </header>
  )
}