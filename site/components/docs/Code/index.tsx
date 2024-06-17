import { Code, Pre } from "nextra/components";
import { useEffect, useState } from "react";
import { codeToHtml, createCssVariablesTheme } from "shiki";

// import theme from "./theme.json";

const myTheme = createCssVariablesTheme({
  name: "css-variables",
  variablePrefix: "--shiki-",
  variableDefaults: {},
  fontStyle: true,
});

export default ({ code, lang }) => {
  const [highlightedCode, setHighlightedCode] = useState(null);

  useEffect(() => {
    codeToHtml(code, {
      lang,
      theme: myTheme,
    }).then((html) => {
      // Trim the html to the content of <code> tag
      const start = html.indexOf("<code>") + 6;
      const end = html.indexOf("</code>");
      setHighlightedCode(html.substring(start, end));
    });
  }, [code, lang]);

  if (!highlightedCode) {
    return null;
  }

  return (
    <Pre data-language={lang} data-theme="default">
      <Code
        data-language={lang}
        data-theme="default"
        dangerouslySetInnerHTML={{ __html: highlightedCode }}
      />
    </Pre>
  );
};
