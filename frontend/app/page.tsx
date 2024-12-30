'use client';
import { useRouter } from 'next/navigation';
import React, { useEffect } from 'react';

const App: React.FC = () => {
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem('authToken');

    if (token) {
      router.push('/store');
    } else {
      router.push('/login');
    }
  }, [router]); 

  return <></>
};

export default App;
