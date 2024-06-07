import { Container } from "@dashotv/components";

import { LibrariesList } from "components/Libraries";
import { LibraryTemplatesList } from "components/LibraryTemplates";
import { LibraryTypesList } from "components/LibraryTypes";
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
        <LibrariesList />
        <LibraryTypesList />
        <LibraryTemplatesList />
      </Container>
    </>
  );
};

export default Libraries;
