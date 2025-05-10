import { Racing_Sans_One, Montserrat_Alternates } from "next/font/google";
import "./globals.css";
import Navbar from "@/components/Navbar";

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
        className={`${montserrat.variable} ${racingSans.variable} antialiased`}
      >
        <Navbar />
          <div>
            {children}          
          </div>
      </body>
    </html>
  );
}
