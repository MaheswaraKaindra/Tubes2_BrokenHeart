import React, { useState } from 'react'
import { CombinationCard } from './CombinationCard'
import Dummy from '../public/water.png'
import TreeVisualizer from './TreeVisualizer'

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

export const MultipleResultCard = ({result}) => {
  const [selectedIndex, setSelectedIndex] = useState(null)
      console.log(result?.recipes);

  return (
    <div className='w-[800px] bg-purple-light rounded-2xl shadow-dark-light p-10 h-full flex flex-col items-center justify-center gap-10'>
      <div className='text-center font-monts text-2xl font-bold text-purple-dark'>Click one of the following to see the Tree of the selected recipe!</div>
      {selectedIndex === null && Array.isArray(result?.recipes) ? (
        result.recipes.map((item, index) => (
          <CombinationCard
            key={item.id || index}
            name1={item.Component1}
            name2={item.Component2}
            onClick={() => setSelectedIndex(index)}
          />
        ))
      ) : (
        <div className='w-full bg-purple-dark text-white rounded-2xl p-4 flex flex-col items-center justify-center gap-4'>
          {selectedIndex !== null && (
            <>
              <div className='w-full h-full flex items-center justify-center'>
                {result?.trees && result.trees[selectedIndex] ? (
                  <TreeVisualizer
                    tree={result.trees[selectedIndex]}
                  />
                ) : (
                  <div className='text-center text-white'>Tree data not available</div>
                )}
              </div>

              <button 
                onClick={() => setSelectedIndex(null)} 
                className='bg-white text-purple-dark px-4 py-2 rounded-lg font-semibold'
              >
                Back
              </button>
            </>
          )}
        </div>
      )}
    </div>
  )
}

