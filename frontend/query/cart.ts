import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axiosCustom from "@/lib/axios";
import { CartItem } from "@/types/cart_item";
import { CartDetail } from "@/types/cart";
import { useToast } from "@/providers/ToastProvider";

export const useQueryCartItems = (cartId: number) => {
  return useQuery<CartItem[],Error>({
    queryKey: [`cartItems`, cartId],
    queryFn: async () => {
      const response = await axiosCustom.get(`/v1/carts/${cartId}/cartItem`);
      return response.data;
    },
    enabled: !!cartId,
  });
};

export const useQueryCartDetail = (userId: number) => {
  return useQuery<CartDetail,Error>({
    queryKey: [`cartDetail`, userId],
    queryFn: async () => {
      const response = await axiosCustom.get(`/v1/carts/user/${userId}`);
      return response.data;
    },
    enabled: !!userId,
  });
};

export const useMutationAddCartItem = (cartId: number) => {
  const queryClient = useQueryClient();
  const {showToast} = useToast();
  return useMutation({
    mutationFn: async (cartItem: CartItem) => {
      const response = await axiosCustom.post(`/v1/carts/${cartId}/cartItem`, cartItem);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:[`cartItems`, cartId]});
      queryClient.invalidateQueries({queryKey:[`cartDetail`]});
      showToast('Add item to cart successful', 'success', 3000);
    },
    onError: () => {showToast('Failed to add item to cart', 'error', 3000);
    }
  });
};

export const useMutationUpdateCartItem = (cartId: number, cartItemId: number) => {
  const queryClient = useQueryClient();
  const {showToast} = useToast();
  return useMutation({
    mutationFn: async (cartItem: CartItem) => {
      const response = await axiosCustom.put(`/v1/carts/${cartId}/cartItem/${cartItemId}`, cartItem);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:[`cartItems`, cartId]});
      queryClient.invalidateQueries({queryKey:[`cartDetail`]});
      showToast('Update item quantity successful', 'success', 1000);
    },
    onError: () => {showToast('Failed to update item quantity', 'error', 1000);
    }
  });
};

export const useMutationDeleteCartItem = (cartId: number, cartItemId: number) => {
  const queryClient = useQueryClient();
  const {showToast} = useToast();
  return useMutation({
    mutationFn: async () => {
      const response = await axiosCustom.delete(`/v1/carts/${cartId}/cartItem/${cartItemId}`);
      return response.data;
    },
    onSuccess: () => {
      showToast('Remove item successful', 'success', 1000);
      queryClient.invalidateQueries({queryKey:[`cartItems`, cartId]});
      queryClient.invalidateQueries({queryKey:[`cartDetail`]});
    },
    onError: () => {showToast('Failed to remove item from cart', 'error', 1000);
    }
  });
};