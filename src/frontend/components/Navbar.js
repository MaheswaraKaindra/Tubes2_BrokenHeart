'use client'
import React from 'react'
import '../app/globals.css'
import Image from 'next/image'
import Link from 'next/link'
import Logo from '../public/brokenheart.png'
import TextLogo from '../public/textlogo.png'
import Home1 from '../public/home1.png'
import Home2 from '../public/home2.png'
import Search1 from '../public/search1.png'
import Search2 from '../public/search2.png'
import { usePathname } from 'next/navigation'
import Home from '@/app/page'

const Navbar = () => {
    const pathname = usePathname();
  return (
    <div className='h-[60px] bg-cream-light sticky top-0 opacity-80'>
      <div className='flex justify-between items-center px-10'>
        {/* logo */}
        <div className='flex items-center justify-center my-3'>
          <Image
          src ={TextLogo}
          alt="Logo"
          width={180}
          />
        </div>

        {/* buttons */}
        <div className='flex items-center justify-between gap-5 text-lg'>
          <Link href="/">
            <button className={`hover:scale-105 transition-transform duration-300 ease-in-out flex items-center justify-center px-5 rounded-full hover:bg-purple-light hover:text-purple-dark  ${pathname ==='/' ? 'bg-purple-dark text-white' : 'bg-cream-light text-purple-dark'}`}>
              <Image
                src={ pathname === '/' ? Home2 : Home1}
                alt="Home"
                width={20}
                className='mr-3 my-3 hover:scale-110 transition-transform duration-300'
              />
              <div className='font-racing'> Home </div>
            </button>
          </Link>
          <Link href="/bfs">
             <button className={`hover:scale-105 transition-transform duration-300 ease-in-out flex items-center justify-center px-5 rounded-full hover:bg-purple-light hover:text-purple-dark  ${pathname ==='/bfs' ? 'bg-purple-dark text-white' : 'bg-cream-light text-purple-dark'}`}>
              <Image
                src={ pathname === '/bfs' ? Search2 : Search1}
                alt="Search"
                width={20}
                className='mr-3 my-3 hover:scale-110 transition-transform duration-300'
              />
              <div className='font-racing'> BFS </div>
            </button>  
          </Link>  
            <Link href="/dfs">
            <button className={`hover:scale-105 transition-transform duration-300 ease-in-out flex items-center justify-center px-5 rounded-full hover:bg-purple-light hover:text-purple-dark  ${pathname ==='/dfs' ? 'bg-purple-dark text-white' : 'bg-cream-light text-purple-dark'}`}>
              <Image
                src={ pathname === '/dfs' ? Search2 : Search1}
                alt="Search"
                width={20}
                className='mr-3 my-3 hover:scale-110 transition-transform duration-300'
              />
              <div className='font-racing'> DFS </div>
            </button>  
          </Link>        
        </div>

      </div>
    </div>
  )
}

export default Navbar
