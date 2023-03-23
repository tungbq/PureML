import { useLoaderData } from "@remix-run/react";
import { Suspense } from "react";
import Card from "~/components/ui/Card";
import Loader from "~/components/ui/Loading";
import Tag from "~/components/ui/Tag";
import type { model } from "~/lib.type";
import { getSession } from "~/session";
import { fetchModels } from "../api/models.server";
import { fetchOrgDetails } from "../api/org.server";
import EmptyModel from "./EmptyModel";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const orgDetails = await fetchOrgDetails(
    session.get("orgId"),
    session.get("accessToken")
  );
  const orgName = orgDetails[0].name;
  const models: model[] = await fetchModels(
    session.get("orgId"),
    session.get("accessToken")
  );
  return { models, orgName };
}

export default function Index() {
  const modelData = useLoaderData();
  return (
    <Suspense fallback={<Loader />}>
      {modelData ? (
        <>
          {modelData.models[0] ? (
            <div className="px-12 pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72">
              {modelData.models.map((model: any, index: number) => (
                <a
                  href={`/org/${modelData.orgName}/models/${model.name}`}
                  key={index}
                >
                  <Card
                    intent="modelCard"
                    name={`${modelData.orgName}/${model.name}`}
                    description={`Updated by ${model.updated_by.handle || "-"}`}
                    tag2={
                      model.is_public ? (
                        <Tag intent="publicTag">Public</Tag>
                      ) : (
                        <Tag intent="privateTag">Private</Tag>
                      )
                    }
                  />
                </a>
              ))}
            </div>
          ) : (
            <EmptyModel />
          )}
        </>
      ) : (
        "All public models shown here"
      )}
    </Suspense>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}
