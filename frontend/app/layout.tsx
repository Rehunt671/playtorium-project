"use client"
import React from "react";
import "@/app/globals.css";
import { TanstackProvider } from "@/providers/TanstackProvider";
import { ToastProvider } from "@/providers/ToastProvider";
import { UserProvider } from "@/providers/UserProvider";
import Navbar from "@/components/Navbar";
import { usePathname } from "next/navigation";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname(); 

  const showNavbar = pathname !== "/login" && pathname !== "/register" ;

  return (
    <html lang="en">
      <body>
        <TanstackProvider>
          <ToastProvider>
            <UserProvider>
              {showNavbar && <Navbar />}
              <div>{children}</div>
            </UserProvider>
          </ToastProvider>
        </TanstackProvider>
      </body>
    </html>
  );
}
