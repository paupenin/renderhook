import Code from "../Code";

const Response = ({ data }) => {
  return <Code code={JSON.stringify(data, null, 2)} lang="json" />;
};

Response.url = (path) => {
  return "https://api.renderhook.com" + path;
};

export default Response;
