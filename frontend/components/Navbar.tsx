"use client";
import { useUserContext } from "@/providers/UserProvider";
import Link from "next/link";

const Navbar: React.FC = () => {
  const { logout } = useUserContext();

  return (
    <nav className="bg-gradient-to-r from-blue-600 to-blue-800 sticky top-0 z-50 shadow-lg p-4">
      <div className="max-w-7xl mx-auto flex justify-between items-center">
        <div className="text-white text-2xl font-bold">
          <Link href="/store" className="hover:text-gray-300">
            Playtorium
          </Link>
        </div>
        <div className="hidden md:flex space-x-6">
          <Link
            href="/store"
            className="text-white hover:text-gray-300 transition-colors duration-300 font-poppins text-lg font-semibold"
          >
            Store
          </Link>
          <Link
            href="/cart"
            className="text-white hover:text-gray-300 transition-colors duration-300 font-poppins text-lg font-semibold"
          >
            My Cart
          </Link>
          <Link
            href="/campaign"
            className="text-white hover:text-gray-300 transition-colors duration-300 font-poppins text-lg font-semibold"
          >
            Campaigns
          </Link>
          {/* Add logout button */}
          <button
            onClick={logout}
            className="text-white hover:text-gray-300 transition-colors duration-300 font-poppins text-lg font-semibold"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
