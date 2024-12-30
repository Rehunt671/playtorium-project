import axiosCustom from "@/lib/axios";
import { DiscountCategory } from "@/types/discount_category";
import { useQuery } from "@tanstack/react-query";

export const useQueryCampaignOnDiscountCategory = (discountCategoryId: number) => {
    return useQuery<DiscountCategory[], Error>({
        queryKey: [`campaign-on-discount-type`, discountCategoryId],
        queryFn: async () => {
            const response = await axiosCustom.get(`/v1/campaigns`, {
                params: { discountCategoryId },
            });
            return response.data;
        },
        enabled: !!discountCategoryId
    });
};