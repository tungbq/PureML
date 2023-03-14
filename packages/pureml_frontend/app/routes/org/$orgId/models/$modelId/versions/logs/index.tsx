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
} from "~/routes/api/models.server";
import { getSession } from "~/session";
import Loader from "~/components/ui/Loading";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import * as SelectPrimitive from "@radix-ui/react-select";
import AvatarIcon from "~/components/ui/Avatar";
import Select from "~/components/ui/Select";
import Breadcrumbs from "~/components/Breadcrumbs";

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
  };
}

export async function action({ params, request }: any) {
  const formData = await request.formData();
  let branch = formData.get("branch");
  const session = await getSession(request.headers.get("Cookie"));
  const versions = await fetchModelVersions(
    session.get("orgId"),
    params.modelId,
    branch,
    session.get("accessToken")
  );
  return {
    versions: versions,
  };
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
  const [branch, setBranch] = useState("main");
  const [metricsData, setMetricsData] = useState({});
  const [metricsData2, setMetricsData2] = useState({});
  const [paramsData, setParamsData] = useState({});
  const [paramsData2, setParamsData2] = useState({});
  const [metricsDict, setMetricsDict] = useState({});
  const [paramsDict, setParamsDict] = useState({});
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
    if (versionData.at(-Number(ver1.slice(1))).logs === null) {
      setMetricsData({});
      setParamsData({});
      setMetricsDict({});
      setParamsDict({});
      return;
    }
    setMetricsData(
      JSON.parse(versionData.at(-Number(ver1.slice(1))).logs[0].data)
    );
    const temp = JSON.parse(
      versionData.at(-Number(ver1.slice(1))).logs[0].data
    );
    Object.keys(temp).forEach((key) => {
      temp[key] = [temp[key], "-"];
    });
    setMetricsDict(temp);

    setParamsData(
      JSON.parse(versionData.at(-Number(ver1.slice(1))).logs[1].data)
    );
    const temp2 = JSON.parse(
      versionData.at(-Number(ver1.slice(1))).logs[1].data
    );
    Object.keys(temp2).forEach((key) => {
      temp2[key] = [temp2[key], "-"];
    });
    setParamsDict(temp2);
  }, [ver1, versionData]);

  // ##### comparing versions #####
  useEffect(() => {
    if (!versionData) return;
    if (ver2 === "") {
      setMetricsData2({});
      setParamsData2({});
      return;
    }
    if (versionData.at(Number(ver2.slice(1)) - 1).logs === null) {
      setMetricsData2({});
      setParamsData2({});
      const emptyTemp = metricsDict;
      const emptyTemp2 = paramsDict;
      Object.keys(emptyTemp).forEach((key) => {
        if (emptyTemp[key][0] === "-") {
          delete emptyTemp[key];
        } else {
          emptyTemp[key] = [emptyTemp[key][0], "-"];
        }
      });
      setMetricsDict(emptyTemp);

      Object.keys(emptyTemp2).forEach((key) => {
        if (emptyTemp2[key][0] === "-") {
          delete emptyTemp2[key];
        } else {
          emptyTemp2[key] = [emptyTemp2[key][0], "-"];
        }
      });
      setParamsDict(emptyTemp2);

      return;
    }
    setMetricsData2(
      JSON.parse(versionData.at(Number(ver2.slice(1)) - 1).logs[0].data)
    );
    setParamsData2(
      JSON.parse(versionData.at(Number(ver2.slice(1)) - 1).logs[1].data)
    );

    const original = metricsDict;
    const original2 = paramsDict;

    Object.keys(original).forEach((key) => {
      if (original[key][0] === "-") {
        delete original[key];
      }
    });

    Object.keys(original2).forEach((key) => {
      if (original2[key][0] === "-") {
        delete original2[key];
      }
    });

    const temp = JSON.parse(
      versionData.at(-Number(ver2.slice(1))).logs[0].data
    );
    const temp2 = JSON.parse(
      versionData.at(-Number(ver2.slice(1))).logs[1].data
    );

    Object.keys(temp).forEach((key) => {
      if (key in original) {
        original[key][1] = temp[key];
      } else {
        original[key] = ["-", temp[key]];
      }
    });
    Object.keys(temp2).forEach((key) => {
      if (key in original2) {
        original2[key][1] = temp2[key];
      } else {
        original2[key] = ["-", temp2[key]];
      }
    });
    setMetricsDict(original);
    setParamsDict(original2);
  }, [metricsDict, paramsDict, ver2, versionData]);

  useEffect(() => {
    if (adata === null || adata === undefined) {
      return;
    }
    setVersionData(adata.versions);
  }, [adata]);

  // ##### dropdown branch switch functionality #####
  function branchChange(event: any) {
    setBranch(event.target.value);
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
                {/* ##### metrics secion ##### */}
                <section>
                  <div
                    className="flex items-center justify-between w-full border-b-slate-300 border-b pb-4"
                    onClick={() => setMetrics(!metrics)}
                  >
                    <h1 className="text-slate-800 font-medium text-sm">
                      Metrics
                    </h1>
                    {metrics ? (
                      <ChevronUp className="text-slate-400" />
                    ) : (
                      <ChevronDown className="text-slate-400" />
                    )}
                  </div>
                  {metrics && (
                    <div className="py-6">
                      {(Object.keys(metricsData).length !== 0 ||
                        Object.keys(metricsData2).length !== 0) &&
                      versionData !== null ? (
                        <>
                          <table className="max-w-[1000px] w-full">
                            {Object.keys(metricsDict).length !== 0 && (
                              <>
                                <thead>
                                  <tr>
                                    <th className="text-slate-600 font-medium text-left border p-4">
                                      {" "}
                                    </th>
                                    <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                      {ver1}
                                    </th>
                                    {ver2 !== "" ? (
                                      <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                        {ver2}
                                      </th>
                                    ) : null}
                                  </tr>
                                </thead>
                                {Object.keys(metricsDict).map((metric, i) => (
                                  <tr key={i}>
                                    <th className="text-slate-600 font-medium text-left border p-4">
                                      {metric}
                                    </th>
                                    <td className="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                                      {metricsDict[metric][0].slice(0, 5)}
                                    </td>
                                    {ver2 !== "" && (
                                      <td className="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                                        {metricsDict[metric][1].slice(0, 5)}
                                      </td>
                                    )}
                                  </tr>
                                ))}
                              </>
                            )}
                          </table>
                        </>
                      ) : (
                        <div>No Metrics available</div>
                      )}
                    </div>
                  )}
                </section>
                {/* ##### parameters section ##### */}
                <section className="pt-16">
                  <div
                    className="flex items-center justify-between w-full border-b-slate-300 border-b pb-4"
                    onClick={() => setParams(!params)}
                  >
                    <h1 className="text-slate-800 font-medium text-sm">
                      Parameters
                    </h1>
                    {params ? (
                      <ChevronUp className="text-slate-400" />
                    ) : (
                      <ChevronDown className="text-slate-400" />
                    )}
                  </div>
                  {params && (
                    <div className="py-6">
                      {(Object.keys(paramsData).length !== 0 ||
                        Object.keys(paramsData2).length !== 0) &&
                      versionData !== null ? (
                        <>
                          <table className=" max-w-[1000px] w-full">
                            {Object.keys(paramsDict).length !== 0 && (
                              <>
                                <thead>
                                  <tr>
                                    <th className="text-slate-600 font-medium text-left border p-4">
                                      {" "}
                                    </th>
                                    <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                      {ver1}
                                    </th>
                                    {ver2 !== "" ? (
                                      <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                        {ver2}
                                      </th>
                                    ) : null}
                                  </tr>
                                </thead>
                                {Object.keys(paramsDict).map(
                                  (metric: any, index: number) => (
                                    <tr key={index}>
                                      <th className="text-slate-600 font-medium text-left border p-4">
                                        {metric}
                                      </th>
                                      <td className="text-slate-600 font-medium text-left border p-4">
                                        {paramsDict[metric][0].slice(0, 5)}
                                      </td>
                                      {ver2 !== "" && (
                                        <td className="text-slate-600 font-medium text-left border p-4">
                                          {paramsDict[metric][1].slice(0, 5)}
                                        </td>
                                      )}
                                    </tr>
                                  )
                                )}
                              </>
                            )}
                          </table>
                        </>
                      ) : (
                        <div>No parameters available</div>
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
                <Select name="branch" title={branch}>
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
                      {/* <DropdownMenu.Root>
                        <DropdownMenu.Trigger className="focus:outline-none">
                          <MoreVertical className="text-slate-400" />
                        </DropdownMenu.Trigger>
                        <DropdownMenu.Content sideOffset={7} align="end">
                          <DropdownMenu.Item
                            className="bg-white flex px-3 py-2 text-slate-600 justify-left items-center rounded outline-none focus:outline-none hover:bg-slate-100 cursor-pointer"
                            onClick={() => {
                              navigate(
                                `/org/${data.params.orgId}/datasets/${data.params.datasetId}/review`
                              );
                            }}
                          >
                            Submit for Review
                          </DropdownMenu.Item>
                        </DropdownMenu.Content>
                      </DropdownMenu.Root> */}
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
