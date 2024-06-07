import {
  Library,
  LibraryCreate,
  LibraryDelete,
  LibraryIndex,
  LibraryUpdate,
} from "client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

export const useQueryLibraries = () => {
  return useQuery<Library[]>({
    queryKey: ["libraries"],
    queryFn: async () => {
      const response = await LibraryIndex({ page: 1, limit: 50 });
      return response?.result;
    },
  });
};

export const useMutationLibrary = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: Library }) => {
      return LibraryUpdate({ id, subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraries"] });
    },
  });
};

export const useMutationLibraryCreate = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: Library) => {
      return LibraryCreate({ subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraries"] });
    },
  });
};

export const useMutationLibraryDelete = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (id: string) => {
      return LibraryDelete({ id });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraries"] });
    },
  });
};
