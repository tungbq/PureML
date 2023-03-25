import Tabbar from "~/components/Tabbar";
import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
  useSubmit,
} from "@remix-run/react";
import { getSession } from "~/session";
import {
  fetchDatasetBranch,
  fetchDatasetVersions,
  submitDatasetReview,
} from "~/routes/api/datasets.server";
import Pipeline from "../Pipeline";
import { Suspense, useEffect, useState } from "react";
import type { MetaFunction } from "@remix-run/node";
import { Check, MoreVertical } from "lucide-react";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import * as SelectPrimitive from "@radix-ui/react-select";
import AvatarIcon from "~/components/ui/Avatar";
import { edgesSchema, nodesSchema, versionDataSchema } from "~/lib.schema";
import Select from "~/components/ui/Select";
import Loader from "~/components/ui/Loading";
import Breadcrumbs from "~/components/Breadcrumbs";
import { toast } from "react-toastify";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Dataset Versions | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const allBranch = await fetchDatasetBranch(
    session.get("orgId"),
    params.datasetId,
    session.get("accessToken")
  );
  const versions = await fetchDatasetVersions(
    session.get("orgId"),
    params.datasetId,
    allBranch[0].value,
    session.get("accessToken")
  );
  return {
    versions: versions,
    branches: allBranch,
    params: params,
  };
}

export async function action({ params, request }: any) {
  const formData = await request.formData();
  let option = Object.fromEntries(formData);
  const session = await getSession(request.headers.get("Cookie"));
  if (option._action === "changeBranch") {
    let branch = formData.get("branch");
    const versions = await fetchDatasetVersions(
      session.get("orgId"),
      params.datasetId,
      branch,
      session.get("accessToken")
    );
    return {
      versions: versions,
    };
  } else if (option._action === "submitReview") {
    const submitReview = await submitDatasetReview(
      session.get("orgId"),
      params.datasetId,
      option.fromBranch,
      option.version,
      session.get("accessToken")
    );
    return { _action: "submitReview", submitReview: submitReview };
  } else {
    return null;
  }
}

export default function DatasetVersions() {
  const data = useLoaderData();
  const adata = useActionData();
  const submit = useSubmit();
  const navigate = useNavigate();
  const [versionData, setVersionData] = useState(data.versions);
  const [submitReviewVersion, setSubmitReviewVersion] = useState("");
  const branchData = data.branches;
  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [dataVersion, setDataVersion] = useState({});
  const [node, setNode] = useState(null);
  const [edge, setEdge] = useState(null);
  const [node2, setNode2] = useState(null);
  const [edge2, setEdge2] = useState(null);
  const [branch, setBranch] = useState("main");

  useEffect(() => {
    if (!adata) return;
    adata.versions !== undefined
      ? setVersionData(adata.versions)
      : setVersionData(null);
  }, [adata]);

  // ##### checking version data #####
  useEffect(() => {
    if (!versionData) return;
    if (!versionData[0]) return;
    setVer1(versionData.at(0).version);
    setVer2("");

    const tempDict = {};
    versionData.forEach((version: { version: any }) => {
      // @ts-ignore
      tempDict[version.version] = version;
    });
    setDataVersion(tempDict);
  }, [versionData]);

  useEffect(() => {
    if (!versionData) return;
    if (Object.keys(dataVersion).length === 0) return;
    if (dataVersion[ver1].lineage) {
      if (dataVersion[ver1].lineage.lineage) {
        let lineageData1 = dataVersion[ver1].lineage.lineage;
        lineageData1 = lineageData1.replace(/'/g, "");
        const validJson = versionDataSchema.safeParse(lineageData1);
        const validNodes = nodesSchema.safeParse(
          JSON.parse(lineageData1).nodes
        );
        const validEdges = edgesSchema.safeParse(
          JSON.parse(lineageData1).edges
        );
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
      } else {
        setNode(null);
        setEdge(null);
      }
    }
  }, [ver1, dataVersion]);

  // ##### fetching & comparing latest version data #####
  useEffect(() => {
    if (!versionData) return;
    if (Object.keys(dataVersion).length === 0) return;

    if (ver2 === "") {
      setNode2(null);
      setEdge2(null);
      return;
    }

    if (dataVersion[ver2].lineage) {
      if (dataVersion[ver2].lineage.lineage) {
        let lineageData2 = dataVersion[ver2].lineage.lineage;
        console.log(typeof lineageData2);
        lineageData2 = lineageData2.replace(/'/g, "");
        const validJson = versionDataSchema.safeParse(lineageData2);
        const validNodes = nodesSchema.safeParse(
          JSON.parse(lineageData2).nodes
        );
        const validEdges = edgesSchema.safeParse(
          JSON.parse(lineageData2).edges
        );
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
      } else {
        setNode2(null);
        setEdge2(null);
      }
    }
  }, [ver2, dataVersion]);

  // ##### submit review functionality #####

  useEffect(() => {
    if (adata === null || adata === undefined) {
      return;
    }
    if (adata._action === "submitReview") {
      toast.success(adata.submitReview.message);
      navigate(
        `/org/${data.params.orgId}/datasets/${data.params.datasetId}/review`
      );
    } else {
      setVersionData(adata.versions);
    }
  }, [adata]);

  function branchChange(event: any) {
    setBranch(event.target.value);
    submit(event.currentTarget, { replace: true });
  }
  function submitReview(event: any) {
    submit(event.currentTarget, { replace: true });
  }

  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200">
        <div className="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
          <Breadcrumbs />
          <Tabbar intent="primaryDatasetTab" tab="versions" fullWidth={false} />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="flex justify-between h-full">
            <div className="w-4/5">
              <Tabbar intent="datasetTab" tab="datalineage" />
              <div className="px-12 pt-2 pb-8 h-full overflow-auto">
                <section>
                  {!versionData && (
                    <div className="text-slate-600">
                      No data lineage available
                    </div>
                  )}
                  {versionData && (
                    <div className="flex pt-6 gap-x-8">
                      {!ver2 && !node2 && (
                        <div className="w-full h-[80vh] max-h-full">
                          {node ? (
                            <Pipeline pnode={node} pedge={edge} />
                          ) : (
                            "No datalineage available"
                          )}
                        </div>
                      )}
                      {ver2 && (
                        <>
                          <div className="w-1/2 h-[80vh] max-h-full">
                            <div className="text-sm text-slate-600 font-medium pb-6">
                              {ver1}
                            </div>
                            {node ? (
                              <Pipeline pnode={node} pedge={edge} />
                            ) : (
                              "No datalineage available"
                            )}
                          </div>
                          <div className="w-1/2 h-[80vh] max-h-full">
                            <div className="text-sm text-slate-600 font-medium pb-6">
                              {ver2}
                            </div>
                            {node2 ? (
                              <Pipeline pnode={node2} pedge={edge2} />
                            ) : (
                              "No datalineage available"
                            )}
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
              <Form
                method="post"
                onChange={branchChange}
                className="flex justify-end"
              >
                <input name="_action" value="changeBranch" type="hidden" />
                <Select intent="primary" name="branch" title={branch}>
                  {branchData.map((branch: any, index: number) => (
                    <SelectPrimitive.Item
                      key={`${branch}-${index}`}
                      value={branch.value}
                      className="flex items-center justify-between px-4 py-2 text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
                    >
                      <SelectPrimitive.ItemText className="text-slate-600 text-base font-medium">
                        {branch.label}
                      </SelectPrimitive.ItemText>
                      <SelectPrimitive.ItemIndicator>
                        <Check className="text-slate-400 w-4" />
                      </SelectPrimitive.ItemIndicator>
                    </SelectPrimitive.Item>
                  ))}
                </Select>
              </Form>
              {versionData ? (
                <ul className="h-3/4 space-y-2 mt-8 overflow-auto">
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
                          <AvatarIcon intent="avatar">
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
                      {branch !== "main" && (
                        <Form
                          method="post"
                          onChange={submitReview}
                          className="flex justify-end"
                        >
                          <input
                            name="_action"
                            value="submitReview"
                            type="hidden"
                          />
                          <input
                            name="fromBranch"
                            value={branch}
                            type="hidden"
                          />
                          <input name="toBranch" value="main" type="hidden" />
                          <Select
                            intent="more"
                            name="version"
                            title={submitReviewVersion}
                          >
                            <SelectPrimitive.Item
                              value={version.version}
                              className="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer  hover:bg-slate-100 hover:border-none focus:outline-none"
                            >
                              <SelectPrimitive.ItemText className="text-slate-600 text-base font-medium">
                                Submit For review
                              </SelectPrimitive.ItemText>
                            </SelectPrimitive.Item>
                          </Select>
                        </Form>
                      )}
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="py-8">No version created yet</div>
              )}
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
