import React from 'react'
import { ElementCard } from './ElementCard'
import Plus from '../public/plus.png'
import Image from 'next/image'

export const CombinationCard = ({onClick, picture1, name1, picture2, name2}) => {
  return (
    <div 
    onClick={onClick}
    className='flex items-center justify-center gap-10 font-monts bg-purple-dark-light rounded-4xl shadow-dark-light px-15 py-8 hover:scale-105 transition-transform duration-300 ease-in-out'>
      <ElementCard picture = {picture1 ? picture1 : null} name = {name1 ? name1 : null} />
      <Image
      src = {Plus}
      alt="Plus Sign"
      width={50}
      />
      <ElementCard picture = {picture2 ? picture2 : null} name = {name2 ? name2 : null}/>
    </div>
  )
}
