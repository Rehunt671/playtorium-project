'use client';
import React, { createContext, useContext, useRef } from 'react';
import { Toast } from 'primereact/toast';
import 'primeicons/primeicons.css';
import 'primereact/resources/primereact.css';
import 'primereact/resources/themes/lara-light-indigo/theme.css';

type ToastContextType = {
  showToast: (message: string, severity: 'success' | 'error', duration: number) => void;
};

const ToastContext = createContext<ToastContextType | undefined>(undefined);

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const toast = useRef<Toast>(null);

  const showToast = (message: string, severity: 'success' | 'error', duration: number) => {
    toast.current?.show({severity:severity, summary: severity ? 'Success': 'Error', detail: message, life: duration });
  };

  return (
      <ToastContext.Provider value={{ showToast }}>
        {children}
        <Toast ref={toast} position="bottom-left" />
      </ToastContext.Provider>

  );
};

export const useToast = (): ToastContextType => {
  const context = useContext(ToastContext);
  if (!context) {
    throw new Error('useToast must be used within a ToastProvider');
  }
  return context;
};
