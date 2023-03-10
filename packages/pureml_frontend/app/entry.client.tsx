import { RemixBrowser } from "@remix-run/react";
import { startTransition, StrictMode } from "react";
import { hydrateRoot } from "react-dom/client";

function hydrate() {
  startTransition(() => {
    hydrateRoot(
      document,
      <StrictMode>
        <RemixBrowser />
      </StrictMode>
    );
  });
}

if (typeof requestIdleCallback === "function") {
  requestIdleCallback(hydrate);
} else {
  // Safari doesn't support requestIdleCallback
  // https://caniuse.com/requestidlecallback
  setTimeout(hydrate, 1);
}

// import { RemixBrowser } from "@remix-run/react";
// import { StrictMode } from "react";
// import { hydrate } from "react-dom";

// function hydrateFunc() {
//   hydrate(
//     <StrictMode>
//       <RemixBrowser />
//     </StrictMode>,
//     document
//   );
// }

// if (typeof requestIdleCallback === "function") {
//   requestIdleCallback(hydrateFunc);
// } else {
//   // Safari doesn't support requestIdleCallback
//   // https://caniuse.com/requestidlecallback
//   setTimeout(hydrateFunc, 1);
// }
