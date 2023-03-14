import type { MetaFunction } from "@remix-run/node";
import {
  Link,
  useActionData,
  useLoaderData,
  useMatches,
  useNavigate,
  useSubmit,
} from "@remix-run/react";
import {
  ArrowLeft,
  ChevronDown,
  ChevronUp,
  GitFork,
  GitPullRequest,
  User,
} from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import Breadcrumbs from "~/components/Breadcrumbs";
import ReviewTabbar from "~/components/ReviewTabbar";
import Tabbar from "~/components/Tabbar";
import AvatarIcon from "~/components/ui/Avatar";
import Button from "~/components/ui/Button";
import Dropdown from "~/components/ui/Dropdown";
import Loader from "~/components/ui/Loading";
import { fetchModelVersions } from "~/routes/api/models.server";
import { getSession } from "~/session";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Review Model Commit | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const versions = await fetchModelVersions(
    session.get("orgId"),
    params.modelId,
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
  const navigate = useNavigate();
  const [metrics, setMetrics] = useState(true);
  const [params, setParams] = useState(true);
  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [metricsData, setMetricsData] = useState({});
  const [metricsData2, setMetricsData2] = useState({});
  const [paramsData, setParamsData] = useState({});
  const [paramsData2, setParamsData2] = useState({});
  const [metricsDict, setMetricsDict] = useState({});
  const [paramsDict, setParamsDict] = useState({});
  const [versionData, setVersionData] = useState(data.versions);

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
  return (
    <Suspense fallback={<Loader />}>
      <div className="w-full max-w-screen-2xl">
        <div className="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200">
          <div className="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
            <Breadcrumbs />
            <Tabbar intent="primaryModelTab" tab="review" fullWidth={false} />
          </div>
        </div>
        <div className="flex justify-center w-full">
          <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
            <div className="flex justify-between h-full">
              <div className="w-4/5">
                <ReviewTabbar intent="modelReviewMetricsTab" tab="metrics" />
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
                <div className="flex justify-end">
                  <select
                    name="branch"
                    className="text-slate-500 border-slate-500 cursor-not-allowed bg-transparent rounded"
                  >
                    <option value="value" selected>
                      dev
                    </option>
                  </select>
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
                              version.version === ver1 ||
                              version.version === ver2
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
      </div>
    </Suspense>
  );
}
