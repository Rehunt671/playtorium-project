"use client";
import { useQuery, UseQueryResult } from "@tanstack/react-query";
import axiosCustom from "@/lib/axios";
import {  ItemCategory } from "@/types/item";

export const useQueryItemCategories = (): UseQueryResult<ItemCategory[], Error> => {
  return useQuery<ItemCategory[], Error>({
    queryKey: [`item-categories`],
    queryFn: async () => {
      const response = await axiosCustom.get(`/v1/item-categories`);
      return response.data;
    },
  });
};

