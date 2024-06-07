import {
  LibraryTemplate,
  LibraryTemplateCreate,
  LibraryTemplateDelete,
  LibraryTemplateIndex,
  LibraryTemplateUpdate,
} from "client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

export const useQueryLibraryTemplates = () => {
  return useQuery<LibraryTemplate[]>({
    queryKey: ["libraryTemplates"],
    queryFn: async () => {
      const response = await LibraryTemplateIndex({ page: 1, limit: 50 });
      return response?.result;
    },
  });
};

export const useMutationLibraryTemplate = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: LibraryTemplate }) => {
      return LibraryTemplateUpdate({ id, subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTemplates"] });
    },
  });
};

export const useMutationLibraryTemplateCreate = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: LibraryTemplate) => {
      return LibraryTemplateCreate({ subject: data });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTemplates"] });
    },
  });
};

export const useMutationLibraryTemplateDelete = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (id: string) => {
      return LibraryTemplateDelete({ id });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["libraryTemplates"] });
    },
  });
};
