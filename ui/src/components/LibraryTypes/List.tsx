import React from 'react';

import { LibraryType } from 'client/tower';

import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
import QueueIcon from '@mui/icons-material/Queue';
import { IconButton, Paper, Typography } from '@mui/material';
import { Stack } from '@mui/system';

import { LoadingIndicator, Row } from '@dashotv/components';

import { LibraryTypesDialog } from './Dialog';
import {
  useMutationLibraryType,
  useMutationLibraryTypeCreate,
  useMutationLibraryTypeDelete,
  useQueryLibraryTypes,
} from './query';

export const LibraryTypesList = () => {
  const [editing, setEditing] = React.useState<LibraryType | null>(null);
  const { data, isFetching } = useQueryLibraryTypes();
  const creator = useMutationLibraryTypeCreate();
  const deleter = useMutationLibraryTypeDelete();
  const updater = useMutationLibraryType();

  const create = (data: LibraryType) => {
    setEditing(null);
    creator.mutate(data);
  };
  const update = (data: LibraryType) => {
    setEditing(null);
    if (!data.id) {
      throw new Error('ID must be set');
    }
    updater.mutate({ id: data.id, data });
  };
  const save = (data: LibraryType | null) => {
    setEditing(null);
    if (!data) {
      return;
    }
    if (!data?.id) {
      create(data);
      return;
    }
    update(data);
  };
  const remove = (id: string) => {
    deleter.mutate(id);
  };

  return (
    <Paper elevation={0} sx={{ p: 1, mb: 2 }}>
      {isFetching && <LoadingIndicator />}
      <Stack direction="row" spacing={1}>
        <Typography variant="h6" color="primary">
          Types
        </Typography>
        <IconButton size="small" onClick={() => setEditing({ name: '' })}>
          <QueueIcon fontSize="small" color="primary" />
        </IconButton>
      </Stack>
      {(data || []).map((libraryType, index) => (
        <Row key={index}>
          <Stack direction="row" spacing={1} alignItems="center" justifyContent="space-between">
            <Typography variant="body1" fontWeight="bolder" color="primary">
              {libraryType.name}
            </Typography>
            <Stack direction="row" spacing={1}>
              <IconButton size="small" onClick={() => setEditing(libraryType)}>
                <EditIcon fontSize="small" color="primary" />
              </IconButton>
              <IconButton size="small" onClick={() => libraryType?.id && remove(libraryType.id)}>
                <DeleteForeverIcon fontSize="small" color="error" />
              </IconButton>
            </Stack>
          </Stack>
        </Row>
      ))}
      {editing && <LibraryTypesDialog libraryType={editing} submit={save} />}
    </Paper>
  );
};
