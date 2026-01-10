import type * as Preset from '@docusaurus/preset-classic';
import type { Config } from '@docusaurus/types';
import { themes as prismThemes } from 'prism-react-renderer';

const config: Config = {
  title: 'Wails New',
  tagline: 'Instantly bootstrap a modern Wails desktop app.',
  favicon: 'img/favicon.ico',

  future: {
    v4: true,
  },

  url: 'https://tidjee-dev.github.io',
  baseUrl: '/wails-new/',
  organizationName: 'tidjee-dev',
  projectName: 'wails-new',
  deploymentBranch: 'gh-pages',
  trailingSlash: true,

  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  scripts: [
    {
      src: "https://buttons.github.io/buttons.js",
      async: true,
      defer: true,
    },
  ],

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          editUrl:
            'https://github.com/tidjee-dev/wails-new/edit/main/docs-site/docs/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    colorMode: {
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: 'Wails New',
      items: [
        {
          to: '/docs/overview',
          position: 'left',
          label: 'Docs',
        },
        // {
        //   to: '/docs/examples/',
        //   position: 'left',
        //   label: 'Examples',
        // },
        {
          href: 'https://github.com/tidjee-dev/wails-new',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        // {
        //   title: 'More',
        //   items: [
        //     {
        //       label: 'Project Source',
        //       href: 'https://github.com/tidjee-dev/wails-new',
        //     },
        //   ],
        // },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} <a href="https://github.com/tidjee-dev" target="_blank" rel="noopener noreferrer">Tidjee</a>. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
