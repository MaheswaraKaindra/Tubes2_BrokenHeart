import React from 'react'
import Image from 'next/image'
import Dummy from '../public/water.png'

export const ElementCard = ({ picture, name }) => {
  return (
    <div className='flex flex-col items-center justify-center gap-3 bg-cream-light rounded-4xl shadow-dark-light w-[150px] text-center h-full p-8 font-monts italic'>
      {picture && (
        <Image
          src={picture}
          alt="Dummy"
          width={100}
          height={100}
          className='rounded-2xl w-full shadow-sm'
        />        
      )}
      <div> {name ? name : "Invalid Element"} </div>
    </div>
  )
}

