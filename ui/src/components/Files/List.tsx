import { useQueryFiles } from "./query";

export interface FilesListProps {
  page: number;
}
export const FilesList = ({ page }: FilesListProps) => {
  const { data, isLoading } = useQueryFiles(page, 50);
  return (
    <div>
      {isLoading && <div>Loading...</div>}
      {data?.map((file) => <div key={file.id}>{file.path}</div>)}
    </div>
  );
};
