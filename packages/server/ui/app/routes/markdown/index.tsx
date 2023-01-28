import { ClientOnly } from "remix-utils";
import Quill from "~/components/quill.client";

import quillCss from "quill/dist/quill.snow.css";
import type { LinksFunction } from "@remix-run/node";

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: quillCss },
];

export default function Index() {
  return (
    <div className="m-2">
      <h1 className="text-2xl font-bold">Remix Quill Example!</h1>
      <ClientOnly fallback={<div style={{ width: 500, height: 300 }}></div>}>
        {() => <Quill defaultValue="Hello <b>Remix!</b>" />}
      </ClientOnly>
    </div>
  );
}
