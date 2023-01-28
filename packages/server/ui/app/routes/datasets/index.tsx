import { Link, useLoaderData } from "@remix-run/react";
import Card from "~/components/Card";
import { getSession } from "~/session";
import { fetchDatasets } from "../api/datasets.server";
import EmptyDataset from "./EmptyDataset";

export type dataset = {
  id: string;
  name: string;
  updated_at: string;
  created_by: string;
  uuid: string;
  updated_by: string;
};

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const orgId = session.get("orgId");
  const datasets: dataset[] = await fetchDatasets(
    session.get("orgId"),
    session.get("accessToken")
  );
  return { datasets, orgId };
}

export default function Index() {
  const datasetData = useLoaderData();
  return (
    <div id="datasets">
      <div className="flex justify-between font-medium text-slate-800 text-base pt-6">
        Datasets
      </div>
      {datasetData ? (
        <>
          {datasetData.datasets[0].length !== 0 ? (
            <div
              key="0"
              className="pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72"
            >
              {datasetData.datasets.map((dataset: any) => (
                <Link
                  to={`/org/${datasetData.orgId}/datasets/${dataset.name}`}
                  key={dataset.id}
                >
                  <Card
                    intent="datasetCard"
                    key={dataset.updated_at}
                    name={dataset.name}
                    description={`Updated by ${dataset.updated_by.handle}`}
                    // tag1={dataset.tag1}
                    tag2={dataset.created_by.handle}
                  />
                </Link>
              ))}
            </div>
          ) : (
            <EmptyDataset />
          )}
        </>
      ) : (
        "All public datasets shown here"
      )}
    </div>
  );
}
