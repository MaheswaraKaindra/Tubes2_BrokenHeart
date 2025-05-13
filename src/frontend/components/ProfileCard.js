"use client";
import React from "react";
import Image from "next/image";
import Link from "next/link";

export default function ProfileCard({ data }) {
  return (
    <div>
      <Link
        href={`${data.link}`}
        target="_blank"
        rel="noopener noreferrer"
        passHref
      >
        <div
          className={`bg-${data.bg} hover:scale-105 transition-transform duration-300 ease-in-out sm:min-w-[550px] md:min-w-[750px] h-full mx-[50px] my-[50px] rounded-2xl py-[50px] px-10 md:px-20 shadow-dark-light flex flex-col md:flex-row items-center max-w-[1200px] justify-center drop-shadow-xl font-monts`}
        >
          <div className="flex-shrink-0 max-w-[350px] h-[250px] overflow-hidden rounded-2xl">
            <Image
              src={data.img}
              width={300}
              height={300}
              alt="Orang Baik"
              className="rounded-2xl object-cover w-full h-full"
            />
          </div>

          <div className="flex flex-col justify-center px-5 md:px-10 text-white">
            <div className="text-xl sm:text-2xl md:text-3xl lg:text-3xl font-bold py-5">
              {data.name}
            </div>
            <div className="text-lg sm:text-lg md:text-xl lg:text-2xl font-semibold">
              {data.nim}
            </div>
            <div className="text-sm sm:text-sm md:text-md lg:text-md">
              {data.pesan}
            </div>
          </div>
        </div>
      </Link>
    </div>
  );
}