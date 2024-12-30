'use client';
import React from 'react';
import ItemCard from '@/components/cards/ItemCard';
import { useQueryItems } from '@/query/item';
import Spinner from '@/components/Spinner';

const Store: React.FC = () => {
  const queryItems = useQueryItems();
  const items = queryItems.data;

  if (queryItems.isLoading) {
    return <Spinner/>
  }

  return (
    <div className="min-h-screen bg-gradient-to-r from-gray-50 via-white to-gray-100 py-8 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <header className="text-center mb-8">
          <h1 className="text-4xl font-extrabold text-gray-800 mb-2 drop-shadow-lg">
            E-Commerce Application
          </h1>
          <p className="text-lg text-gray-600">Find the best items at amazing prices!</p>
        </header>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
          {items?.map((item) => (
            <ItemCard
              key={item.id}
              item={item}
            />
          ))}
        </div>
      </div>
    </div>
  );
};

export default Store;
