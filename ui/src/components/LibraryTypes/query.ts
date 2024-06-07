import {
  LibraryType,
  LibraryTypeCreate,
  LibraryTypeDelete,
  LibraryTypeIndex,
  LibraryTypeUpdate,
} from "client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

export const useQueryLibraryTypes = () => {
  return useQuery<LibraryType[]>({
    queryKey: ["libraryTypes"],
    queryFn: async () => {
      const response = await LibraryTypeIndex({ page: 1, limit: 50 });
      return response?.result;
    },
  });
};

export const useMutationLibraryType = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: LibraryType }) => {
      return LibraryTypeUpdate({ id, subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTypes"] });
    },
  });
};

export const useMutationLibraryTypeCreate = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: LibraryType) => {
      return LibraryTypeCreate({ subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTypes"] });
    },
  });
};

export const useMutationLibraryTypeDelete = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (id: string) => {
      return LibraryTypeDelete({ id });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTypes"] });
    },
  });
};
