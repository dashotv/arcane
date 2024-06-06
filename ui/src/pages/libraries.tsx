import { Helmet } from "react-helmet-async";

const Libraries = () => {
  // limit, skip, queries, etc
  // const [page] = useState(1);
  // const handleCancel = (id: string) => {
  //   console.log('cancel', id);
  // };

  return (
    <>
      <Helmet>
        <title>Arcane - Libraries</title>
        <meta name="description" content="arcane" />
      </Helmet>
      {/* <JobsList {...{ page, handleCancel }} /> */}
    </>
  );
};

export default Libraries;
