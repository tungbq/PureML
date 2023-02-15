var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf, __hasOwnProp = Object.prototype.hasOwnProperty;
var __commonJS = (cb, mod) => function() {
  return mod || (0, cb[__getOwnPropNames(cb)[0]])((mod = { exports: {} }).exports, mod), mod.exports;
};
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: !0 });
}, __copyProps = (to, from, except, desc) => {
  if (from && typeof from == "object" || typeof from == "function")
    for (let key of __getOwnPropNames(from))
      !__hasOwnProp.call(to, key) && key !== except && __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  return to;
};
var __toESM = (mod, isNodeMode, target) => (target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
  isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", { value: mod, enumerable: !0 }) : target,
  mod
)), __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: !0 }), mod);

// empty-module:~/components/quill.client
var require_quill = __commonJS({
  "empty-module:~/components/quill.client"(exports, module2) {
    module2.exports = {};
  }
});

// <stdin>
var stdin_exports = {};
__export(stdin_exports, {
  assets: () => assets_manifest_default,
  assetsBuildDirectory: () => assetsBuildDirectory,
  entry: () => entry,
  future: () => future,
  publicPath: () => publicPath,
  routes: () => routes
});
module.exports = __toCommonJS(stdin_exports);

// app/entry.server.tsx
var entry_server_exports = {};
__export(entry_server_exports, {
  default: () => handleRequest
});
var import_stream = require("stream"), import_node = require("@remix-run/node"), import_react = require("@remix-run/react"), import_isbot = __toESM(require("isbot")), import_server = require("react-dom/server"), import_jsx_runtime = require("react/jsx-runtime"), ABORT_DELAY = 5e3;
function handleRequest(request, responseStatusCode, responseHeaders, remixContext) {
  return (0, import_isbot.default)(request.headers.get("user-agent")) ? handleBotRequest(
    request,
    responseStatusCode,
    responseHeaders,
    remixContext
  ) : handleBrowserRequest(
    request,
    responseStatusCode,
    responseHeaders,
    remixContext
  );
}
function handleBotRequest(request, responseStatusCode, responseHeaders, remixContext) {
  return new Promise((resolve, reject) => {
    let didError = !1, { pipe, abort } = (0, import_server.renderToPipeableStream)(
      /* @__PURE__ */ (0, import_jsx_runtime.jsx)(import_react.RemixServer, { context: remixContext, url: request.url }),
      {
        onAllReady() {
          let body = new import_stream.PassThrough();
          responseHeaders.set("Content-Type", "text/html"), resolve(
            new import_node.Response(body, {
              headers: responseHeaders,
              status: didError ? 500 : responseStatusCode
            })
          ), pipe(body);
        },
        onShellError(error) {
          reject(error);
        },
        onError(error) {
          didError = !0, console.error(error);
        }
      }
    );
    setTimeout(abort, ABORT_DELAY);
  });
}
function handleBrowserRequest(request, responseStatusCode, responseHeaders, remixContext) {
  return new Promise((resolve, reject) => {
    let didError = !1, { pipe, abort } = (0, import_server.renderToPipeableStream)(
      /* @__PURE__ */ (0, import_jsx_runtime.jsx)(import_react.RemixServer, { context: remixContext, url: request.url }),
      {
        onShellReady() {
          let body = new import_stream.PassThrough();
          responseHeaders.set("Content-Type", "text/html"), resolve(
            new import_node.Response(body, {
              headers: responseHeaders,
              status: didError ? 500 : responseStatusCode
            })
          ), pipe(body);
        },
        onShellError(err) {
          reject(err);
        },
        onError(error) {
          didError = !0, console.error(error);
        }
      }
    );
    setTimeout(abort, ABORT_DELAY);
  });
}

// app/root.tsx
var root_exports = {};
__export(root_exports, {
  CatchBoundary: () => CatchBoundary,
  ErrorBoundary: () => ErrorBoundary,
  default: () => App,
  links: () => links,
  loader: () => loader,
  meta: () => meta
});
var import_react3 = require("@remix-run/react"), import_react4 = require("react");

// app/components/ui/Tag.tsx
var import_class_variance_authority = require("class-variance-authority"), import_lucide_react = require("lucide-react"), import_jsx_runtime2 = require("react/jsx-runtime"), tagStyles = (0, import_class_variance_authority.cva)("rounded flex items-center", {
  variants: {
    intent: {
      primary: "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100",
      landingpg: "text-xs text-slate-600 w-fit h-fit px-4 py-1 border border-slate-200",
      modelTag: "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100",
      datasetTag: "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100"
    }
  },
  defaultVariants: {
    intent: "primary"
  }
});
function Tag({ intent, children }) {
  return /* @__PURE__ */ (0, import_jsx_runtime2.jsx)("div", { className: "pr-2", children: /* @__PURE__ */ (0, import_jsx_runtime2.jsx)("div", { className: tagStyles({ intent }), children: intent === "modelTag" || intent === "datasetTag" ? /* @__PURE__ */ (0, import_jsx_runtime2.jsx)(import_jsx_runtime2.Fragment, { children: intent === "modelTag" ? /* @__PURE__ */ (0, import_jsx_runtime2.jsxs)(import_jsx_runtime2.Fragment, { children: [
    /* @__PURE__ */ (0, import_jsx_runtime2.jsx)(import_lucide_react.Database, { className: "w-3 h-3" }),
    /* @__PURE__ */ (0, import_jsx_runtime2.jsx)("div", { className: "pl-2", children })
  ] }) : /* @__PURE__ */ (0, import_jsx_runtime2.jsxs)(import_jsx_runtime2.Fragment, { children: [
    /* @__PURE__ */ (0, import_jsx_runtime2.jsx)(import_lucide_react.File, { className: "w-3 h-3" }),
    /* @__PURE__ */ (0, import_jsx_runtime2.jsx)("div", { className: "pl-2", children })
  ] }) }) : /* @__PURE__ */ (0, import_jsx_runtime2.jsx)(import_jsx_runtime2.Fragment, { children: intent === "primary" ? /* @__PURE__ */ (0, import_jsx_runtime2.jsx)("div", { children }) : /* @__PURE__ */ (0, import_jsx_runtime2.jsxs)("div", { className: "flex justify-center items-center", children: [
    /* @__PURE__ */ (0, import_jsx_runtime2.jsx)(
      "img",
      {
        src: "/imgs/landingPage/ComingSoonIcon.svg",
        alt: "ComingSoon",
        className: "pr-2"
      }
    ),
    "Coming soon..."
  ] }) }) }) });
}
var Tag_default = Tag;

// app/components/landingPage/RawDataSection.tsx
var import_jsx_runtime3 = require("react/jsx-runtime");
function RawDataSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime3.jsxs)("div", { className: "h-fit flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime3.jsx)("div", { className: "flex justify-start items-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "Raw Data" }),
    /* @__PURE__ */ (0, import_jsx_runtime3.jsxs)("div", { className: "flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600", children: [
      /* @__PURE__ */ (0, import_jsx_runtime3.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime3.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Load from different sources"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime3.jsxs)("div", { className: "flex text-slate-600 items-center", children: [
        /* @__PURE__ */ (0, import_jsx_runtime3.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        /* @__PURE__ */ (0, import_jsx_runtime3.jsxs)("div", { children: [
          /* @__PURE__ */ (0, import_jsx_runtime3.jsx)("div", { className: "pr-2", children: "Works with any orchestrator" }),
          /* @__PURE__ */ (0, import_jsx_runtime3.jsx)(Tag_default, { intent: "landingpg" })
        ] })
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime3.jsx)("div", { className: "flex justify-end", children: /* @__PURE__ */ (0, import_jsx_runtime3.jsx)(
      "img",
      {
        src: "/imgs/landingPage/RawData.svg",
        alt: "RawData",
        className: "w-64 sm:hidden"
      }
    ) })
  ] });
}

// app/components/ui/Button.tsx
var import_class_variance_authority2 = require("class-variance-authority"), import_jsx_runtime4 = require("react/jsx-runtime"), buttonStyles = (0, import_class_variance_authority2.cva)(
  "flex items-center font-medium focus:outline-none rounded-lg",
  {
    variants: {
      intent: {
        primary: "bg-blue-750 text-white px-4 py-2 h-9 hover:bg-blue-600 justify-center",
        landingpg: "bg-blue-750 text-white px-12 py-2 h-10 hover:bg-blue-600 justify-center",
        secondary: "bg-white text-slate-600 border border-slate-400 px-4 py-2 h-9 hover:bg-slate-100 justify-center",
        danger: "bg-red-400 text-white hover:bg-red-500 justify-center px-4 py-2 h-9",
        icon: "px-4 py-2 border border-slate-400 hover:bg-slate-100 text-slate-600",
        org: "px-4 py-2 hover:bg-slate-100 hover:text-slate-600"
      },
      fullWidth: {
        true: "w-full"
      }
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: !0
    }
  }
);
function Button({ intent, fullWidth, children, icon, onClick }) {
  return /* @__PURE__ */ (0, import_jsx_runtime4.jsx)("button", { onClick, className: buttonStyles({ intent, fullWidth }), children: intent !== "icon" ? /* @__PURE__ */ (0, import_jsx_runtime4.jsx)("div", { className: "w-max", children }) : /* @__PURE__ */ (0, import_jsx_runtime4.jsxs)("div", { className: "flex items-center", children: [
    icon,
    /* @__PURE__ */ (0, import_jsx_runtime4.jsx)("div", { className: "px-4 w-max", children })
  ] }) });
}
var Button_default = Button;

// app/components/landingPage/CTASection.tsx
var import_jsx_runtime5 = require("react/jsx-runtime");
function CTASection() {
  return /* @__PURE__ */ (0, import_jsx_runtime5.jsx)("div", { className: "p-8 pb-32 h-fit flex flex-col justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime5.jsxs)("div", { className: "lg:w-[60rem] flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime5.jsx)("div", { className: "flex justify-center items-center md:pt-28 text-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "PureML empowers everyone work together, and build knowledge." }),
    /* @__PURE__ */ (0, import_jsx_runtime5.jsx)("div", { className: "py-18 !leading-normal text-lg md:text-2xl text-slate-600 justify-self-center items-center text-center font-medium", children: "No more jumping between tools, struggling with versions, compare changes or sharing via cloud providers." }),
    /* @__PURE__ */ (0, import_jsx_runtime5.jsx)("div", { className: "flex justify-center items-center pt-10", children: /* @__PURE__ */ (0, import_jsx_runtime5.jsx)(
      Button_default,
      {
        intent: "landingpg",
        icon: "",
        fullWidth: !1,
        className: "!w-24",
        children: "Schedule Demo"
      }
    ) })
  ] }) });
}

// app/components/landingPage/DatasetSection.tsx
var import_jsx_runtime6 = require("react/jsx-runtime");
function DatasetSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "h-fit flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime6.jsx)("div", { className: "flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "Datasets" }),
    /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600", children: [
      /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime6.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Create your own branching for experimentation"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime6.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Track versions of dataset"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime6.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Compare different versions of dataset"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime6.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime6.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Review the submitted commit for versioning"
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime6.jsx)("div", { className: "flex justify-end", children: /* @__PURE__ */ (0, import_jsx_runtime6.jsx)(
      "img",
      {
        src: "/imgs/landingPage/Datasets.svg",
        alt: "Datasets",
        className: "w-64 sm:hidden"
      }
    ) })
  ] });
}

// app/components/landingPage/Footer.tsx
var import_jsx_runtime7 = require("react/jsx-runtime");
function Footer() {
  return /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("div", { className: "p-8 md:p-24 lg:p-14 bg-slate-850 xl:h-fit flex md:justify-center", children: /* @__PURE__ */ (0, import_jsx_runtime7.jsxs)("div", { className: "md:flex justify-center text-lg text-slate-400 font-medium gap-y-2 md:gap-y-0", children: [
    /* @__PURE__ */ (0, import_jsx_runtime7.jsxs)("span", { className: "px-4", children: [
      "\xA9 2022,",
      " ",
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "https://pureml.com", className: "text-brand-200", children: "PureML Inc" })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime7.jsxs)("div", { className: "py-4 md:py-0 px-4 flex flex-col md:flex md:flex-row", children: [
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "https://docs.pureml.com", children: "Docs" }),
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("span", { className: "px-2 hidden md:block", children: " | " }),
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "https://discord.com/invite/xNUHt9yguJ", children: "Join Discord" })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime7.jsxs)("div", { className: "flex justify-between px-4", children: [
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "https://twitter.com/getPureML", children: /* @__PURE__ */ (0, import_jsx_runtime7.jsx)(
        "img",
        {
          src: "/imgs/landingPage/Twitter.svg",
          alt: "Twitter",
          width: "36",
          height: "36",
          className: "px-2"
        }
      ) }),
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "mailto:contact@pureml.com", children: /* @__PURE__ */ (0, import_jsx_runtime7.jsx)(
        "img",
        {
          src: "/imgs/landingPage/Mail.svg",
          alt: "Mail",
          width: "36",
          height: "36",
          className: "px-2"
        }
      ) }),
      /* @__PURE__ */ (0, import_jsx_runtime7.jsx)("a", { href: "https://www.linkedin.com/company/pureml-inc/", children: /* @__PURE__ */ (0, import_jsx_runtime7.jsx)(
        "img",
        {
          src: "/imgs/landingPage/Linkedin.svg",
          alt: "Linkedin",
          width: "36",
          height: "36",
          className: "px-2"
        }
      ) })
    ] })
  ] }) });
}

// app/components/landingPage/HeroSection.tsx
var import_jsx_runtime8 = require("react/jsx-runtime");
function HeroSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime8.jsxs)("div", { className: "h-fit sm:flex gap-y-16 justify-between", "data-aos": "fade-up", children: [
    /* @__PURE__ */ (0, import_jsx_runtime8.jsxs)("div", { className: "flex flex-col gap-y-6 sm:w-1/2 justify-center md:pb-8", children: [
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)("div", { className: "flex justify-center items-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "A bridge connecting Data Engineer and Data Scientist" }),
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)("div", { className: "py-18 text-lg md:text-2xl !leading-normal text-slate-600 justify-self-center items-center", children: "We reduce friction between data engineer and data scientist to facilitate seamless collaboration and efficient model development." }),
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)("div", { className: "flex items-center", children: /* @__PURE__ */ (0, import_jsx_runtime8.jsx)("a", { href: "https://calendly.com/pureml-inc/pureml", children: /* @__PURE__ */ (0, import_jsx_runtime8.jsx)(
        Button_default,
        {
          intent: "landingpg",
          icon: "",
          fullWidth: !1,
          className: "!w-24",
          children: "Schedule Demo"
        }
      ) }) })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime8.jsxs)("div", { className: "flex flex-col justify-center items-center py-8 overflow-visible", children: [
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)(
        "img",
        {
          src: "/imgs/landingPage/HeroDataEngineer.svg",
          alt: "Hero",
          className: "absolute move-avatar w-32 -ml-72 -mt-20 md:w-36 md:-ml-96 lg:w-64 lg:-ml-[30rem] z-10"
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)("img", { src: "/imgs/landingPage/HeroImg.svg", alt: "Hero" }),
      /* @__PURE__ */ (0, import_jsx_runtime8.jsx)(
        "img",
        {
          src: "/imgs/landingPage/HeroDataScientist.svg",
          alt: "Hero",
          className: "absolute move-avatar w-32 mt-16 ml-48 md:w-36 md:ml-56 lg:w-56 lg:ml-72 z-10"
        }
      )
    ] })
  ] });
}

// app/components/landingPage/Navbar.tsx
var import_react2 = require("react"), import_lucide_react2 = require("lucide-react");
var import_jsx_runtime9 = require("react/jsx-runtime");
function Navbar() {
  let [open, setOpen] = (0, import_react2.useState)(!0);
  return open ? /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "bg-slate-50 flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsxs)("div", { className: "flex p-4 w-full justify-between max-w-7xl", children: [
    /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("img", { src: "/LogoWText.svg", alt: "", width: "96", height: "96" }),
    /* @__PURE__ */ (0, import_jsx_runtime9.jsx)(
      import_lucide_react2.Menu,
      {
        className: "sm:hidden text-slate-900 cursor-pointer w-8 h-8",
        onClick: () => setOpen(!open)
      }
    ),
    /* @__PURE__ */ (0, import_jsx_runtime9.jsxs)("div", { className: "hidden sm:flex flex font-medium", children: [
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "px-4 flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("a", { href: "https://docs.pureml.com", className: "w-full text-slate-850", children: "Docs" }) }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "px-4 flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)(
        "a",
        {
          href: "https://discord.gg/xNUHt9yguJ",
          className: "w-max text-slate-850",
          children: "Join Discord"
        }
      ) }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "px-4 flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)(
        "a",
        {
          className: "github-button",
          href: "https://github.com/pureml-inc/pureml",
          "data-color-scheme": "no-preference: dark_dimmed; light: light_high_contrast; dark: light;",
          "data-size": "large",
          "data-show-count": "true",
          "aria-label": "Star pureml-inc/pureml on GitHub",
          children: "Star"
        }
      ) }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "px-4 flex justify-center items-center text-blue-600 hover:text-black", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("a", { href: "https://app.pureml.com", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)(Button_default, { intent: "primary", icon: "", children: "Sign in" }) }) })
    ] })
  ] }) }) : /* @__PURE__ */ (0, import_jsx_runtime9.jsxs)("div", { className: "max-w-full sm:max-w-7xl sm:px-24", children: [
    /* @__PURE__ */ (0, import_jsx_runtime9.jsxs)("div", { className: "flex p-4 md:px-12 w-full justify-between max-w-7xl", children: [
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("img", { src: "/LogoWText.svg", alt: "", width: "96", height: "96" }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)(
        import_lucide_react2.X,
        {
          className: "sm:hidden text-slate-900 cursor-pointer w-8 h-8",
          onClick: () => setOpen(!open)
        }
      )
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime9.jsxs)("div", { className: "flex flex-col gap-y-2 p-4 font-medium text-slate-850", children: [
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "flex items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("a", { href: "https://docs.pureml.com", className: "w-max", children: "Docs" }) }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "flex items-center", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("a", { href: "https://discord.gg/xNUHt9yguJ", className: "w-max", children: "Join Discord" }) }),
      /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("div", { className: "flex items-center text-blue-600 hover:text-black pb-2 border-b border-slate-200", children: /* @__PURE__ */ (0, import_jsx_runtime9.jsx)("a", { href: "https://app.pureml.com", children: "Sign in" }) })
    ] })
  ] });
}

// app/components/landingPage/ModelSection.tsx
var import_jsx_runtime10 = require("react/jsx-runtime");
function ModelSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "h-fit flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime10.jsx)("div", { className: "flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "Models" }),
    /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600", children: [
      /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime10.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Create your own branching for experimentation"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime10.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Track model version history"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime10.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Compare different versions of the model"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime10.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime10.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Review submitted commit for versioning"
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime10.jsx)("div", { className: "flex justify-end", children: /* @__PURE__ */ (0, import_jsx_runtime10.jsx)(
      "img",
      {
        src: "/imgs/landingPage/Models.svg",
        alt: "Models",
        className: "w-64 sm:hidden"
      }
    ) })
  ] });
}

// app/components/landingPage/TransformerSection.tsx
var import_jsx_runtime11 = require("react/jsx-runtime");
function TransformerSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { className: "h-fit flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime11.jsx)("div", { className: "flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "Transformers" }),
    /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { className: "flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600", children: [
      /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime11.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Track data lineage from raw data to dataset"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime11.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Save code snippets for the transformations"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { className: "flex text-slate-600 items-center", children: [
        /* @__PURE__ */ (0, import_jsx_runtime11.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        /* @__PURE__ */ (0, import_jsx_runtime11.jsxs)("div", { children: [
          "Automate data pipeline with changes in any transformation step.",
          /* @__PURE__ */ (0, import_jsx_runtime11.jsx)(Tag_default, { intent: "landingpg" })
        ] })
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime11.jsx)("div", { className: "flex justify-end", children: /* @__PURE__ */ (0, import_jsx_runtime11.jsx)(
      "img",
      {
        src: "/imgs/landingPage/Transformer.svg",
        alt: "Transformer",
        className: "w-[18rem] sm:hidden"
      }
    ) })
  ] });
}

// app/routes/api/auth.server.ts
var auth_server_exports = {};
__export(auth_server_exports, {
  fetchPublicProfile: () => fetchPublicProfile,
  fetchSignIn: () => fetchSignIn,
  fetchSignUp: () => fetchSignUp,
  fetchUserOrg: () => fetchUserOrg,
  fetchUserSettings: () => fetchUserSettings
});
var backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL, makeUrl = (path) => `${backendUrl}${path}`;
async function fetchSignIn(email, password, username) {
  let url = makeUrl("user/login");
  return await fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  }).then((res2) => res2.json());
}
async function fetchSignUp(name, username, email, password, bio, avatar) {
  let url = makeUrl("user/signup");
  return await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json"
    },
    body: new URLSearchParams({ name, username, email, password, bio })
  }).then((res2) => res2.json());
}
async function fetchUserOrg(accessToken) {
  let url = makeUrl("org/");
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchUserSettings(accessToken) {
  let url = makeUrl("user/profile");
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchPublicProfile(username) {
  let url = makeUrl(`user/profile/${username}`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json"
    }
  }).then((res2) => res2.json())).Data;
}

// app/session.ts
var import_node2 = require("@remix-run/node"), sessionCookie = (0, import_node2.createCookie)("__session", {
  secrets: ["r3m1xr0ck5"],
  sameSite: !0,
  httpOnly: !0
}), { getSession, commitSession, destroySession } = (0, import_node2.createCookieSessionStorage)({
  cookie: sessionCookie
});

// app/styles/app.css
var app_default = "/build/_assets/app-FL5MQUIR.css";

// node_modules/.pnpm/reactflow@11.5.1_biqbaboplfbrettd7655fr4n2y/node_modules/reactflow/dist/style.css
var style_default = "/build/_assets/style-LWIJXALZ.css";

// app/root.tsx
var import_react_hot_toast = require("react-hot-toast");

// app/components/landingPage/PackageSection.tsx
var import_jsx_runtime12 = require("react/jsx-runtime");
function PackageSection() {
  return /* @__PURE__ */ (0, import_jsx_runtime12.jsxs)("div", { className: "h-fit flex flex-col gap-y-6", children: [
    /* @__PURE__ */ (0, import_jsx_runtime12.jsx)("div", { className: "flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850", children: "Packages" }),
    /* @__PURE__ */ (0, import_jsx_runtime12.jsxs)("div", { className: "flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600", children: [
      /* @__PURE__ */ (0, import_jsx_runtime12.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime12.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Docker for shipping"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime12.jsxs)("div", { className: "flex text-slate-600", children: [
        /* @__PURE__ */ (0, import_jsx_runtime12.jsx)(
          "img",
          {
            src: "/imgs/landingPage/Bullet400.svg",
            alt: "",
            className: "pr-3 w-14 lg:w-20"
          }
        ),
        "Streamlit & gradio for sharing with your team"
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime12.jsx)("div", { className: "flex justify-end", children: /* @__PURE__ */ (0, import_jsx_runtime12.jsx)(
      "img",
      {
        src: "/imgs/landingPage/Packages.svg",
        alt: "Packages",
        className: "w-64 sm:hidden"
      }
    ) })
  ] });
}

// app/root.tsx
var import_jsx_runtime13 = require("react/jsx-runtime");
function links() {
  return [
    { rel: "stylesheet", href: app_default },
    { rel: "stylesheet", href: style_default }
  ];
}
var meta = () => ({
  charset: "utf-8",
  title: "PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader({ request }) {
  let cookieHeader = request.headers.get("cookie"), cacheControlHeader = request.headers.get("cache-control");
  if (cookieHeader && cacheControlHeader) {
    let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken"), profile = await fetchUserSettings(accesstoken);
    if (accesstoken)
      return profile.Data, profile;
  } else
    request.headers.delete("cookie");
  return null;
}
function App() {
  let navigate = (0, import_react3.useNavigate)(), pathname = (0, import_react3.useMatches)()[1].pathname, prof = (0, import_react3.useLoaderData)();
  return (0, import_react4.useEffect)(() => {
    if (pathname === "/" && prof)
      return navigate("/models");
  }, []), /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)(Document, { children: [
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react_hot_toast.Toaster, { position: "bottom-right", reverseOrder: !0 }),
    !prof && pathname === "/" ? /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_jsx_runtime13.Fragment, { children: !prof && /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "flex flex-col justify-center !font-outfit bg-mobBG md:bg-tabBG lg:bg-desktopBG 2xl:bg-largeBG bg-no-repeat bg-cover bg-center bg-fixed", children: [
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(Navbar, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "flex justify-center bg-slate-50", children: /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "w-full md:max-w-7xl px-6", children: /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(HeroSection, {}) }) }),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "flex justify-center", children: /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "flex md:max-w-7xl p-8 md:py-16", children: [
        /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "flex flex-col md:gap-y-24 lg:gap-y-48 xl:gap-y-80 sm:w-1/2", children: [
          /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "flex flex-col md:gap-y-32 lg:gap-y-64 xl:gap-y-80 2xl:gap-y-96", children: [
            /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(RawDataSection, {}),
            /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(TransformerSection, {}),
            /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(DatasetSection, {})
          ] }),
          /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "flex flex-col md:gap-y-24 lg:gap-y-48 xl:gap-y-64 2xl:gap-y-72", children: [
            /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(ModelSection, {}),
            /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(PackageSection, {})
          ] })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(
          "img",
          {
            src: "/imgs/landingPage/LongPipeline.svg",
            alt: "",
            className: "hidden sm:block sm:w-3/5 xl:w-1/2"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(CTASection, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(Footer, {})
    ] }) }) : "",
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.Outlet, {})
  ] });
}
function Document({ children }) {
  return /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("html", { lang: "en", children: [
    /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("head", { children: [
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.Meta, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.Links, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("link", { rel: "preconnect", href: "https://fonts.googleapis.com" }),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(
        "link",
        {
          href: "https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;1,100;1,200;1,300;1,400;1,500;1,600;1,700&family=Outfit:wght@100;200;300;400;500;600;700;800;900&display=swap",
          rel: "stylesheet"
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("script", { async: !0, defer: !0, src: "https://buttons.github.io/buttons.js" }),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.Scripts, {})
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("body", { children: [
      children,
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.ScrollRestoration, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.Scripts, {}),
      /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(import_react3.LiveReload, {})
    ] })
  ] });
}
function ErrorBoundary({ error }) {
  return /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(Document, { children: /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "p-12", children: [
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("span", { className: "text-3xl font-medium", children: "Error" }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("p", { children: error.message }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "text-xl pt-8 font-medium", children: "The stack trace is:" }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("pre", { className: "whitespace-pre-wrap", children: error.stack }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("img", { src: "/error/FunctionalError.gif", alt: "Error", width: "500" }) })
  ] }) });
}
function CatchBoundary() {
  let caught = (0, import_react3.useCatch)();
  return /* @__PURE__ */ (0, import_jsx_runtime13.jsx)(Document, { children: /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "p-12", children: [
    /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("span", { className: "text-3xl font-medium", children: [
      "Status: ",
      caught == null ? void 0 : caught.status
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsxs)("div", { className: "text-xl pt-8 font-medium", children: [
      "Data: ",
      caught == null ? void 0 : caught.data
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("div", { className: "flex justify-center items-center", children: /* @__PURE__ */ (0, import_jsx_runtime13.jsx)("img", { src: "/error/FunctionalError.gif", alt: "Error", width: "500" }) })
  ] }) });
}

// app/routes/api/datasets.server.ts
var datasets_server_exports = {};
__export(datasets_server_exports, {
  fetchDatasetBranch: () => fetchDatasetBranch,
  fetchDatasetReadme: () => fetchDatasetReadme,
  fetchDatasetVersions: () => fetchDatasetVersions,
  fetchDatasets: () => fetchDatasets,
  writeDatasetReadme: () => writeDatasetReadme
});
var backendUrl2 = process.env.NEXT_PUBLIC_BACKEND_URL, makeUrl2 = (path) => `${backendUrl2}${path}`;
async function fetchDatasets(orgId, accessToken) {
  let url = makeUrl2(`org/${orgId}/dataset/all`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchDatasetBranch(orgId, datasetName, accessToken) {
  let url = makeUrl2(`org/${orgId}/dataset/${datasetName}/branch`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchDatasetVersions(orgId, datasetName, accessToken) {
  let url = makeUrl2(`org/${orgId}/dataset/${datasetName}/branch/dev/version`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchDatasetReadme(orgId, datasetName, accessToken) {
  let url = makeUrl2(`org/${orgId}/dataset/${datasetName}/readme/version`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function writeDatasetReadme(orgId, datasetName, content, accessToken) {
  let url = makeUrl2(`org/${orgId}/dataset/${datasetName}/readme`);
  return await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    },
    body: JSON.stringify({
      content,
      file_type: "html"
    })
  }).then((res2) => res2.json());
}

// app/routes/api/models.server.ts
var models_server_exports = {};
__export(models_server_exports, {
  fetchModel: () => fetchModel,
  fetchModelMetrics: () => fetchModelMetrics,
  fetchModelParameters: () => fetchModelParameters,
  fetchModelReadme: () => fetchModelReadme,
  fetchModelVersions: () => fetchModelVersions,
  fetchModels: () => fetchModels,
  writeModelReadme: () => writeModelReadme
});
var backendUrl3 = process.env.NEXT_PUBLIC_BACKEND_URL, makeUrl3 = (path) => `${backendUrl3}${path}`;
async function fetchModels(orgId, accessToken) {
  let url = makeUrl3(`org/${orgId}/model/all`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function fetchModelReadme(orgId, modelName, accessToken) {
  let url = makeUrl3(`org/${orgId}/model/${modelName}/readme/version`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}
async function writeModelReadme(orgId, modelName, content, accessToken) {
  let url = makeUrl3(`org/${orgId}/model/${modelName}/readme`);
  return await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    },
    body: JSON.stringify({
      content,
      file_type: "html"
    })
  }).then((res2) => res2.json());
}
async function fetchModel(orgId, projectId, modelId, accessToken) {
  let url = makeUrl3(
    `${orgId}/project/${projectId}/model/${modelId}/latest/details`
  );
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).data;
}
async function fetchModelVersions(orgId, modelId, accessToken) {
  let url = makeUrl3(`org/${orgId}/model/${modelId}/branch/dev/version`), res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json()), versions = [];
  return res.Data.forEach((version) => {
    versions.push(version.version);
  }), versions;
}
async function fetchModelMetrics(orgId, modelId, version, accessToken) {
  let url = makeUrl3(
    `org/${orgId}/model/${modelId}/branch/dev/version/${version}/log`
  );
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data[0].data;
}
async function fetchModelParameters(orgId, projectId, modelId, version, accessToken) {
  let url = makeUrl3(
    `${orgId}/project/${projectId}/model/${modelId}/${version}/params`
  );
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).data;
}

// app/routes/api/org.server.ts
var org_server_exports = {};
__export(org_server_exports, {
  fetchOrgDetails: () => fetchOrgDetails
});
var backendUrl4 = process.env.NEXT_PUBLIC_BACKEND_URL, makeUrl4 = (path) => `${backendUrl4}${path}`;
async function fetchOrgDetails(orgId, accessToken) {
  let url = makeUrl4(`org/id/${orgId}`);
  return (await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`
    }
  }).then((res2) => res2.json())).Data;
}

// app/routes/markdown/index.tsx
var markdown_exports = {};
__export(markdown_exports, {
  default: () => Index,
  links: () => links2
});
var import_remix_utils = require("remix-utils"), import_quill = __toESM(require_quill());

// node_modules/.pnpm/quill@1.3.7/node_modules/quill/dist/quill.snow.css
var quill_snow_default = "/build/_assets/quill.snow-4ADGMK2W.css";

// app/routes/markdown/index.tsx
var import_jsx_runtime14 = require("react/jsx-runtime"), links2 = () => [
  { rel: "stylesheet", href: quill_snow_default }
];
function Index() {
  return /* @__PURE__ */ (0, import_jsx_runtime14.jsxs)("div", { className: "m-2", children: [
    /* @__PURE__ */ (0, import_jsx_runtime14.jsx)("h1", { className: "text-2xl font-bold", children: "Remix Quill Example!" }),
    /* @__PURE__ */ (0, import_jsx_runtime14.jsx)(import_remix_utils.ClientOnly, { fallback: /* @__PURE__ */ (0, import_jsx_runtime14.jsx)("div", { style: { width: 500, height: 300 } }), children: () => /* @__PURE__ */ (0, import_jsx_runtime14.jsx)(import_quill.default, { defaultValue: "Hello <b>Remix!</b>" }) })
  ] });
}

// app/routes/$username.tsx
var username_exports = {};
__export(username_exports, {
  default: () => UserProfile,
  loader: () => loader2,
  meta: () => meta2
});
var import_react7 = require("@remix-run/react");

// app/components/ui/Avatar.tsx
var Avatar = __toESM(require("@radix-ui/react-avatar")), import_class_variance_authority3 = require("class-variance-authority"), import_jsx_runtime15 = require("react/jsx-runtime"), avatarStyles = (0, import_class_variance_authority3.cva)(
  "flex items-center px-3 py-2 font-medium focus:outline-none",
  {
    variants: {
      intent: {
        primary: "w-6 h-6 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-full capitalize",
        profile: "text-black h-9 rounded justify-center capitalize",
        org: "bg-brand-100 text-black h-9 rounded-full justify-center capitalize"
      },
      fullWidth: {
        true: "w-full"
      }
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: !0
    }
  }
);
function AvatarIcon({ intent, fullWidth, children }) {
  return /* @__PURE__ */ (0, import_jsx_runtime15.jsx)("div", { children: intent === "primary" ? /* @__PURE__ */ (0, import_jsx_runtime15.jsx)("div", { className: "h-full", children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(Avatar.Root, { children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(Avatar.Fallback, { className: avatarStyles({ intent, fullWidth }), children }) }) }) : /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(import_jsx_runtime15.Fragment, { children: intent === "profile" ? /* @__PURE__ */ (0, import_jsx_runtime15.jsx)("div", { className: "px-1", children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(Avatar.Root, { children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(
    Avatar.Fallback,
    {
      className: avatarStyles({ intent, fullWidth }),
      children
    }
  ) }) }) : /* @__PURE__ */ (0, import_jsx_runtime15.jsx)("div", { className: "h-full", children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(Avatar.Root, { children: /* @__PURE__ */ (0, import_jsx_runtime15.jsx)(
    Avatar.Fallback,
    {
      className: avatarStyles({ intent, fullWidth }),
      children
    }
  ) }) }) }) });
}

// app/routes/$username.tsx
var import_lucide_react6 = require("lucide-react");

// app/components/ProfileCard.tsx
var import_class_variance_authority4 = require("class-variance-authority"), import_jsx_runtime16 = require("react/jsx-runtime"), cardStyles = (0, import_class_variance_authority4.cva)(
  "items-center justify-center p-8 text-lg font-medium focus:outline-none",
  {
    variants: {
      intent: {
        profile: "bg-slate-0 text-slate-800 rounded-xl border border-slate-200"
      },
      fullWidth: {
        true: "w-full"
      }
    },
    defaultVariants: {
      intent: "profile",
      fullWidth: !0
    }
  }
);
function ProfileCard({
  intent,
  fullWidth,
  count,
  title
}) {
  return /* @__PURE__ */ (0, import_jsx_runtime16.jsx)("div", { children: /* @__PURE__ */ (0, import_jsx_runtime16.jsxs)("div", { className: cardStyles({ intent, fullWidth }), children: [
    /* @__PURE__ */ (0, import_jsx_runtime16.jsx)("div", { className: "text-3xl", children: count }),
    /* @__PURE__ */ (0, import_jsx_runtime16.jsx)("div", { className: "text-lg", children: title })
  ] }) });
}

// app/components/Navbar.tsx
var import_clsx = __toESM(require("clsx")), import_lucide_react5 = require("lucide-react");

// app/components/ui/Input.tsx
var import_class_variance_authority5 = require("class-variance-authority"), import_lucide_react3 = require("lucide-react"), import_jsx_runtime17 = require("react/jsx-runtime"), inputStyles = (0, import_class_variance_authority5.cva)(
  "flex items-center justify-center px-4 py-2 focus:outline-none",
  {
    variants: {
      intent: {
        primary: "bg-slate-100 text-slate-600 rounded h-9 border-slate-600 border hover:border-blue-750 focus:border-blue-750",
        search: "bg-transparent text-sm text-slate-400 !justify-start rounded border-slate-200 border hover:border-blue-750 focus:border-blue-750",
        read: "bg-slate-100 text-slate-900 rounded h-9 border-slate-600 border hover:border-blue-750 focus:border-blue-750"
      },
      fullWidth: {
        true: "w-full",
        false: "w-[18rem]"
      },
      type: {
        text: "text",
        password: "password",
        email: "email",
        number: "number"
      }
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: !0,
      type: "text"
    }
  }
);
function Input({
  intent,
  fullWidth,
  type,
  name,
  placeholder,
  onChange,
  ariaLabel,
  dataTestid,
  required
}) {
  return /* @__PURE__ */ (0, import_jsx_runtime17.jsx)("div", { children: intent === "read" ? /* @__PURE__ */ (0, import_jsx_runtime17.jsx)(
    "input",
    {
      required,
      type,
      className: inputStyles({ intent, fullWidth }),
      onChange,
      "aria-label": ariaLabel,
      "data-testid": dataTestid,
      defaultValue: placeholder,
      disabled: !0
    }
  ) : /* @__PURE__ */ (0, import_jsx_runtime17.jsx)("div", { children: intent === "search" ? /* @__PURE__ */ (0, import_jsx_runtime17.jsxs)("div", { className: inputStyles({ intent, fullWidth }), children: [
    /* @__PURE__ */ (0, import_jsx_runtime17.jsx)(import_lucide_react3.Search, { className: "w-4 h-4" }),
    /* @__PURE__ */ (0, import_jsx_runtime17.jsx)(
      "input",
      {
        required,
        type,
        className: "border-none focus:outline-none pl-2 w-full",
        placeholder,
        onChange,
        "aria-label": ariaLabel,
        "data-testid": dataTestid
      }
    )
  ] }) : /* @__PURE__ */ (0, import_jsx_runtime17.jsx)(import_jsx_runtime17.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime17.jsx)(
    "input",
    {
      required,
      type,
      name,
      className: inputStyles({ intent, fullWidth }),
      placeholder,
      onChange,
      "aria-label": ariaLabel,
      "data-testid": dataTestid
    }
  ) }) }) });
}
var Input_default = Input;

// app/components/Navbar.tsx
var import_react6 = require("@remix-run/react"), import_class_variance_authority7 = require("class-variance-authority");

// app/components/ui/Dropdown.tsx
var import_class_variance_authority6 = require("class-variance-authority"), DropdownMenu = __toESM(require("@radix-ui/react-dropdown-menu")), import_lucide_react4 = require("lucide-react"), import_react5 = require("@remix-run/react"), import_jsx_runtime18 = require("react/jsx-runtime"), dropdownTrigger = (0, import_class_variance_authority6.cva)("focus:outline-none z-50", {
  variants: {
    intent: {
      primary: "flex justify-between items-center z-50",
      orgType: "flex justify-between items-center text-sm text-slate-600 rounded border border-slate-600 h-8 px-2 z-50",
      branch: "flex justify-between items-center text-sm text-slate-600 rounded border border-slate-600 h-8 px-2 z-50"
    },
    fullWidth: {
      true: "w-full",
      false: "w-28"
    }
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: !0
  }
}), dropdownContent = (0, import_class_variance_authority6.cva)("focus:outline-none", {
  variants: {
    color: {
      primary: "bg-slate-100 justify-center items-center text-sm text-slate-600 rounded border border-slate-200 z-50 shadow"
    },
    contentWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    color: "primary",
    contentWidth: !0
  }
}), dropdownItems = (0, import_class_variance_authority6.cva)("focus:outline-none", {
  variants: {
    space: {
      primary: "flex px-3 py-2 text-sm text-base justify-left items-center rounded outline-none hover:bg-slate-200 cursor-pointer"
    }
  },
  defaultVariants: {
    space: "primary"
  }
});
function Dropdown({
  intent,
  fullWidth,
  color,
  contentWidth,
  space,
  children
}) {
  let navigate = (0, import_react5.useNavigate)();
  return /* @__PURE__ */ (0, import_jsx_runtime18.jsx)("div", { children: /* @__PURE__ */ (0, import_jsx_runtime18.jsxs)(DropdownMenu.Root, { children: [
    /* @__PURE__ */ (0, import_jsx_runtime18.jsxs)(
      DropdownMenu.Trigger,
      {
        className: dropdownTrigger({ intent, fullWidth }),
        children: [
          children,
          intent !== "primary" ? /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(import_lucide_react4.ChevronDown, { className: "text-slate-400" }) : ""
        ]
      }
    ),
    intent === "primary" ? /* @__PURE__ */ (0, import_jsx_runtime18.jsxs)(
      DropdownMenu.Content,
      {
        sideOffset: 7,
        align: "end",
        className: dropdownContent({ color, contentWidth }),
        children: [
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(
            DropdownMenu.Item,
            {
              className: dropdownItems({ space }),
              onClick: () => {
                navigate("/profile");
              },
              children: "Profile"
            }
          ),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(
            DropdownMenu.Item,
            {
              className: dropdownItems({ space }),
              onClick: () => {
                navigate("/settings");
              },
              children: "Settings"
            }
          ),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(
            DropdownMenu.Item,
            {
              className: dropdownItems({ space }),
              onClick: () => {
                navigate("/logout");
              },
              children: "Sign out"
            }
          )
        ]
      }
    ) : /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(import_jsx_runtime18.Fragment, { children: intent === "orgType" ? /* @__PURE__ */ (0, import_jsx_runtime18.jsxs)(
      DropdownMenu.Content,
      {
        sideOffset: 7,
        className: dropdownContent({ color, contentWidth }),
        children: [
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(DropdownMenu.Item, { className: dropdownItems({ space }), children: "Company" }),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(DropdownMenu.Item, { className: dropdownItems({ space }), children: "University" }),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(DropdownMenu.Item, { className: dropdownItems({ space }), children: "Classroom" }),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(DropdownMenu.Item, { className: dropdownItems({ space }), children: "Non-profit" }),
          /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(DropdownMenu.Item, { className: dropdownItems({ space }), children: "Community" })
        ]
      }
    ) : /* @__PURE__ */ (0, import_jsx_runtime18.jsx)(
      DropdownMenu.Content,
      {
        sideOffset: 7,
        className: dropdownContent({ color, contentWidth })
      }
    ) })
  ] }) });
}

// app/components/Navbar.tsx
var import_jsx_runtime19 = require("react/jsx-runtime");
function linkCss(currentPage) {
  return (0, import_clsx.default)(
    currentPage ? " text-blue-700 " : " text-slate-600 ",
    " hover:text-blue-700 flex justify-center items-center px-5 cursor-pointer "
  );
}
var navbarStyles = (0, import_class_variance_authority7.cva)(
  "fixed z-20 h-18 px-12 py-4 w-full bg-slate-0 flex justify-between text-sm font-medium border-b-2 border-slate-200",
  {
    variants: {
      intent: {
        loggedIn: "",
        loggedOut: ""
      },
      fullWidth: {
        true: "w-full"
      }
    },
    defaultVariants: {
      intent: "loggedOut",
      fullWidth: !0
    }
  }
);
function NavBar({ intent, fullWidth, user }) {
  let navigate = (0, import_react6.useNavigate)(), pathname = (0, import_react6.useMatches)()[1].pathname;
  return /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(import_jsx_runtime19.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)("div", { className: navbarStyles({ intent, fullWidth }), children: [
    /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)("div", { className: "flex", children: [
      /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("a", { href: "/models", className: "flex items-center justify-center pr-8", children: /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("img", { src: "/LogoWText.svg", alt: "Logo", width: "140", height: "96" }) }),
      /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(
        Input_default,
        {
          intent: "search",
          placeholder: "Search models, datasets, users...",
          fullWidth: !1
        }
      )
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)("div", { className: "flex justify-center items-center", children: [
      /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)(
        "div",
        {
          onClick: () => {
            navigate("/models");
          },
          className: `${linkCss(pathname === "/models")}`,
          children: [
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(import_lucide_react5.Box, { className: "w-4 h-4" }),
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("span", { className: "pl-2", children: "Models" })
          ]
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)(
        "div",
        {
          onClick: () => {
            navigate("/datasets");
          },
          className: `${linkCss(pathname === "/datasets")}`,
          children: [
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(import_lucide_react5.Database, { className: "w-4 h-4" }),
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("span", { className: "pl-2", children: "Datasets" })
          ]
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)(
        "a",
        {
          href: "https://docs.pureml.com",
          className: "flex justify-center items-center cursor-pointer px-5 hover:text-blue-700 border-r-2 border-slate-slate-200 font-medium text-slate-600",
          children: [
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(import_lucide_react5.File, { className: "w-4 h-4" }),
            /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("span", { className: "pl-2", children: "Docs" })
          ]
        }
      ),
      intent === "loggedOut" ? /* @__PURE__ */ (0, import_jsx_runtime19.jsxs)(import_jsx_runtime19.Fragment, { children: [
        /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("div", { className: "w-full flex justify-center items-center px-5", children: /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("a", { href: "/auth/signin", children: "Sign in" }) }),
        /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(Button_default, { intent: "primary", icon: "", children: "Sign up" })
      ] }) : /* @__PURE__ */ (0, import_jsx_runtime19.jsx)("div", { className: "w-full flex justify-center items-center px-5", children: /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(Dropdown, { intent: "primary", children: /* @__PURE__ */ (0, import_jsx_runtime19.jsx)(AvatarIcon, { children: user }) }) })
    ] })
  ] }) });
}

// app/Error404.tsx
var import_jsx_runtime20 = require("react/jsx-runtime");
function Error2() {
  return /* @__PURE__ */ (0, import_jsx_runtime20.jsx)("div", { className: "w-screen h-full bg-slate-0 text-slate-800 font-medium", children: /* @__PURE__ */ (0, import_jsx_runtime20.jsx)("div", { className: "flex px-12 pt-48 pb-12 md:justify-center items-center bg-slate-0", children: /* @__PURE__ */ (0, import_jsx_runtime20.jsx)("img", { src: "/error/Error404.gif", alt: "Error", width: "500", height: "400" }) }) });
}

// app/routes/$username.tsx
var import_jsx_runtime21 = require("react/jsx-runtime"), meta2 = () => ({
  charset: "utf-8",
  title: "Profile | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader2({ params, request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken"), publicProfile = await fetchPublicProfile(params.username), userProfile = await fetchUserSettings(accesstoken);
  return publicProfile ? { publicProfile, userProfile } : null;
}
function UserProfile() {
  var _a, _b, _c, _d, _e, _f, _g;
  let userProfileData = (0, import_react7.useLoaderData)();
  return userProfileData ? /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(import_react7.Meta, {}) }),
    userProfileData.userProfile ? /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(
      NavBar,
      {
        intent: "loggedIn",
        user: userProfileData.userProfile[0].name.charAt(0).toUpperCase()
      }
    ) : /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { className: "flex px-12 pt-24 pb-12 text-slate-800 font-medium", children: [
      /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { className: "h-full w-28 md:w-36 lg:w-56 2xl:w-96", children: [
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg", children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(AvatarIcon, { intent: "profile", children: ((_a = userProfileData.publicProfile[0]) == null ? void 0 : _a.name.charAt(0).toUpperCase()) || "User" }) }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "pt-6 font-semibold text-base text-slate-900", children: ((_b = userProfileData.publicProfile[0]) == null ? void 0 : _b.name) || "Name" }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "pb-6 text-base font-normal", children: ((_c = userProfileData.publicProfile[0]) == null ? void 0 : _c.email) || "Email" }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(Button_default, { "aria-label": "follow", intent: "primary", icon: "", children: "Follow" }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { className: "flex justify-between text-base pt-8", children: [
          /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("span", { children: "Bio" }),
          /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(import_lucide_react6.Edit2, {})
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "font-medium text-base text-slate-600", children: ((_d = userProfileData.publicProfile[0]) == null ? void 0 : _d.bio) || "Add your bio" })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { className: "pl-12 w-full", children: [
        /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "pb-6", children: "Overview" }),
        /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { className: "flex w-full", children: [
          /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "pr-4 w-full", children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(
            ProfileCard,
            {
              title: "Projects",
              count: ((_e = userProfileData.publicProfile[0]) == null ? void 0 : _e.number_of_projects) || "0"
            }
          ) }),
          /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "pr-4 w-full", children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(
            ProfileCard,
            {
              title: "Models",
              count: ((_f = userProfileData.publicProfile[0]) == null ? void 0 : _f.number_of_models) || "0"
            }
          ) }),
          /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("div", { className: "w-full", children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(
            ProfileCard,
            {
              title: "Datasets",
              count: ((_g = userProfileData.publicProfile[0]) == null ? void 0 : _g.number_of_datasets) || "0"
            }
          ) })
        ] })
      ] })
    ] })
  ] }) : /* @__PURE__ */ (0, import_jsx_runtime21.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime21.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(import_react7.Meta, {}) }),
    userProfileData.userProfile ? /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(
      NavBar,
      {
        intent: "loggedIn",
        user: userProfileData.userProfile[0].name.charAt(0).toUpperCase()
      }
    ) : /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime21.jsx)(Error2, {})
  ] });
}

// app/routes/datasets.tsx
var datasets_exports = {};
__export(datasets_exports, {
  default: () => DatasetsLayout,
  loader: () => loader3,
  meta: () => meta3
});
var import_react8 = require("@remix-run/react");
var import_jsx_runtime22 = require("react/jsx-runtime"), meta3 = () => ({
  charset: "utf-8",
  title: "Datasets | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader3({ request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken");
  return await fetchUserSettings(accesstoken);
}
function DatasetsLayout() {
  let prof = (0, import_react8.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime22.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime22.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime22.jsx)(import_react8.Meta, {}) }),
    prof ? /* @__PURE__ */ (0, import_jsx_runtime22.jsx)(NavBar, { intent: "loggedIn", user: prof[0].name.charAt(0).toUpperCase() }) : /* @__PURE__ */ (0, import_jsx_runtime22.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime22.jsx)("div", { className: "px-12 pt-16 pb-12 h-full w-screen", children: /* @__PURE__ */ (0, import_jsx_runtime22.jsx)(import_react8.Outlet, {}) })
  ] });
}

// app/routes/datasets/EmptyDataset.tsx
var EmptyDataset_exports = {};
__export(EmptyDataset_exports, {
  default: () => EmptyDataset
});
var import_class_variance_authority8 = require("class-variance-authority"), import_jsx_runtime23 = require("react/jsx-runtime"), emptyStyles = (0, import_class_variance_authority8.cva)("h-full flex justify-center items-center", {
  variants: {
    intent: {
      primary: "flex justify-center items-center bg-slate-150 drop-shadow-3xl rounded-lg"
    },
    fullWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: !0
  }
}), codeStyles = (0, import_class_variance_authority8.cva)("h-fit ml-8 bg-slate-150 text-slate-400", {
  variants: {
    intent: {
      primary: ""
    },
    fullWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: !0
  }
});
function EmptyDataset() {
  return /* @__PURE__ */ (0, import_jsx_runtime23.jsx)("div", { className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 min-w-72", children: /* @__PURE__ */ (0, import_jsx_runtime23.jsxs)("div", { className: "rounded-lg border-2 border-slate-200 px-6 py-4", children: [
    /* @__PURE__ */ (0, import_jsx_runtime23.jsx)("div", { className: "font-medium text-sm pb-6", children: "There are no datasets yet" }),
    /* @__PURE__ */ (0, import_jsx_runtime23.jsx)("div", { className: "rounded-lg h-2 bg-slate-200 w-1/3" }),
    /* @__PURE__ */ (0, import_jsx_runtime23.jsx)("div", { className: "pt-2" }),
    /* @__PURE__ */ (0, import_jsx_runtime23.jsx)("div", { className: "rounded-lg h-2 bg-slate-200 w-2/3" })
  ] }) });
}

// app/routes/datasets/index.tsx
var datasets_exports2 = {};
__export(datasets_exports2, {
  default: () => Index2,
  loader: () => loader4
});
var import_react9 = require("@remix-run/react");

// app/components/Card.tsx
var import_class_variance_authority9 = require("class-variance-authority"), import_lucide_react7 = require("lucide-react");
var import_jsx_runtime24 = require("react/jsx-runtime"), cardStyles2 = (0, import_class_variance_authority9.cva)(
  "items-center justify-center px-6 py-4 text-lg font-normal focus:outline-none cursor-pointer",
  {
    variants: {
      intent: {
        modelCard: "text-slate-600 rounded-lg border-2 border-slate-200 hover:bg-slate-50",
        datasetCard: "text-slate-600 rounded-lg border-2 border-slate-200 hover:bg-slate-50"
      },
      fullWidth: {
        true: "w-full"
      }
    },
    defaultVariants: {
      intent: "modelCard",
      fullWidth: !0
    }
  }
);
function Card({
  intent,
  fullWidth,
  name,
  description,
  tag2,
  onClick
}) {
  return /* @__PURE__ */ (0, import_jsx_runtime24.jsxs)("div", { onClick, className: cardStyles2({ intent, fullWidth }), children: [
    /* @__PURE__ */ (0, import_jsx_runtime24.jsx)("header", { className: "pb-0 text-slate-800", children: /* @__PURE__ */ (0, import_jsx_runtime24.jsxs)("div", { className: "flex items-center", children: [
      intent === "modelCard" ? /* @__PURE__ */ (0, import_jsx_runtime24.jsx)(import_lucide_react7.Box, { className: "w-4" }) : /* @__PURE__ */ (0, import_jsx_runtime24.jsx)(import_lucide_react7.Database, { className: "w-4" }),
      /* @__PURE__ */ (0, import_jsx_runtime24.jsx)("span", { className: "ml-2 truncate text-sm font-medium", children: name })
    ] }) }),
    /* @__PURE__ */ (0, import_jsx_runtime24.jsx)("div", { className: "text-xs font-normal truncate pt-2", children: description }),
    /* @__PURE__ */ (0, import_jsx_runtime24.jsx)("div", { className: "flex pt-4", children: intent === "modelCard" ? /* @__PURE__ */ (0, import_jsx_runtime24.jsx)(Tag_default, { intent: "modelTag", children: tag2 }) : /* @__PURE__ */ (0, import_jsx_runtime24.jsx)(Tag_default, { intent: "datasetTag", children: tag2 }) })
  ] });
}

// app/routes/datasets/index.tsx
var import_jsx_runtime25 = require("react/jsx-runtime");
async function loader4({ request }) {
  let session = await getSession(request.headers.get("Cookie")), orgId = session.get("orgId");
  return { datasets: await fetchDatasets(
    session.get("orgId"),
    session.get("accessToken")
  ), orgId };
}
function Index2() {
  let datasetData = (0, import_react9.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime25.jsxs)("div", { id: "datasets", children: [
    /* @__PURE__ */ (0, import_jsx_runtime25.jsx)("div", { className: "flex justify-between font-medium text-slate-800 text-base pt-6", children: "Datasets" }),
    datasetData ? /* @__PURE__ */ (0, import_jsx_runtime25.jsx)(import_jsx_runtime25.Fragment, { children: datasetData.datasets[0].length !== 0 ? /* @__PURE__ */ (0, import_jsx_runtime25.jsx)(
      "div",
      {
        className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72",
        children: datasetData.datasets.map((dataset) => /* @__PURE__ */ (0, import_jsx_runtime25.jsx)(
          import_react9.Link,
          {
            to: `/org/${datasetData.orgId}/datasets/${dataset.name}`,
            children: /* @__PURE__ */ (0, import_jsx_runtime25.jsx)(
              Card,
              {
                intent: "datasetCard",
                name: dataset.name,
                description: `Updated by ${dataset.updated_by.handle}`,
                tag2: dataset.created_by.handle
              },
              dataset.updated_at
            )
          },
          dataset.id
        ))
      },
      "0"
    ) : /* @__PURE__ */ (0, import_jsx_runtime25.jsx)(EmptyDataset, {}) }) : "All public datasets shown here"
  ] });
}

// app/routes/settings.tsx
var settings_exports = {};
__export(settings_exports, {
  default: () => SettingsLayout,
  loader: () => loader5,
  meta: () => meta4
});
var import_react10 = require("@remix-run/react");
var import_jsx_runtime26 = require("react/jsx-runtime"), meta4 = () => ({
  charset: "utf-8",
  title: "Settings | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader5({ request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken");
  return await fetchUserSettings(accesstoken);
}
function SettingsLayout() {
  let prof = (0, import_react10.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime26.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime26.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime26.jsx)(import_react10.Meta, {}) }),
    prof ? /* @__PURE__ */ (0, import_jsx_runtime26.jsx)(NavBar, { intent: "loggedIn", user: prof[0].name.charAt(0).toUpperCase() }) : /* @__PURE__ */ (0, import_jsx_runtime26.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime26.jsxs)("div", { className: "pt-20 pb-12 h-full w-screen", children: [
      /* @__PURE__ */ (0, import_jsx_runtime26.jsx)("div", { className: "px-12 flex justify-between font-medium text-slate-800 text-base", children: "Settings" }),
      /* @__PURE__ */ (0, import_jsx_runtime26.jsx)(import_react10.Outlet, {})
    ] })
  ] });
}

// app/routes/settings/account/index.tsx
var account_exports = {};
__export(account_exports, {
  action: () => action,
  default: () => Index3
});
var import_node3 = require("@remix-run/node");

// app/components/Tabbar.tsx
var import_clsx2 = __toESM(require("clsx")), import_class_variance_authority10 = require("class-variance-authority"), import_react11 = require("@remix-run/react"), import_jsx_runtime27 = require("react/jsx-runtime"), tabStyles = (0, import_class_variance_authority10.cva)("text-zinc-400 font-medium flex bg-slate-0 sticky z-10", {
  variants: {
    intent: {
      primaryModelTab: "pt-4 border-b-2 border-slate-100 top-44",
      primaryDatasetTab: "pt-4 border-b-2 border-slate-100 top-28",
      primarySettingTab: "pt-4 border-b-2 border-slate-100 top-28",
      modelTab: "pt-8 top-[16rem]",
      datasetTab: "pt-8 top-[11.7rem]",
      modelReviewTab: "pt-8 top-[16rem]",
      datasetReviewTab: "pt-8 top-[11.7rem]"
    },
    fullWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    intent: "primaryModelTab",
    fullWidth: !0
  }
});
function secondaryLinkCss(currentPage) {
  return (0, import_clsx2.default)(
    currentPage ? "text-white" : "text-slate-600",
    "flex justify-center items-center"
  );
}
function TabBar({ intent, fullWidth, tab }) {
  let path = (0, import_react11.useMatches)()[2].pathname, pathname = decodeURI(path.slice(1)), orgId = pathname.split("/")[1], modelId = pathname.split("/")[3], datasetId = pathname.split("/")[3], primaryModelTabs = [
    {
      id: "modelCard",
      name: "Model Card",
      hyperlink: `/org/${orgId}/models/${modelId}`
    },
    {
      id: "versions",
      name: "Versions",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/metrics`
    },
    {
      id: "review",
      name: "Review",
      hyperlink: `/org/${orgId}/models/${modelId}/review`
    }
  ], primaryDatasetTabs = [
    {
      id: "datasetCard",
      name: "Dataset Card",
      hyperlink: `/org/${orgId}/datasets/${datasetId}`
    },
    {
      id: "versions",
      name: "Versions",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/datalineage`
    },
    {
      id: "review",
      name: "Review",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review`
    }
  ], primarySettingTabs = [
    {
      id: "profile",
      name: "Profile",
      hyperlink: "/settings"
    },
    {
      id: "account",
      name: "Account",
      hyperlink: "/settings/account"
    },
    {
      id: "members",
      name: "Members",
      hyperlink: "/settings/members"
    }
  ], modelTabs = [
    {
      id: "metrics",
      name: "Metrics",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/metrics`
    },
    {
      id: "graphs",
      name: "Graphs",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/graphs`
    }
  ], datasetTabs = [
    {
      id: "datalineage",
      name: "Data Lineage",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/datalineage`
    }
  ], modelReviewTabs = [
    {
      id: "metrics",
      name: "Metrics",
      hyperlink: `/org/${orgId}/models/${modelId}/review/commit`
    }
  ], datasetReviewTabs = [
    {
      id: "datalineage",
      name: "Data Lineage",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review/commit`
    }
  ];
  return /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("div", { className: tabStyles({ intent, fullWidth }), children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("div", { className: "flex px-10", children: intent === "primaryModelTab" || intent === "primaryDatasetTab" || intent === "primarySettingTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: intent === "primaryModelTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(primaryModelTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === primaryModelTabs[key].id ? "text-blue-700 border-b-2 border-blue-700" : "text-slate-600"} p-4`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_react11.Link, { to: primaryModelTabs[key].hyperlink, children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: primaryModelTabs[key].name }) })
    },
    key
  )) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: intent === "primaryDatasetTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(primaryDatasetTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === primaryDatasetTabs[key].id ? "text-blue-700 border-b-2 border-blue-700" : "text-slate-600"} p-4`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_react11.Link, { to: primaryDatasetTabs[key].hyperlink, children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: primaryDatasetTabs[key].name }) })
    },
    key
  )) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(primarySettingTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === primarySettingTabs[key].id ? "text-blue-700 border-b-2 border-blue-700" : "text-slate-600"} p-4`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_react11.Link, { to: primarySettingTabs[key].hyperlink, children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: primarySettingTabs[key].name }) })
    },
    key
  )) }) }) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: intent === "modelTab" || intent === "datasetTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("div", { className: "flex", children: intent === "modelTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(modelTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === modelTabs[key].id ? "bg-blue-700 rounded text-white" : ""} px-4 py-2`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
        import_react11.Link,
        {
          to: modelTabs[key].hyperlink,
          className: `${secondaryLinkCss(
            tab === modelTabs[key].id
          )}`,
          children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: modelTabs[key].name })
        }
      )
    },
    key
  )) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(datasetTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === datasetTabs[key].id ? "bg-blue-700 rounded text-white" : ""} px-4 py-2`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
        import_react11.Link,
        {
          to: datasetTabs[key].hyperlink,
          className: `${secondaryLinkCss(
            tab === datasetTabs[key].id
          )}`,
          children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: datasetTabs[key].name })
        }
      )
    },
    key
  )) }) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("div", { className: "flex", children: intent === "modelReviewTab" ? /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(modelReviewTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === modelReviewTabs[key].id ? "bg-blue-700 rounded text-white" : ""} px-4 py-2`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
        import_react11.Link,
        {
          to: modelReviewTabs[key].hyperlink,
          className: `${secondaryLinkCss(
            tab === modelReviewTabs[key].id
          )}`,
          children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: modelReviewTabs[key].name })
        }
      )
    },
    key
  )) }) : /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(import_jsx_runtime27.Fragment, { children: Object.keys(datasetReviewTabs).map((key) => /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
    "div",
    {
      className: `${tab === datasetReviewTabs[key].id ? "bg-blue-700 rounded text-white" : ""} px-4 py-2`,
      children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)(
        import_react11.Link,
        {
          to: datasetReviewTabs[key].hyperlink,
          className: `${secondaryLinkCss(
            tab === datasetReviewTabs[key].id
          )}`,
          children: /* @__PURE__ */ (0, import_jsx_runtime27.jsx)("span", { children: datasetReviewTabs[key].name })
        }
      )
    },
    key
  )) }) }) }) }) });
}

// app/routes/settings/account/index.tsx
var import_jsx_runtime28 = require("react/jsx-runtime"), action = async ({ request }) => {
  let form = await request.formData(), orgname = form.get("orgname"), orgdesc = form.get("desc");
  return (0, import_node3.redirect)("/settings");
};
function Index3() {
  return /* @__PURE__ */ (0, import_jsx_runtime28.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime28.jsx)(TabBar, { intent: "primarySettingTab", tab: "account" }),
    /* @__PURE__ */ (0, import_jsx_runtime28.jsxs)("form", { method: "post", className: "pt-8 px-12 w-2/3", children: [
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime28.jsxs)("label", { htmlFor: "orgname", className: "text-base pb-1", children: [
        "Email",
        /* @__PURE__ */ (0, import_jsx_runtime28.jsx)(
          Input_default,
          {
            intent: "primary",
            type: "email",
            name: "email",
            placeholder: "Enter email...",
            "aria-label": "email",
            "data-testid": "email",
            required: !0
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)("div", { className: "pb-8", children: /* @__PURE__ */ (0, import_jsx_runtime28.jsxs)("label", { htmlFor: "orgdesc", className: "text-base pb-1", children: [
        "Organization domain name",
        /* @__PURE__ */ (0, import_jsx_runtime28.jsx)(
          Input_default,
          {
            intent: "primary",
            type: "text",
            name: "orgdesc",
            placeholder: "Enter Org Description...",
            "aria-label": "org-desc",
            "data-testid": "org-desc",
            required: !0
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)(Button_default, { icon: "", fullWidth: !1, children: "Save changes" }),
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)("div", { className: "pt-12 text-slate-800", children: "Delete Organization" }),
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)("div", { className: "text-base pb-8 text-slate-400", children: "Delete this organization permanently, this action is irreversible All its repositories (models, datasets) will be deleted." }),
      /* @__PURE__ */ (0, import_jsx_runtime28.jsx)(Button_default, { intent: "danger", icon: "", fullWidth: !1, children: "Delete Organization" })
    ] })
  ] });
}

// app/routes/settings/members/index.tsx
var members_exports = {};
__export(members_exports, {
  action: () => action2,
  default: () => Index4
});
var import_node4 = require("@remix-run/node");
var import_jsx_runtime29 = require("react/jsx-runtime"), action2 = async ({ request }) => {
  let form = await request.formData(), name = form.get("name"), desc = form.get("desc");
  return (0, import_node4.redirect)("/org/oegId/settings");
};
function Index4() {
  return /* @__PURE__ */ (0, import_jsx_runtime29.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime29.jsx)(TabBar, { intent: "primarySettingTab", tab: "members" }),
    /* @__PURE__ */ (0, import_jsx_runtime29.jsx)("div", { className: "pt-8 px-12 w-2/3", children: "Members List will be shown here" })
  ] });
}

// app/routes/settings/index.tsx
var settings_exports2 = {};
__export(settings_exports2, {
  action: () => action3,
  default: () => Index5
});
var import_node5 = require("@remix-run/node");
var import_jsx_runtime30 = require("react/jsx-runtime"), action3 = async ({ request }) => {
  let form = await request.formData(), orgname = form.get("orgname"), orgdesc = form.get("desc");
  return (0, import_node5.redirect)("/settings");
};
function Index5() {
  return /* @__PURE__ */ (0, import_jsx_runtime30.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime30.jsx)(TabBar, { intent: "primarySettingTab", tab: "profile" }),
    /* @__PURE__ */ (0, import_jsx_runtime30.jsxs)("form", { method: "post", className: "p-12 w-2/3", children: [
      /* @__PURE__ */ (0, import_jsx_runtime30.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime30.jsxs)("label", { htmlFor: "orgname", className: "text-base pb-1", children: [
        "Org Name",
        /* @__PURE__ */ (0, import_jsx_runtime30.jsx)(
          Input_default,
          {
            intent: "primary",
            type: "text",
            name: "orgname",
            placeholder: "Enter Org Name...",
            "aria-label": "org-name",
            "data-testid": "org-name",
            required: !0
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime30.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime30.jsxs)("label", { htmlFor: "orgdesc", className: "text-base pb-1", children: [
        "Org Description",
        /* @__PURE__ */ (0, import_jsx_runtime30.jsx)(
          Input_default,
          {
            intent: "primary",
            type: "text",
            name: "orgdesc",
            placeholder: "Enter Org Description...",
            "aria-label": "org-desc",
            "data-testid": "org-desc",
            required: !0
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime30.jsx)("div", { className: "pb-8", children: /* @__PURE__ */ (0, import_jsx_runtime30.jsxs)("label", { htmlFor: "orgtype", className: "text-base pb-1", children: [
        "Org Type",
        /* @__PURE__ */ (0, import_jsx_runtime30.jsx)(Dropdown, { intent: "orgType", fullWidth: !1, children: "Choose" })
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime30.jsx)(Button_default, { icon: "", fullWidth: !1, children: "Save changes" })
    ] })
  ] });
}

// app/routes/contact.tsx
var contact_exports = {};
__export(contact_exports, {
  action: () => action4,
  default: () => Contact,
  meta: () => meta5
});
var import_node6 = require("@remix-run/node"), import_react12 = require("@remix-run/react");
var import_jsx_runtime31 = require("react/jsx-runtime"), action4 = async ({ request }) => {
  let comment = (await request.formData()).get("comment"), accessToken = (await getSession(request.headers.get("Cookie"))).get("accessToken"), email = (await fetchUserSettings(accessToken))[0].email;
  return (0, import_node6.redirect)("/models", {});
}, meta5 = () => ({
  charset: "utf-8",
  title: "Contact Us | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function Contact() {
  let navigate = (0, import_react12.useNavigate)(), data = (0, import_react12.useActionData)();
  return /* @__PURE__ */ (0, import_jsx_runtime31.jsxs)("div", { className: "flex h-screen justify-center items-center bg-zinc-800 opacity-60", children: [
    /* @__PURE__ */ (0, import_jsx_runtime31.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime31.jsx)(import_react12.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime31.jsxs)("div", { className: "bg-slate-0 p-4 rounded-lg w-[20rem]", children: [
      /* @__PURE__ */ (0, import_jsx_runtime31.jsx)("div", { className: "text-slate-800 font-medium pb-4", children: "Contact Us" }),
      /* @__PURE__ */ (0, import_jsx_runtime31.jsxs)("form", { method: "post", children: [
        /* @__PURE__ */ (0, import_jsx_runtime31.jsx)("label", { htmlFor: "comment", children: /* @__PURE__ */ (0, import_jsx_runtime31.jsx)(
          "textarea",
          {
            typeof: "text",
            name: "comment",
            required: !0,
            className: "whitespace-pre-line w-full bg-transparent text-sm border border-slate-600 rounded-md h-full hover:border-blue-750 focus:outline-none focus:border-blue-750 max-h-[200px] p-4",
            placeholder: "Add your query or feedback here..."
          }
        ) }),
        /* @__PURE__ */ (0, import_jsx_runtime31.jsx)("div", { className: "pt-12 grid justify-items-end w-full", children: /* @__PURE__ */ (0, import_jsx_runtime31.jsxs)("div", { className: "flex justify-between w-1/2", children: [
          /* @__PURE__ */ (0, import_jsx_runtime31.jsx)(
            Button_default,
            {
              icon: "",
              fullWidth: !1,
              intent: "secondary",
              onClick: () => {
                navigate("/models");
              },
              children: "No"
            }
          ),
          /* @__PURE__ */ (0, import_jsx_runtime31.jsx)(Button_default, { icon: "", intent: "danger", fullWidth: !1, type: "submit", children: "Submit" })
        ] }) })
      ] })
    ] })
  ] });
}

// app/routes/profile.tsx
var profile_exports = {};
__export(profile_exports, {
  default: () => ProfileRedirection,
  loader: () => loader6
});
var import_node7 = require("@remix-run/node"), import_react13 = require("@remix-run/react");
var import_jsx_runtime32 = require("react/jsx-runtime");
async function loader6({ request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken"), username = (await fetchUserSettings(accesstoken))[0].handle, profile = await fetchPublicProfile(username);
  return (0, import_node7.redirect)(`/${profile[0].handle}`, {});
}
function ProfileRedirection() {
  let prof = (0, import_react13.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime32.jsx)(import_jsx_runtime32.Fragment, {});
}

// app/routes/logout.tsx
var logout_exports = {};
__export(logout_exports, {
  action: () => action5,
  default: () => Logout,
  meta: () => meta6
});
var import_node8 = require("@remix-run/node"), import_react14 = require("@remix-run/react"), import_react_hot_toast2 = __toESM(require("react-hot-toast"));
var import_jsx_runtime33 = require("react/jsx-runtime"), action5 = async ({ request }) => {
  let session = await getSession(request.headers.get("Cookie"));
  return import_react_hot_toast2.default.success("Signed Out Successfully!"), (0, import_node8.redirect)("/", {
    headers: {
      "Set-Cookie": await destroySession(session)
    }
  });
}, meta6 = () => ({
  charset: "utf-8",
  title: "Logout | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function Logout() {
  let navigate = (0, import_react14.useNavigate)();
  return /* @__PURE__ */ (0, import_jsx_runtime33.jsxs)("div", { className: "flex h-screen justify-center items-center bg-zinc-800 opacity-60", children: [
    /* @__PURE__ */ (0, import_jsx_runtime33.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime33.jsx)(import_react14.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime33.jsxs)("div", { className: "bg-slate-0 p-4 rounded-lg", children: [
      /* @__PURE__ */ (0, import_jsx_runtime33.jsx)("div", { className: "text-slate-800 font-medium", children: "Are you sure you want to Sign out?" }),
      /* @__PURE__ */ (0, import_jsx_runtime33.jsx)("div", { className: "pt-12 grid justify-items-end w-full", children: /* @__PURE__ */ (0, import_jsx_runtime33.jsxs)("div", { className: "flex justify-between w-1/2", children: [
        /* @__PURE__ */ (0, import_jsx_runtime33.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            intent: "secondary",
            onClick: () => {
              navigate("/");
            },
            children: "No"
          }
        ),
        /* @__PURE__ */ (0, import_jsx_runtime33.jsx)(import_react14.Form, { method: "post", children: /* @__PURE__ */ (0, import_jsx_runtime33.jsx)(Button_default, { icon: "", intent: "danger", fullWidth: !1, type: "submit", children: "Yes" }) })
      ] }) })
    ] })
  ] });
}

// app/routes/models.tsx
var models_exports = {};
__export(models_exports, {
  default: () => ModelsLayout,
  loader: () => loader7,
  meta: () => meta7
});
var import_react15 = require("@remix-run/react");
var import_jsx_runtime34 = require("react/jsx-runtime"), meta7 = () => ({
  charset: "utf-8",
  title: "Models | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader7({ request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken");
  return await fetchUserSettings(accesstoken);
}
function ModelsLayout() {
  let prof = (0, import_react15.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime34.jsxs)("div", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime34.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime34.jsx)(import_react15.Meta, {}) }),
    prof ? /* @__PURE__ */ (0, import_jsx_runtime34.jsx)(NavBar, { intent: "loggedIn", user: prof[0].name.charAt(0).toUpperCase() }) : /* @__PURE__ */ (0, import_jsx_runtime34.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime34.jsx)("div", { className: "px-12 pt-16 pb-12 h-full w-screen", children: /* @__PURE__ */ (0, import_jsx_runtime34.jsx)(import_react15.Outlet, {}) })
  ] });
}

// app/routes/models/EmptyModel.tsx
var EmptyModel_exports = {};
__export(EmptyModel_exports, {
  default: () => EmptyModel
});
var import_class_variance_authority11 = require("class-variance-authority"), import_jsx_runtime35 = require("react/jsx-runtime"), emptyStyles2 = (0, import_class_variance_authority11.cva)("h-full flex justify-center items-center", {
  variants: {
    intent: {
      primary: "flex justify-center items-center bg-slate-150 drop-shadow-3xl rounded-lg"
    },
    fullWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: !0
  }
}), codeStyles2 = (0, import_class_variance_authority11.cva)("h-fit ml-8 bg-slate-150 text-slate-400", {
  variants: {
    intent: {
      primary: ""
    },
    fullWidth: {
      true: "w-full"
    }
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: !0
  }
});
function EmptyModel() {
  return /* @__PURE__ */ (0, import_jsx_runtime35.jsx)("div", { className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 min-w-72", children: /* @__PURE__ */ (0, import_jsx_runtime35.jsxs)("div", { className: "rounded-lg border-2 border-slate-200 px-6 py-4", children: [
    /* @__PURE__ */ (0, import_jsx_runtime35.jsx)("div", { className: "font-medium text-sm pb-6", children: "There are no models yet" }),
    /* @__PURE__ */ (0, import_jsx_runtime35.jsx)("div", { className: "rounded-lg h-2 bg-slate-200 w-1/3" }),
    /* @__PURE__ */ (0, import_jsx_runtime35.jsx)("div", { className: "pt-2" }),
    /* @__PURE__ */ (0, import_jsx_runtime35.jsx)("div", { className: "rounded-lg h-2 bg-slate-200 w-2/3" })
  ] }) });
}

// app/routes/models/index.tsx
var models_exports2 = {};
__export(models_exports2, {
  default: () => Index6,
  loader: () => loader8
});
var import_react16 = require("@remix-run/react");
var import_jsx_runtime36 = require("react/jsx-runtime");
async function loader8({ request }) {
  let session = await getSession(request.headers.get("Cookie")), orgId = session.get("orgId");
  return { models: await fetchModels(
    session.get("orgId"),
    session.get("accessToken")
  ), orgId };
}
function Index6() {
  let modelData = (0, import_react16.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime36.jsxs)("div", { id: "models", children: [
    /* @__PURE__ */ (0, import_jsx_runtime36.jsx)("div", { className: "flex justify-between font-medium text-slate-800 text-base pt-6", children: "Models" }),
    modelData ? /* @__PURE__ */ (0, import_jsx_runtime36.jsx)(import_jsx_runtime36.Fragment, { children: modelData.models[0].length !== 0 ? /* @__PURE__ */ (0, import_jsx_runtime36.jsx)("div", { className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72", children: modelData.models.map((model) => /* @__PURE__ */ (0, import_jsx_runtime36.jsx)(
      import_react16.Link,
      {
        to: `/org/${modelData.orgId}/models/${model.name}`,
        children: /* @__PURE__ */ (0, import_jsx_runtime36.jsx)(
          Card,
          {
            intent: "modelCard",
            name: model.name,
            description: `Updated by ${model.updated_by.handle || "-"}`,
            tag2: model.created_by.handle
          }
        )
      },
      model.id
    )) }) : /* @__PURE__ */ (0, import_jsx_runtime36.jsx)(EmptyModel, {}) }) : "All public models shown here"
  ] });
}

// app/routes/index.tsx
var routes_exports = {};
__export(routes_exports, {
  default: () => Index7
});
var import_react17 = require("@remix-run/react"), import_jsx_runtime37 = require("react/jsx-runtime");
function Index7() {
  return /* @__PURE__ */ (0, import_jsx_runtime37.jsx)("div", { children: /* @__PURE__ */ (0, import_jsx_runtime37.jsx)(import_react17.Outlet, {}) });
}

// app/routes/auth.tsx
var auth_exports = {};
__export(auth_exports, {
  default: () => AuthLayout,
  meta: () => meta8
});
var import_react18 = require("@remix-run/react"), import_jsx_runtime38 = require("react/jsx-runtime"), meta8 = () => ({
  charset: "utf-8",
  title: "Join Us | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function AuthLayout() {
  return /* @__PURE__ */ (0, import_jsx_runtime38.jsxs)("main", { className: "w-screen h-screen flex justify-center items-center bg-slate-0", children: [
    /* @__PURE__ */ (0, import_jsx_runtime38.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime38.jsx)(import_react18.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime38.jsxs)("div", { className: "hidden h-screen md:w-2/5 bg-slate-50 md:flex md:flex-col md:justify-center 4xl:items-center md:pl-24 xl:pl-44 md:overflow-visible", children: [
      /* @__PURE__ */ (0, import_jsx_runtime38.jsx)("img", { src: "/Logo.svg", alt: "logo", className: "w-24" }),
      /* @__PURE__ */ (0, import_jsx_runtime38.jsxs)("h1", { className: "text-slate-600 text-5xl font-medium mt-6", children: [
        "Welcome to",
        " ",
        /* @__PURE__ */ (0, import_jsx_runtime38.jsx)("span", { className: "font-medium text-slate-900 mt-11", children: "PureML" })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime38.jsx)(
        "img",
        {
          src: "/imgs/AuthCodeSnippet.svg",
          alt: "SignInCode",
          className: "md:w-80 lg:w-96 xl:w-[30rem] 2xl:w-[42rem] 4xl:w-[54rem] max-w-[710px] mt-12 -ml-4 z-10"
        }
      )
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime38.jsx)(import_react18.Outlet, {})
  ] });
}

// app/routes/auth/forgot_password/index.tsx
var forgot_password_exports = {};
__export(forgot_password_exports, {
  action: () => action6,
  default: () => Posts
});
var import_node9 = require("@remix-run/node");

// app/components/ui/Link.tsx
var import_react19 = require("@remix-run/react"), import_class_variance_authority12 = require("class-variance-authority"), import_jsx_runtime39 = require("react/jsx-runtime"), linkStyles = (0, import_class_variance_authority12.cva)("", {
  variants: {
    intent: {
      primary: "text-base text-blue-700 hover:text-blue-700 font-medium w-fit h-fit",
      secondary: "text-sm text-brand-200 hover:text-brand-300 font-medium"
    }
  },
  defaultVariants: {
    intent: "primary"
  }
});
function link({ intent, hyperlink, children }) {
  return /* @__PURE__ */ (0, import_jsx_runtime39.jsx)(import_react19.Link, { to: hyperlink, className: linkStyles({ intent }), children });
}

// app/routes/auth/forgot_password/index.tsx
var import_jsx_runtime40 = require("react/jsx-runtime"), action6 = async ({ request }) => {
  let email = (await request.formData()).get("email");
  return (0, import_node9.redirect)("/auth/reset_password");
};
function Posts() {
  return /* @__PURE__ */ (0, import_jsx_runtime40.jsx)("div", { className: "md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white", children: /* @__PURE__ */ (0, import_jsx_runtime40.jsxs)("div", { className: "md:max-w-[450px] w-96 text-center", children: [
    /* @__PURE__ */ (0, import_jsx_runtime40.jsx)("h2", { className: "font-semibold text-2xl text-slate-800 mb-12", children: "Forgot Password" }),
    /* @__PURE__ */ (0, import_jsx_runtime40.jsxs)("form", { method: "post", className: "text-slate-400 flex flex-col text-left", children: [
      /* @__PURE__ */ (0, import_jsx_runtime40.jsxs)("div", { className: "pb-6", children: [
        /* @__PURE__ */ (0, import_jsx_runtime40.jsx)("label", { htmlFor: "email", className: "text-base pb-1", children: "Email ID" }),
        /* @__PURE__ */ (0, import_jsx_runtime40.jsx)(
          Input_default,
          {
            intent: "primary",
            type: "email",
            name: "email",
            placeholder: "Enter email ID...",
            "aria-label": "emailid",
            "data-testid": "email-input3",
            required: !0
          }
        )
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime40.jsx)(Button_default, { "aria-label": "signin", intent: "primary", icon: "", children: "Send Link" }),
      /* @__PURE__ */ (0, import_jsx_runtime40.jsx)("span", { className: "text-sm text-zinc-400 pt-6", children: "Reset password link will be sent to your mail ID given above. Click link to change or reset password." })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime40.jsx)("div", { className: "flex items-center text-slate-600 space-x-2 justify-center mt-6", children: /* @__PURE__ */ (0, import_jsx_runtime40.jsx)(link, { intent: "secondary", hyperlink: "/auth/signin", children: "Go back" }) })
  ] }) });
}

// app/routes/auth/reset_password/index.tsx
var reset_password_exports = {};
__export(reset_password_exports, {
  action: () => action7,
  default: () => Posts2
});
var import_node10 = require("@remix-run/node");
var import_jsx_runtime41 = require("react/jsx-runtime"), action7 = async ({ request }) => {
  let form = await request.formData(), password = form.get("password"), cpassword = form.get("cpassword");
  return (0, import_node10.redirect)("/auth/signin");
};
function Posts2() {
  return /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("div", { className: "md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white", children: /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("div", { className: "md:max-w-[450px] w-96 text-center", children: [
    /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("h2", { className: "font-semibold text-2xl text-slate-800 mb-12", children: "Reset Password" }),
    /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("form", { method: "post", className: "text-slate-400 flex flex-col text-left", children: [
      /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("label", { htmlFor: "password", className: "text-base pb-1", children: [
        "New Password",
        /* @__PURE__ */ (0, import_jsx_runtime41.jsx)(
          Input_default,
          {
            intent: "primary",
            required: !0,
            type: "password",
            name: "password",
            placeholder: "Enter password...",
            "aria-label": "password",
            "data-testid": "password-input"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("label", { htmlFor: "cpassword", className: "text-base pb-1", children: [
        "Confirm Password",
        /* @__PURE__ */ (0, import_jsx_runtime41.jsx)(
          Input_default,
          {
            intent: "primary",
            required: !0,
            type: "password",
            name: "cpassword",
            placeholder: "Enter password...",
            "aria-label": "password",
            "data-testid": "password-input"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime41.jsx)(Button_default, { "aria-label": "signin", intent: "primary", icon: "", children: "Sign In" })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("div", { className: "flex items-center text-slate-600 space-x-2 justify-center mt-6", children: [
      /* @__PURE__ */ (0, import_jsx_runtime41.jsx)(link, { intent: "secondary", hyperlink: "/forgot_password", children: "Forgot Password?" }),
      /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("p", { children: "|" }),
      /* @__PURE__ */ (0, import_jsx_runtime41.jsxs)("div", { className: "flex items-center space-x-1", children: [
        /* @__PURE__ */ (0, import_jsx_runtime41.jsx)("span", { className: "text-sm", children: "Already have an account?" }),
        /* @__PURE__ */ (0, import_jsx_runtime41.jsx)(link, { intent: "secondary", hyperlink: "/signin", children: "Sign In" })
      ] })
    ] })
  ] }) });
}

// app/routes/auth/signin/index.tsx
var signin_exports = {};
__export(signin_exports, {
  ErrorBoundary: () => ErrorBoundary2,
  action: () => action8,
  default: () => SignIn
});
var import_node11 = require("@remix-run/node"), import_react20 = require("@remix-run/react"), import_react_hot_toast3 = __toESM(require("react-hot-toast"));
var import_jsx_runtime42 = require("react/jsx-runtime"), ErrorBoundary2 = ({ error }) => SignIn(error.message), action8 = async ({ request }) => {
  let form = await request.formData(), email = form.get("email"), password = form.get("password"), data = await fetchSignIn(email, password), session = await getSession(request.headers.get("Cookie"));
  if (data.Message === "User logged in") {
    let accessToken = data.Data[0].accessToken;
    session.set("accessToken", accessToken);
    let org = await fetchUserOrg(accessToken);
    return session.set("orgId", org[0].org.uuid), session.set("orgName", org[0].org.name), import_react_hot_toast3.default.success("Here is your toast."), (0, import_node11.redirect)("/models", {
      headers: { "Set-Cookie": await commitSession(session) }
    });
  } else
    return data.Message === "User not found" ? (0, import_node11.json)({ message: "User not found" }) : data.Message === "Invalid username" ? (0, import_node11.json)({ message: "Invalid username" }) : data.Message === "Invalid credentials" ? (0, import_node11.json)({ message: "Invalid credentials" }) : (0, import_node11.json)({ message: "Something went Wrong!" });
};
function SignIn(err) {
  let data = (0, import_react20.useActionData)(), errorComp = /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("div", {});
  return err.length > 0 && (errorComp = /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("p", { className: "errorBox", children: [
    "There was an error with your data: ",
    /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("i", { className: "errorMsg", children: err })
  ] })), /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("div", { className: "md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white", children: [
    errorComp,
    /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("div", { className: "md:max-w-[450px] w-96 text-center", children: [
      /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("h2", { className: "font-semibold text-2xl text-slate-800 mb-12", children: "Sign In to PureML" }),
      /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("form", { method: "post", className: "text-slate-400 flex flex-col text-left", children: [
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("p", { className: "text-red-500", children: data ? data.message : null }),
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("label", { htmlFor: "email", className: "text-base pb-1", children: [
          "Email ID",
          /* @__PURE__ */ (0, import_jsx_runtime42.jsx)(
            Input_default,
            {
              intent: "primary",
              type: "email",
              name: "email",
              placeholder: "Enter email ID...",
              "aria-label": "emailid",
              "data-testid": "email-input1",
              required: !0
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("label", { htmlFor: "password", className: "text-base pb-1", children: [
          "Password",
          /* @__PURE__ */ (0, import_jsx_runtime42.jsx)(
            Input_default,
            {
              intent: "primary",
              required: !0,
              type: "password",
              name: "password",
              placeholder: "Enter password...",
              "aria-label": "password",
              "data-testid": "password-input1"
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)(Button_default, { intent: "primary", icon: "", type: "submit", children: "Sign in" })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("div", { className: "flex items-center text-slate-600 space-x-3 justify-center mt-6", children: [
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)(link, { intent: "secondary", hyperlink: "/auth/forgot_password", children: "Forgot Password?" }),
        /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("p", { children: "|" }),
        /* @__PURE__ */ (0, import_jsx_runtime42.jsxs)("div", { className: "flex items-center space-x-1", children: [
          /* @__PURE__ */ (0, import_jsx_runtime42.jsx)("span", { className: "text-sm", children: "Dont have an account?" }),
          /* @__PURE__ */ (0, import_jsx_runtime42.jsx)(link, { intent: "secondary", hyperlink: "/auth/signup", children: "Sign Up" })
        ] })
      ] })
    ] })
  ] });
}

// app/routes/auth/signup/index.tsx
var signup_exports = {};
__export(signup_exports, {
  ErrorBoundary: () => ErrorBoundary3,
  action: () => action9,
  default: () => SignUp
});
var import_node12 = require("@remix-run/node"), import_react21 = require("@remix-run/react"), import_react_hot_toast4 = __toESM(require("react-hot-toast"));
var import_jsx_runtime43 = require("react/jsx-runtime"), ErrorBoundary3 = ({ error }) => SignUp(error.message), action9 = async ({ request }) => {
  let form = await request.formData();
  console.log("form=", form);
  let name = form.get("name"), username = form.get("username"), email = form.get("email"), password = form.get("password"), bio = form.get("bio");
  console.log(name, username, email, password, bio);
  let data = await fetchSignUp(name, username, email, password, bio);
  return console.log("data=", data), data.Message === "User created" ? (import_react_hot_toast4.default.success("Here is your toast."), (0, import_node12.json)({ message: "Account created successfully" }), (0, import_node12.redirect)("/auth/signin", {})) : data.Message === "User already exists" ? (0, import_node12.json)({ message: "User not found" }) : (0, import_node12.json)({ message: "Something went Wrong!" });
};
function SignUp(err) {
  let data = (0, import_react21.useActionData)(), errorComp = /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", {});
  return err.length > 0 && (errorComp = /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("p", { className: "errorBox", children: [
    "There was an error with your data: ",
    /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("i", { className: "errorMsg", children: err })
  ] })), /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("div", { className: "md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white", children: [
    errorComp,
    /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("div", { className: "md:max-w-[450px] w-96 text-center", children: [
      /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("h2", { className: "font-semibold text-2xl text-slate-800 mb-12", children: "Sign up to PureML" }),
      /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("form", { method: "post", className: "text-slate-400 flex flex-col text-left", children: [
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("label", { htmlFor: "name", className: "text-base pb-1", children: [
          "Name",
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(
            Input_default,
            {
              intent: "primary",
              type: "text",
              name: "name",
              placeholder: "Enter name...",
              "aria-label": "name",
              "data-testid": "name-input",
              required: !0
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("label", { htmlFor: "username", className: "text-base pb-1", children: [
          "Username",
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(
            Input_default,
            {
              intent: "primary",
              type: "text",
              name: "username",
              placeholder: "Enter your username...",
              "aria-label": "username",
              "data-testid": "username-input",
              required: !0
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("label", { htmlFor: "email", className: "text-base pb-1", children: [
          "Email ID",
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(
            Input_default,
            {
              intent: "primary",
              required: !0,
              type: "email",
              name: "email",
              placeholder: "Enter email ID...",
              "aria-label": "emalid",
              "data-testid": "email-input2"
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("label", { htmlFor: "password", className: "text-base pb-1", children: [
          "Password",
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(
            Input_default,
            {
              intent: "primary",
              type: "password",
              name: "password",
              required: !0,
              placeholder: "Enter password...",
              "aria-label": "password",
              "data-testid": "password-input2"
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("label", { htmlFor: "bio", className: "text-base pb-1", children: [
          "Short Bio",
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(
            Input_default,
            {
              intent: "primary",
              type: "text",
              name: "bio",
              placeholder: "Enter your short bio...",
              "aria-label": "bio",
              "data-testid": "bio-input",
              required: !0
            }
          )
        ] }) }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(Button_default, { intent: "primary", icon: "", children: "Sign Up" })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("div", { className: "flex items-center text-slate-600 space-x-2 justify-center mt-6", children: [
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(link, { intent: "secondary", hyperlink: "/forgot_password", children: "Forgot Password?" }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("p", { children: "|" }),
        /* @__PURE__ */ (0, import_jsx_runtime43.jsxs)("div", { className: "flex items-center space-x-1", children: [
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)("span", { className: "text-sm", children: "Already have an account?" }),
          /* @__PURE__ */ (0, import_jsx_runtime43.jsx)(link, { intent: "secondary", hyperlink: "/auth/signin", children: "Sign In" })
        ] })
      ] })
    ] })
  ] });
}

// app/routes/org.tsx
var org_exports = {};
__export(org_exports, {
  default: () => OrgLayout,
  loader: () => loader9
});
var import_react22 = require("@remix-run/react");
var import_jsx_runtime44 = require("react/jsx-runtime");
async function loader9({ request }) {
  let accesstoken = (await getSession(request.headers.get("Cookie"))).get("accessToken");
  return await fetchUserSettings(accesstoken);
}
function OrgLayout() {
  let prof = (0, import_react22.useLoaderData)();
  return /* @__PURE__ */ (0, import_jsx_runtime44.jsxs)("div", { children: [
    prof ? /* @__PURE__ */ (0, import_jsx_runtime44.jsx)(NavBar, { intent: "loggedIn", user: prof[0].name.charAt(0).toUpperCase() }) : /* @__PURE__ */ (0, import_jsx_runtime44.jsx)(NavBar, { intent: "loggedOut" }),
    /* @__PURE__ */ (0, import_jsx_runtime44.jsx)("div", { className: "pt-16 h-screen w-screen", children: /* @__PURE__ */ (0, import_jsx_runtime44.jsx)(import_react22.Outlet, {}) })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId.tsx
var datasetId_exports = {};
__export(datasetId_exports, {
  default: () => DatasetIndex,
  meta: () => meta9
});
var import_react24 = require("@remix-run/react");

// app/components/Breadcrumbs.tsx
var import_react23 = require("@remix-run/react"), import_jsx_runtime45 = require("react/jsx-runtime");
function Breadcrumbs() {
  let pathname = (0, import_react23.useMatches)()[2].pathname, url = decodeURI(pathname.slice(1)).split("/"), urlitems = url.filter(function(val, idx) {
    if ((idx + 1) % 2 == 0)
      return val;
  });
  return /* @__PURE__ */ (0, import_jsx_runtime45.jsx)("ul", { className: "font-medium flex pt-6", children: urlitems.map((item, index) => /* @__PURE__ */ (0, import_jsx_runtime45.jsxs)("li", { children: [
    /* @__PURE__ */ (0, import_jsx_runtime45.jsx)(
      import_react23.Link,
      {
        to: `/${url.slice(0, index + 2).join("/")}`,
        className: "text-slate-600",
        children: item
      }
    ),
    index !== url.length - 1 && /* @__PURE__ */ (0, import_jsx_runtime45.jsx)("span", { className: "text-slate-400 mx-1", children: "/" })
  ] }, item)) });
}
var Breadcrumbs_default = Breadcrumbs;

// app/routes/org/$orgId/datasets/$datasetId.tsx
var import_jsx_runtime46 = require("react/jsx-runtime"), meta9 = () => ({
  charset: "utf-8",
  title: "Dataset Details | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function DatasetIndex() {
  return /* @__PURE__ */ (0, import_jsx_runtime46.jsxs)("div", { id: "datasets", children: [
    /* @__PURE__ */ (0, import_jsx_runtime46.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime46.jsx)(import_react24.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime46.jsxs)("div", { className: "flex flex-col", children: [
      /* @__PURE__ */ (0, import_jsx_runtime46.jsx)("div", { className: "px-12 sticky top-16 bg-slate-0 w-full z-10", children: /* @__PURE__ */ (0, import_jsx_runtime46.jsx)(Breadcrumbs_default, {}) }),
      /* @__PURE__ */ (0, import_jsx_runtime46.jsx)(import_react24.Outlet, {})
    ] })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/versions/datalineage.tsx
var datalineage_exports = {};
__export(datalineage_exports, {
  default: () => Metrics
});
var import_react25 = require("@remix-run/react");
var import_jsx_runtime47 = require("react/jsx-runtime");
function Metrics() {
  return /* @__PURE__ */ (0, import_jsx_runtime47.jsxs)("div", { className: "", children: [
    /* @__PURE__ */ (0, import_jsx_runtime47.jsx)(TabBar, { intent: "primaryDatasetTab", tab: "versions" }),
    /* @__PURE__ */ (0, import_jsx_runtime47.jsx)(import_react25.Outlet, {})
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/versions/datalineage/index.tsx
var datalineage_exports2 = {};
__export(datalineage_exports2, {
  default: () => DatasetVersions,
  loader: () => loader10,
  meta: () => meta10
});
var import_react27 = require("@remix-run/react");

// app/routes/org/$orgId/datasets/$datasetId/versions/Pipeline.tsx
var Pipeline_exports = {};
__export(Pipeline_exports, {
  default: () => Pipeline_default
});
var import_react26 = require("react"), import_reactflow = __toESM(require("reactflow")), import_dagre = __toESM(require("dagre")), import_jsx_runtime48 = require("react/jsx-runtime"), dagreGraph = new import_dagre.default.graphlib.Graph();
dagreGraph.setDefaultEdgeLabel(() => ({}));
var nodeWidth = 172, nodeHeight = 36, getLayoutedElements = (nodes, edges, direction = "TB") => {
  let isHorizontal = direction === "LR";
  return dagreGraph.setGraph({ rankdir: direction }), nodes.forEach((node) => {
    dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
  }), edges.forEach((edge) => {
    dagreGraph.setEdge(edge.source, edge.target);
  }), import_dagre.default.layout(dagreGraph), nodes.forEach((node) => {
    let nodeWithPosition = dagreGraph.node(node.id);
    return node.targetPosition = isHorizontal ? "left" : "top", node.sourcePosition = isHorizontal ? "right" : "bottom", node.position = {
      x: nodeWithPosition.x - nodeWidth / 2,
      y: nodeWithPosition.y - nodeHeight / 2
    }, node;
  }), { nodes, edges };
};
function Pipeline({ pnode, pedge }) {
  let { nodes: layoutedNodes, edges: layoutedEdges } = getLayoutedElements(
    pnode,
    pedge
  ), [nodes, setNodes, onNodesChange] = (0, import_reactflow.useNodesState)(layoutedNodes), [edges, setEdges, onEdgesChange] = (0, import_reactflow.useEdgesState)(layoutedEdges), onConnect = (0, import_react26.useCallback)(
    (params) => setEdges(
      (eds) => (0, import_reactflow.addEdge)(
        { ...params, type: import_reactflow.ConnectionLineType.SmoothStep, animated: !0 },
        eds
      )
    ),
    []
  ), onLayout = (0, import_react26.useCallback)(
    (direction) => {
      let { nodes: layoutedNodes2, edges: layoutedEdges2 } = getLayoutedElements(nodes, edges, direction);
      setNodes([...layoutedNodes2]), setEdges([...layoutedEdges2]);
    },
    [nodes, edges]
  );
  return /* @__PURE__ */ (0, import_jsx_runtime48.jsx)("div", { className: "h-full", children: /* @__PURE__ */ (0, import_jsx_runtime48.jsxs)(
    import_reactflow.default,
    {
      nodes,
      edges,
      onNodesChange,
      onEdgesChange,
      onConnect,
      connectionLineType: import_reactflow.ConnectionLineType.SmoothStep,
      fitView: !0,
      children: [
        /* @__PURE__ */ (0, import_jsx_runtime48.jsx)(import_reactflow.Controls, {}),
        /* @__PURE__ */ (0, import_jsx_runtime48.jsx)(import_reactflow.MiniMap, { zoomable: !0, pannable: !0 })
      ]
    }
  ) });
}
var Pipeline_default = Pipeline;

// app/routes/org/$orgId/datasets/$datasetId/versions/datalineage/index.tsx
var import_react28 = require("react"), import_lucide_react8 = require("lucide-react"), import_clsx3 = __toESM(require("clsx")), import_jsx_runtime49 = require("react/jsx-runtime"), meta10 = () => ({
  charset: "utf-8",
  title: "Dataset Versions | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader10({ params, request }) {
  let session = await getSession(request.headers.get("Cookie"));
  return {
    versions: await fetchDatasetVersions(
      params.orgId,
      params.datasetId,
      session.get("accessToken")
    )
  };
}
function DatasetVersions({ tab }) {
  let data = (0, import_react27.useLoaderData)(), path = (0, import_react27.useMatches)()[2].pathname, pathname = decodeURI(path.slice(1)), orgId = pathname.split("/")[1], datasetId = pathname.split("/")[3], [lineage, setLineage] = (0, import_react28.useState)(!0), [graphs, setGraphs] = (0, import_react28.useState)(!0), versionData = data.versions, [ver1, setVer1] = (0, import_react28.useState)(""), [ver2, setVer2] = (0, import_react28.useState)(""), [node, setNode] = (0, import_react28.useState)(null), [edge, setEdge] = (0, import_react28.useState)(null), [node2, setNode2] = (0, import_react28.useState)(null), [edge2, setEdge2] = (0, import_react28.useState)(null);
  return (0, import_react28.useEffect)(() => {
    !versionData[0] || !versionData[0].lineage.lineage || (setVer1(versionData.at(-1).version), setVer2(""));
  }, [versionData]), (0, import_react28.useEffect)(() => {
    if (!versionData)
      return;
    if (ver2 === "") {
      setNode2(null), setEdge2(null);
      return;
    }
    let n = JSON.parse(
      versionData.at(Number(ver2.slice(1)) - 1).lineage.lineage
    ).nodes;
    n.forEach((e) => {
      e.data = { label: e.text };
    }), setNode2(null), setTimeout(() => {
      setNode2(n);
    }, 10);
    let ed = JSON.parse(
      versionData.at(Number(ver2.slice(1)) - 1).lineage.lineage
    ).edges;
    ed.forEach((e) => {
      e.source = e.from, e.target = e.to;
    }), setEdge2(null), setTimeout(() => {
      setEdge2(ed);
    }, 10);
  }, [ver2, versionData]), (0, import_react28.useEffect)(() => {
    if (!versionData)
      return;
    let n = JSON.parse(
      versionData.at(Number(ver1.slice(1)) - 1).lineage.lineage
    ).nodes;
    n.forEach((e) => {
      e.data = { label: e.text };
    }), setNode(null), setTimeout(() => {
      setNode(n);
    }, 10);
    let ed = JSON.parse(
      versionData.at(Number(ver1.slice(1)) - 1).lineage.lineage
    ).edges;
    ed.forEach((e) => {
      e.source = e.from, e.target = e.to;
    }), setEdge(null), setTimeout(() => {
      setEdge(ed);
    }, 10);
  }, [ver1, versionData]), /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("main", { className: "flex px-2", children: [
    /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("div", { className: "w-full", id: "main", children: [
      /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(TabBar, { intent: "datasetTab", tab: "datalineage" }),
      /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("section", { className: "w-full", children: [
        /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)(
          "div",
          {
            className: "flex items-center justify-between w-full border-b-slate-300 border-b",
            onClick: () => setLineage(!lineage),
            children: [
              /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("h1", { className: "text-slate-900 font-medium text-sm rounded-lg", children: "Data Lineage" }),
              lineage ? /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(import_lucide_react8.ChevronUp, { className: "text-slate-400" }) : /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(import_lucide_react8.ChevronDown, { className: "text-slate-400" })
            ]
          }
        ),
        lineage && /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("div", { className: "flex", children: [
          /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("div", { className: "w-full h-screen max-h-[600px]", children: node && /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(Pipeline_default, { pnode: node, pedge: edge }) }),
          /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("div", { className: "w-full h-screen max-h-[600px]", children: node2 && /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(Pipeline_default, { pnode: node2, pedge: edge2 }) })
        ] }),
        !lineage && /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("div", { children: "Add Data lineage" })
      ] }) })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("aside", { className: "sticky top-40 bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: [
      /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("p", { className: "mb-8", children: "Main" }),
      versionData ? /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("ul", { className: "space-y-2", children: versionData.map((version) => /* @__PURE__ */ (0, import_jsx_runtime49.jsxs)("li", { className: "flex items-center space-x-2", children: [
        /* @__PURE__ */ (0, import_jsx_runtime49.jsx)(
          "input",
          {
            name: "version2",
            value: version.version,
            type: "checkbox",
            checked: version.version === ver1 || version.version === ver2,
            onChange: (e) => {
              e.target.checked ? setVer2(version.version) : ver1 === version.version && ver2 === "" ? new Error("You can't uncheck the present version") : (ver1 === version.version && setVer1(ver2), setVer2(""));
            }
          }
        ),
        /* @__PURE__ */ (0, import_jsx_runtime49.jsx)("p", { children: version.version })
      ] }, version.version)) }) : "No side"
    ] })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/versions/nodes-edges.js
var nodes_edges_exports = {};
__export(nodes_edges_exports, {
  initialEdges: () => initialEdges,
  initialNodes: () => initialNodes
});
var edgeType = "smoothstep", initialNodes = [
  {
    id: "1",
    type: "input",
    data: { label: "input" }
  },
  {
    id: "2",
    data: { label: "node 2" }
  },
  {
    id: "2a",
    data: { label: "node 2a" }
  },
  {
    id: "2b",
    data: { label: "node 2b" }
  },
  {
    id: "2c",
    data: { label: "node 2c" }
  },
  {
    id: "2d",
    data: { label: "node 2d" }
  },
  {
    id: "3",
    data: { label: "node 3" }
  },
  {
    id: "4",
    data: { label: "node 4" }
  },
  {
    id: "5",
    data: { label: "node 5" }
  },
  {
    id: "6",
    type: "output",
    data: { label: "output" }
  },
  { id: "7", type: "output", data: { label: "output" } }
], initialEdges = [
  { id: "e12", source: "1", target: "2", type: edgeType, animated: !0 },
  { id: "e13", source: "1", target: "3", type: edgeType, animated: !0 },
  { id: "e22a", source: "2", target: "2a", type: edgeType, animated: !0 },
  { id: "e22b", source: "2", target: "2b", type: edgeType, animated: !1 },
  { id: "e22c", source: "2", target: "2c", type: edgeType, animated: !1 },
  { id: "e2c2d", source: "2c", target: "2d", type: edgeType, animated: !1 },
  { id: "e45", source: "4", target: "5", type: edgeType, animated: !0 },
  { id: "e56", source: "5", target: "6", type: edgeType, animated: !0 },
  { id: "e57", source: "5", target: "7", type: edgeType, animated: !0 }
];

// app/routes/org/$orgId/datasets/$datasetId/versions/graphs.tsx
var graphs_exports = {};
__export(graphs_exports, {
  default: () => Graphs
});
var import_react29 = require("@remix-run/react");
var import_jsx_runtime50 = require("react/jsx-runtime");
function Graphs() {
  return /* @__PURE__ */ (0, import_jsx_runtime50.jsxs)("div", { className: "", children: [
    /* @__PURE__ */ (0, import_jsx_runtime50.jsx)(TabBar, { intent: "primaryDatasetTab", tab: "versions" }),
    /* @__PURE__ */ (0, import_jsx_runtime50.jsx)(import_react29.Outlet, {})
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/versions/graphs/index.tsx
var graphs_exports2 = {};
__export(graphs_exports2, {
  default: () => DatasetGraphs
});
var import_jsx_runtime51 = require("react/jsx-runtime");
function DatasetGraphs() {
  return /* @__PURE__ */ (0, import_jsx_runtime51.jsxs)("main", { className: "flex", children: [
    /* @__PURE__ */ (0, import_jsx_runtime51.jsxs)("div", { className: "w-full", id: "main", children: [
      /* @__PURE__ */ (0, import_jsx_runtime51.jsx)(TabBar, { intent: "datasetTab", tab: "graphs" }),
      /* @__PURE__ */ (0, import_jsx_runtime51.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime51.jsx)("section", { className: "w-full" }) })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime51.jsx)("aside", { className: "bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: /* @__PURE__ */ (0, import_jsx_runtime51.jsx)(Dropdown, { fullWidth: !1, intent: "branch", children: "dev" }) })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/review/$commit.tsx
var commit_exports = {};
__export(commit_exports, {
  default: () => Review,
  meta: () => meta11
});
var import_react30 = require("@remix-run/react"), import_lucide_react9 = require("lucide-react"), import_react31 = require("react");
var import_jsx_runtime52 = require("react/jsx-runtime"), meta11 = () => ({
  charset: "utf-8",
  title: "Review Dataset Commit | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function Review() {
  let path = (0, import_react30.useMatches)()[2].pathname, pathname = decodeURI(path.slice(1)), orgId = pathname.split("/")[1], datasetId = pathname.split("/")[3], [lineage, setLineage] = (0, import_react31.useState)(!0);
  return /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)("div", { id: "reviewCommit", children: [
    /* @__PURE__ */ (0, import_jsx_runtime52.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(import_react30.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(TabBar, { intent: "primaryDatasetTab", tab: "review" }),
    /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)(
      import_react30.Link,
      {
        to: `/org/${orgId}/datasets/${datasetId}/review`,
        className: "flex font-medium text-sm text-slate-600 items-center px-12 pt-8",
        children: [
          /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(import_lucide_react9.ArrowLeft, {}),
          " View Commit"
        ]
      }
    ),
    /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)("div", { className: "flex px-4", children: [
      /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)("div", { className: "w-full", id: "main", children: [
        /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(TabBar, { intent: "datasetReviewTab", tab: "datalineage" }),
        /* @__PURE__ */ (0, import_jsx_runtime52.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)("section", { className: "w-full", children: [
          /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)(
            "div",
            {
              className: "flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4",
              onClick: () => setLineage(!lineage),
              children: [
                /* @__PURE__ */ (0, import_jsx_runtime52.jsx)("h1", { className: "text-slate-900 font-medium text-sm", children: "Data Lineage" }),
                lineage ? /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(import_lucide_react9.ChevronUp, { className: "text-slate-400" }) : /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(import_lucide_react9.ChevronDown, { className: "text-slate-400" })
              ]
            }
          ),
          lineage && /* @__PURE__ */ (0, import_jsx_runtime52.jsx)("div", { className: "py-6", children: "Data Lineage will be shown here" })
        ] }) })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime52.jsxs)("aside", { className: "bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: [
        /* @__PURE__ */ (0, import_jsx_runtime52.jsx)(Dropdown, { fullWidth: !1, intent: "branch", children: "dev" }),
        "List of versions"
      ] })
    ] })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/review/index.tsx
var review_exports = {};
__export(review_exports, {
  default: () => DatasetReview,
  meta: () => meta12
});
var import_react32 = require("@remix-run/react");
var import_jsx_runtime53 = require("react/jsx-runtime"), meta12 = () => ({
  charset: "utf-8",
  title: "Dataset Review | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function DatasetReview() {
  let navigate = (0, import_react32.useNavigate)();
  return /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { id: "datasetsReview", children: [
    /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(import_react32.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(TabBar, { intent: "primaryDatasetTab", tab: "review" }),
    /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "px-12 pt-8 w-2/3", children: [
      /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(AvatarIcon, { children: "A" }),
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(AvatarIcon, { children: "B" }),
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime53.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(AvatarIcon, { children: "C" }),
          /* @__PURE__ */ (0, import_jsx_runtime53.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime53.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) })
    ] })
  ] });
}

// app/routes/org/$orgId/datasets/$datasetId/index.tsx
var datasetId_exports2 = {};
__export(datasetId_exports2, {
  action: () => action10,
  default: () => DatasetIndex2,
  links: () => links3,
  loader: () => loader11,
  meta: () => meta13
});
var import_react33 = require("@remix-run/react");

// node_modules/.pnpm/marked@4.2.12/node_modules/marked/lib/marked.esm.js
function getDefaults() {
  return {
    async: !1,
    baseUrl: null,
    breaks: !1,
    extensions: null,
    gfm: !0,
    headerIds: !0,
    headerPrefix: "",
    highlight: null,
    langPrefix: "language-",
    mangle: !0,
    pedantic: !1,
    renderer: null,
    sanitize: !1,
    sanitizer: null,
    silent: !1,
    smartypants: !1,
    tokenizer: null,
    walkTokens: null,
    xhtml: !1
  };
}
var defaults = getDefaults();
function changeDefaults(newDefaults) {
  defaults = newDefaults;
}
var escapeTest = /[&<>"']/, escapeReplace = new RegExp(escapeTest.source, "g"), escapeTestNoEncode = /[<>"']|&(?!(#\d{1,7}|#[Xx][a-fA-F0-9]{1,6}|\w+);)/, escapeReplaceNoEncode = new RegExp(escapeTestNoEncode.source, "g"), escapeReplacements = {
  "&": "&amp;",
  "<": "&lt;",
  ">": "&gt;",
  '"': "&quot;",
  "'": "&#39;"
}, getEscapeReplacement = (ch) => escapeReplacements[ch];
function escape(html, encode) {
  if (encode) {
    if (escapeTest.test(html))
      return html.replace(escapeReplace, getEscapeReplacement);
  } else if (escapeTestNoEncode.test(html))
    return html.replace(escapeReplaceNoEncode, getEscapeReplacement);
  return html;
}
var unescapeTest = /&(#(?:\d+)|(?:#x[0-9A-Fa-f]+)|(?:\w+));?/ig;
function unescape(html) {
  return html.replace(unescapeTest, (_, n) => (n = n.toLowerCase(), n === "colon" ? ":" : n.charAt(0) === "#" ? n.charAt(1) === "x" ? String.fromCharCode(parseInt(n.substring(2), 16)) : String.fromCharCode(+n.substring(1)) : ""));
}
var caret = /(^|[^\[])\^/g;
function edit(regex, opt) {
  regex = typeof regex == "string" ? regex : regex.source, opt = opt || "";
  let obj = {
    replace: (name, val) => (val = val.source || val, val = val.replace(caret, "$1"), regex = regex.replace(name, val), obj),
    getRegex: () => new RegExp(regex, opt)
  };
  return obj;
}
var nonWordAndColonTest = /[^\w:]/g, originIndependentUrl = /^$|^[a-z][a-z0-9+.-]*:|^[?#]/i;
function cleanUrl(sanitize, base2, href) {
  if (sanitize) {
    let prot;
    try {
      prot = decodeURIComponent(unescape(href)).replace(nonWordAndColonTest, "").toLowerCase();
    } catch {
      return null;
    }
    if (prot.indexOf("javascript:") === 0 || prot.indexOf("vbscript:") === 0 || prot.indexOf("data:") === 0)
      return null;
  }
  base2 && !originIndependentUrl.test(href) && (href = resolveUrl(base2, href));
  try {
    href = encodeURI(href).replace(/%25/g, "%");
  } catch {
    return null;
  }
  return href;
}
var baseUrls = {}, justDomain = /^[^:]+:\/*[^/]*$/, protocol = /^([^:]+:)[\s\S]*$/, domain = /^([^:]+:\/*[^/]*)[\s\S]*$/;
function resolveUrl(base2, href) {
  baseUrls[" " + base2] || (justDomain.test(base2) ? baseUrls[" " + base2] = base2 + "/" : baseUrls[" " + base2] = rtrim(base2, "/", !0)), base2 = baseUrls[" " + base2];
  let relativeBase = base2.indexOf(":") === -1;
  return href.substring(0, 2) === "//" ? relativeBase ? href : base2.replace(protocol, "$1") + href : href.charAt(0) === "/" ? relativeBase ? href : base2.replace(domain, "$1") + href : base2 + href;
}
var noopTest = { exec: function() {
} };
function merge(obj) {
  let i = 1, target, key;
  for (; i < arguments.length; i++) {
    target = arguments[i];
    for (key in target)
      Object.prototype.hasOwnProperty.call(target, key) && (obj[key] = target[key]);
  }
  return obj;
}
function splitCells(tableRow, count) {
  let row = tableRow.replace(/\|/g, (match, offset, str) => {
    let escaped = !1, curr = offset;
    for (; --curr >= 0 && str[curr] === "\\"; )
      escaped = !escaped;
    return escaped ? "|" : " |";
  }), cells = row.split(/ \|/), i = 0;
  if (cells[0].trim() || cells.shift(), cells.length > 0 && !cells[cells.length - 1].trim() && cells.pop(), cells.length > count)
    cells.splice(count);
  else
    for (; cells.length < count; )
      cells.push("");
  for (; i < cells.length; i++)
    cells[i] = cells[i].trim().replace(/\\\|/g, "|");
  return cells;
}
function rtrim(str, c, invert) {
  let l = str.length;
  if (l === 0)
    return "";
  let suffLen = 0;
  for (; suffLen < l; ) {
    let currChar = str.charAt(l - suffLen - 1);
    if (currChar === c && !invert)
      suffLen++;
    else if (currChar !== c && invert)
      suffLen++;
    else
      break;
  }
  return str.slice(0, l - suffLen);
}
function findClosingBracket(str, b) {
  if (str.indexOf(b[1]) === -1)
    return -1;
  let l = str.length, level = 0, i = 0;
  for (; i < l; i++)
    if (str[i] === "\\")
      i++;
    else if (str[i] === b[0])
      level++;
    else if (str[i] === b[1] && (level--, level < 0))
      return i;
  return -1;
}
function checkSanitizeDeprecation(opt) {
  opt && opt.sanitize && !opt.silent && console.warn("marked(): sanitize and sanitizer parameters are deprecated since version 0.7.0, should not be used and will be removed in the future. Read more here: https://marked.js.org/#/USING_ADVANCED.md#options");
}
function repeatString(pattern, count) {
  if (count < 1)
    return "";
  let result = "";
  for (; count > 1; )
    count & 1 && (result += pattern), count >>= 1, pattern += pattern;
  return result + pattern;
}
function outputLink(cap, link2, raw, lexer2) {
  let href = link2.href, title = link2.title ? escape(link2.title) : null, text = cap[1].replace(/\\([\[\]])/g, "$1");
  if (cap[0].charAt(0) !== "!") {
    lexer2.state.inLink = !0;
    let token = {
      type: "link",
      raw,
      href,
      title,
      text,
      tokens: lexer2.inlineTokens(text)
    };
    return lexer2.state.inLink = !1, token;
  }
  return {
    type: "image",
    raw,
    href,
    title,
    text: escape(text)
  };
}
function indentCodeCompensation(raw, text) {
  let matchIndentToCode = raw.match(/^(\s+)(?:```)/);
  if (matchIndentToCode === null)
    return text;
  let indentToCode = matchIndentToCode[1];
  return text.split(`
`).map((node) => {
    let matchIndentInNode = node.match(/^\s+/);
    if (matchIndentInNode === null)
      return node;
    let [indentInNode] = matchIndentInNode;
    return indentInNode.length >= indentToCode.length ? node.slice(indentToCode.length) : node;
  }).join(`
`);
}
var Tokenizer = class {
  constructor(options2) {
    this.options = options2 || defaults;
  }
  space(src) {
    let cap = this.rules.block.newline.exec(src);
    if (cap && cap[0].length > 0)
      return {
        type: "space",
        raw: cap[0]
      };
  }
  code(src) {
    let cap = this.rules.block.code.exec(src);
    if (cap) {
      let text = cap[0].replace(/^ {1,4}/gm, "");
      return {
        type: "code",
        raw: cap[0],
        codeBlockStyle: "indented",
        text: this.options.pedantic ? text : rtrim(text, `
`)
      };
    }
  }
  fences(src) {
    let cap = this.rules.block.fences.exec(src);
    if (cap) {
      let raw = cap[0], text = indentCodeCompensation(raw, cap[3] || "");
      return {
        type: "code",
        raw,
        lang: cap[2] ? cap[2].trim().replace(this.rules.inline._escapes, "$1") : cap[2],
        text
      };
    }
  }
  heading(src) {
    let cap = this.rules.block.heading.exec(src);
    if (cap) {
      let text = cap[2].trim();
      if (/#$/.test(text)) {
        let trimmed = rtrim(text, "#");
        (this.options.pedantic || !trimmed || / $/.test(trimmed)) && (text = trimmed.trim());
      }
      return {
        type: "heading",
        raw: cap[0],
        depth: cap[1].length,
        text,
        tokens: this.lexer.inline(text)
      };
    }
  }
  hr(src) {
    let cap = this.rules.block.hr.exec(src);
    if (cap)
      return {
        type: "hr",
        raw: cap[0]
      };
  }
  blockquote(src) {
    let cap = this.rules.block.blockquote.exec(src);
    if (cap) {
      let text = cap[0].replace(/^ *>[ \t]?/gm, ""), top = this.lexer.state.top;
      this.lexer.state.top = !0;
      let tokens = this.lexer.blockTokens(text);
      return this.lexer.state.top = top, {
        type: "blockquote",
        raw: cap[0],
        tokens,
        text
      };
    }
  }
  list(src) {
    let cap = this.rules.block.list.exec(src);
    if (cap) {
      let raw, istask, ischecked, indent, i, blankLine, endsWithBlankLine, line, nextLine, rawLine, itemContents, endEarly, bull = cap[1].trim(), isordered = bull.length > 1, list = {
        type: "list",
        raw: "",
        ordered: isordered,
        start: isordered ? +bull.slice(0, -1) : "",
        loose: !1,
        items: []
      };
      bull = isordered ? `\\d{1,9}\\${bull.slice(-1)}` : `\\${bull}`, this.options.pedantic && (bull = isordered ? bull : "[*+-]");
      let itemRegex = new RegExp(`^( {0,3}${bull})((?:[	 ][^\\n]*)?(?:\\n|$))`);
      for (; src && (endEarly = !1, !(!(cap = itemRegex.exec(src)) || this.rules.block.hr.test(src))); ) {
        if (raw = cap[0], src = src.substring(raw.length), line = cap[2].split(`
`, 1)[0].replace(/^\t+/, (t) => " ".repeat(3 * t.length)), nextLine = src.split(`
`, 1)[0], this.options.pedantic ? (indent = 2, itemContents = line.trimLeft()) : (indent = cap[2].search(/[^ ]/), indent = indent > 4 ? 1 : indent, itemContents = line.slice(indent), indent += cap[1].length), blankLine = !1, !line && /^ *$/.test(nextLine) && (raw += nextLine + `
`, src = src.substring(nextLine.length + 1), endEarly = !0), !endEarly) {
          let nextBulletRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}(?:[*+-]|\\d{1,9}[.)])((?:[ 	][^\\n]*)?(?:\\n|$))`), hrRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}((?:- *){3,}|(?:_ *){3,}|(?:\\* *){3,})(?:\\n+|$)`), fencesBeginRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}(?:\`\`\`|~~~)`), headingBeginRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}#`);
          for (; src && (rawLine = src.split(`
`, 1)[0], nextLine = rawLine, this.options.pedantic && (nextLine = nextLine.replace(/^ {1,4}(?=( {4})*[^ ])/g, "  ")), !(fencesBeginRegex.test(nextLine) || headingBeginRegex.test(nextLine) || nextBulletRegex.test(nextLine) || hrRegex.test(src))); ) {
            if (nextLine.search(/[^ ]/) >= indent || !nextLine.trim())
              itemContents += `
` + nextLine.slice(indent);
            else {
              if (blankLine || line.search(/[^ ]/) >= 4 || fencesBeginRegex.test(line) || headingBeginRegex.test(line) || hrRegex.test(line))
                break;
              itemContents += `
` + nextLine;
            }
            !blankLine && !nextLine.trim() && (blankLine = !0), raw += rawLine + `
`, src = src.substring(rawLine.length + 1), line = nextLine.slice(indent);
          }
        }
        list.loose || (endsWithBlankLine ? list.loose = !0 : /\n *\n *$/.test(raw) && (endsWithBlankLine = !0)), this.options.gfm && (istask = /^\[[ xX]\] /.exec(itemContents), istask && (ischecked = istask[0] !== "[ ] ", itemContents = itemContents.replace(/^\[[ xX]\] +/, ""))), list.items.push({
          type: "list_item",
          raw,
          task: !!istask,
          checked: ischecked,
          loose: !1,
          text: itemContents
        }), list.raw += raw;
      }
      list.items[list.items.length - 1].raw = raw.trimRight(), list.items[list.items.length - 1].text = itemContents.trimRight(), list.raw = list.raw.trimRight();
      let l = list.items.length;
      for (i = 0; i < l; i++)
        if (this.lexer.state.top = !1, list.items[i].tokens = this.lexer.blockTokens(list.items[i].text, []), !list.loose) {
          let spacers = list.items[i].tokens.filter((t) => t.type === "space"), hasMultipleLineBreaks = spacers.length > 0 && spacers.some((t) => /\n.*\n/.test(t.raw));
          list.loose = hasMultipleLineBreaks;
        }
      if (list.loose)
        for (i = 0; i < l; i++)
          list.items[i].loose = !0;
      return list;
    }
  }
  html(src) {
    let cap = this.rules.block.html.exec(src);
    if (cap) {
      let token = {
        type: "html",
        raw: cap[0],
        pre: !this.options.sanitizer && (cap[1] === "pre" || cap[1] === "script" || cap[1] === "style"),
        text: cap[0]
      };
      if (this.options.sanitize) {
        let text = this.options.sanitizer ? this.options.sanitizer(cap[0]) : escape(cap[0]);
        token.type = "paragraph", token.text = text, token.tokens = this.lexer.inline(text);
      }
      return token;
    }
  }
  def(src) {
    let cap = this.rules.block.def.exec(src);
    if (cap) {
      let tag = cap[1].toLowerCase().replace(/\s+/g, " "), href = cap[2] ? cap[2].replace(/^<(.*)>$/, "$1").replace(this.rules.inline._escapes, "$1") : "", title = cap[3] ? cap[3].substring(1, cap[3].length - 1).replace(this.rules.inline._escapes, "$1") : cap[3];
      return {
        type: "def",
        tag,
        raw: cap[0],
        href,
        title
      };
    }
  }
  table(src) {
    let cap = this.rules.block.table.exec(src);
    if (cap) {
      let item = {
        type: "table",
        header: splitCells(cap[1]).map((c) => ({ text: c })),
        align: cap[2].replace(/^ *|\| *$/g, "").split(/ *\| */),
        rows: cap[3] && cap[3].trim() ? cap[3].replace(/\n[ \t]*$/, "").split(`
`) : []
      };
      if (item.header.length === item.align.length) {
        item.raw = cap[0];
        let l = item.align.length, i, j, k, row;
        for (i = 0; i < l; i++)
          /^ *-+: *$/.test(item.align[i]) ? item.align[i] = "right" : /^ *:-+: *$/.test(item.align[i]) ? item.align[i] = "center" : /^ *:-+ *$/.test(item.align[i]) ? item.align[i] = "left" : item.align[i] = null;
        for (l = item.rows.length, i = 0; i < l; i++)
          item.rows[i] = splitCells(item.rows[i], item.header.length).map((c) => ({ text: c }));
        for (l = item.header.length, j = 0; j < l; j++)
          item.header[j].tokens = this.lexer.inline(item.header[j].text);
        for (l = item.rows.length, j = 0; j < l; j++)
          for (row = item.rows[j], k = 0; k < row.length; k++)
            row[k].tokens = this.lexer.inline(row[k].text);
        return item;
      }
    }
  }
  lheading(src) {
    let cap = this.rules.block.lheading.exec(src);
    if (cap)
      return {
        type: "heading",
        raw: cap[0],
        depth: cap[2].charAt(0) === "=" ? 1 : 2,
        text: cap[1],
        tokens: this.lexer.inline(cap[1])
      };
  }
  paragraph(src) {
    let cap = this.rules.block.paragraph.exec(src);
    if (cap) {
      let text = cap[1].charAt(cap[1].length - 1) === `
` ? cap[1].slice(0, -1) : cap[1];
      return {
        type: "paragraph",
        raw: cap[0],
        text,
        tokens: this.lexer.inline(text)
      };
    }
  }
  text(src) {
    let cap = this.rules.block.text.exec(src);
    if (cap)
      return {
        type: "text",
        raw: cap[0],
        text: cap[0],
        tokens: this.lexer.inline(cap[0])
      };
  }
  escape(src) {
    let cap = this.rules.inline.escape.exec(src);
    if (cap)
      return {
        type: "escape",
        raw: cap[0],
        text: escape(cap[1])
      };
  }
  tag(src) {
    let cap = this.rules.inline.tag.exec(src);
    if (cap)
      return !this.lexer.state.inLink && /^<a /i.test(cap[0]) ? this.lexer.state.inLink = !0 : this.lexer.state.inLink && /^<\/a>/i.test(cap[0]) && (this.lexer.state.inLink = !1), !this.lexer.state.inRawBlock && /^<(pre|code|kbd|script)(\s|>)/i.test(cap[0]) ? this.lexer.state.inRawBlock = !0 : this.lexer.state.inRawBlock && /^<\/(pre|code|kbd|script)(\s|>)/i.test(cap[0]) && (this.lexer.state.inRawBlock = !1), {
        type: this.options.sanitize ? "text" : "html",
        raw: cap[0],
        inLink: this.lexer.state.inLink,
        inRawBlock: this.lexer.state.inRawBlock,
        text: this.options.sanitize ? this.options.sanitizer ? this.options.sanitizer(cap[0]) : escape(cap[0]) : cap[0]
      };
  }
  link(src) {
    let cap = this.rules.inline.link.exec(src);
    if (cap) {
      let trimmedUrl = cap[2].trim();
      if (!this.options.pedantic && /^</.test(trimmedUrl)) {
        if (!/>$/.test(trimmedUrl))
          return;
        let rtrimSlash = rtrim(trimmedUrl.slice(0, -1), "\\");
        if ((trimmedUrl.length - rtrimSlash.length) % 2 === 0)
          return;
      } else {
        let lastParenIndex = findClosingBracket(cap[2], "()");
        if (lastParenIndex > -1) {
          let linkLen = (cap[0].indexOf("!") === 0 ? 5 : 4) + cap[1].length + lastParenIndex;
          cap[2] = cap[2].substring(0, lastParenIndex), cap[0] = cap[0].substring(0, linkLen).trim(), cap[3] = "";
        }
      }
      let href = cap[2], title = "";
      if (this.options.pedantic) {
        let link2 = /^([^'"]*[^\s])\s+(['"])(.*)\2/.exec(href);
        link2 && (href = link2[1], title = link2[3]);
      } else
        title = cap[3] ? cap[3].slice(1, -1) : "";
      return href = href.trim(), /^</.test(href) && (this.options.pedantic && !/>$/.test(trimmedUrl) ? href = href.slice(1) : href = href.slice(1, -1)), outputLink(cap, {
        href: href && href.replace(this.rules.inline._escapes, "$1"),
        title: title && title.replace(this.rules.inline._escapes, "$1")
      }, cap[0], this.lexer);
    }
  }
  reflink(src, links5) {
    let cap;
    if ((cap = this.rules.inline.reflink.exec(src)) || (cap = this.rules.inline.nolink.exec(src))) {
      let link2 = (cap[2] || cap[1]).replace(/\s+/g, " ");
      if (link2 = links5[link2.toLowerCase()], !link2) {
        let text = cap[0].charAt(0);
        return {
          type: "text",
          raw: text,
          text
        };
      }
      return outputLink(cap, link2, cap[0], this.lexer);
    }
  }
  emStrong(src, maskedSrc, prevChar = "") {
    let match = this.rules.inline.emStrong.lDelim.exec(src);
    if (!match || match[3] && prevChar.match(/[\p{L}\p{N}]/u))
      return;
    let nextChar = match[1] || match[2] || "";
    if (!nextChar || nextChar && (prevChar === "" || this.rules.inline.punctuation.exec(prevChar))) {
      let lLength = match[0].length - 1, rDelim, rLength, delimTotal = lLength, midDelimTotal = 0, endReg = match[0][0] === "*" ? this.rules.inline.emStrong.rDelimAst : this.rules.inline.emStrong.rDelimUnd;
      for (endReg.lastIndex = 0, maskedSrc = maskedSrc.slice(-1 * src.length + lLength); (match = endReg.exec(maskedSrc)) != null; ) {
        if (rDelim = match[1] || match[2] || match[3] || match[4] || match[5] || match[6], !rDelim)
          continue;
        if (rLength = rDelim.length, match[3] || match[4]) {
          delimTotal += rLength;
          continue;
        } else if ((match[5] || match[6]) && lLength % 3 && !((lLength + rLength) % 3)) {
          midDelimTotal += rLength;
          continue;
        }
        if (delimTotal -= rLength, delimTotal > 0)
          continue;
        rLength = Math.min(rLength, rLength + delimTotal + midDelimTotal);
        let raw = src.slice(0, lLength + match.index + (match[0].length - rDelim.length) + rLength);
        if (Math.min(lLength, rLength) % 2) {
          let text2 = raw.slice(1, -1);
          return {
            type: "em",
            raw,
            text: text2,
            tokens: this.lexer.inlineTokens(text2)
          };
        }
        let text = raw.slice(2, -2);
        return {
          type: "strong",
          raw,
          text,
          tokens: this.lexer.inlineTokens(text)
        };
      }
    }
  }
  codespan(src) {
    let cap = this.rules.inline.code.exec(src);
    if (cap) {
      let text = cap[2].replace(/\n/g, " "), hasNonSpaceChars = /[^ ]/.test(text), hasSpaceCharsOnBothEnds = /^ /.test(text) && / $/.test(text);
      return hasNonSpaceChars && hasSpaceCharsOnBothEnds && (text = text.substring(1, text.length - 1)), text = escape(text, !0), {
        type: "codespan",
        raw: cap[0],
        text
      };
    }
  }
  br(src) {
    let cap = this.rules.inline.br.exec(src);
    if (cap)
      return {
        type: "br",
        raw: cap[0]
      };
  }
  del(src) {
    let cap = this.rules.inline.del.exec(src);
    if (cap)
      return {
        type: "del",
        raw: cap[0],
        text: cap[2],
        tokens: this.lexer.inlineTokens(cap[2])
      };
  }
  autolink(src, mangle2) {
    let cap = this.rules.inline.autolink.exec(src);
    if (cap) {
      let text, href;
      return cap[2] === "@" ? (text = escape(this.options.mangle ? mangle2(cap[1]) : cap[1]), href = "mailto:" + text) : (text = escape(cap[1]), href = text), {
        type: "link",
        raw: cap[0],
        text,
        href,
        tokens: [
          {
            type: "text",
            raw: text,
            text
          }
        ]
      };
    }
  }
  url(src, mangle2) {
    let cap;
    if (cap = this.rules.inline.url.exec(src)) {
      let text, href;
      if (cap[2] === "@")
        text = escape(this.options.mangle ? mangle2(cap[0]) : cap[0]), href = "mailto:" + text;
      else {
        let prevCapZero;
        do
          prevCapZero = cap[0], cap[0] = this.rules.inline._backpedal.exec(cap[0])[0];
        while (prevCapZero !== cap[0]);
        text = escape(cap[0]), cap[1] === "www." ? href = "http://" + cap[0] : href = cap[0];
      }
      return {
        type: "link",
        raw: cap[0],
        text,
        href,
        tokens: [
          {
            type: "text",
            raw: text,
            text
          }
        ]
      };
    }
  }
  inlineText(src, smartypants2) {
    let cap = this.rules.inline.text.exec(src);
    if (cap) {
      let text;
      return this.lexer.state.inRawBlock ? text = this.options.sanitize ? this.options.sanitizer ? this.options.sanitizer(cap[0]) : escape(cap[0]) : cap[0] : text = escape(this.options.smartypants ? smartypants2(cap[0]) : cap[0]), {
        type: "text",
        raw: cap[0],
        text
      };
    }
  }
}, block = {
  newline: /^(?: *(?:\n|$))+/,
  code: /^( {4}[^\n]+(?:\n(?: *(?:\n|$))*)?)+/,
  fences: /^ {0,3}(`{3,}(?=[^`\n]*\n)|~{3,})([^\n]*)\n(?:|([\s\S]*?)\n)(?: {0,3}\1[~`]* *(?=\n|$)|$)/,
  hr: /^ {0,3}((?:-[\t ]*){3,}|(?:_[ \t]*){3,}|(?:\*[ \t]*){3,})(?:\n+|$)/,
  heading: /^ {0,3}(#{1,6})(?=\s|$)(.*)(?:\n+|$)/,
  blockquote: /^( {0,3}> ?(paragraph|[^\n]*)(?:\n|$))+/,
  list: /^( {0,3}bull)([ \t][^\n]+?)?(?:\n|$)/,
  html: "^ {0,3}(?:<(script|pre|style|textarea)[\\s>][\\s\\S]*?(?:</\\1>[^\\n]*\\n+|$)|comment[^\\n]*(\\n+|$)|<\\?[\\s\\S]*?(?:\\?>\\n*|$)|<![A-Z][\\s\\S]*?(?:>\\n*|$)|<!\\[CDATA\\[[\\s\\S]*?(?:\\]\\]>\\n*|$)|</?(tag)(?: +|\\n|/?>)[\\s\\S]*?(?:(?:\\n *)+\\n|$)|<(?!script|pre|style|textarea)([a-z][\\w-]*)(?:attribute)*? */?>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$)|</(?!script|pre|style|textarea)[a-z][\\w-]*\\s*>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$))",
  def: /^ {0,3}\[(label)\]: *(?:\n *)?([^<\s][^\s]*|<.*?>)(?:(?: +(?:\n *)?| *\n *)(title))? *(?:\n+|$)/,
  table: noopTest,
  lheading: /^((?:.|\n(?!\n))+?)\n {0,3}(=+|-+) *(?:\n+|$)/,
  _paragraph: /^([^\n]+(?:\n(?!hr|heading|lheading|blockquote|fences|list|html|table| +\n)[^\n]+)*)/,
  text: /^[^\n]+/
};
block._label = /(?!\s*\])(?:\\.|[^\[\]\\])+/;
block._title = /(?:"(?:\\"?|[^"\\])*"|'[^'\n]*(?:\n[^'\n]+)*\n?'|\([^()]*\))/;
block.def = edit(block.def).replace("label", block._label).replace("title", block._title).getRegex();
block.bullet = /(?:[*+-]|\d{1,9}[.)])/;
block.listItemStart = edit(/^( *)(bull) */).replace("bull", block.bullet).getRegex();
block.list = edit(block.list).replace(/bull/g, block.bullet).replace("hr", "\\n+(?=\\1?(?:(?:- *){3,}|(?:_ *){3,}|(?:\\* *){3,})(?:\\n+|$))").replace("def", "\\n+(?=" + block.def.source + ")").getRegex();
block._tag = "address|article|aside|base|basefont|blockquote|body|caption|center|col|colgroup|dd|details|dialog|dir|div|dl|dt|fieldset|figcaption|figure|footer|form|frame|frameset|h[1-6]|head|header|hr|html|iframe|legend|li|link|main|menu|menuitem|meta|nav|noframes|ol|optgroup|option|p|param|section|source|summary|table|tbody|td|tfoot|th|thead|title|tr|track|ul";
block._comment = /<!--(?!-?>)[\s\S]*?(?:-->|$)/;
block.html = edit(block.html, "i").replace("comment", block._comment).replace("tag", block._tag).replace("attribute", / +[a-zA-Z:_][\w.:-]*(?: *= *"[^"\n]*"| *= *'[^'\n]*'| *= *[^\s"'=<>`]+)?/).getRegex();
block.paragraph = edit(block._paragraph).replace("hr", block.hr).replace("heading", " {0,3}#{1,6} ").replace("|lheading", "").replace("|table", "").replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", block._tag).getRegex();
block.blockquote = edit(block.blockquote).replace("paragraph", block.paragraph).getRegex();
block.normal = merge({}, block);
block.gfm = merge({}, block.normal, {
  table: "^ *([^\\n ].*\\|.*)\\n {0,3}(?:\\| *)?(:?-+:? *(?:\\| *:?-+:? *)*)(?:\\| *)?(?:\\n((?:(?! *\\n|hr|heading|blockquote|code|fences|list|html).*(?:\\n|$))*)\\n*|$)"
});
block.gfm.table = edit(block.gfm.table).replace("hr", block.hr).replace("heading", " {0,3}#{1,6} ").replace("blockquote", " {0,3}>").replace("code", " {4}[^\\n]").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", block._tag).getRegex();
block.gfm.paragraph = edit(block._paragraph).replace("hr", block.hr).replace("heading", " {0,3}#{1,6} ").replace("|lheading", "").replace("table", block.gfm.table).replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", block._tag).getRegex();
block.pedantic = merge({}, block.normal, {
  html: edit(
    `^ *(?:comment *(?:\\n|\\s*$)|<(tag)[\\s\\S]+?</\\1> *(?:\\n{2,}|\\s*$)|<tag(?:"[^"]*"|'[^']*'|\\s[^'"/>\\s]*)*?/?> *(?:\\n{2,}|\\s*$))`
  ).replace("comment", block._comment).replace(/tag/g, "(?!(?:a|em|strong|small|s|cite|q|dfn|abbr|data|time|code|var|samp|kbd|sub|sup|i|b|u|mark|ruby|rt|rp|bdi|bdo|span|br|wbr|ins|del|img)\\b)\\w+(?!:|[^\\w\\s@]*@)\\b").getRegex(),
  def: /^ *\[([^\]]+)\]: *<?([^\s>]+)>?(?: +(["(][^\n]+[")]))? *(?:\n+|$)/,
  heading: /^(#{1,6})(.*)(?:\n+|$)/,
  fences: noopTest,
  lheading: /^(.+?)\n {0,3}(=+|-+) *(?:\n+|$)/,
  paragraph: edit(block.normal._paragraph).replace("hr", block.hr).replace("heading", ` *#{1,6} *[^
]`).replace("lheading", block.lheading).replace("blockquote", " {0,3}>").replace("|fences", "").replace("|list", "").replace("|html", "").getRegex()
});
var inline = {
  escape: /^\\([!"#$%&'()*+,\-./:;<=>?@\[\]\\^_`{|}~])/,
  autolink: /^<(scheme:[^\s\x00-\x1f<>]*|email)>/,
  url: noopTest,
  tag: "^comment|^</[a-zA-Z][\\w:-]*\\s*>|^<[a-zA-Z][\\w-]*(?:attribute)*?\\s*/?>|^<\\?[\\s\\S]*?\\?>|^<![a-zA-Z]+\\s[\\s\\S]*?>|^<!\\[CDATA\\[[\\s\\S]*?\\]\\]>",
  link: /^!?\[(label)\]\(\s*(href)(?:\s+(title))?\s*\)/,
  reflink: /^!?\[(label)\]\[(ref)\]/,
  nolink: /^!?\[(ref)\](?:\[\])?/,
  reflinkSearch: "reflink|nolink(?!\\()",
  emStrong: {
    lDelim: /^(?:\*+(?:([punct_])|[^\s*]))|^_+(?:([punct*])|([^\s_]))/,
    rDelimAst: /^(?:[^_*\\]|\\.)*?\_\_(?:[^_*\\]|\\.)*?\*(?:[^_*\\]|\\.)*?(?=\_\_)|(?:[^*\\]|\\.)+(?=[^*])|[punct_](\*+)(?=[\s]|$)|(?:[^punct*_\s\\]|\\.)(\*+)(?=[punct_\s]|$)|[punct_\s](\*+)(?=[^punct*_\s])|[\s](\*+)(?=[punct_])|[punct_](\*+)(?=[punct_])|(?:[^punct*_\s\\]|\\.)(\*+)(?=[^punct*_\s])/,
    rDelimUnd: /^(?:[^_*\\]|\\.)*?\*\*(?:[^_*\\]|\\.)*?\_(?:[^_*\\]|\\.)*?(?=\*\*)|(?:[^_\\]|\\.)+(?=[^_])|[punct*](\_+)(?=[\s]|$)|(?:[^punct*_\s\\]|\\.)(\_+)(?=[punct*\s]|$)|[punct*\s](\_+)(?=[^punct*_\s])|[\s](\_+)(?=[punct*])|[punct*](\_+)(?=[punct*])/
  },
  code: /^(`+)([^`]|[^`][\s\S]*?[^`])\1(?!`)/,
  br: /^( {2,}|\\)\n(?!\s*$)/,
  del: noopTest,
  text: /^(`+|[^`])(?:(?= {2,}\n)|[\s\S]*?(?:(?=[\\<!\[`*_]|\b_|$)|[^ ](?= {2,}\n)))/,
  punctuation: /^([\spunctuation])/
};
inline._punctuation = "!\"#$%&'()+\\-.,/:;<=>?@\\[\\]`^{|}~";
inline.punctuation = edit(inline.punctuation).replace(/punctuation/g, inline._punctuation).getRegex();
inline.blockSkip = /\[[^\]]*?\]\([^\)]*?\)|`[^`]*?`|<[^>]*?>/g;
inline.escapedEmSt = /(?:^|[^\\])(?:\\\\)*\\[*_]/g;
inline._comment = edit(block._comment).replace("(?:-->|$)", "-->").getRegex();
inline.emStrong.lDelim = edit(inline.emStrong.lDelim).replace(/punct/g, inline._punctuation).getRegex();
inline.emStrong.rDelimAst = edit(inline.emStrong.rDelimAst, "g").replace(/punct/g, inline._punctuation).getRegex();
inline.emStrong.rDelimUnd = edit(inline.emStrong.rDelimUnd, "g").replace(/punct/g, inline._punctuation).getRegex();
inline._escapes = /\\([!"#$%&'()*+,\-./:;<=>?@\[\]\\^_`{|}~])/g;
inline._scheme = /[a-zA-Z][a-zA-Z0-9+.-]{1,31}/;
inline._email = /[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+(@)[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+(?![-_])/;
inline.autolink = edit(inline.autolink).replace("scheme", inline._scheme).replace("email", inline._email).getRegex();
inline._attribute = /\s+[a-zA-Z:_][\w.:-]*(?:\s*=\s*"[^"]*"|\s*=\s*'[^']*'|\s*=\s*[^\s"'=<>`]+)?/;
inline.tag = edit(inline.tag).replace("comment", inline._comment).replace("attribute", inline._attribute).getRegex();
inline._label = /(?:\[(?:\\.|[^\[\]\\])*\]|\\.|`[^`]*`|[^\[\]\\`])*?/;
inline._href = /<(?:\\.|[^\n<>\\])+>|[^\s\x00-\x1f]*/;
inline._title = /"(?:\\"?|[^"\\])*"|'(?:\\'?|[^'\\])*'|\((?:\\\)?|[^)\\])*\)/;
inline.link = edit(inline.link).replace("label", inline._label).replace("href", inline._href).replace("title", inline._title).getRegex();
inline.reflink = edit(inline.reflink).replace("label", inline._label).replace("ref", block._label).getRegex();
inline.nolink = edit(inline.nolink).replace("ref", block._label).getRegex();
inline.reflinkSearch = edit(inline.reflinkSearch, "g").replace("reflink", inline.reflink).replace("nolink", inline.nolink).getRegex();
inline.normal = merge({}, inline);
inline.pedantic = merge({}, inline.normal, {
  strong: {
    start: /^__|\*\*/,
    middle: /^__(?=\S)([\s\S]*?\S)__(?!_)|^\*\*(?=\S)([\s\S]*?\S)\*\*(?!\*)/,
    endAst: /\*\*(?!\*)/g,
    endUnd: /__(?!_)/g
  },
  em: {
    start: /^_|\*/,
    middle: /^()\*(?=\S)([\s\S]*?\S)\*(?!\*)|^_(?=\S)([\s\S]*?\S)_(?!_)/,
    endAst: /\*(?!\*)/g,
    endUnd: /_(?!_)/g
  },
  link: edit(/^!?\[(label)\]\((.*?)\)/).replace("label", inline._label).getRegex(),
  reflink: edit(/^!?\[(label)\]\s*\[([^\]]*)\]/).replace("label", inline._label).getRegex()
});
inline.gfm = merge({}, inline.normal, {
  escape: edit(inline.escape).replace("])", "~|])").getRegex(),
  _extended_email: /[A-Za-z0-9._+-]+(@)[a-zA-Z0-9-_]+(?:\.[a-zA-Z0-9-_]*[a-zA-Z0-9])+(?![-_])/,
  url: /^((?:ftp|https?):\/\/|www\.)(?:[a-zA-Z0-9\-]+\.?)+[^\s<]*|^email/,
  _backpedal: /(?:[^?!.,:;*_'"~()&]+|\([^)]*\)|&(?![a-zA-Z0-9]+;$)|[?!.,:;*_'"~)]+(?!$))+/,
  del: /^(~~?)(?=[^\s~])([\s\S]*?[^\s~])\1(?=[^~]|$)/,
  text: /^([`~]+|[^`~])(?:(?= {2,}\n)|(?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)|[\s\S]*?(?:(?=[\\<!\[`*~_]|\b_|https?:\/\/|ftp:\/\/|www\.|$)|[^ ](?= {2,}\n)|[^a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-](?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)))/
});
inline.gfm.url = edit(inline.gfm.url, "i").replace("email", inline.gfm._extended_email).getRegex();
inline.breaks = merge({}, inline.gfm, {
  br: edit(inline.br).replace("{2,}", "*").getRegex(),
  text: edit(inline.gfm.text).replace("\\b_", "\\b_| {2,}\\n").replace(/\{2,\}/g, "*").getRegex()
});
function smartypants(text) {
  return text.replace(/---/g, "\u2014").replace(/--/g, "\u2013").replace(/(^|[-\u2014/(\[{"\s])'/g, "$1\u2018").replace(/'/g, "\u2019").replace(/(^|[-\u2014/(\[{\u2018\s])"/g, "$1\u201C").replace(/"/g, "\u201D").replace(/\.{3}/g, "\u2026");
}
function mangle(text) {
  let out = "", i, ch, l = text.length;
  for (i = 0; i < l; i++)
    ch = text.charCodeAt(i), Math.random() > 0.5 && (ch = "x" + ch.toString(16)), out += "&#" + ch + ";";
  return out;
}
var Lexer = class {
  constructor(options2) {
    this.tokens = [], this.tokens.links = /* @__PURE__ */ Object.create(null), this.options = options2 || defaults, this.options.tokenizer = this.options.tokenizer || new Tokenizer(), this.tokenizer = this.options.tokenizer, this.tokenizer.options = this.options, this.tokenizer.lexer = this, this.inlineQueue = [], this.state = {
      inLink: !1,
      inRawBlock: !1,
      top: !0
    };
    let rules = {
      block: block.normal,
      inline: inline.normal
    };
    this.options.pedantic ? (rules.block = block.pedantic, rules.inline = inline.pedantic) : this.options.gfm && (rules.block = block.gfm, this.options.breaks ? rules.inline = inline.breaks : rules.inline = inline.gfm), this.tokenizer.rules = rules;
  }
  static get rules() {
    return {
      block,
      inline
    };
  }
  static lex(src, options2) {
    return new Lexer(options2).lex(src);
  }
  static lexInline(src, options2) {
    return new Lexer(options2).inlineTokens(src);
  }
  lex(src) {
    src = src.replace(/\r\n|\r/g, `
`), this.blockTokens(src, this.tokens);
    let next;
    for (; next = this.inlineQueue.shift(); )
      this.inlineTokens(next.src, next.tokens);
    return this.tokens;
  }
  blockTokens(src, tokens = []) {
    this.options.pedantic ? src = src.replace(/\t/g, "    ").replace(/^ +$/gm, "") : src = src.replace(/^( *)(\t+)/gm, (_, leading, tabs) => leading + "    ".repeat(tabs.length));
    let token, lastToken, cutSrc, lastParagraphClipped;
    for (; src; )
      if (!(this.options.extensions && this.options.extensions.block && this.options.extensions.block.some((extTokenizer) => (token = extTokenizer.call({ lexer: this }, src, tokens)) ? (src = src.substring(token.raw.length), tokens.push(token), !0) : !1))) {
        if (token = this.tokenizer.space(src)) {
          src = src.substring(token.raw.length), token.raw.length === 1 && tokens.length > 0 ? tokens[tokens.length - 1].raw += `
` : tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.code(src)) {
          src = src.substring(token.raw.length), lastToken = tokens[tokens.length - 1], lastToken && (lastToken.type === "paragraph" || lastToken.type === "text") ? (lastToken.raw += `
` + token.raw, lastToken.text += `
` + token.text, this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text) : tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.fences(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.heading(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.hr(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.blockquote(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.list(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.html(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.def(src)) {
          src = src.substring(token.raw.length), lastToken = tokens[tokens.length - 1], lastToken && (lastToken.type === "paragraph" || lastToken.type === "text") ? (lastToken.raw += `
` + token.raw, lastToken.text += `
` + token.raw, this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text) : this.tokens.links[token.tag] || (this.tokens.links[token.tag] = {
            href: token.href,
            title: token.title
          });
          continue;
        }
        if (token = this.tokenizer.table(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.lheading(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (cutSrc = src, this.options.extensions && this.options.extensions.startBlock) {
          let startIndex = 1 / 0, tempSrc = src.slice(1), tempStart;
          this.options.extensions.startBlock.forEach(function(getStartIndex) {
            tempStart = getStartIndex.call({ lexer: this }, tempSrc), typeof tempStart == "number" && tempStart >= 0 && (startIndex = Math.min(startIndex, tempStart));
          }), startIndex < 1 / 0 && startIndex >= 0 && (cutSrc = src.substring(0, startIndex + 1));
        }
        if (this.state.top && (token = this.tokenizer.paragraph(cutSrc))) {
          lastToken = tokens[tokens.length - 1], lastParagraphClipped && lastToken.type === "paragraph" ? (lastToken.raw += `
` + token.raw, lastToken.text += `
` + token.text, this.inlineQueue.pop(), this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text) : tokens.push(token), lastParagraphClipped = cutSrc.length !== src.length, src = src.substring(token.raw.length);
          continue;
        }
        if (token = this.tokenizer.text(src)) {
          src = src.substring(token.raw.length), lastToken = tokens[tokens.length - 1], lastToken && lastToken.type === "text" ? (lastToken.raw += `
` + token.raw, lastToken.text += `
` + token.text, this.inlineQueue.pop(), this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text) : tokens.push(token);
          continue;
        }
        if (src) {
          let errMsg = "Infinite loop on byte: " + src.charCodeAt(0);
          if (this.options.silent) {
            console.error(errMsg);
            break;
          } else
            throw new Error(errMsg);
        }
      }
    return this.state.top = !0, tokens;
  }
  inline(src, tokens = []) {
    return this.inlineQueue.push({ src, tokens }), tokens;
  }
  inlineTokens(src, tokens = []) {
    let token, lastToken, cutSrc, maskedSrc = src, match, keepPrevChar, prevChar;
    if (this.tokens.links) {
      let links5 = Object.keys(this.tokens.links);
      if (links5.length > 0)
        for (; (match = this.tokenizer.rules.inline.reflinkSearch.exec(maskedSrc)) != null; )
          links5.includes(match[0].slice(match[0].lastIndexOf("[") + 1, -1)) && (maskedSrc = maskedSrc.slice(0, match.index) + "[" + repeatString("a", match[0].length - 2) + "]" + maskedSrc.slice(this.tokenizer.rules.inline.reflinkSearch.lastIndex));
    }
    for (; (match = this.tokenizer.rules.inline.blockSkip.exec(maskedSrc)) != null; )
      maskedSrc = maskedSrc.slice(0, match.index) + "[" + repeatString("a", match[0].length - 2) + "]" + maskedSrc.slice(this.tokenizer.rules.inline.blockSkip.lastIndex);
    for (; (match = this.tokenizer.rules.inline.escapedEmSt.exec(maskedSrc)) != null; )
      maskedSrc = maskedSrc.slice(0, match.index + match[0].length - 2) + "++" + maskedSrc.slice(this.tokenizer.rules.inline.escapedEmSt.lastIndex), this.tokenizer.rules.inline.escapedEmSt.lastIndex--;
    for (; src; )
      if (keepPrevChar || (prevChar = ""), keepPrevChar = !1, !(this.options.extensions && this.options.extensions.inline && this.options.extensions.inline.some((extTokenizer) => (token = extTokenizer.call({ lexer: this }, src, tokens)) ? (src = src.substring(token.raw.length), tokens.push(token), !0) : !1))) {
        if (token = this.tokenizer.escape(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.tag(src)) {
          src = src.substring(token.raw.length), lastToken = tokens[tokens.length - 1], lastToken && token.type === "text" && lastToken.type === "text" ? (lastToken.raw += token.raw, lastToken.text += token.text) : tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.link(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.reflink(src, this.tokens.links)) {
          src = src.substring(token.raw.length), lastToken = tokens[tokens.length - 1], lastToken && token.type === "text" && lastToken.type === "text" ? (lastToken.raw += token.raw, lastToken.text += token.text) : tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.emStrong(src, maskedSrc, prevChar)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.codespan(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.br(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.del(src)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (token = this.tokenizer.autolink(src, mangle)) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (!this.state.inLink && (token = this.tokenizer.url(src, mangle))) {
          src = src.substring(token.raw.length), tokens.push(token);
          continue;
        }
        if (cutSrc = src, this.options.extensions && this.options.extensions.startInline) {
          let startIndex = 1 / 0, tempSrc = src.slice(1), tempStart;
          this.options.extensions.startInline.forEach(function(getStartIndex) {
            tempStart = getStartIndex.call({ lexer: this }, tempSrc), typeof tempStart == "number" && tempStart >= 0 && (startIndex = Math.min(startIndex, tempStart));
          }), startIndex < 1 / 0 && startIndex >= 0 && (cutSrc = src.substring(0, startIndex + 1));
        }
        if (token = this.tokenizer.inlineText(cutSrc, smartypants)) {
          src = src.substring(token.raw.length), token.raw.slice(-1) !== "_" && (prevChar = token.raw.slice(-1)), keepPrevChar = !0, lastToken = tokens[tokens.length - 1], lastToken && lastToken.type === "text" ? (lastToken.raw += token.raw, lastToken.text += token.text) : tokens.push(token);
          continue;
        }
        if (src) {
          let errMsg = "Infinite loop on byte: " + src.charCodeAt(0);
          if (this.options.silent) {
            console.error(errMsg);
            break;
          } else
            throw new Error(errMsg);
        }
      }
    return tokens;
  }
}, Renderer = class {
  constructor(options2) {
    this.options = options2 || defaults;
  }
  code(code, infostring, escaped) {
    let lang = (infostring || "").match(/\S*/)[0];
    if (this.options.highlight) {
      let out = this.options.highlight(code, lang);
      out != null && out !== code && (escaped = !0, code = out);
    }
    return code = code.replace(/\n$/, "") + `
`, lang ? '<pre><code class="' + this.options.langPrefix + escape(lang) + '">' + (escaped ? code : escape(code, !0)) + `</code></pre>
` : "<pre><code>" + (escaped ? code : escape(code, !0)) + `</code></pre>
`;
  }
  blockquote(quote) {
    return `<blockquote>
${quote}</blockquote>
`;
  }
  html(html) {
    return html;
  }
  heading(text, level, raw, slugger) {
    if (this.options.headerIds) {
      let id = this.options.headerPrefix + slugger.slug(raw);
      return `<h${level} id="${id}">${text}</h${level}>
`;
    }
    return `<h${level}>${text}</h${level}>
`;
  }
  hr() {
    return this.options.xhtml ? `<hr/>
` : `<hr>
`;
  }
  list(body, ordered, start) {
    let type = ordered ? "ol" : "ul", startatt = ordered && start !== 1 ? ' start="' + start + '"' : "";
    return "<" + type + startatt + `>
` + body + "</" + type + `>
`;
  }
  listitem(text) {
    return `<li>${text}</li>
`;
  }
  checkbox(checked) {
    return "<input " + (checked ? 'checked="" ' : "") + 'disabled="" type="checkbox"' + (this.options.xhtml ? " /" : "") + "> ";
  }
  paragraph(text) {
    return `<p>${text}</p>
`;
  }
  table(header, body) {
    return body && (body = `<tbody>${body}</tbody>`), `<table>
<thead>
` + header + `</thead>
` + body + `</table>
`;
  }
  tablerow(content) {
    return `<tr>
${content}</tr>
`;
  }
  tablecell(content, flags) {
    let type = flags.header ? "th" : "td";
    return (flags.align ? `<${type} align="${flags.align}">` : `<${type}>`) + content + `</${type}>
`;
  }
  strong(text) {
    return `<strong>${text}</strong>`;
  }
  em(text) {
    return `<em>${text}</em>`;
  }
  codespan(text) {
    return `<code>${text}</code>`;
  }
  br() {
    return this.options.xhtml ? "<br/>" : "<br>";
  }
  del(text) {
    return `<del>${text}</del>`;
  }
  link(href, title, text) {
    if (href = cleanUrl(this.options.sanitize, this.options.baseUrl, href), href === null)
      return text;
    let out = '<a href="' + href + '"';
    return title && (out += ' title="' + title + '"'), out += ">" + text + "</a>", out;
  }
  image(href, title, text) {
    if (href = cleanUrl(this.options.sanitize, this.options.baseUrl, href), href === null)
      return text;
    let out = `<img src="${href}" alt="${text}"`;
    return title && (out += ` title="${title}"`), out += this.options.xhtml ? "/>" : ">", out;
  }
  text(text) {
    return text;
  }
}, TextRenderer = class {
  strong(text) {
    return text;
  }
  em(text) {
    return text;
  }
  codespan(text) {
    return text;
  }
  del(text) {
    return text;
  }
  html(text) {
    return text;
  }
  text(text) {
    return text;
  }
  link(href, title, text) {
    return "" + text;
  }
  image(href, title, text) {
    return "" + text;
  }
  br() {
    return "";
  }
}, Slugger = class {
  constructor() {
    this.seen = {};
  }
  serialize(value) {
    return value.toLowerCase().trim().replace(/<[!\/a-z].*?>/ig, "").replace(/[\u2000-\u206F\u2E00-\u2E7F\\'!"#$%&()*+,./:;<=>?@[\]^`{|}~]/g, "").replace(/\s/g, "-");
  }
  getNextSafeSlug(originalSlug, isDryRun) {
    let slug = originalSlug, occurenceAccumulator = 0;
    if (this.seen.hasOwnProperty(slug)) {
      occurenceAccumulator = this.seen[originalSlug];
      do
        occurenceAccumulator++, slug = originalSlug + "-" + occurenceAccumulator;
      while (this.seen.hasOwnProperty(slug));
    }
    return isDryRun || (this.seen[originalSlug] = occurenceAccumulator, this.seen[slug] = 0), slug;
  }
  slug(value, options2 = {}) {
    let slug = this.serialize(value);
    return this.getNextSafeSlug(slug, options2.dryrun);
  }
}, Parser = class {
  constructor(options2) {
    this.options = options2 || defaults, this.options.renderer = this.options.renderer || new Renderer(), this.renderer = this.options.renderer, this.renderer.options = this.options, this.textRenderer = new TextRenderer(), this.slugger = new Slugger();
  }
  static parse(tokens, options2) {
    return new Parser(options2).parse(tokens);
  }
  static parseInline(tokens, options2) {
    return new Parser(options2).parseInline(tokens);
  }
  parse(tokens, top = !0) {
    let out = "", i, j, k, l2, l3, row, cell, header, body, token, ordered, start, loose, itemBody, item, checked, task, checkbox, ret, l = tokens.length;
    for (i = 0; i < l; i++) {
      if (token = tokens[i], this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[token.type] && (ret = this.options.extensions.renderers[token.type].call({ parser: this }, token), ret !== !1 || !["space", "hr", "heading", "code", "table", "blockquote", "list", "html", "paragraph", "text"].includes(token.type))) {
        out += ret || "";
        continue;
      }
      switch (token.type) {
        case "space":
          continue;
        case "hr": {
          out += this.renderer.hr();
          continue;
        }
        case "heading": {
          out += this.renderer.heading(
            this.parseInline(token.tokens),
            token.depth,
            unescape(this.parseInline(token.tokens, this.textRenderer)),
            this.slugger
          );
          continue;
        }
        case "code": {
          out += this.renderer.code(
            token.text,
            token.lang,
            token.escaped
          );
          continue;
        }
        case "table": {
          for (header = "", cell = "", l2 = token.header.length, j = 0; j < l2; j++)
            cell += this.renderer.tablecell(
              this.parseInline(token.header[j].tokens),
              { header: !0, align: token.align[j] }
            );
          for (header += this.renderer.tablerow(cell), body = "", l2 = token.rows.length, j = 0; j < l2; j++) {
            for (row = token.rows[j], cell = "", l3 = row.length, k = 0; k < l3; k++)
              cell += this.renderer.tablecell(
                this.parseInline(row[k].tokens),
                { header: !1, align: token.align[k] }
              );
            body += this.renderer.tablerow(cell);
          }
          out += this.renderer.table(header, body);
          continue;
        }
        case "blockquote": {
          body = this.parse(token.tokens), out += this.renderer.blockquote(body);
          continue;
        }
        case "list": {
          for (ordered = token.ordered, start = token.start, loose = token.loose, l2 = token.items.length, body = "", j = 0; j < l2; j++)
            item = token.items[j], checked = item.checked, task = item.task, itemBody = "", item.task && (checkbox = this.renderer.checkbox(checked), loose ? item.tokens.length > 0 && item.tokens[0].type === "paragraph" ? (item.tokens[0].text = checkbox + " " + item.tokens[0].text, item.tokens[0].tokens && item.tokens[0].tokens.length > 0 && item.tokens[0].tokens[0].type === "text" && (item.tokens[0].tokens[0].text = checkbox + " " + item.tokens[0].tokens[0].text)) : item.tokens.unshift({
              type: "text",
              text: checkbox
            }) : itemBody += checkbox), itemBody += this.parse(item.tokens, loose), body += this.renderer.listitem(itemBody, task, checked);
          out += this.renderer.list(body, ordered, start);
          continue;
        }
        case "html": {
          out += this.renderer.html(token.text);
          continue;
        }
        case "paragraph": {
          out += this.renderer.paragraph(this.parseInline(token.tokens));
          continue;
        }
        case "text": {
          for (body = token.tokens ? this.parseInline(token.tokens) : token.text; i + 1 < l && tokens[i + 1].type === "text"; )
            token = tokens[++i], body += `
` + (token.tokens ? this.parseInline(token.tokens) : token.text);
          out += top ? this.renderer.paragraph(body) : body;
          continue;
        }
        default: {
          let errMsg = 'Token with "' + token.type + '" type was not found.';
          if (this.options.silent) {
            console.error(errMsg);
            return;
          } else
            throw new Error(errMsg);
        }
      }
    }
    return out;
  }
  parseInline(tokens, renderer) {
    renderer = renderer || this.renderer;
    let out = "", i, token, ret, l = tokens.length;
    for (i = 0; i < l; i++) {
      if (token = tokens[i], this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[token.type] && (ret = this.options.extensions.renderers[token.type].call({ parser: this }, token), ret !== !1 || !["escape", "html", "link", "image", "strong", "em", "codespan", "br", "del", "text"].includes(token.type))) {
        out += ret || "";
        continue;
      }
      switch (token.type) {
        case "escape": {
          out += renderer.text(token.text);
          break;
        }
        case "html": {
          out += renderer.html(token.text);
          break;
        }
        case "link": {
          out += renderer.link(token.href, token.title, this.parseInline(token.tokens, renderer));
          break;
        }
        case "image": {
          out += renderer.image(token.href, token.title, token.text);
          break;
        }
        case "strong": {
          out += renderer.strong(this.parseInline(token.tokens, renderer));
          break;
        }
        case "em": {
          out += renderer.em(this.parseInline(token.tokens, renderer));
          break;
        }
        case "codespan": {
          out += renderer.codespan(token.text);
          break;
        }
        case "br": {
          out += renderer.br();
          break;
        }
        case "del": {
          out += renderer.del(this.parseInline(token.tokens, renderer));
          break;
        }
        case "text": {
          out += renderer.text(token.text);
          break;
        }
        default: {
          let errMsg = 'Token with "' + token.type + '" type was not found.';
          if (this.options.silent) {
            console.error(errMsg);
            return;
          } else
            throw new Error(errMsg);
        }
      }
    }
    return out;
  }
};
function marked(src, opt, callback) {
  if (typeof src > "u" || src === null)
    throw new Error("marked(): input parameter is undefined or null");
  if (typeof src != "string")
    throw new Error("marked(): input parameter is of type " + Object.prototype.toString.call(src) + ", string expected");
  if (typeof opt == "function" && (callback = opt, opt = null), opt = merge({}, marked.defaults, opt || {}), checkSanitizeDeprecation(opt), callback) {
    let highlight = opt.highlight, tokens;
    try {
      tokens = Lexer.lex(src, opt);
    } catch (e) {
      return callback(e);
    }
    let done = function(err) {
      let out;
      if (!err)
        try {
          opt.walkTokens && marked.walkTokens(tokens, opt.walkTokens), out = Parser.parse(tokens, opt);
        } catch (e) {
          err = e;
        }
      return opt.highlight = highlight, err ? callback(err) : callback(null, out);
    };
    if (!highlight || highlight.length < 3 || (delete opt.highlight, !tokens.length))
      return done();
    let pending = 0;
    marked.walkTokens(tokens, function(token) {
      token.type === "code" && (pending++, setTimeout(() => {
        highlight(token.text, token.lang, function(err, code) {
          if (err)
            return done(err);
          code != null && code !== token.text && (token.text = code, token.escaped = !0), pending--, pending === 0 && done();
        });
      }, 0));
    }), pending === 0 && done();
    return;
  }
  function onError(e) {
    if (e.message += `
Please report this to https://github.com/markedjs/marked.`, opt.silent)
      return "<p>An error occurred:</p><pre>" + escape(e.message + "", !0) + "</pre>";
    throw e;
  }
  try {
    let tokens = Lexer.lex(src, opt);
    if (opt.walkTokens) {
      if (opt.async)
        return Promise.all(marked.walkTokens(tokens, opt.walkTokens)).then(() => Parser.parse(tokens, opt)).catch(onError);
      marked.walkTokens(tokens, opt.walkTokens);
    }
    return Parser.parse(tokens, opt);
  } catch (e) {
    onError(e);
  }
}
marked.options = marked.setOptions = function(opt) {
  return merge(marked.defaults, opt), changeDefaults(marked.defaults), marked;
};
marked.getDefaults = getDefaults;
marked.defaults = defaults;
marked.use = function(...args) {
  let extensions = marked.defaults.extensions || { renderers: {}, childTokens: {} };
  args.forEach((pack) => {
    let opts = merge({}, pack);
    if (opts.async = marked.defaults.async || opts.async, pack.extensions && (pack.extensions.forEach((ext) => {
      if (!ext.name)
        throw new Error("extension name required");
      if (ext.renderer) {
        let prevRenderer = extensions.renderers[ext.name];
        prevRenderer ? extensions.renderers[ext.name] = function(...args2) {
          let ret = ext.renderer.apply(this, args2);
          return ret === !1 && (ret = prevRenderer.apply(this, args2)), ret;
        } : extensions.renderers[ext.name] = ext.renderer;
      }
      if (ext.tokenizer) {
        if (!ext.level || ext.level !== "block" && ext.level !== "inline")
          throw new Error("extension level must be 'block' or 'inline'");
        extensions[ext.level] ? extensions[ext.level].unshift(ext.tokenizer) : extensions[ext.level] = [ext.tokenizer], ext.start && (ext.level === "block" ? extensions.startBlock ? extensions.startBlock.push(ext.start) : extensions.startBlock = [ext.start] : ext.level === "inline" && (extensions.startInline ? extensions.startInline.push(ext.start) : extensions.startInline = [ext.start]));
      }
      ext.childTokens && (extensions.childTokens[ext.name] = ext.childTokens);
    }), opts.extensions = extensions), pack.renderer) {
      let renderer = marked.defaults.renderer || new Renderer();
      for (let prop in pack.renderer) {
        let prevRenderer = renderer[prop];
        renderer[prop] = (...args2) => {
          let ret = pack.renderer[prop].apply(renderer, args2);
          return ret === !1 && (ret = prevRenderer.apply(renderer, args2)), ret;
        };
      }
      opts.renderer = renderer;
    }
    if (pack.tokenizer) {
      let tokenizer = marked.defaults.tokenizer || new Tokenizer();
      for (let prop in pack.tokenizer) {
        let prevTokenizer = tokenizer[prop];
        tokenizer[prop] = (...args2) => {
          let ret = pack.tokenizer[prop].apply(tokenizer, args2);
          return ret === !1 && (ret = prevTokenizer.apply(tokenizer, args2)), ret;
        };
      }
      opts.tokenizer = tokenizer;
    }
    if (pack.walkTokens) {
      let walkTokens2 = marked.defaults.walkTokens;
      opts.walkTokens = function(token) {
        let values = [];
        return values.push(pack.walkTokens.call(this, token)), walkTokens2 && (values = values.concat(walkTokens2.call(this, token))), values;
      };
    }
    marked.setOptions(opts);
  });
};
marked.walkTokens = function(tokens, callback) {
  let values = [];
  for (let token of tokens)
    switch (values = values.concat(callback.call(marked, token)), token.type) {
      case "table": {
        for (let cell of token.header)
          values = values.concat(marked.walkTokens(cell.tokens, callback));
        for (let row of token.rows)
          for (let cell of row)
            values = values.concat(marked.walkTokens(cell.tokens, callback));
        break;
      }
      case "list": {
        values = values.concat(marked.walkTokens(token.items, callback));
        break;
      }
      default:
        marked.defaults.extensions && marked.defaults.extensions.childTokens && marked.defaults.extensions.childTokens[token.type] ? marked.defaults.extensions.childTokens[token.type].forEach(function(childTokens) {
          values = values.concat(marked.walkTokens(token[childTokens], callback));
        }) : token.tokens && (values = values.concat(marked.walkTokens(token.tokens, callback)));
    }
  return values;
};
marked.parseInline = function(src, opt) {
  if (typeof src > "u" || src === null)
    throw new Error("marked.parseInline(): input parameter is undefined or null");
  if (typeof src != "string")
    throw new Error("marked.parseInline(): input parameter is of type " + Object.prototype.toString.call(src) + ", string expected");
  opt = merge({}, marked.defaults, opt || {}), checkSanitizeDeprecation(opt);
  try {
    let tokens = Lexer.lexInline(src, opt);
    return opt.walkTokens && marked.walkTokens(tokens, opt.walkTokens), Parser.parseInline(tokens, opt);
  } catch (e) {
    if (e.message += `
Please report this to https://github.com/markedjs/marked.`, opt.silent)
      return "<p>An error occurred:</p><pre>" + escape(e.message + "", !0) + "</pre>";
    throw e;
  }
};
marked.Parser = Parser;
marked.parser = Parser.parse;
marked.Renderer = Renderer;
marked.TextRenderer = TextRenderer;
marked.Lexer = Lexer;
marked.lexer = Lexer.lex;
marked.Tokenizer = Tokenizer;
marked.Slugger = Slugger;
marked.parse = marked;
var options = marked.options, setOptions = marked.setOptions, use = marked.use, walkTokens = marked.walkTokens, parseInline = marked.parseInline;
var parser = Parser.parse, lexer = Lexer.lex;

// app/routes/org/$orgId/datasets/$datasetId/index.tsx
var import_react34 = require("react");
var import_remix_utils2 = require("remix-utils"), import_quill2 = __toESM(require_quill()), import_jsx_runtime54 = require("react/jsx-runtime"), meta13 = () => ({
  charset: "utf-8",
  title: "Dataset Card | PureML",
  viewport: "width=device-width,initial-scale=1"
}), links3 = () => [
  { rel: "stylesheet", href: quill_snow_default }
];
async function loader11({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), readme = await fetchDatasetReadme(
    params.orgId,
    params.datasetId,
    session.get("accessToken")
  ), html = marked(readme.at(-1).content);
  return { readme: readme.at(-1).content, html };
}
async function action10({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), content = (await request.formData()).get("content");
  console.log("fromData", content);
  let res = await writeDatasetReadme(
    params.orgId,
    params.datasetId,
    content,
    session.get("accessToken")
  );
  return null;
}
function DatasetIndex2() {
  let { readme, html } = (0, import_react33.useLoaderData)(), submit = (0, import_react33.useSubmit)(), [edit2, setEdit] = (0, import_react34.useState)(!1), [content, setContent] = (0, import_react34.useState)("");
  return /* @__PURE__ */ (0, import_jsx_runtime54.jsxs)("div", { id: "datasets", children: [
    /* @__PURE__ */ (0, import_jsx_runtime54.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(import_react33.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(TabBar, { intent: "primaryDatasetTab", tab: "datasetCard" }),
    /* @__PURE__ */ (0, import_jsx_runtime54.jsx)("div", { className: "px-12 pt-8 space-y-4", children: edit2 ? /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(import_jsx_runtime54.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime54.jsxs)(import_react33.Form, { method: "post", reloadDocument: !0, className: "flex justify-between", children: [
      /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(
        import_remix_utils2.ClientOnly,
        {
          fallback: /* @__PURE__ */ (0, import_jsx_runtime54.jsx)("div", { className: "w-2/3", style: { width: 500, height: 300 }, children: "Editor Failed to Load!" }),
          children: () => /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(import_quill2.default, { defaultValue: html, setContent })
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(
        Button_default,
        {
          intent: "secondary",
          fullWidth: !1,
          type: "submit",
          icon: "",
          children: "Save"
        }
      )
    ] }) }) : /* @__PURE__ */ (0, import_jsx_runtime54.jsxs)("div", { className: "flex justify-between", children: [
      /* @__PURE__ */ (0, import_jsx_runtime54.jsx)("div", { dangerouslySetInnerHTML: { __html: html } }),
      /* @__PURE__ */ (0, import_jsx_runtime54.jsx)(
        Button_default,
        {
          intent: "secondary",
          fullWidth: !1,
          icon: "",
          onClick: () => {
            setEdit(!0);
          },
          children: "Edit"
        }
      )
    ] }) })
  ] });
}

// app/routes/org/$orgId/models/$modelId.tsx
var modelId_exports = {};
__export(modelId_exports, {
  default: () => ModelIndex,
  meta: () => meta14
});
var import_react35 = require("@remix-run/react");
var import_jsx_runtime55 = require("react/jsx-runtime"), meta14 = () => ({
  charset: "utf-8",
  title: "Model Details | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function ModelIndex() {
  return /* @__PURE__ */ (0, import_jsx_runtime55.jsxs)("div", { id: "models", children: [
    /* @__PURE__ */ (0, import_jsx_runtime55.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(import_react35.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime55.jsxs)("div", { className: "flex flex-col", children: [
      /* @__PURE__ */ (0, import_jsx_runtime55.jsxs)("div", { className: "px-12 sticky top-16 bg-slate-0 w-full z-10", children: [
        /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(Breadcrumbs_default, {}),
        /* @__PURE__ */ (0, import_jsx_runtime55.jsxs)("div", { className: "flex pt-6 py-4", children: [
          /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(Tag_default, { children: "Dummy Tag 1" }),
          /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(Tag_default, { children: "Tag 2" }),
          /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(Tag_default, { children: "Dummy" })
        ] })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime55.jsx)(import_react35.Outlet, {})
    ] })
  ] });
}

// app/routes/org/$orgId/models/$modelId/versions/metrics.tsx
var metrics_exports = {};
__export(metrics_exports, {
  default: () => Metrics2
});
var import_react36 = require("@remix-run/react");
var import_jsx_runtime56 = require("react/jsx-runtime");
function Metrics2() {
  return /* @__PURE__ */ (0, import_jsx_runtime56.jsxs)("div", { className: "", children: [
    /* @__PURE__ */ (0, import_jsx_runtime56.jsx)(TabBar, { intent: "primaryModelTab", tab: "versions" }),
    /* @__PURE__ */ (0, import_jsx_runtime56.jsx)(import_react36.Outlet, {})
  ] });
}

// app/routes/org/$orgId/models/$modelId/versions/metrics/index.tsx
var metrics_exports2 = {};
__export(metrics_exports2, {
  action: () => action11,
  default: () => ModelMetrics,
  loader: () => loader12
});
var import_react37 = require("@remix-run/react"), import_react38 = require("react"), import_lucide_react10 = require("lucide-react");
var import_react39 = require("@remix-run/react");
var import_jsx_runtime57 = require("react/jsx-runtime");
async function loader12({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), versions = await fetchModelVersions(
    params.orgId,
    params.modelId,
    session.get("accessToken")
  );
  return {
    metrics1: await fetchModelMetrics(
      params.orgId,
      params.modelId,
      versions.at(-1),
      session.get("accessToken")
    ),
    versions
  };
}
async function action11({ params, request }) {
  let formData = await request.formData(), version1 = formData.get("version1"), version2 = formData.get("version2"), v = formData.get("v"), session = await getSession(request.headers.get("Cookie"));
  v === "true" && (version1 = version2, version2 = null);
  let metrics1 = version1 !== null ? await fetchModelMetrics(
    params.orgId,
    params.modelId,
    version1,
    session.get("accessToken")
  ) : null, metrics2 = version2 !== null ? await fetchModelMetrics(
    params.orgId,
    params.modelId,
    version2,
    session.get("accessToken")
  ) : null;
  return console.log(version1, version2), console.log(metrics1, metrics2), {
    metrics1,
    metrics2,
    version1,
    version2
  };
}
function ModelMetrics() {
  let data = (0, import_react37.useLoaderData)(), adata = (0, import_react37.useActionData)(), submit = (0, import_react39.useSubmit)(), transition = (0, import_react39.useTransition)(), [metrics, setMetrics] = (0, import_react38.useState)(!0), versions = data.versions, metricsData = JSON.parse(data.metrics1).metrics[1], v1 = versions.at(-1), v2 = "", metricsData2;
  adata && (metricsData = adata.metrics1 !== null ? adata.metrics1 : data.metrics1, v1 = adata.version1 !== null ? adata.version1 : versions.at(-1), v2 = adata.version2 !== null ? adata.version2 : "", metricsData2 = adata.version2 !== null ? adata.metrics2 : []);
  function versionChange(event) {
    submit(event.currentTarget, { replace: !0 });
  }
  return /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("main", { className: "flex px-2", children: [
    /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("div", { className: "w-full", id: "main", children: [
      /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(TabBar, { intent: "modelTab", tab: "metrics" }),
      /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("section", { className: "w-full", children: [
        /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)(
          "div",
          {
            className: "flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4",
            onClick: () => setMetrics(!metrics),
            children: [
              /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("h1", { className: "text-slate-900 font-medium text-sm", children: "Metrics" }),
              metrics ? /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(import_lucide_react10.ChevronUp, { className: "text-slate-400" }) : /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(import_lucide_react10.ChevronDown, { className: "text-slate-400" })
            ]
          }
        ),
        metrics && /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("div", { className: "py-6", children: Object.keys(metricsData).length !== 0 ? /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(import_jsx_runtime57.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("table", { className: " max-w-[1000px] w-full", children: Object.keys(metricsData).map((metric, i) => /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(import_jsx_runtime57.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("tr", { children: [
          /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("th", { className: "text-slate-600 font-medium text-left border p-4", children: metric.charAt(0).toUpperCase() + metric.slice(1) }),
          /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("td", { className: "text-slate-600 font-medium text-left border p-4", children: metricsData[metric].slice(0, 5) }),
          v2 !== "" && Object.keys(metricsData2).length > 0 ? /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("td", { className: "text-slate-600 font-medium text-left border p-4", children: metricsData2[metric].value.slice(0, 5) }) : null
        ] }) })) }) }) : /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("div", { children: "nothing" }) })
      ] }) })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("aside", { className: "bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: [
      /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(Dropdown, { fullWidth: !1, intent: "branch", children: "dev" }),
      /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("div", { className: "py-4", children: [
        "Status: ",
        JSON.stringify(transition.state)
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("ul", { className: "space-y-2", children: versions.map((version) => /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)("li", { className: "flex items-center space-x-2", children: [
        /* @__PURE__ */ (0, import_jsx_runtime57.jsxs)(import_react37.Form, { method: "post", onChange: versionChange, children: [
          /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("input", { hidden: !0, name: "version1", value: v1 }),
          /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("input", { hidden: !0, name: "v", value: version === v1 }),
          /* @__PURE__ */ (0, import_jsx_runtime57.jsx)(
            "input",
            {
              name: "version2",
              value: version,
              type: "checkbox",
              checked: version === v1 || version === v2
            }
          )
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime57.jsx)("p", { children: version })
      ] }, version)) })
    ] })
  ] });
}

// app/routes/org/$orgId/models/$modelId/versions/graphs.tsx
var graphs_exports3 = {};
__export(graphs_exports3, {
  default: () => Graphs2
});
var import_react40 = require("@remix-run/react");
var import_jsx_runtime58 = require("react/jsx-runtime");
function Graphs2() {
  return /* @__PURE__ */ (0, import_jsx_runtime58.jsxs)("div", { className: "", children: [
    /* @__PURE__ */ (0, import_jsx_runtime58.jsx)(TabBar, { intent: "primaryModelTab", tab: "versions" }),
    /* @__PURE__ */ (0, import_jsx_runtime58.jsx)(import_react40.Outlet, {})
  ] });
}

// app/routes/org/$orgId/models/$modelId/versions/graphs/index.tsx
var graphs_exports4 = {};
__export(graphs_exports4, {
  action: () => action12,
  default: () => ModelGraphs,
  loader: () => loader13
});
var import_react41 = require("@remix-run/react"), import_react42 = require("react"), import_lucide_react11 = require("lucide-react");
var import_react43 = require("@remix-run/react");
var import_jsx_runtime59 = require("react/jsx-runtime");
async function loader13({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), versions = await fetchModelVersions(
    params.orgId,
    params.modelId,
    session.get("accessToken")
  );
  return {
    metrics1: await fetchModelMetrics(
      params.orgId,
      params.modelId,
      versions.at(-1),
      session.get("accessToken")
    ),
    versions
  };
}
async function action12({ params, request }) {
  let formData = await request.formData(), version1 = formData.get("version1"), version2 = formData.get("version2"), v = formData.get("v"), session = await getSession(request.headers.get("Cookie"));
  v === "true" && (version1 = version2, version2 = null);
  let metrics1 = version1 !== null ? await fetchModelMetrics(
    params.orgId,
    params.modelId,
    version1,
    session.get("accessToken")
  ) : null, metrics2 = version2 !== null ? await fetchModelMetrics(
    params.orgId,
    params.modelId,
    version2,
    session.get("accessToken")
  ) : null;
  return console.log(version1, version2), console.log(metrics1, metrics2), {
    metrics1,
    metrics2,
    version1,
    version2
  };
}
function ModelGraphs() {
  let data = (0, import_react41.useLoaderData)(), adata = (0, import_react41.useActionData)(), submit = (0, import_react43.useSubmit)(), transition = (0, import_react43.useTransition)(), [graphs, setGraphs] = (0, import_react42.useState)(!0), versions = data.versions, metricsData = JSON.parse(data.metrics1).metrics[1], v1 = versions.at(-1), v2 = "", metricsData2;
  adata && (metricsData = adata.metrics1 !== null ? adata.metrics1 : data.metrics1, v1 = adata.version1 !== null ? adata.version1 : versions.at(-1), v2 = adata.version2 !== null ? adata.version2 : "", metricsData2 = adata.version2 !== null ? adata.metrics2 : []);
  function versionChange(event) {
    submit(event.currentTarget, { replace: !0 });
  }
  return /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("main", { className: "flex", children: [
    /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("div", { className: "w-full", id: "main", children: [
      /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(TabBar, { intent: "modelTab", tab: "graphs" }),
      /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("section", { className: "w-full", children: [
        /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)(
          "div",
          {
            className: "flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4",
            onClick: () => setGraphs(!graphs),
            children: [
              /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("h1", { className: "text-slate-900 font-medium text-sm", children: "Graphs" }),
              graphs ? /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(import_lucide_react11.ChevronUp, { className: "text-slate-400" }) : /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(import_lucide_react11.ChevronDown, { className: "text-slate-400" })
            ]
          }
        ),
        graphs && /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("div", { className: "pt-2", children: [
          /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("div", { className: "py-6", children: /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("div", { className: "px-12 py-6 border-2 border-slate-200 rounded-lg", children: [
            /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("div", { className: "text-slate-900 text-sm font-medium", children: "Confusion Matrix" }),
            /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(
              "img",
              {
                src: "/imgs/ConfusionMatrix.svg",
                alt: "ConfusionMatrix"
              }
            )
          ] }) }),
          /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("div", { className: "py-6", children: /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("div", { className: "px-12 py-6 border-2 border-slate-200 rounded-lg", children: [
            /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("div", { className: "text-slate-900 text-sm font-medium", children: "Classification Report" }),
            /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(
              "img",
              {
                src: "/imgs/ClassificationReport.svg",
                alt: "ClassificationReport"
              }
            )
          ] }) })
        ] })
      ] }) })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("aside", { className: "bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: [
      /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(Dropdown, { fullWidth: !1, intent: "branch", children: "dev" }),
      /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("div", { className: "py-4", children: [
        "Status: ",
        JSON.stringify(transition.state)
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("ul", { className: "space-y-2", children: versions.map((version) => /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)("li", { className: "flex items-center space-x-2", children: [
        /* @__PURE__ */ (0, import_jsx_runtime59.jsxs)(import_react41.Form, { method: "post", onChange: versionChange, children: [
          /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("input", { hidden: !0, name: "version1", value: v1 }),
          /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("input", { hidden: !0, name: "v", value: version === v1 }),
          /* @__PURE__ */ (0, import_jsx_runtime59.jsx)(
            "input",
            {
              name: "version2",
              value: version,
              type: "checkbox",
              checked: version === v1 || version === v2
            }
          )
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime59.jsx)("p", { children: version })
      ] }, version)) })
    ] })
  ] });
}

// app/routes/org/$orgId/models/$modelId/review/$commit.tsx
var commit_exports2 = {};
__export(commit_exports2, {
  default: () => Review2,
  meta: () => meta15
});
var import_react44 = require("@remix-run/react"), import_lucide_react12 = require("lucide-react"), import_react45 = require("react");
var import_jsx_runtime60 = require("react/jsx-runtime"), meta15 = () => ({
  charset: "utf-8",
  title: "Review Model Commit | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function Review2() {
  let path = (0, import_react44.useMatches)()[2].pathname, pathname = decodeURI(path.slice(1)), orgId = pathname.split("/")[1], modelId = pathname.split("/")[3], [metrics, setMetrics] = (0, import_react45.useState)(!0);
  return /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)("div", { id: "reviewCommit", children: [
    /* @__PURE__ */ (0, import_jsx_runtime60.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(import_react44.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(TabBar, { intent: "primaryModelTab", tab: "review" }),
    /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)(
      import_react44.Link,
      {
        to: `/org/${orgId}/models/${modelId}/review`,
        className: "flex font-medium text-sm text-slate-600 items-center px-12 pt-8",
        children: [
          /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(import_lucide_react12.ArrowLeft, {}),
          " View Commit"
        ]
      }
    ),
    /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)("div", { className: "flex px-4", children: [
      /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)("div", { className: "w-full", id: "main", children: [
        /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(TabBar, { intent: "modelReviewTab", tab: "metrics" }),
        /* @__PURE__ */ (0, import_jsx_runtime60.jsx)("div", { className: "px-10 py-8", children: /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)("section", { className: "w-full", children: [
          /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)(
            "div",
            {
              className: "flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4",
              onClick: () => setMetrics(!metrics),
              children: [
                /* @__PURE__ */ (0, import_jsx_runtime60.jsx)("h1", { className: "text-slate-900 font-medium text-sm", children: "Metrics" }),
                metrics ? /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(import_lucide_react12.ChevronUp, { className: "text-slate-400" }) : /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(import_lucide_react12.ChevronDown, { className: "text-slate-400" })
              ]
            }
          ),
          metrics && /* @__PURE__ */ (0, import_jsx_runtime60.jsx)("div", { className: "py-6", children: "Metrics will be shown here" })
        ] }) })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime60.jsxs)("aside", { className: "bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6", children: [
        /* @__PURE__ */ (0, import_jsx_runtime60.jsx)(Dropdown, { fullWidth: !1, intent: "branch", children: "dev" }),
        "List of versions"
      ] })
    ] })
  ] });
}

// app/routes/org/$orgId/models/$modelId/review/index.tsx
var review_exports2 = {};
__export(review_exports2, {
  default: () => Review3,
  meta: () => meta16
});
var import_react46 = require("@remix-run/react");
var import_jsx_runtime61 = require("react/jsx-runtime"), meta16 = () => ({
  charset: "utf-8",
  title: "Model Review | PureML",
  viewport: "width=device-width,initial-scale=1"
});
function Review3() {
  let navigate = (0, import_react46.useNavigate)();
  return /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { id: "modelsReview", children: [
    /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(import_react46.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(TabBar, { intent: "primaryModelTab", tab: "review" }),
    /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "px-12 pt-8 w-2/3", children: [
      /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(AvatarIcon, { children: "A" }),
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(AvatarIcon, { children: "B" }),
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) }),
      /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "pb-6", children: /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "bg-slate-100 rounded-md flex justify-between p-4", children: [
        /* @__PURE__ */ (0, import_jsx_runtime61.jsxs)("div", { className: "flex items-center", children: [
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(AvatarIcon, { children: "C" }),
          /* @__PURE__ */ (0, import_jsx_runtime61.jsx)("div", { className: "text-slate-600 pl-4", children: "Dason J. submitted V1.2 of \u201CHousing_prediction\u201D from Dev" })
        ] }),
        /* @__PURE__ */ (0, import_jsx_runtime61.jsx)(
          Button_default,
          {
            icon: "",
            fullWidth: !1,
            onClick: () => {
              navigate("commit");
            },
            children: "View Commit"
          }
        )
      ] }) })
    ] })
  ] });
}

// app/routes/org/$orgId/models/$modelId/index.tsx
var modelId_exports2 = {};
__export(modelId_exports2, {
  action: () => action13,
  default: () => ModelIndex2,
  links: () => links4,
  loader: () => loader14,
  meta: () => meta17
});
var import_react47 = require("@remix-run/react");
var import_react48 = require("react"), import_remix_utils3 = require("remix-utils"), import_quill3 = __toESM(require_quill());
var import_jsx_runtime62 = require("react/jsx-runtime"), meta17 = () => ({
  charset: "utf-8",
  title: "Model Card | PureML",
  viewport: "width=device-width,initial-scale=1"
}), links4 = () => [
  { rel: "stylesheet", href: quill_snow_default }
];
async function loader14({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), readme = await fetchModelReadme(
    params.orgId,
    params.modelId,
    session.get("accessToken")
  ), html = marked(readme.at(-1).content);
  return { readme: readme.at(-1).content, html };
}
async function action13({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), content = (await request.formData()).get("content");
  console.log("fromData", content);
  let res = await writeModelReadme(
    params.orgId,
    params.modelId,
    content,
    session.get("accessToken")
  );
  return null;
}
function ModelIndex2() {
  let { readme, html } = (0, import_react47.useLoaderData)(), submit = (0, import_react47.useSubmit)(), [edit2, setEdit] = (0, import_react48.useState)(!1), [content, setContent] = (0, import_react48.useState)("");
  return /* @__PURE__ */ (0, import_jsx_runtime62.jsxs)("div", { id: "models", children: [
    /* @__PURE__ */ (0, import_jsx_runtime62.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(import_react47.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(TabBar, { intent: "primaryModelTab", tab: "modelCard" }),
    /* @__PURE__ */ (0, import_jsx_runtime62.jsx)("div", { className: "px-12 pt-8 space-y-4", children: edit2 ? /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(import_jsx_runtime62.Fragment, { children: /* @__PURE__ */ (0, import_jsx_runtime62.jsxs)(import_react47.Form, { method: "post", reloadDocument: !0, className: "flex justify-between", children: [
      /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(
        import_remix_utils3.ClientOnly,
        {
          fallback: /* @__PURE__ */ (0, import_jsx_runtime62.jsx)("div", { className: "w-2/3", style: { width: 500, height: 300 }, children: "Editor Failed to Load!" }),
          children: () => /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(import_quill3.default, { defaultValue: html, setContent })
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(
        Button_default,
        {
          intent: "secondary",
          fullWidth: !1,
          type: "submit",
          icon: "",
          children: "Save"
        }
      )
    ] }) }) : /* @__PURE__ */ (0, import_jsx_runtime62.jsxs)("div", { className: "flex justify-between", children: [
      /* @__PURE__ */ (0, import_jsx_runtime62.jsx)("div", { dangerouslySetInnerHTML: { __html: html } }),
      /* @__PURE__ */ (0, import_jsx_runtime62.jsx)(
        Button_default,
        {
          intent: "secondary",
          fullWidth: !1,
          icon: "",
          onClick: () => {
            setEdit(!0);
          },
          children: "Edit"
        }
      )
    ] }) })
  ] });
}

// app/routes/org/$orgId/datasets/index.tsx
var datasets_exports3 = {};
__export(datasets_exports3, {
  default: () => Index8
});
var import_node13 = require("@remix-run/node");
function Index8() {
  return (0, import_node13.redirect)("/datasets");
}

// app/routes/org/$orgId/models/index.tsx
var models_exports3 = {};
__export(models_exports3, {
  default: () => Index9
});
var import_node14 = require("@remix-run/node");
function Index9() {
  return (0, import_node14.redirect)("/models");
}

// app/routes/org/$orgId/index.tsx
var orgId_exports = {};
__export(orgId_exports, {
  default: () => OrgIndex,
  loader: () => loader15,
  meta: () => meta18
});
var import_react49 = require("@remix-run/react"), import_lucide_react13 = require("lucide-react");
var import_jsx_runtime63 = require("react/jsx-runtime"), meta18 = () => ({
  charset: "utf-8",
  title: "Organization Details | PureML",
  viewport: "width=device-width,initial-scale=1"
});
async function loader15({ params, request }) {
  let session = await getSession(request.headers.get("Cookie")), orgDetails = await fetchOrgDetails(
    params.orgId,
    session.get("accessToken")
  );
  if (!orgDetails)
    return null;
  let modelDetails = await fetchModels(
    orgDetails[0].uuid,
    session.get("accessToken")
  ), datasetDetails = await fetchDatasets(
    orgDetails[0].uuid,
    session.get("accessToken")
  );
  return { orgDetails, modelDetails, datasetDetails };
}
function OrgIndex() {
  let orgData = (0, import_react49.useLoaderData)();
  return orgData === null ? /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(Error2, {}) : /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { className: "flex h-full", children: [
    /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("head", { children: /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(import_react49.Meta, {}) }),
    /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { className: "w-4/5 pt-6 px-12", children: [
      /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)(
        import_react49.Link,
        {
          to: "/models",
          className: "flex text-sm font-medium text-slate-600 pb-6",
          children: [
            /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(import_lucide_react13.ArrowLeft, {}),
            " Go back"
          ]
        }
      ),
      /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { className: "p-6 bg-brand-50 flex items-center rounded-lg", children: [
        /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(AvatarIcon, { intent: "org", children: orgData.orgDetails[0].name.charAt(0) }),
        /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { className: "pl-8", children: [
          /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "text-sm font-medium text-slate-800", children: orgData.orgDetails[0].name }),
          /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "text-xs text-slate-600", children: orgData.orgDetails[0].description })
        ] })
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { children: [
        /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "flex justify-between font-medium text-slate-800 text-base pt-6", children: "Models" }),
        orgData ? /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(import_jsx_runtime63.Fragment, { children: orgData.modelDetails.length !== 0 ? /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72", children: orgData.modelDetails.map((model) => /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(
          import_react49.Link,
          {
            to: `/org/${orgData.orgDetails[0].uuid}/models/${model.name}`,
            children: /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(
              Card,
              {
                intent: "modelCard",
                name: model.name,
                description: `Updated by ${model.updated_by.handle}`,
                tag2: model.created_by.handle
              }
            )
          },
          model.id
        )) }) : /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(EmptyModel, {}) }) : "All public models shown here"
      ] }),
      /* @__PURE__ */ (0, import_jsx_runtime63.jsxs)("div", { children: [
        /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "flex justify-between font-medium text-slate-800 text-base pt-6", children: "Datasets" }),
        orgData ? /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(import_jsx_runtime63.Fragment, { children: orgData.datasetDetails.length !== 0 ? /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72", children: orgData.datasetDetails.map((dataset) => /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(
          import_react49.Link,
          {
            to: `/org/${orgData.orgDetails[0].uuid}/datasets/${dataset.name}`,
            children: /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(
              Card,
              {
                intent: "datasetCard",
                name: dataset.name,
                description: `Updated by ${dataset.updated_by.handle}`,
                tag2: dataset.created_by.handle
              },
              dataset.updated_at
            )
          },
          dataset.id
        )) }) : /* @__PURE__ */ (0, import_jsx_runtime63.jsx)(EmptyDataset, {}) }) : "All public datasets shown here"
      ] })
    ] }),
    /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "w-1/5 border-l-2 border-slate-200", children: /* @__PURE__ */ (0, import_jsx_runtime63.jsx)("div", { className: "text-slate-900 font-medium text-base px-8 py-6", children: "Activity" }) })
  ] });
}

// server-assets-manifest:@remix-run/dev/assets-manifest
var assets_manifest_default = { version: "ff80121c", entry: { module: "/build/entry.client-75SH5MKH.js", imports: ["/build/_shared/chunk-GA5CE6QK.js", "/build/_shared/chunk-6ERGQZP6.js", "/build/_shared/chunk-WNC7G6UM.js", "/build/_shared/chunk-LQHMM3AA.js", "/build/_shared/chunk-CXVWUV7G.js", "/build/_shared/chunk-ADMCF34Z.js"] }, routes: { root: { id: "root", parentId: void 0, path: "", index: void 0, caseSensitive: void 0, module: "/build/root-VCBOLX3L.js", imports: ["/build/_shared/chunk-HK5S4EQI.js", "/build/_shared/chunk-GERQPFJA.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !0, hasErrorBoundary: !0 }, "routes/$username": { id: "routes/$username", parentId: "root", path: ":username", index: void 0, caseSensitive: void 0, module: "/build/routes/$username-U5GBDUT6.js", imports: ["/build/_shared/chunk-MBBO5ZCQ.js", "/build/_shared/chunk-UDABMJW7.js", "/build/_shared/chunk-3J2FPQZ6.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/api/auth.server": { id: "routes/api/auth.server", parentId: "root", path: "api/auth/server", index: void 0, caseSensitive: void 0, module: "/build/routes/api/auth.server-36JBG2O6.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/api/datasets.server": { id: "routes/api/datasets.server", parentId: "root", path: "api/datasets/server", index: void 0, caseSensitive: void 0, module: "/build/routes/api/datasets.server-Q2X4HYSR.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/api/models.server": { id: "routes/api/models.server", parentId: "root", path: "api/models/server", index: void 0, caseSensitive: void 0, module: "/build/routes/api/models.server-HPRJR4RN.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/api/org.server": { id: "routes/api/org.server", parentId: "root", path: "api/org/server", index: void 0, caseSensitive: void 0, module: "/build/routes/api/org.server-OA7OSTXJ.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/auth": { id: "routes/auth", parentId: "root", path: "auth", index: void 0, caseSensitive: void 0, module: "/build/routes/auth-Q7ZZQRVW.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/auth/forgot_password/index": { id: "routes/auth/forgot_password/index", parentId: "routes/auth", path: "forgot_password", index: !0, caseSensitive: void 0, module: "/build/routes/auth/forgot_password/index-AEXDFIBO.js", imports: ["/build/_shared/chunk-CKVKVYRM.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/auth/reset_password/index": { id: "routes/auth/reset_password/index", parentId: "routes/auth", path: "reset_password", index: !0, caseSensitive: void 0, module: "/build/routes/auth/reset_password/index-BHF47AHP.js", imports: ["/build/_shared/chunk-CKVKVYRM.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/auth/signin/index": { id: "routes/auth/signin/index", parentId: "routes/auth", path: "signin", index: !0, caseSensitive: void 0, module: "/build/routes/auth/signin/index-WTWMDKY2.js", imports: ["/build/_shared/chunk-CKVKVYRM.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-HK5S4EQI.js", "/build/_shared/chunk-UDABMJW7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !0 }, "routes/auth/signup/index": { id: "routes/auth/signup/index", parentId: "routes/auth", path: "signup", index: !0, caseSensitive: void 0, module: "/build/routes/auth/signup/index-R2A33ERB.js", imports: ["/build/_shared/chunk-CKVKVYRM.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-HK5S4EQI.js", "/build/_shared/chunk-UDABMJW7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !0 }, "routes/contact": { id: "routes/contact", parentId: "root", path: "contact", index: void 0, caseSensitive: void 0, module: "/build/routes/contact-IVF7S6QJ.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-Q6RXI4JL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/datasets": { id: "routes/datasets", parentId: "root", path: "datasets", index: void 0, caseSensitive: void 0, module: "/build/routes/datasets-R2GPBLJB.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-3J2FPQZ6.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/datasets/EmptyDataset": { id: "routes/datasets/EmptyDataset", parentId: "routes/datasets", path: "EmptyDataset", index: void 0, caseSensitive: void 0, module: "/build/routes/datasets/EmptyDataset-HL55JOQC.js", imports: ["/build/_shared/chunk-TP56N7TT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/datasets/index": { id: "routes/datasets/index", parentId: "routes/datasets", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/datasets/index-RKPNC6YD.js", imports: ["/build/_shared/chunk-D2BMCLLG.js", "/build/_shared/chunk-TP56N7TT.js", "/build/_shared/chunk-GERQPFJA.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/index": { id: "routes/index", parentId: "root", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/index-YNI3BFC7.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/logout": { id: "routes/logout", parentId: "root", path: "logout", index: void 0, caseSensitive: void 0, module: "/build/routes/logout-SD5VWD7L.js", imports: ["/build/_shared/chunk-Q6RXI4JL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/markdown/index": { id: "routes/markdown/index", parentId: "root", path: "markdown", index: !0, caseSensitive: void 0, module: "/build/routes/markdown/index-3FYHBFCI.js", imports: ["/build/_shared/chunk-OZ6DWB6O.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/models": { id: "routes/models", parentId: "root", path: "models", index: void 0, caseSensitive: void 0, module: "/build/routes/models-SPEMWKJW.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-3J2FPQZ6.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/models/EmptyModel": { id: "routes/models/EmptyModel", parentId: "routes/models", path: "EmptyModel", index: void 0, caseSensitive: void 0, module: "/build/routes/models/EmptyModel-WJLIDLDQ.js", imports: ["/build/_shared/chunk-XUGTD4NM.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/models/index": { id: "routes/models/index", parentId: "routes/models", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/models/index-4H3PLBFV.js", imports: ["/build/_shared/chunk-XUGTD4NM.js", "/build/_shared/chunk-D2BMCLLG.js", "/build/_shared/chunk-GERQPFJA.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org": { id: "routes/org", parentId: "root", path: "org", index: void 0, caseSensitive: void 0, module: "/build/routes/org-23BP3LDQ.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-3J2FPQZ6.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId": { id: "routes/org/$orgId/datasets/$datasetId", parentId: "routes/org", path: ":orgId/datasets/:datasetId", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId-6KCWFWJK.js", imports: ["/build/_shared/chunk-JRNARZYO.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/index": { id: "routes/org/$orgId/datasets/$datasetId/index", parentId: "routes/org/$orgId/datasets/$datasetId", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/index-X7ELTTUT.js", imports: ["/build/_shared/chunk-43LPSCJY.js", "/build/_shared/chunk-FNGYAMLS.js", "/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-OZ6DWB6O.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/review/$commit": { id: "routes/org/$orgId/datasets/$datasetId/review/$commit", parentId: "routes/org/$orgId/datasets/$datasetId", path: "review/:commit", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/review/$commit-B63A2ZZY.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/review/index": { id: "routes/org/$orgId/datasets/$datasetId/review/index", parentId: "routes/org/$orgId/datasets/$datasetId", path: "review", index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/review/index-JSESRYZY.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/Pipeline": { id: "routes/org/$orgId/datasets/$datasetId/versions/Pipeline", parentId: "routes/org/$orgId/datasets/$datasetId", path: "versions/Pipeline", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/Pipeline-66J7K3FX.js", imports: ["/build/_shared/chunk-HMIFHG5U.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/datalineage": { id: "routes/org/$orgId/datasets/$datasetId/versions/datalineage", parentId: "routes/org/$orgId/datasets/$datasetId", path: "versions/datalineage", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/datalineage-XMZLXMWP.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/datalineage/index": { id: "routes/org/$orgId/datasets/$datasetId/versions/datalineage/index", parentId: "routes/org/$orgId/datasets/$datasetId/versions/datalineage", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/datalineage/index-WBHQRTML.js", imports: ["/build/_shared/chunk-HMIFHG5U.js", "/build/_shared/chunk-FNGYAMLS.js", "/build/_shared/chunk-4MN4GN7E.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/graphs": { id: "routes/org/$orgId/datasets/$datasetId/versions/graphs", parentId: "routes/org/$orgId/datasets/$datasetId", path: "versions/graphs", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/graphs-KKMACVUC.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/graphs/index": { id: "routes/org/$orgId/datasets/$datasetId/versions/graphs/index", parentId: "routes/org/$orgId/datasets/$datasetId/versions/graphs", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/graphs/index-RG6I3PKJ.js", imports: ["/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-4MN4GN7E.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/$datasetId/versions/nodes-edges": { id: "routes/org/$orgId/datasets/$datasetId/versions/nodes-edges", parentId: "routes/org/$orgId/datasets/$datasetId", path: "versions/nodes-edges", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/$datasetId/versions/nodes-edges-S2LWG3RU.js", imports: void 0, hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/datasets/index": { id: "routes/org/$orgId/datasets/index", parentId: "routes/org", path: ":orgId/datasets", index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/datasets/index-FMCZGPSR.js", imports: ["/build/_shared/chunk-Q6RXI4JL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/index": { id: "routes/org/$orgId/index", parentId: "routes/org", path: ":orgId", index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/index-XSM7IRLS.js", imports: ["/build/_shared/chunk-53UJ32ZV.js", "/build/_shared/chunk-FNGYAMLS.js", "/build/_shared/chunk-XUGTD4NM.js", "/build/_shared/chunk-D2BMCLLG.js", "/build/_shared/chunk-TP56N7TT.js", "/build/_shared/chunk-GERQPFJA.js", "/build/_shared/chunk-MBBO5ZCQ.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId": { id: "routes/org/$orgId/models/$modelId", parentId: "routes/org", path: ":orgId/models/:modelId", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId-A2QN56L2.js", imports: ["/build/_shared/chunk-JRNARZYO.js", "/build/_shared/chunk-GERQPFJA.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/index": { id: "routes/org/$orgId/models/$modelId/index", parentId: "routes/org/$orgId/models/$modelId", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/index-W4WWNI5Q.js", imports: ["/build/_shared/chunk-53UJ32ZV.js", "/build/_shared/chunk-43LPSCJY.js", "/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-OZ6DWB6O.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-5TQH6BYT.js"], hasAction: !0, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/review/$commit": { id: "routes/org/$orgId/models/$modelId/review/$commit", parentId: "routes/org/$orgId/models/$modelId", path: "review/:commit", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/review/$commit-YMSSZ35D.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/review/index": { id: "routes/org/$orgId/models/$modelId/review/index", parentId: "routes/org/$orgId/models/$modelId", path: "review", index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/review/index-XQS6FOGN.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-XZZCBT56.js", "/build/_shared/chunk-5TQH6BYT.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/versions/graphs": { id: "routes/org/$orgId/models/$modelId/versions/graphs", parentId: "routes/org/$orgId/models/$modelId", path: "versions/graphs", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/versions/graphs-C2IX2LOF.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/versions/graphs/index": { id: "routes/org/$orgId/models/$modelId/versions/graphs/index", parentId: "routes/org/$orgId/models/$modelId/versions/graphs", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/versions/graphs/index-SWN4KJMN.js", imports: ["/build/_shared/chunk-53UJ32ZV.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/versions/metrics": { id: "routes/org/$orgId/models/$modelId/versions/metrics", parentId: "routes/org/$orgId/models/$modelId", path: "versions/metrics", index: void 0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/versions/metrics-VFJFFHVT.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/$modelId/versions/metrics/index": { id: "routes/org/$orgId/models/$modelId/versions/metrics/index", parentId: "routes/org/$orgId/models/$modelId/versions/metrics", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/$modelId/versions/metrics/index-XXFSNJXY.js", imports: ["/build/_shared/chunk-53UJ32ZV.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/org/$orgId/models/index": { id: "routes/org/$orgId/models/index", parentId: "routes/org", path: ":orgId/models", index: !0, caseSensitive: void 0, module: "/build/routes/org/$orgId/models/index-DQWFW4YV.js", imports: ["/build/_shared/chunk-Q6RXI4JL.js"], hasAction: !1, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/profile": { id: "routes/profile", parentId: "root", path: "profile", index: void 0, caseSensitive: void 0, module: "/build/routes/profile-TBXCZFYQ.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-Q6RXI4JL.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/settings": { id: "routes/settings", parentId: "root", path: "settings", index: void 0, caseSensitive: void 0, module: "/build/routes/settings-FXYBRJ6Y.js", imports: ["/build/_shared/chunk-VBJDDZTZ.js", "/build/_shared/chunk-3J2FPQZ6.js", "/build/_shared/chunk-I2RIK3CX.js", "/build/_shared/chunk-FNDXDYOH.js", "/build/_shared/chunk-4BJBRKM7.js", "/build/_shared/chunk-Z2L4KIN7.js", "/build/_shared/chunk-XZZCBT56.js"], hasAction: !1, hasLoader: !0, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/settings/account/index": { id: "routes/settings/account/index", parentId: "routes/settings", path: "account", index: !0, caseSensitive: void 0, module: "/build/routes/settings/account/index-LNN62BMJ.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/settings/index": { id: "routes/settings/index", parentId: "routes/settings", path: void 0, index: !0, caseSensitive: void 0, module: "/build/routes/settings/index-YUSMRRVZ.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-4MN4GN7E.js", "/build/_shared/chunk-5TQH6BYT.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 }, "routes/settings/members/index": { id: "routes/settings/members/index", parentId: "routes/settings", path: "members", index: !0, caseSensitive: void 0, module: "/build/routes/settings/members/index-QAGMZGC6.js", imports: ["/build/_shared/chunk-66BWBXYA.js", "/build/_shared/chunk-Q6RXI4JL.js", "/build/_shared/chunk-P6UAVHBL.js"], hasAction: !0, hasLoader: !1, hasCatchBoundary: !1, hasErrorBoundary: !1 } }, cssBundleHref: void 0, url: "/build/manifest-FF80121C.js" };

// server-entry-module:@remix-run/dev/server-build
var assetsBuildDirectory = "public\\build", future = { unstable_cssModules: !1, unstable_cssSideEffectImports: !1, unstable_vanillaExtract: !1, v2_errorBoundary: !1, v2_meta: !1, v2_routeConvention: !1 }, publicPath = "/build/", entry = { module: entry_server_exports }, routes = {
  root: {
    id: "root",
    parentId: void 0,
    path: "",
    index: void 0,
    caseSensitive: void 0,
    module: root_exports
  },
  "routes/api/datasets.server": {
    id: "routes/api/datasets.server",
    parentId: "root",
    path: "api/datasets/server",
    index: void 0,
    caseSensitive: void 0,
    module: datasets_server_exports
  },
  "routes/api/models.server": {
    id: "routes/api/models.server",
    parentId: "root",
    path: "api/models/server",
    index: void 0,
    caseSensitive: void 0,
    module: models_server_exports
  },
  "routes/api/auth.server": {
    id: "routes/api/auth.server",
    parentId: "root",
    path: "api/auth/server",
    index: void 0,
    caseSensitive: void 0,
    module: auth_server_exports
  },
  "routes/api/org.server": {
    id: "routes/api/org.server",
    parentId: "root",
    path: "api/org/server",
    index: void 0,
    caseSensitive: void 0,
    module: org_server_exports
  },
  "routes/markdown/index": {
    id: "routes/markdown/index",
    parentId: "root",
    path: "markdown",
    index: !0,
    caseSensitive: void 0,
    module: markdown_exports
  },
  "routes/$username": {
    id: "routes/$username",
    parentId: "root",
    path: ":username",
    index: void 0,
    caseSensitive: void 0,
    module: username_exports
  },
  "routes/datasets": {
    id: "routes/datasets",
    parentId: "root",
    path: "datasets",
    index: void 0,
    caseSensitive: void 0,
    module: datasets_exports
  },
  "routes/datasets/EmptyDataset": {
    id: "routes/datasets/EmptyDataset",
    parentId: "routes/datasets",
    path: "EmptyDataset",
    index: void 0,
    caseSensitive: void 0,
    module: EmptyDataset_exports
  },
  "routes/datasets/index": {
    id: "routes/datasets/index",
    parentId: "routes/datasets",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: datasets_exports2
  },
  "routes/settings": {
    id: "routes/settings",
    parentId: "root",
    path: "settings",
    index: void 0,
    caseSensitive: void 0,
    module: settings_exports
  },
  "routes/settings/account/index": {
    id: "routes/settings/account/index",
    parentId: "routes/settings",
    path: "account",
    index: !0,
    caseSensitive: void 0,
    module: account_exports
  },
  "routes/settings/members/index": {
    id: "routes/settings/members/index",
    parentId: "routes/settings",
    path: "members",
    index: !0,
    caseSensitive: void 0,
    module: members_exports
  },
  "routes/settings/index": {
    id: "routes/settings/index",
    parentId: "routes/settings",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: settings_exports2
  },
  "routes/contact": {
    id: "routes/contact",
    parentId: "root",
    path: "contact",
    index: void 0,
    caseSensitive: void 0,
    module: contact_exports
  },
  "routes/profile": {
    id: "routes/profile",
    parentId: "root",
    path: "profile",
    index: void 0,
    caseSensitive: void 0,
    module: profile_exports
  },
  "routes/logout": {
    id: "routes/logout",
    parentId: "root",
    path: "logout",
    index: void 0,
    caseSensitive: void 0,
    module: logout_exports
  },
  "routes/models": {
    id: "routes/models",
    parentId: "root",
    path: "models",
    index: void 0,
    caseSensitive: void 0,
    module: models_exports
  },
  "routes/models/EmptyModel": {
    id: "routes/models/EmptyModel",
    parentId: "routes/models",
    path: "EmptyModel",
    index: void 0,
    caseSensitive: void 0,
    module: EmptyModel_exports
  },
  "routes/models/index": {
    id: "routes/models/index",
    parentId: "routes/models",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: models_exports2
  },
  "routes/index": {
    id: "routes/index",
    parentId: "root",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: routes_exports
  },
  "routes/auth": {
    id: "routes/auth",
    parentId: "root",
    path: "auth",
    index: void 0,
    caseSensitive: void 0,
    module: auth_exports
  },
  "routes/auth/forgot_password/index": {
    id: "routes/auth/forgot_password/index",
    parentId: "routes/auth",
    path: "forgot_password",
    index: !0,
    caseSensitive: void 0,
    module: forgot_password_exports
  },
  "routes/auth/reset_password/index": {
    id: "routes/auth/reset_password/index",
    parentId: "routes/auth",
    path: "reset_password",
    index: !0,
    caseSensitive: void 0,
    module: reset_password_exports
  },
  "routes/auth/signin/index": {
    id: "routes/auth/signin/index",
    parentId: "routes/auth",
    path: "signin",
    index: !0,
    caseSensitive: void 0,
    module: signin_exports
  },
  "routes/auth/signup/index": {
    id: "routes/auth/signup/index",
    parentId: "routes/auth",
    path: "signup",
    index: !0,
    caseSensitive: void 0,
    module: signup_exports
  },
  "routes/org": {
    id: "routes/org",
    parentId: "root",
    path: "org",
    index: void 0,
    caseSensitive: void 0,
    module: org_exports
  },
  "routes/org/$orgId/datasets/$datasetId": {
    id: "routes/org/$orgId/datasets/$datasetId",
    parentId: "routes/org",
    path: ":orgId/datasets/:datasetId",
    index: void 0,
    caseSensitive: void 0,
    module: datasetId_exports
  },
  "routes/org/$orgId/datasets/$datasetId/versions/datalineage": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/datalineage",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "versions/datalineage",
    index: void 0,
    caseSensitive: void 0,
    module: datalineage_exports
  },
  "routes/org/$orgId/datasets/$datasetId/versions/datalineage/index": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/datalineage/index",
    parentId: "routes/org/$orgId/datasets/$datasetId/versions/datalineage",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: datalineage_exports2
  },
  "routes/org/$orgId/datasets/$datasetId/versions/nodes-edges": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/nodes-edges",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "versions/nodes-edges",
    index: void 0,
    caseSensitive: void 0,
    module: nodes_edges_exports
  },
  "routes/org/$orgId/datasets/$datasetId/versions/Pipeline": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/Pipeline",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "versions/Pipeline",
    index: void 0,
    caseSensitive: void 0,
    module: Pipeline_exports
  },
  "routes/org/$orgId/datasets/$datasetId/versions/graphs": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/graphs",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "versions/graphs",
    index: void 0,
    caseSensitive: void 0,
    module: graphs_exports
  },
  "routes/org/$orgId/datasets/$datasetId/versions/graphs/index": {
    id: "routes/org/$orgId/datasets/$datasetId/versions/graphs/index",
    parentId: "routes/org/$orgId/datasets/$datasetId/versions/graphs",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: graphs_exports2
  },
  "routes/org/$orgId/datasets/$datasetId/review/$commit": {
    id: "routes/org/$orgId/datasets/$datasetId/review/$commit",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "review/:commit",
    index: void 0,
    caseSensitive: void 0,
    module: commit_exports
  },
  "routes/org/$orgId/datasets/$datasetId/review/index": {
    id: "routes/org/$orgId/datasets/$datasetId/review/index",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: "review",
    index: !0,
    caseSensitive: void 0,
    module: review_exports
  },
  "routes/org/$orgId/datasets/$datasetId/index": {
    id: "routes/org/$orgId/datasets/$datasetId/index",
    parentId: "routes/org/$orgId/datasets/$datasetId",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: datasetId_exports2
  },
  "routes/org/$orgId/models/$modelId": {
    id: "routes/org/$orgId/models/$modelId",
    parentId: "routes/org",
    path: ":orgId/models/:modelId",
    index: void 0,
    caseSensitive: void 0,
    module: modelId_exports
  },
  "routes/org/$orgId/models/$modelId/versions/metrics": {
    id: "routes/org/$orgId/models/$modelId/versions/metrics",
    parentId: "routes/org/$orgId/models/$modelId",
    path: "versions/metrics",
    index: void 0,
    caseSensitive: void 0,
    module: metrics_exports
  },
  "routes/org/$orgId/models/$modelId/versions/metrics/index": {
    id: "routes/org/$orgId/models/$modelId/versions/metrics/index",
    parentId: "routes/org/$orgId/models/$modelId/versions/metrics",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: metrics_exports2
  },
  "routes/org/$orgId/models/$modelId/versions/graphs": {
    id: "routes/org/$orgId/models/$modelId/versions/graphs",
    parentId: "routes/org/$orgId/models/$modelId",
    path: "versions/graphs",
    index: void 0,
    caseSensitive: void 0,
    module: graphs_exports3
  },
  "routes/org/$orgId/models/$modelId/versions/graphs/index": {
    id: "routes/org/$orgId/models/$modelId/versions/graphs/index",
    parentId: "routes/org/$orgId/models/$modelId/versions/graphs",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: graphs_exports4
  },
  "routes/org/$orgId/models/$modelId/review/$commit": {
    id: "routes/org/$orgId/models/$modelId/review/$commit",
    parentId: "routes/org/$orgId/models/$modelId",
    path: "review/:commit",
    index: void 0,
    caseSensitive: void 0,
    module: commit_exports2
  },
  "routes/org/$orgId/models/$modelId/review/index": {
    id: "routes/org/$orgId/models/$modelId/review/index",
    parentId: "routes/org/$orgId/models/$modelId",
    path: "review",
    index: !0,
    caseSensitive: void 0,
    module: review_exports2
  },
  "routes/org/$orgId/models/$modelId/index": {
    id: "routes/org/$orgId/models/$modelId/index",
    parentId: "routes/org/$orgId/models/$modelId",
    path: void 0,
    index: !0,
    caseSensitive: void 0,
    module: modelId_exports2
  },
  "routes/org/$orgId/datasets/index": {
    id: "routes/org/$orgId/datasets/index",
    parentId: "routes/org",
    path: ":orgId/datasets",
    index: !0,
    caseSensitive: void 0,
    module: datasets_exports3
  },
  "routes/org/$orgId/models/index": {
    id: "routes/org/$orgId/models/index",
    parentId: "routes/org",
    path: ":orgId/models",
    index: !0,
    caseSensitive: void 0,
    module: models_exports3
  },
  "routes/org/$orgId/index": {
    id: "routes/org/$orgId/index",
    parentId: "routes/org",
    path: ":orgId",
    index: !0,
    caseSensitive: void 0,
    module: orgId_exports
  }
};
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  assets,
  assetsBuildDirectory,
  entry,
  future,
  publicPath,
  routes
});
