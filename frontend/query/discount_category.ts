import axiosCustom from "@/lib/axios";
import { DiscountCategory } from "@/types/discount_category";
import { useQuery } from "@tanstack/react-query";

export const useQueryAvailableDiscountCategories = () => {
    return useQuery<DiscountCategory[],Error>({
      queryKey: [`discount-categories`],
      queryFn: async () => {
        const response = await axiosCustom.get(`/v1/discount-categories`);
        return response.data;
      },
    });
};

  
