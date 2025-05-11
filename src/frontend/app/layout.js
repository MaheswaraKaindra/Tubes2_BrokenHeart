import { Racing_Sans_One, Montserrat_Alternates } from "next/font/google";
import "./globals.css";
import Navbar from "@/components/Navbar";
import Image from "next/image";
import Logo from "../public/brokenheart.png";

const racingSans = Racing_Sans_One({
  variable: "--font-racing-sans",
  subsets: ["latin"],
  weight: ["400"],
});

const montserrat = Montserrat_Alternates({
  variable: "--font-sans",
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
});

export const metadata = {
  title: "Broken Heart",
  description: "Recipe finder for Little Alchemy 2",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body
        className={`${montserrat.variable} ${racingSans.variable} antialiased overflow-x-hidden `}
      >
        <Image
          src={Logo}
          alt="Logo"
          width={100}
          height={100}
          className="absolute top-120 -left-15 opacity-30 -z-20 rotate-30 w-[350px]"
        />

        <Image
          src={Logo}
          alt="Logo"
          width={100}
          height={100}
          className="absolute top-30 -right-15 opacity-30 -z-20 -rotate-30 w-[200px]"
        />

        <Navbar />
          <div>
            {children}          
          </div>
      </body>
    </html>
  );
}
