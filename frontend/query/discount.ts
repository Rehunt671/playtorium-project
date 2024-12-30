"use client"
import axiosCustom from "@/lib/axios";
import { useToast } from "@/providers/ToastProvider";
import { CreateDiscount, Discount} from "@/types/discount";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

export const useMutationCreateDiscount = () => {
  const queryClient = useQueryClient();
  const {showToast} = useToast();
    return useMutation({
      mutationFn:async (discountBody : CreateDiscount) => {
        const response = await axiosCustom.post('/v1/discounts', discountBody);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({queryKey:[`discounts`]});
        queryClient.invalidateQueries({queryKey:[`discount-categories`]});
        showToast('Add discount successful', 'success', 1000);
      },
      onError: () => {showToast('Failed to add discount', 'error', 1000);
      }}
    );
  };

  export const useMutationDeleteDiscount = () => {
    const queryClient = useQueryClient();
    const {showToast} = useToast();
      return useMutation({
        mutationFn:async (discountId : number) => {
          const response = await axiosCustom.delete(`/v1/discounts/${discountId}`);
          return response.data;
        },
        onSuccess: () => {
          queryClient.invalidateQueries({queryKey:[`discounts`]});
          queryClient.invalidateQueries({queryKey:[`discount-categories`]});
          showToast('Remove discount successful', 'success', 1000);
        },
        onError: () => {showToast('Failed to remove discount', 'error', 1000);
        }}
      );
    };

    
export const useQueryDiscounts = () => {
    return useQuery<Discount[],Error>({
      queryKey: [`discounts`],
      queryFn: async () => {
        const response = await axiosCustom.get(`/v1/discounts`);
        return response.data;
      },
    });
};
