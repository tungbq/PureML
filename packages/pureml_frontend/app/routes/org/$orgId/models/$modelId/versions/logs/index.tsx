import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
} from "@remix-run/react";
import { Suspense, useEffect, useState } from "react";
import { Check, ChevronDown, ChevronUp, MoreVertical } from "lucide-react";
import Tabbar from "~/components/Tabbar";
import { useSubmit } from "@remix-run/react";
import {
  fetchModelBranch,
  fetchModelVersions,
  submitModelReview,
} from "~/routes/api/models.server";
import { getSession } from "~/session";
import Loader from "~/components/ui/Loading";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import * as SelectPrimitive from "@radix-ui/react-select";
import AvatarIcon from "~/components/ui/Avatar";
import Select from "~/components/ui/Select";
import Breadcrumbs from "~/components/Breadcrumbs";
import { toast } from "react-toastify";
import ComparisionTable from "~/components/ComparisionTable";

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const allBranch = await fetchModelBranch(
    session.get("orgId"),
    params.modelId,
    session.get("accessToken")
  );
  const versions = await fetchModelVersions(
    session.get("orgId"),
    params.modelId,
    allBranch.at(0).value,
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
    const versions = await fetchModelVersions(
      session.get("orgId"),
      params.modelId,
      branch,
      session.get("accessToken")
    );
    return {
      _action: "changeBranch",
      versions: versions,
    };
  } else if (option._action === "submitReview") {
    const submitReview = await submitModelReview(
      session.get("orgId"),
      params.modelId,
      option.fromBranch,
      option.version,
      session.get("accessToken")
    );
    return { _action: "submitReview", submitReview: submitReview };
  } else {
    return null;
  }
}

function isJson(item: string | object) {
  let value = typeof item !== "string" ? JSON.stringify(item) : item;
  try {
    value = JSON.parse(value);
  } catch (e) {
    return false;
  }

  return typeof value === "object" && value !== null;
}

export default function ModelMetrics() {
  const data = useLoaderData();
  const adata = useActionData();
  const submit = useSubmit();
  const navigate = useNavigate();
  const [metrics, setMetrics] = useState(true);
  const [params, setParams] = useState(true);
  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [submitReviewVersion, setSubmitReviewVersion] = useState("");
  const [branch, setBranch] = useState("main");
  const [dataVersion, setDataVersion] = useState({});
  const [ver1Logs, setVer1Logs] = useState<{ [key: string]: string }>({});
  const [ver2Logs, setVer2Logs] = useState<{ [key: string]: string }>({});
  const [commonMetrics, setCommonMetrics] = useState<string[]>([]);
  const [versionData, setVersionData] = useState(data.versions);
  const branchData = data.branches;

  // ##### checking version data #####
  useEffect(() => {
    if (!versionData) return;
    if (!versionData[0]) return;

    setVer1(versionData.at(0).version);
    setVer2("");
  }, [versionData]);

  // ##### fetching & displaying latest version data #####
  useEffect(() => {
    if (!versionData) return;

    const tempDict = {};
    versionData.forEach((version: { version: any }) => {
      // @ts-ignore
      tempDict[version.version] = version;
    });
    setDataVersion(tempDict);
    const tt = dataVersion[ver1];
    // console.log('tt=', tt);
    if (tt) {
      if (tt.logs === null) {
        setVer1Logs({});
        setCommonMetrics([]);
        return;
      } else {
        const tempDictv1 = {};
        tt.logs.forEach((log: { key: string; data: any }) => {
          if (isJson(log.data)) {
            tempDictv1[log.key] = JSON.parse(log.data);
            if (!commonMetrics.includes(log.key)) {
              setCommonMetrics((prev) => [...prev, log.key]);
            }
          }
        });
        setVer1Logs(tempDictv1);
        // console.log(tt.logs);
      }
    }
  }, [ver1, versionData]);
  // ##### comparing versions #####
  useEffect(() => {
    if (!versionData) return;

    const t1 = dataVersion[ver1];
    // console.log('t1=', t1);
    if (t1) {
      if (t1.logs === null) {
        setCommonMetrics([]);
      } else {
        t1.logs.forEach((log: { key: string; data: any }) => {
          if (isJson(log.data)) {
            if (!commonMetrics.includes(log.key)) {
              setCommonMetrics((prev) => [...prev, log.key]);
            }
          }
        });
      }
    }
    if (ver2 === "") {
      setVer2Logs({});
      // console.log('ver2 is empty');

      return;
    }
    const tt = dataVersion[ver2];
    // console.log('tt2=', tt);
    if (tt) {
      if (tt.logs === null) {
        setVer2Logs({});
        return;
      } else {
        const tempDictv2 = {};
        // console.log('tt.logs=', tt.logs);

        tt.logs.forEach((log: { data: any }) => {
          try {
            tempDictv2[log.key] = JSON.parse(log.data);
            if (!commonMetrics.includes(log.key)) {
              setCommonMetrics((prev) => [...prev, log.key]);
            }
          } catch {
            console.log("Invalid log.key=", log.key);
          }
        });
        // console.log(tempDictv2);

        setVer2Logs(tempDictv2);
      }
    }
  }, [ver2, versionData]);

  // ##### submit review functionality #####
  useEffect(() => {
    if (adata === null || adata === undefined) {
      return;
    }
    if (adata._action === "submitReview") {
      toast.success(adata.submitReview.message);
      navigate(
        `/org/${data.params.orgId}/models/${data.params.modelId}/review`
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
    // console.log(event.target.value);
    submit(event.currentTarget, { replace: true });
  }

  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200">
        <div className="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
          <Breadcrumbs />
          <Tabbar intent="primaryModelTab" tab="versions" fullWidth={false} />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="flex justify-between h-full">
            <div className="w-4/5">
              <Tabbar intent="modelTab" tab="metrics" />
              <div className="px-12 pt-2 pb-8 h-[70vh] overflow-auto">
                {/* ##### version section ##### */}
                {commonMetrics.length !== 0 && versionData !== null ? (
                  <>
                    {commonMetrics.map((key) => {
                      return (
                        <ComparisionTable
                          metric={key}
                          ver1={ver1}
                          ver2={ver2}
                          dataVer1={
                            ver1Logs[key] as unknown as {
                              [key: string]: string;
                            }
                          }
                          dataVer2={
                            ver2Logs[key] as unknown as {
                              [key: string]: string;
                            }
                          }
                        />
                      );
                    })}
                  </>
                ) : null}
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
                      className="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer  hover:bg-slate-100 hover:border-none focus:outline-none"
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
