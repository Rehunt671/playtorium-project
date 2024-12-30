import { useQuery } from "@tanstack/react-query";
import axiosCustom from "@/lib/axios"; 

export const useQueryUser = (token: string) => {
  return useQuery({
    queryKey: ['user'],
    queryFn: async () => {
      if (token === '') return null;
      const response = await axiosCustom.get('/v1/users/me', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return response.data;
    },
    enabled: !!token, 
  });
};
