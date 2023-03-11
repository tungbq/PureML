<a href="https://pureml.com" align="center">
  <img
    align="center"
    alt="PureML"
    src="/assets/FrontendCoverImg.svg"
    style="width:100%;"
  />
</a>

<br/>
<br/>

<div align="center">
  <a
    href="https://pypi.org/project/pureml/"
  >
    <img alt="Remix" src="https://img.shields.io/badge/remix-1.11.1-green?style=flat&logo=remix&logoColor=white" />
  </a>
  &nbsp;
  <a
    href="https://python-poetry.org/"
  >
    <img src="https://img.shields.io/badge/react-18.2.0-red?style=flat&logo=react&logoColor=white" />
  </a>
  &nbsp;
  <a
    href="https://opensource.org/licenses/Apache-2.0"
  >
    <img alt="License" src="https://img.shields.io/badge/tailwindcss-3.1.2-blue?style=flat&logo=tailwindcss&logoColor=white" />
  </a>
  &nbsp;
  <a
    href="https://discord.gg/xNUHt9yguJ"
  >
    <img alt="Discord" src="https://img.shields.io/badge/Discord-Join%20Discord-blueviolet?style=flat&logo=discord&logoColor=white" />
  </a>
  &nbsp;
</div>

<br/>

## Quick start

PureML UI helps you to visualize all the details and versions of your models and datasets you want to work with. It takes no time to run PureML UI on your local system. Follow below steps to run:

Start the server:

```bash
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) to use the UI.

<br/>

## Directory Structure

```
pureml_frontend
├─ app/
│  ├─ analytics/             # Analytics
│  │  ├─ gtags.client.ts     # Google Tag
│  │  ├─ reportWebVitals.ts  # Vercel web vitals
│  │  └─ vitals.ts           # Vitals
|  |
|  ├─ components/            # Components
│  │  ├─ landingPage/        # Landing Page files
│  │  ├─ ui/                 # Minimal UI components used in all
│  │  └─ ...                 # other components used
│  │
│  ├─ routes/                # Route pages
│  │  ├─ api/                # All apis used in app
│  │  ├─ auth/               # Pages under Authentication
│  │  ├─ datasets/           # Pages showing all datasets
│  │  ├─ models/             # Pages showing all models
│  │  ├─ org/                # Pages under Org feature
│  │  ├─ settings/           # Pages under settings
│  │  └─ ...                 # …that has layouts of each page
│  │
│  ├─ entry.client.ts        # Entry.client file by remix
│  ├─ entry.server.ts        # Entry.server file by remix
│  ├─ lib.type.d.ts          # Type file
│  ├─ session.ts             # Session used for authentication
│  └─ root.tsx               # Root Index
│
├─ public/                   # Public
|  ├─ error/                 # Error images
|  └─ imgs/                  # Images used in application
|
├─ styles/                   # Styles
│  └─ app.css                # CSS file
|
├─ .gitignore                # List of files and folders not tracked by Git
├─ .eslintrc                 # Linting preferences for JavasScript
├─ remix.config.js           # Remix configuration file
├─ tailwind.config.js        # Tailwind configuration file
├─ package.json              # Project manifest
└─ README.md                 # This file
```

## Technology used

1. [Remix framework](https://remix.run/)
2. [Tailwind CSS](https://remix.run/)
3. [Tailwind Variants](https://www.tailwind-variants.org/docs/introduction)
4. [Radix components](https://www.radix-ui.com/docs/primitives/overview/introduction)
5. [Reactflow](https://reactflow.dev/)

## Reporting Bugs

To report any bugs you have faced while using PureML package, please

1. Report it in [Discord](https://discord.gg/xNUHt9yguJ) channel
2. Open an [issue](https://github.com/PureMLHQ/PureML/issues)

<br />

## Contributing and Developing

Lets work together to improve the features for everyone. Here's step one for you to go through our [Contributing Guide](./CONTRIBUTING.md). We are already waiting for amazing ideas and features which you all have got.

Work with mutual respect. Please take a look at our public [Roadmap here](https://pureml.notion.site/7de13568835a4cf18913307503a2cdd4?v=82199f96833a48e5907023c8a8d565c6).

<br />

## Community

To get quick updates of feature releases of PureML, follow us on:

[<img alt="Twitter" height="20" src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" />](https://twitter.com/getPureML) [<img alt="LinkedIn" height="20" src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/company/PuremlHQ/) [<img alt="GitHub" height="20" src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white" />](https://github.com/PureMLHQ/PureML) [<img alt="GitHub" height="20" src="https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white" />](https://discord.gg/DBvedzGu)

<br/>

## 📄 License

See the [Apache-2.0](./License) file for licensing information.
