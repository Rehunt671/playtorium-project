'use client';

import React, { useState, useEffect } from 'react';
import { useMutationLogin } from '@/query/auth';
import { useRouter } from 'next/navigation';
import { useUserContext } from '@/providers/UserProvider';
import { useQueryUser } from '@/query/user';
import Link from 'next/link';
import Spinner from '@/components/Spinner';

const LoginPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const mutationLogin = useMutationLogin();
  const { setUser } = useUserContext();
  const router = useRouter();

  const storedToken = typeof window !== 'undefined' ? localStorage.getItem('authToken') : null;
  const queryUser = useQueryUser(storedToken ?? '');

  useEffect(() => {
    if (storedToken && queryUser.data) {
      setUser(queryUser.data);
    }
  }, [storedToken, queryUser.data, setUser]);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    try {
      const { token } = await mutationLogin.mutateAsync({ username, password });
      localStorage.setItem('authToken', token);

      if (token) {
        await queryUser.refetch();
        router.push('/store');
      }
    } catch (err) {
      setError('Invalid username or password');
    }
  };

  if (queryUser.isLoading) {
    return <Spinner />;
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="max-w-md w-full bg-white shadow-md rounded-lg p-8">
        <h1 className="text-2xl font-bold text-center text-gray-800 mb-6">Login</h1>
        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label htmlFor="username" className="block text-sm font-medium text-gray-700">
              Username
            </label>
            <input
              id="username"
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password
            </label>
            <input
              id="password"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>
          {error && <p className="text-red-600 text-sm mb-4">{error}</p>}
          <button
            type="submit"
            className={`w-full py-2 px-4 text-white font-semibold rounded-md shadow ${
              mutationLogin.isPending ? 'bg-blue-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'
            }`}
            disabled={mutationLogin.isPending}
          >
            {mutationLogin.isPending ? 'Logging in...' : 'Login'}
          </button>
        </form>
        <div className="text-center mt-6">
          <p className="text-gray-600">
            Don’t have an account?{' '}
            <Link href="/register" className="text-blue-600 hover:underline">
              Register here
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
