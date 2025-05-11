"use client"
import { useState } from 'react'
import React from 'react'
import { ElementCard } from '@/components/ElementCard'
import { MultipleResultCard } from './MultipleResultCard'
import { CombinationCard } from '@/components/CombinationCard'
import Image from 'next/image'
import Dummy from '../public/water.png'

export const SearchPage = () => {
  const [isToggled, setIsToggled] = useState(false);
  const [isClicked, setIsClicked] = useState(false);
  const [input, setInput] = useState("");
const [result, setResult] = useState(null);

const handleSubmit = async () => {
  const endpoint = isToggled ? "dfs" : "bfs";
  try {
    const res = await fetch(`http://localhost:8080/api/${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ target: input, index: 0 }), // Add index if needed
    });

    if (!res.ok) {
      const errorText = await res.text();
      console.error("Server error:", errorText);
      alert("Server error: " + errorText);
      return;
    }

    const data = await res.json();
    setResult(data);
    console.log("Tree:", data);
  } catch (err) {
    console.error("Failed to fetch:", err);
  }
};


  const handleToggle = () => {
    setIsToggled(!isToggled);
  };

  return (
    <div className='flex flex-col items-center h-full gap-10 py-10 px-10'>
      {/* titlenya */}
      <div className='font-monts text-4xl font-bold text-purple-dark'>Enter Your Element!</div>

      {/* input */}
        <form
          className="flex items-center justify-center gap-5 w-[800px] font-monts"
          onSubmit={(e) => {
            e.preventDefault();
            handleSubmit(); 
          }}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Enter your element"
            className="rounded-full w-full p-2 text-center bg-orange-bright shadow-orange focus:outline-none focus:ring-0"
          />

          <button
            type="submit"
            className="bg-purple-dark text-white font-monts font-bold px-4 py-2 rounded-full shadow-lg hover:scale-105 transition-transform duration-300 ease-in-out focus:outline-none focus:border-2 focus:border-[#380028]"
          >
            Go!
          </button>
        </form>

        {/* toggle */}
      <div className='flex justify-center items-center gap-5 w-full font-monts text-purple-dark'>
        <div className='flex items-center justify-center gap-2'>
          <div className='text-center '>Shortest Route</div>
          <button
            onClick={handleToggle}
            className={`w-14 h-8 flex items-center rounded-full p-1 transition-colors duration-300 ${
              isToggled ? "bg-purple-dark" : "bg-purple-light"
            } focus:outline-none`}
          >
            <div
              className={`bg-white w-6 h-6 rounded-full shadow-md transform transition-transform duration-300 ${
                isToggled ? "translate-x-6" : "translate-x-0"
              }`}
            />
          </button>
          <div className='text-center text-sm'>Multiple Recipes</div>            
        </div>
        {isToggled &&
        <div>
          <form
            className="flex items-center justify-center gap-5 w-full font-monts"
            onSubmit={(e) => {
              e.preventDefault();
              // Nanti masuk logic search
            }}>
            <input
              type="text"
              placeholder="Enter max. results"
              className="rounded-full max-w-[300px] p-2 text-center text-sm bg-purple-light shadow-lg focus:outline-none focus:ring-0"
            />
            <button
              type="submit"
              className="bg-purple-dark text-white text-sm font-monts px-4 py-2 rounded-full shadow-lg hover:scale-105 transition-transform duration-300 ease-in-out focus:outline-none focus:border-2 focus:border-[#380028]"
            >
              Submit
            </button>
          </form> 
        </div>}
      </div>


      {/* result */}
      <div className='w-full font-bold flex flex-col items-center justify-center gap-10 font-monts text-purple-dark mt-10 '>
          <div className=' px-10 py-5 text-center flex flex-col gap-5 items-center justify-center w-full h-full rounded-2xl bg-purple-light shadow-dark-light text-2xl'>   
            Here are the recipes to find: 
          </div>
          {result && (
            <div className='mb-10'>
              <ElementCard 
                picture={Dummy}
                name={result.Name}
              />
            </div>
          )}

          {/* kalau shortest, lgsg tampilin result treenya */}
          {!isToggled && (
            <div></div>
          )}

          {/* kalau multiple, lgsg tampilin result treenya */}
          {isToggled &&(
          <MultipleResultCard/>)
          }

      </div>
    </div>
  )
}