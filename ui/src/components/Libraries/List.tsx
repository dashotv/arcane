import React from "react";

import { Library } from "client";

import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import EditIcon from "@mui/icons-material/Edit";
import QueueIcon from "@mui/icons-material/Queue";
import { IconButton, Paper, Typography } from "@mui/material";
import { Stack } from "@mui/system";

import { LoadingIndicator, Row } from "@dashotv/components";

import { LibrariesDialog } from "./Dialog";
import {
  useMutationLibrary,
  useMutationLibraryCreate,
  useMutationLibraryDelete,
  useQueryLibraries,
} from "./query";

export const LibrariesList = () => {
  const [editing, setEditing] = React.useState<Library | null>(null);
  const { data, isFetching } = useQueryLibraries();
  const creator = useMutationLibraryCreate();
  const deleter = useMutationLibraryDelete();
  const updater = useMutationLibrary();

  const create = (data: Library) => {
    setEditing(null);
    creator.mutate(data);
  };
  const update = (data: Library) => {
    setEditing(null);
    if (!data.id) {
      throw new Error("ID must be set");
    }
    updater.mutate({ id: data.id, data });
  };
  const save = (data: Library | null) => {
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
          Libraries (Kinds)
        </Typography>
        <IconButton size="small" onClick={() => setEditing({ name: "" })}>
          <QueueIcon fontSize="small" color="primary" />
        </IconButton>
      </Stack>
      {(data || []).map((library, index) => (
        <Row key={index}>
          <Stack
            direction="row"
            spacing={1}
            alignItems="center"
            justifyContent="space-between"
          >
            <Stack direction="row" spacing={1} alignItems="baseline">
              <Typography
                variant="body1"
                fontWeight="bolder"
                minWidth="100px"
                color="primary"
              >
                {library.name}
              </Typography>
              <Typography
                variant="body2"
                fontWeight="bolder"
                color="primary.dark"
              >
                {library.path}
              </Typography>
              <Stack direction="row" spacing={0.5} alignItems="baseline">
                <Typography
                  variant="body2"
                  fontWeight="bolder"
                  color="secondary"
                >
                  type:
                </Typography>
                <Typography
                  variant="body2"
                  fontWeight="bolder"
                  color="secondary.dark"
                >
                  {library.library_type?.name}
                </Typography>
              </Stack>
              <Stack direction="row" spacing={0.5} alignItems="baseline">
                <Typography
                  variant="body2"
                  fontWeight="bolder"
                  color="secondary"
                >
                  template:
                </Typography>
                <Typography
                  variant="body2"
                  fontWeight="bolder"
                  color="secondary.dark"
                >
                  {library.library_template?.name}
                </Typography>
              </Stack>
            </Stack>
            <Stack direction="row" spacing={1} alignItems="center">
              <IconButton size="small" onClick={() => setEditing(library)}>
                <EditIcon fontSize="small" color="primary" />
              </IconButton>
              <IconButton
                size="small"
                onClick={() => library?.id && remove(library.id)}
              >
                <DeleteForeverIcon fontSize="small" color="error" />
              </IconButton>
            </Stack>
          </Stack>
        </Row>
      ))}
      {editing && <LibrariesDialog library={editing} submit={save} />}
    </Paper>
  );
};
