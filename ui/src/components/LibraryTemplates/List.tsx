import React from 'react';

import { LibraryTemplate } from 'client/tower';

import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
import QueueIcon from '@mui/icons-material/Queue';
import { IconButton, Paper, Typography } from '@mui/material';
import { Stack } from '@mui/system';

import { LoadingIndicator, Row } from '@dashotv/components';

import { LibraryTemplatesDialog } from './Dialog';
import {
  useMutationLibraryTemplate,
  useMutationLibraryTemplateCreate,
  useMutationLibraryTemplateDelete,
  useQueryLibraryTemplates,
} from './query';

export const LibraryTemplatesList = () => {
  const [editing, setEditing] = React.useState<LibraryTemplate | null>(null);
  const { data, isFetching } = useQueryLibraryTemplates();
  const creator = useMutationLibraryTemplateCreate();
  const deleter = useMutationLibraryTemplateDelete();
  const updater = useMutationLibraryTemplate();

  const create = (data: LibraryTemplate) => {
    setEditing(null);
    creator.mutate(data);
  };
  const update = (data: LibraryTemplate) => {
    setEditing(null);
    if (!data.id) {
      throw new Error('ID must be set');
    }
    updater.mutate({ id: data.id, data });
  };
  const save = (data: LibraryTemplate | null) => {
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
          Templates
        </Typography>
        <IconButton size="small" onClick={() => setEditing({ name: '' })}>
          <QueueIcon fontSize="small" color="primary" />
        </IconButton>
      </Stack>
      {(data || []).map((libraryTemplate, index) => (
        <Row key={index}>
          <Stack direction="row" spacing={1} alignItems="center" justifyContent="space-between">
            <Stack direction="row" spacing={1} alignItems="baseline">
              <Typography variant="body1" fontWeight="bolder" minWidth="100px" color="primary">
                {libraryTemplate.name}
              </Typography>
              <Typography variant="body2" fontWeight="bolder" color="primary.dark">
                {libraryTemplate.template}
              </Typography>
            </Stack>
            <Stack direction="row" spacing={1}>
              <IconButton size="small" onClick={() => setEditing(libraryTemplate)}>
                <EditIcon fontSize="small" color="primary" />
              </IconButton>
              <IconButton size="small" onClick={() => libraryTemplate?.id && remove(libraryTemplate.id)}>
                <DeleteForeverIcon fontSize="small" color="error" />
              </IconButton>
            </Stack>
          </Stack>
        </Row>
      ))}
      {editing && <LibraryTemplatesDialog libraryTemplate={editing} submit={save} />}
    </Paper>
  );
};
