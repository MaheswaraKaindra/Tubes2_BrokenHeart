import React, { useState } from 'react'
import { CombinationCard } from './CombinationCard'
import Dummy from '../public/water.png'

const DummyData = [
  {
    id: 1,
    name: 'Water',
    picture: Dummy,
    result: "result1",
  },
  {
    id: 2,
    name: 'Fire',
    picture: Dummy,
    result: "result2",
  },
  {
    id: 3,
    name: 'Earth',
    picture: Dummy,
    result: "result3",
  },
]

export const MultipleResultCard = () => {
  const [selectedIndex, setSelectedIndex] = useState(null)

  // if (!WholeResult) return null // hanya muncul jika parameter terpenuhi

  return (
    <div className='w-[800px] bg-purple-light rounded-2xl shadow-dark-light p-10 h-full flex flex-col items-center justify-center gap-10'>
      <div className='font-monts text-2xl font-bold text-purple-dark shadow-2xl'>Click one of the following to see the Tree of the recipe!</div>
      {selectedIndex === null ? (
        DummyData.map((item, index) => (
          <CombinationCard
            key={item.id}
            picture1={item.picture}
            name1={item.name}
            onClick={() => setSelectedIndex(index)}
          />
        ))
      ) : (
        <div 
        onClick={() => setSelectedIndex(null)}
        className='w-[500px] h-[500px] bg-purple-dark text-white rounded-2xl'>
          {/* nanti kasi button back woi onclicknya di button aja*/}
          {DummyData[selectedIndex].result}
        </div>
      )}
    </div>
  )
}

