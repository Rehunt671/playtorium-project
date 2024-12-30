'use client';
import { useQueryUser } from '@/query/user';
import { User } from '@/types/user';
import { redirect, usePathname, useRouter } from 'next/navigation';
import React, { createContext, useContext, useState, ReactNode, useEffect } from 'react';

interface UserContextType {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
  logout: () => void;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

interface UserProviderProps {
  children: ReactNode;
}

export const UserProvider: React.FC<UserProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null); 
  const queryUser = useQueryUser(token ?? '');
  const pathname = usePathname(); 
  const router = useRouter();

  useEffect(() => {
    const storedToken = localStorage.getItem('authToken');
    if (storedToken) {
      setToken(storedToken);
    }else if(pathname !== "/login"){
      router.push("/login")
    }
  }, []);

  useEffect(() => {
    if (queryUser.data) {
      setUser(queryUser.data); 
    }
  }, [queryUser.data]);

  const logout = () => {
    localStorage.removeItem('authToken'); 
    setUser(null);
    setToken(null);
    redirect('/login');
  };

  return (
    <UserContext.Provider value={{ user, setUser, logout }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUserContext = (): UserContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUserContext must be used within a UserProvider');
  }
  return context;
};
