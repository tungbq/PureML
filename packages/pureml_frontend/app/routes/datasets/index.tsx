import { useLoaderData } from "@remix-run/react";
import { Suspense } from "react";
import Card from "~/components/ui/Card";
import Loader from "~/components/ui/Loading";
import Tag from "~/components/ui/Tag";
import type { dataset } from "~/lib.type";
import { getSession } from "~/session";
import { fetchDatasets } from "../api/datasets.server";
import { fetchOrgDetails } from "../api/org.server";
import EmptyDataset from "./EmptyDataset";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const orgDetails = await fetchOrgDetails(
    session.get("orgId"),
    session.get("accessToken")
  );
  const orgName = orgDetails[0].name;
  const datasets: dataset[] = await fetchDatasets(
    session.get("orgId"),
    session.get("accessToken")
  );
  return { datasets, orgName };
}

export default function Index() {
  const datasetData = useLoaderData();
  return (
    <Suspense fallback={<Loader />}>
      {datasetData ? (
        <>
          {datasetData.datasets[0] ? (
            <div className="px-12 pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72">
              {datasetData.datasets.map((dataset: any, index: number) => (
                <a
                  href={`/org/${datasetData.orgName}/datasets/${dataset.name}`}
                  key={index}
                >
                  <Card
                    intent="datasetCard"
                    key={dataset.updated_at}
                    name={`${datasetData.orgName}/${dataset.name}`}
                    description={`Updated by ${dataset.updated_by.handle}`}
                    // tag1={dataset.tag1}
                    tag2={
                      dataset.is_public ? (
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
            <EmptyDataset />
          )}
        </>
      ) : (
        "All public datasets shown here"
      )}
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
