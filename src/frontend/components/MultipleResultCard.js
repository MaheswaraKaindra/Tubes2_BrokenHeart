import React, { useState } from 'react';
import { CombinationCard } from './CombinationCard';
import Dummy from '../public/water.png';
import TreeVisualizer from './TreeVisualizer';

export const MultipleResultCard = ({ result, max }) => {
  const [selectedIndex, setSelectedIndex] = useState(null);

  return (
    <div className='w-[800px] bg-purple-light rounded-2xl shadow-dark-light p-10 h-full flex flex-col items-center justify-center gap-10'>
      {selectedIndex === null && Array.isArray(result?.recipes) ? (
        <>
          <div className='text-center font-monts text-2xl font-bold text-purple-dark'>
            Click one of the following to see the visualized Tree of the selected recipe!
          </div>
          {result.recipes.slice(0, max ?? result.recipes.length).map((item, index) => (
            <CombinationCard
              key={item.id || index}
              name1={item.Component1}
              name2={item.Component2}
              picture1={`/data/${item.Component1}.svg`}
              picture2={`/data/${item.Component2}.svg`}
              onClick={() => setSelectedIndex(index)}
            />
          ))}
        </>
      ) : (
        <div className='w-full bg-purple-dark text-2xl rounded-2xl p-8 flex flex-col items-center justify-center gap-10'>
          {selectedIndex !== null && (
            <>
            <div className='w-full h-full flex flex-col gap-10 items-center justify-center'>
              <div className=' text-white'>Here is the tree from the Combination:</div>
              <CombinationCard
                name1 = {result.recipes[selectedIndex].Component1}
                name2 = {result.recipes[selectedIndex].Component2}
                picture1={`/data/${result.recipes[selectedIndex].Component1}.svg`}
                picture2={`/data/${result.recipes[selectedIndex].Component2}.svg`}
              />
                {result?.trees && result.trees[selectedIndex] ? (
                  <TreeVisualizer
                    key={selectedIndex}
                    tree={structuredClone(result.trees[selectedIndex])}
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
  );
};
