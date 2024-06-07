import { useQuery } from "@tanstack/react-query";
import { File, FileIndex } from "client";

export const useQueryFiles = (page: number, limit: number) => {
  return useQuery<File[]>({
    queryKey: ["files"],
    queryFn: async () => {
      const response = await FileIndex({ page: page, limit: limit });
      return response?.result;
    },
  });
};
