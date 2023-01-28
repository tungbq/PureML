import { Form, Meta, useLoaderData, useSubmit } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";
import {
  fetchModelReadme,
  writeModelReadme,
} from "~/routes/api/models.server";
import { getSession } from "~/session";
import { marked } from "marked";
import Button from "~/components/ui/Button";
import { useState } from "react";
import { ClientOnly } from "remix-utils";
import Quill from "~/components/quill.client";

import quillCss from "quill/dist/quill.snow.css";
import type { LinksFunction , MetaFunction} from "@remix-run/node";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Model Card | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export type model = {
  id: string;
  name: string;
  status: string;
};

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: quillCss },
];

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const readme = await fetchModelReadme(
    params.orgId,
    params.modelId,
    session.get("accessToken")
  );
  const html = marked(readme.at(-1).content);
  return { readme: readme.at(-1).content, html };
}

export async function action({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const formData = await request.formData();
  const content = formData.get("content");
  console.log("fromData", content);
  const res = await writeModelReadme(
    params.orgId,
    params.modelId,
    content,
    session.get("accessToken")
  );
  return null;
}

export default function ModelIndex() {
  const { readme, html } = useLoaderData();
  const submit = useSubmit();
  const [edit, setEdit] = useState(false);
  const [content, setContent] = useState("");
  // console.log('readme=', readme);
  return (
    <div id="models">
      <head>
        <Meta />
      </head>
      <Tabbar intent="primaryModelTab" tab="modelCard" />
      <div className="px-12 pt-8 space-y-4">
        {edit ? (
          <>
            <Form method="post" reloadDocument className="flex justify-between">
              {/* <input name="content" value={content} /> */}
              <ClientOnly
                fallback={
                  <div className="w-2/3" style={{ width: 500, height: 300 }}>
                    Editor Failed to Load!
                  </div>
                }
              >
                {() => <Quill defaultValue={html} setContent={setContent} />}
              </ClientOnly>
              <Button
                intent="secondary"
                fullWidth={false}
                type="submit"
                icon=""
              >
                Save
              </Button>
            </Form>
          </>
        ) : (
          <div className="flex justify-between">
            <div dangerouslySetInnerHTML={{ __html: html }} />
            <Button
              intent="secondary"
              fullWidth={false}
              icon=""
              onClick={() => {
                setEdit(true);
              }}
            >
              Edit
            </Button>
          </div>
        )}
      </div>
    </div>
  );
}
