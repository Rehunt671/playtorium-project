"use client";
import { useQuery, UseQueryResult } from "@tanstack/react-query";
import axiosCustom from "@/lib/axios";
import { Item } from "@/types/item";

export const useQueryItems = (): UseQueryResult<Item[], Error> => {
  return useQuery<Item[], Error>({
    queryKey: [`items`],
    queryFn: async () => {
      const response = await axiosCustom.get(`/v1/items`);
      return response.data;
    },
  });
};

