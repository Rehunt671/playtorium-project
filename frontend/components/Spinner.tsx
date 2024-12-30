// components/Spinner.tsx

import React from 'react';

const Spinner: React.FC = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="spinner-border animate-spin inline-block w-8 h-8 border-4 border-solid border-blue-600 border-t-transparent rounded-full"></div>
    </div>
  );
};

export default Spinner;
