import Code from "../Code";

// Code JS component
const CodeJs = ({ method, path, params }) => {
  // Generate body code with 2 spaces at the end
  const body = JSON.stringify(params, null, 4).replace(/}$/, "  }");

  // Generate the fetch code
  const code = `fetch("${path}", {
  method: "${method}",
  headers: {
    "Content-Type": "application/json",
    "Authorization": "Bearer API_KEY"
  },
  body: ${body}
})`;

  return <Code code={code} lang="javascript" />;
};

// Code Curl component
const CodeCurl = ({ method, path, params }) => {
  let code = `curl -X ${method} "https://api.renderhook.com${path}"`;

  // Add the Authorization header to the curl command
  code += ` \\\n     -H "Content-Type: application/json"`;

  // Add the Authorization header to the curl command
  code += ` \\\n     -H "Authorization: Bearer API_KEY"`;

  // If there are params, add them to the curl command
  if (params) {
    code += ` \\\n     -d '${JSON.stringify(params)}'`;
  }

  return <Code code={code} lang="bash" />;
};

// Request component that takes in method, path, and params
export default function Request({ method, path, params }) {
  return (
    <>
      <CodeJs method={method} path={path} params={params} />
      <CodeCurl method={method} path={path} params={params} />
    </>
  );
}
