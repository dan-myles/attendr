import { ClerkProvider } from "@clerk/nextjs/dist/components.server";
import "./globals.css";
import { Inter } from "next/font/google";
import Image from "next/image";
import TopButtons from "../components/TopButtons";
import styles from "./layout.module.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "attendr",
  description: "Never miss a class again!",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ClerkProvider>
      <html lang="en">
        <body className={inter.className}>
          <div className="flex justify-between">
            <div>
              <a href="/">
                <Image
                  src="/assets/images/logo.png"
                  alt="attendr Logo"
                  width={160}
                  height={160}
                  style={{ objectFit: "contain" }}
                />
              </a>
            </div>
            <TopButtons />
          </div>
          {children}
        </body>
      </html>
    </ClerkProvider>
  );
}
