import React, { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';

import { Library } from 'client/tower';

import { Box, Button, Dialog, DialogContent, DialogTitle } from '@mui/material';
import { Stack } from '@mui/system';

import { Option, Select, Text } from '@dashotv/components';

import { useQueryLibraryTemplates } from 'components/LibraryTemplates';
import { useQueryLibraryTypes } from 'components/LibraryTypes';

export const LibrariesDialog = ({ library, submit }: { library: Library; submit: (data: Library | null) => void }) => {
  const [open, setOpen] = useState(true);
  const [libraryTypes, setLibraryTypes] = useState<Option[]>([]);
  const [templates, setDestTemplates] = useState<Option[]>([]);
  const { control, handleSubmit } = useForm<Library>({ values: library });
  const { data: libraryTypesData } = useQueryLibraryTypes();
  const { data: templatesData } = useQueryLibraryTemplates();

  const close = (data: Library | null) => {
    submit(data);
    setOpen(false);
  };

  useEffect(() => {
    if (!libraryTypesData) {
      return;
    }
    setLibraryTypes(libraryTypesData.map(item => ({ label: item.name!, value: item.id! }) as Option));
  }, [libraryTypesData]);

  useEffect(() => {
    if (!templatesData) {
      return;
    }
    setDestTemplates(templatesData.map(item => ({ label: item.name!, value: item.id! }) as Option));
  }, [templatesData]);

  return (
    <Dialog open={open} maxWidth="md" fullWidth>
      <DialogTitle>Edit</DialogTitle>
      <DialogContent>
        <Box component="form" onSubmit={handleSubmit(close)}>
          <Stack direction="column" spacing={2}>
            <Text control={control} name="name" />
            <Text control={control} name="path" />
            <Select control={control} name="library_type_id" options={libraryTypes} />
            <Select control={control} name="library_template_id" options={templates} />
          </Stack>
          <Stack direction="row" spacing={1} pt={3} justifyContent="end">
            <Button variant="contained" onClick={() => close(null)}>
              Cancel
            </Button>
            <Button variant="contained" type="submit">
              Submit
            </Button>
          </Stack>
        </Box>
      </DialogContent>
    </Dialog>
  );
};
