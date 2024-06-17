import Table from "../Table";

const Endpoint = ({ method, path, params }) => {
  return (
    <>
      <Table
        rows={[
          ["Method", method],
          ["Path", path],
        ]}
      />

      <h3>Options:</h3>
      {params.map((param, i) => (
        <Table
          key={i}
          rows={[
            ["Key", param.key],
            ["Type", param.type],
            param.required !== undefined
              ? ["Required", param.required ? "Yes" : "No"]
              : ["Required", "Yes"],
            param.default !== undefined ? ["Default", param.default] : null,
            param.description ? ["Description", param.description] : null,
          ].filter((v) => !!v)}
        />
      ))}
    </>
  );
};

Endpoint.string = "string";
Endpoint.boolean = "boolean";

export default Endpoint;
