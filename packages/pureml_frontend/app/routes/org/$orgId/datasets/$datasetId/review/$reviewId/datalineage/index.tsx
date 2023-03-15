import type { MetaFunction } from "@remix-run/node";
import {
  Form,
  useActionData,
  useLoaderData,
  useSubmit,
} from "@remix-run/react";
import {
  GitFork,
  GitPullRequest,
  User,
} from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import Breadcrumbs from "~/components/Breadcrumbs";
import ReviewTabbar from "~/components/ReviewTabbar";
import Tabbar from "~/components/Tabbar";
import Loader from "~/components/ui/Loading";
import { edgesSchema, nodesSchema, versionDataSchema } from "~/lib.schema";
import Pipeline from "../../../versions/Pipeline";
import AvatarIcon from "~/components/ui/Avatar";
import { getSession } from "~/session";
import {
  fetchDatasetVersions,
} from "~/routes/api/datasets.server";
import Button from "~/components/ui/Button";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Review Dataset Commit | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const versions = await fetchDatasetVersions(
    session.get("orgId"),
    params.datasetId,
    "dev",
    session.get("accessToken")
  );
  // "dev" branch should be replaced by to_branch of submitted review
  return {
    versions: versions,
    params: params,
  };
}

// export async function action({ params, request }: any) {

// }

export default function Review() {
  const data = useLoaderData();
  const adata = useActionData();
  const submit = useSubmit();
  // function branchChange(event: any) {
  //   setBranch(event.target.value);
  //   submit(event.currentTarget, { replace: true });
  // }
  const [versionData, setVersionData] = useState(data.versions);
  useEffect(() => {
    if (!adata) return;
    adata.versions !== undefined
      ? setVersionData(adata.versions)
      : setVersionData(null);
  }, [adata]);

  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [node, setNode] = useState(null);
  const [edge, setEdge] = useState(null);
  const [node2, setNode2] = useState(null);
  const [edge2, setEdge2] = useState(null);

  // ##### checking version data #####
  useEffect(() => {
    if (!versionData) return;
    if (!versionData[0]) return;
    if (!versionData[0].lineage.lineage) return;
    setVer1(versionData.at(0).version);
    setVer2("");
  }, [versionData]);

  // ##### fetching & comparing latest version data #####
  useEffect(() => {
    if (!versionData) return;
    if (ver2 === "") {
      setNode2(null);
      setEdge2(null);
      return;
    }

    let lineageData2 = versionData.at(-Number(ver2.slice(1))).lineage.lineage;
    lineageData2 = lineageData2.replace(/'/g, "");
    const validJson = versionDataSchema.safeParse(lineageData2);
    const validNodes = nodesSchema.safeParse(JSON.parse(lineageData2).nodes);
    const validEdges = edgesSchema.safeParse(JSON.parse(lineageData2).edges);
    if (validJson.success && validNodes.success && validEdges.success) {
      const n = JSON.parse(lineageData2).nodes;
      n.forEach((e: any) => {
        e.data = { label: e.text };
      });
      setNode2(null);
      setTimeout(() => {
        setNode2(n);
      }, 10);
      const ed = JSON.parse(lineageData2).edges;
      ed.forEach((e: any) => {
        e.source = e.from;
        e.target = e.to;
      });
      setEdge2(null);
      setTimeout(() => {
        setEdge2(ed);
      }, 10);
    } else {
      setNode2(null);
      setEdge2(null);
    }
  }, [ver2, versionData]);

  useEffect(() => {
    if (!versionData) return;
    let lineageData1 = versionData.at(-Number(ver1.slice(1))).lineage.lineage;
    lineageData1 = lineageData1.replace(/'/g, "");
    const validJson = versionDataSchema.safeParse(lineageData1);
    const validNodes = nodesSchema.safeParse(JSON.parse(lineageData1).nodes);
    const validEdges = edgesSchema.safeParse(JSON.parse(lineageData1).edges);
    if (validJson.success && validNodes.success && validEdges.success) {
      const n = JSON.parse(lineageData1).nodes;
      n.forEach((e: any) => {
        e.data = { label: e.text };
      });
      setNode(null);
      setTimeout(() => {
        setNode(n);
      }, 10);
      const ed = JSON.parse(lineageData1).edges;
      ed.forEach((e: any) => {
        e.source = e.from;
        e.target = e.to;
      });
      setEdge(null);
      setTimeout(() => {
        setEdge(ed);
      }, 10);
    } else {
      setNode(null);
      setEdge(null);
    }
  }, [ver1, versionData]);

  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200">
        <div className="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
          <Breadcrumbs />
          <Tabbar intent="primaryDatasetTab" tab="review" fullWidth={false} />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="flex justify-between h-full">
            <div className="w-4/5">
              <ReviewTabbar
                intent="datasetReviewLineageTab"
                tab="datalineage"
              />
              <div className="px-12 pt-2 pb-8 h-[100vh] overflow-auto">
                <section>
                  {!versionData && (
                    <div className="text-slate-600 pt-6">
                      No data lineage available
                    </div>
                  )}
                  {versionData && (
                    <div className="flex pt-6 gap-x-8">
                      {!node2 && (
                        <div className="w-full h-screen max-h-[600px]">
                          {node && <Pipeline pnode={node} pedge={edge} />}
                        </div>
                      )}
                      {node2 && (
                        <>
                          <div className="w-1/2 h-screen max-h-[600px]">
                            <div className="text-sm text-slate-600 font-medium pb-6">
                              {ver1}
                            </div>
                            {node && <Pipeline pnode={node} pedge={edge} />}
                          </div>
                          <div className="w-1/2 h-screen max-h-[600px]">
                            <div className="text-sm text-slate-600 font-medium pb-6">
                              {ver2}
                            </div>
                            {node2 && <Pipeline pnode={node2} pedge={edge2} />}
                          </div>
                        </>
                      )}
                    </div>
                  )}
                </section>
              </div>
            </div>
            {/* ##### versions list right sidebar ##### */}
            <aside className="bg-slate-50 border-l-2 border-slate-100 h-full w-1/4 max-w-[400px] py-8 px-12 z-10">
              <div className="flex justify-end">
                <select
                  name="branch"
                  className="text-slate-500 border-slate-500 cursor-not-allowed bg-transparent rounded"
                ></select>
              </div>
              {/* incoming branch for review will be displayed here */}
              {versionData ? (
                <ul className="h-2/5 space-y-2 mt-8 overflow-auto">
                  <li className="flex items-center justify-between py-2">
                    <div className="flex items-center">
                      <input
                        className="checkbox checkbox-primary checkbox-sm"
                        name={"version2"}
                        type="checkbox"
                        defaultChecked
                      />
                      <div className="flex items-center justify-center pl-4 text-slate-600">
                        <AvatarIcon>S</AvatarIcon>
                        <div>
                          <p className="px-4 font-medium">Submitted review</p>
                          <p className="px-4">Created by Kessshhhaaa</p>
                        </div>
                      </div>
                    </div>
                  </li>
                  {versionData.map((version: any) => (
                    <li
                      className="flex items-center justify-between py-2"
                      key={version.version}
                    >
                      <div className="flex items-center">
                        <input
                          className="checkbox checkbox-primary checkbox-sm"
                          name={"version2"}
                          value={version.version}
                          type="checkbox"
                          checked={
                            version.version === ver1 || version.version === ver2
                          }
                          onChange={(e) => {
                            if (e.target.checked) {
                              setVer2(version.version);
                            } else if (
                              ver1 === version.version &&
                              ver2 === ""
                            ) {
                              new Error(
                                "You can't uncheck the present version"
                              );
                            } else if (ver1 === version.version) {
                              setVer1(ver2);
                              setVer2("");
                            } else {
                              setVer2("");
                            }
                          }}
                        />
                        <div className="flex items-center justify-center pl-4 text-slate-600">
                          <AvatarIcon>
                            {version.created_by.name.charAt(0).toUpperCase()}
                          </AvatarIcon>
                          <div>
                            <p className="px-4 font-medium">
                              {version.version}
                            </p>
                            <p className="px-4">
                              Created by {version.created_by.name}
                            </p>
                          </div>
                        </div>
                      </div>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="h-2/5 py-8 mt-8">No version created yet</div>
              )}
              <div className="h-2/5 mt-8">
                <div className="text-base text-slate-800 font-medium pb-4">
                  Details
                </div>
                <div className="h-4/5 overflow-auto">
                  <div className="text-slate-400">
                    <div className="flex justify-between items-center py-1">
                      <span className="w-1/7 flex items-center">
                        <GitFork className="w-4" />
                      </span>
                      <span className="w-1/2 pl-2">Commit version</span>
                      <span className="w-1/2 pl-2 text-slate-600 font-medium">
                        {/* {`${datasetDetails[0].created_by.name}` || "Created By"} */}
                        vx
                      </span>
                    </div>
                    <div className="flex justify-between items-center py-1">
                      <span className="w-1/7 flex items-center">
                        <GitPullRequest className="w-4" />
                      </span>
                      <span className="w-1/2 pl-2">Commit branch</span>
                      <span className="w-1/2 pl-2 text-slate-600 font-medium">
                        {/* {datasetDetails[0].updated_by.name || "User X"} */}
                        dev
                      </span>
                    </div>
                    <div className="flex justify-between items-center py-1">
                      <span className="w-1/7 flex items-center">
                        <User className="w-4" />
                      </span>
                      <span className="w-1/2 pl-2">Pushed by</span>
                      <span className="w-1/2 pl-2 text-slate-600 font-medium">
                        {/* {`${datasetDetails[0].is_public ? "Yes" : "No"}`} */}
                        Name
                      </span>
                    </div>
                  </div>
                  <div className="flex gap-x-4 mt-6">
                    <Button intent="secondary" fullWidth={false}>
                      Reject
                    </Button>
                    <Button intent="primary" fullWidth={false}>
                      Approve
                    </Button>
                  </div>
                </div>
              </div>
            </aside>
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
