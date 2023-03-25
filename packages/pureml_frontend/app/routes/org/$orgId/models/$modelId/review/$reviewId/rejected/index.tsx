import type { MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import {
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
import Loader from "~/components/ui/Loading";
import {
  fetchModelVersions,
  fetchOneModelVersion,
} from "~/routes/api/models.server";
import { getSession } from "~/session";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Review Model Commit | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  // fetching review model version from it's branch
  const reviewData = await fetchOneModelVersion(
    session.get("orgId"),
    params.modelId,
    session.get("fromBranch"),
    session.get("version"),
    session.get("accessToken")
  );
  const versions = await fetchModelVersions(
    session.get("orgId"),
    params.modelId,
    "dev",
    session.get("accessToken")
  );
  // "dev" branch should be replaced by to_branch of submitted review
  return {
    reviewData,
    reviewVersion: session.get("version"),
    reviewBranch: session.get("fromBranch"),
    toBranch: session.get("toBranch"),
    versions: versions,
    params: params,
  };
}

export default function Review() {
  const data = useLoaderData();
  const [metrics, setMetrics] = useState(true);
  const [params, setParams] = useState(true);
  const [ver1, setVer1] = useState("");
  const [metricsData, setMetricsData] = useState({});
  const [paramsData, setParamsData] = useState({});
  const [metricsDict, setMetricsDict] = useState({});
  const [paramsDict, setParamsDict] = useState({});
  const [versionData, setVersionData] = useState(data.versions);
  const reviewData = data.reviewData;
  const reviewVersion = data.reviewVersion;
  const reviewBranch = data.reviewBranch;
  const toBranch = data.toBranch;

  // ##### checking version data #####
  useEffect(() => {
    if (!versionData) return;
    if (!versionData[0]) return;
    setVer1(data.reviewVersion);
  }, [data.reviewVersion, versionData]);

  // ##### fetching & displaying latest version data #####
  useEffect(() => {
    if (!versionData) return;
    if (!reviewData) return;
    if (reviewData === null) {
      setMetricsData({});
      setParamsData({});
      setMetricsDict({});
      setParamsDict({});
      return;
    }

    setMetricsData(JSON.parse(reviewData[0].data));
    const temp = JSON.parse(reviewData[0].data);
    Object.keys(temp).forEach((key) => {
      temp[key] = [temp[key], "-"];
    });
    setMetricsDict(temp);

    setParamsData(JSON.parse(reviewData[1].data));
    const temp2 = JSON.parse(reviewData[1].data);
    Object.keys(temp2).forEach((key) => {
      temp2[key] = [temp2[key], "-"];
    });
    setParamsDict(temp2);
  }, [reviewData, ver1, versionData]);

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
                <div className="px-12 pt-8 pb-8 h-[70vh] overflow-auto">
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
                        {Object.keys(metricsData).length !== 0 &&
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
                        {Object.keys(paramsData).length !== 0 &&
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
                    className="text-slate-500 border-slate-400 bg-transparent rounded"
                    disabled
                  >
                    <option value="value" selected>
                      {toBranch}
                    </option>
                  </select>
                </div>
                {/* incoming branch for review will be displayed here */}
                {versionData ? (
                  <ul className="space-y-2 mt-8 overflow-auto">
                    <li className="flex items-center justify-between py-2">
                      <div className="flex items-center">
                        <input
                          className="checkbox checkbox-primary checkbox-sm"
                          name={"version2"}
                          type="checkbox"
                          defaultChecked
                          disabled
                        />
                        <div className="flex items-center justify-center pl-4 text-slate-600">
                          <AvatarIcon>R</AvatarIcon>
                          <div>
                            <p className="px-4 font-medium">
                              Rejected review {reviewVersion}
                            </p>
                            <p className="px-4">from branch {reviewBranch}</p>
                          </div>
                        </div>
                      </div>
                    </li>
                  </ul>
                ) : (
                  <div className="py-8 mt-8">No version created yet</div>
                )}
                <div className="h-3/5 mt-8">
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
                          {reviewVersion || "Version"}
                        </span>
                      </div>
                      <div className="flex justify-between items-center py-1">
                        <span className="w-1/7 flex items-center">
                          <GitPullRequest className="w-4" />
                        </span>
                        <span className="w-1/2 pl-2">Commit branch</span>
                        <span className="w-1/2 pl-2 text-slate-600 font-medium">
                          {reviewBranch || "Branch"}
                        </span>
                      </div>
                      {/* <div className="flex justify-between items-center py-1">
                        <span className="w-1/7 flex items-center">
                          <User className="w-4" />
                        </span>
                        <span className="w-1/2 pl-2">Pushed by</span>
                        <span className="w-1/2 pl-2 text-slate-600 font-medium">
                          Name
                        </span>
                      </div> */}
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
