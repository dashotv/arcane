import React from 'react';
import { useForm } from 'react-hook-form';

import { LibraryTemplate } from 'client/tower';

import { Box, Button, Dialog, DialogContent, DialogTitle } from '@mui/material';
import { Stack } from '@mui/system';

import { Text } from '@dashotv/components';

export const LibraryTemplatesDialog = ({
  libraryTemplate,
  submit,
}: {
  libraryTemplate: LibraryTemplate;
  submit: (data: LibraryTemplate | null) => void;
}) => {
  const [open, setOpen] = React.useState(true);
  const { control, handleSubmit } = useForm<LibraryTemplate>({ values: libraryTemplate });

  const close = (data: LibraryTemplate | null) => {
    submit(data);
    setOpen(false);
  };

  return (
    <Dialog open={open} maxWidth="md" fullWidth>
      <DialogTitle>Edit</DialogTitle>
      <DialogContent>
        <Box component="form" onSubmit={handleSubmit(close)}>
          <Stack direction="column" spacing={2}>
            <Text control={control} name="name" />
            <Text control={control} name="template" />
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
