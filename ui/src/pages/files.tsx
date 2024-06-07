import { Container } from "@dashotv/components";
import { FilesList } from "components/Files/List";

import { Helmet } from "react-helmet-async";

const Libraries = () => {
  // limit, skip, queries, etc
  // const [page] = useState(1);
  // const handleCancel = (id: string) => {
  //   console.log("cancel", id);
  // };

  return (
    <>
      <Helmet>
        <title>Arcane - Libraries</title>
        <meta name="description" content="arcane" />
      </Helmet>
      <Container>
        <FilesList page={1} />
      </Container>
    </>
  );
};

export default Libraries;
