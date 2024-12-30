import axiosCustom from "@/lib/axios";
import { useToast } from "@/providers/ToastProvider";
import { LoginCredentials, RegisterCredentials } from "@/types/auth";
import { useMutation } from "@tanstack/react-query";

export const useMutationLogin = () => {
  const {showToast} = useToast();
  return useMutation({
    mutationFn:async (credentials: LoginCredentials) => {
      const response = await axiosCustom.post('/v1/auth/login', credentials);
      return response.data;
    },
    onSuccess: () => {showToast('Login successful', 'success', 3000);},
    onError: () => {showToast('Login failed. Please try again.', 'error', 3000);
    }
  }
  );
};

export const useMutationRegister = () => {
  const {showToast} = useToast();
  return useMutation({
    mutationFn:async (credentials: RegisterCredentials) => {
      const response = await axiosCustom.post('/v1/auth/register', credentials);
      return response.data;
      
    },
    onSuccess: () => {showToast('Registration successful', 'success', 3000);},
    onError: () => {showToast('Registration failed. Please try again.', 'error', 3000);
    }
    }
  );
};