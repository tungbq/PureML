import type { MetaFunction } from "@remix-run/node";
import { Link, useLoaderData, useNavigate } from "@remix-run/react";
import { ArrowLeft } from "lucide-react";
import { Suspense } from "react";
import Card from "~/components/ui/Card";
import AvatarIcon from "~/components/ui/Avatar";
import Loader from "~/components/ui/Loading";
import Tag from "~/components/ui/Tag";
import Error from "~/Error404";
import type { dataset, model, org } from "~/lib.type";
import { fetchDatasets } from "~/routes/api/datasets.server";
import { fetchModels } from "~/routes/api/models.server";
import { fetchOrgDetails } from "~/routes/api/org.server";
import EmptyDataset from "~/routes/datasets/EmptyDataset";
import EmptyModel from "~/routes/models/EmptyModel";
import { getSession } from "~/session";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Organization Details | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const orgDetails: org[] = await fetchOrgDetails(
    session.get("orgId"),
    session.get("accessToken")
  );
  if (!orgDetails) {
    return null;
  }
  const modelDetails: model[] = await fetchModels(
    orgDetails[0].uuid,
    session.get("accessToken")
  );
  const datasetDetails: dataset[] = await fetchDatasets(
    orgDetails[0].uuid,
    session.get("accessToken")
  );
  return { orgDetails, modelDetails, datasetDetails };
}

export default function OrgIndex() {
  const orgData = useLoaderData();
  const navigate = useNavigate();
  if (orgData === null) {
    return <Error />;
  }
  return (
    <Suspense fallback={<Loader />}>
      <div className="flex h-[93%] overflow-auto">
        <div className="w-full pt-8 px-12">
          {/* ###### org details ##### */}

          <div className="p-6 bg-orange-50 flex items-center rounded-lg">
            <AvatarIcon intent="org">
              {orgData.orgDetails[0].name.charAt(0)}
            </AvatarIcon>
            <div className="pl-8">
              <div className="font-medium text-slate-600">
                {orgData.orgDetails[0].name}
              </div>
              <div className="text-sm text-slate-600">
                {orgData.orgDetails[0].description}
              </div>
            </div>
          </div>

          {/* ###### org models displayed here ##### */}
          <div>
            <div className="flex justify-between font-medium text-slate-800 text-base pt-6">
              Models
            </div>
            {orgData ? (
              <>
                {orgData.modelDetails.length !== 0 ? (
                  <div className="py-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72">
                    {orgData.modelDetails.map((model: any, index: number) => (
                      <Link
                        to={`/org/${orgData.orgDetails[0].name}/models/${model.name}`}
                        key={index}
                      >
                        <Card
                          intent="modelCard"
                          name={`${orgData.orgDetails[0].name}/${model.name}`}
                          description={`Updated by ${model.updated_by.handle}`}
                          // tag1={model.tag1}
                          tag2={
                            model.is_public ? (
                              <Tag intent="publicTag">Public</Tag>
                            ) : (
                              <Tag intent="privateTag">Private</Tag>
                            )
                          }
                        />
                      </Link>
                    ))}
                  </div>
                ) : (
                  <EmptyModel />
                )}
              </>
            ) : (
              "All public models shown here"
            )}
          </div>

          {/* ###### org datasets displayed here ##### */}
          <div>
            <div className="flex justify-between font-medium text-slate-800 text-base pt-6">
              Datasets
            </div>
            {orgData ? (
              <>
                {orgData.datasetDetails.length !== 0 ? (
                  <div className="py-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72">
                    {orgData.datasetDetails.map(
                      (dataset: any, index: number) => (
                        <Link
                          to={`/org/${orgData.orgDetails[0].name}/datasets/${dataset.name}`}
                          key={index}
                        >
                          <Card
                            intent="datasetCard"
                            key={dataset.updated_at}
                            name={`${orgData.orgDetails[0].name}/${dataset.name}`}
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
                        </Link>
                      )
                    )}
                  </div>
                ) : (
                  <EmptyDataset />
                )}
              </>
            ) : (
              "All public datasets shown here"
            )}
          </div>
        </div>

        {/* ###### org activity section ##### */}
        {/* <div className="hidden md:block fixed right-0 h-full w-1/5 border-l-2 border-slate-200">
          <div className="text-slate-900 font-medium text-base px-8 py-6">
            Activity
          </div>
        </div> */}
      </div>
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
