/* eslint-disable @typescript-eslint/no-unused-vars */
import type { LinksFunction, MetaFunction } from "@remix-run/node";
import { Form, useLoaderData, useSubmit } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";
import {
  fetchDatasetByName,
  fetchDatasetReadme,
  writeDatasetReadme,
} from "~/routes/api/datasets.server";
import { getSession } from "~/session";
import { marked } from "marked";
import quillCss from "quill/dist/quill.snow.css";
import { Suspense, useState } from "react";
import Button from "~/components/ui/Button";
import { ClientOnly } from "remix-utils";
import Quill from "~/components/quill.client";
import {
  Clock,
  Copy,
  Database,
  Edit,
  FileCheck,
  Globe,
  Loader,
  Pencil,
  Save,
} from "lucide-react";
import { toast } from "react-toastify";
import Breadcrumbs from "~/components/Breadcrumbs";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Dataset Card | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: quillCss },
];

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const readme = await fetchDatasetReadme(
    session.get("orgId"),
    params.datasetId,
    session.get("accessToken")
  );
  const datasetDetails = await fetchDatasetByName(
    session.get("orgId"),
    params.datasetId,
    session.get("accessToken")
  );
  const html = marked(readme.at(-1).content);
  return { readme: readme.at(-1).content, html, datasetDetails };
}

export async function action({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const formData = await request.formData();
  const content = formData.get("content");
  const res = await writeDatasetReadme(
    session.get("orgId"),
    params.datasetId,
    content,
    session.get("accessToken")
  );
  return null;
}

export default function DatasetIndex() {
  const { readme, html, datasetDetails } = useLoaderData();
  const submit = useSubmit();
  const [edit, setEdit] = useState(false);
  const [content, setContent] = useState("");
  function copy() {
    navigator.clipboard.writeText(datasetDetails[0].uuid);
    toast.success("Copied to clipboard!");
  }
  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200">
        <div className="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
          <Breadcrumbs />
          <Tabbar
            intent="primaryDatasetTab"
            tab="datasetCard"
            fullWidth={false}
          />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="flex justify-between h-full">
            <div className="pl-12 pr-8 pt-8 space-y-4 h-[50vh] w-full overflow-auto">
              {edit ? (
                <>
                  <Form method="post" reloadDocument className="h-full">
                    <div className="flex justify-between h-full">
                      <ClientOnly
                        fallback={
                          <div className="w-2/5">Editor Failed to Load!</div>
                        }
                      >
                        {() => (
                          <Quill defaultValue={html} setContent={setContent} />
                        )}
                      </ClientOnly>
                      <Button
                        intent="secondary"
                        fullWidth={false}
                        type="submit"
                        onClick={() => {
                          toast.success("Saved your model card!");
                        }}
                      >
                        <Save className="w-4" />{" "}
                        <div className="pl-2">Save</div>
                      </Button>
                    </div>
                    <input name="content" value={content} className="hidden" />
                  </Form>
                </>
              ) : (
                <>
                  {html ? (
                    <div className="flex justify-between">
                      <div
                        dangerouslySetInnerHTML={{ __html: html }}
                        className="pr-6 text-slate-600 text-sm"
                      />
                      <Button
                        intent="icon"
                        fullWidth={false}
                        onClick={() => {
                          setEdit(true);
                        }}
                      >
                        <Pencil className="w-4" />{" "}
                        <div className="pl-2">Edit</div>
                      </Button>
                    </div>
                  ) : (
                    <div className="flex flex-col justify-center items-center gap-y-1">
                      <div className="text-base font-medium pt-44">
                        Readme not available
                      </div>
                      <div>
                        No details available in dataset card. Please contact
                        author of the dataset.
                      </div>
                    </div>
                  )}
                </>
              )}
            </div>
            <div className="px-6 pt-8 w-1/4 max-w-[400px] bg-slate-50 border-l-2 border-slate-100">
              <div className="text-slate-800 font-medium text-base">
                Dataset details
              </div>
              <div className="pt-2 text-sm text-slate-400">
                <div className="flex justify-between items-center py-1">
                  <span className="w-1/7 flex items-center">
                    <Database className="w-4" />
                  </span>
                  <span className="w-1/2 pl-2">Name</span>
                  <span className="w-1/2 pl-2 text-slate-600 font-medium">
                    {`${datasetDetails[0].name}` || "Model Name"}
                  </span>
                </div>
                <div className="flex justify-between items-center py-1">
                  <span className="w-1/7 flex items-center">
                    <FileCheck className="w-4" />
                  </span>
                  <span className="w-1/2 pl-2">Dataset Id</span>
                  <span className="w-[48%] text-slate-600 font-medium flex justify-between">
                    <div className="w-4/5 pl-2 overflow-hidden truncate">
                      {`${datasetDetails[0].uuid}` || "Model Id"}
                    </div>
                    <Copy
                      className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                      onClick={() => copy()}
                    />
                  </span>
                </div>
                <div className="flex justify-between items-center py-1">
                  <span className="w-1/7 flex items-center">
                    <Clock className="w-4" />
                  </span>
                  <span className="w-1/2 pl-2">Created By</span>
                  <span className="w-1/2 pl-2 text-slate-600 font-medium">
                    {`${datasetDetails[0].created_by.name}` || "Created By"}
                  </span>
                </div>
                <div className="flex justify-between items-center py-1">
                  <span className="w-1/7 flex items-center">
                    <Edit className="w-4" />
                  </span>
                  <span className="w-1/2 pl-2">Last updated by</span>
                  <span className="w-1/2 pl-2 text-slate-600 font-medium">
                    {datasetDetails[0].updated_by.name || "User X"}
                  </span>
                </div>
                <div className="flex justify-between items-center py-1">
                  <span className="w-1/7 flex items-center">
                    <Globe className="w-4" />
                  </span>
                  <span className="w-1/2 pl-2">Public</span>
                  <span className="w-1/2 pl-2 text-slate-600 font-medium">
                    {`${datasetDetails[0].is_public ? "Yes" : "No"}`}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Suspense>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/FunctionalError.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/FunctionalError.gif" alt="Error" width="500" />
    </div>
  );
}
