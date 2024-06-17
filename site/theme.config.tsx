import React from "react";
import { DocsThemeConfig } from "nextra-theme-docs";

import LogoSite from "./components/logos/LogoSite";

const config: DocsThemeConfig = {
  // logo: <span>My Project</span>,

  // project: {
  //   link: 'https://github.com/shuding/nextra-docs-template',
  // },
  // docsRepositoryBase: 'https://github.com/shuding/nextra-docs-template',
  // Logo in the navbar
  logo: <LogoSite height="1.5rem" />,
  // Footer text
  footer: {
    text: "Renderhook",
  },
  // Hide the last updated time
  gitTimestamp: null,
  // Hide feedback link
  feedback: {
    content: null,
  },
  // Hide the edit link
  editLink: {
    component: null,
  },
  useNextSeoProps: () => ({ titleTemplate: "%s – renderHOOK" }),
};

export default config;
