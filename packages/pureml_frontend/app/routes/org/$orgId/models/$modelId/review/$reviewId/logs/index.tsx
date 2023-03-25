import type { MetaFunction } from "@remix-run/node";
import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
} from "@remix-run/react";
import {
  ChevronDown,
  ChevronUp,
  GitFork,
  GitPullRequest,
  User,
} from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import { toast } from "react-toastify";
import Breadcrumbs from "~/components/Breadcrumbs";
import ReviewTabbar from "~/components/ReviewTabbar";
import Tabbar from "~/components/Tabbar";
import AvatarIcon from "~/components/ui/Avatar";
import Button from "~/components/ui/Button";
import Loader from "~/components/ui/Loading";
import {
  fetchModelVersions,
  fetchOneModelVersion,
  updateModelReview,
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
    session.get("toBranch"),
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

export async function action({ params, request }: any) {
  const formData = await request.formData();
  let option = Object.fromEntries(formData);
  const session = await getSession(request.headers.get("Cookie"));
  if (option._action) {
    const updatedReview = await updateModelReview(
      session.get("orgId"),
      params.modelId,
      params.reviewId,
      option.isAccepted,
      session.get("accessToken")
    );
    return {
      _action: option._action,
      updatedReview: updatedReview,
    };
  } else {
    return null;
  }
}

export default function Review() {
  const data = useLoaderData();
  const adata = useActionData();
  const navigate = useNavigate();
  const [metrics, setMetrics] = useState(true);
  const [params, setParams] = useState(true);
  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [dataVersion, setDataVersion] = useState({});
  const [reviewMetrics, setReviewMetrics] = useState({});
  const [reviewParams, setReviewParams] = useState({});
  const [ver2Metrics, setVer2Metrics] = useState({});
  const [ver2Params, setVer2Params] = useState({});
  const [metricsKeys, setMetricsKeys] = useState<string[]>([]);
  const [paramsKeys, setParamsKeys] = useState<string[]>([]);
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
    setVer2("");
  }, [versionData]);

  useEffect(() => {
    if (!reviewData) return;
    if (reviewData === null) return;
    setReviewMetrics(JSON.parse(reviewData[0].data));
    setMetricsKeys(Object.keys(JSON.parse(reviewData[0].data)));
    setReviewParams(JSON.parse(reviewData[1].data));
    setParamsKeys(Object.keys(JSON.parse(reviewData[1].data)));
  }, [reviewData]);

  useEffect(() => {
    if (!versionData) return;
    if (!versionData[0]) return;
    const tempDict = {};
    versionData.forEach((version: { version: any }) => {
      // @ts-ignore
      tempDict[version.version] = version;
    });
    setDataVersion(tempDict);

    if (ver2 === "") {
      setVer2Metrics({});
      setMetricsKeys(Object.keys(JSON.parse(reviewData[0].data)));
      setVer2Params({});
      setParamsKeys(Object.keys(JSON.parse(reviewData[1].data)));
      return;
    }

    const tt = dataVersion[ver2];
    if (tt) {
      if (tt.logs[0] === null) {
        setVer2Metrics({});
        setVer2Params({});
        return;
      }

      setMetricsKeys(Object.keys(JSON.parse(reviewData[0].data)));
      setParamsKeys(Object.keys(JSON.parse(reviewData[1].data)));
      setVer2Metrics(JSON.parse(tt.logs[0].data));
      setVer2Params(JSON.parse(tt.logs[1].data));
      const metKeys = Object.keys(JSON.parse(tt.logs[0].data));
      const parKeys = Object.keys(JSON.parse(tt.logs[1].data));
      metKeys.forEach((key) => {
        if (!metricsKeys.includes(key)) {
          setMetricsKeys((prev) => [...prev, key]);
        }
      });
      parKeys.forEach((key) => {
        if (!paramsKeys.includes(key)) {
          setParamsKeys((prev) => [...prev, key]);
        }
      });
    }
  }, [ver2]);

  // ##### review action functionality #####
  useEffect(() => {
    if (adata === null || adata === undefined) {
      return;
    }
    if (adata._action === "rejected") {
      if (adata.updatedReview.message === "Model review updated")
        toast.success("Model review version rejected!");
      else toast.error("Something went wrong!");
    } else if (adata._action === "accepted") {
      if (adata.updatedReview.message === "Model review updated")
        toast.success("Model review version accepted!");
      else toast.error("Something went wrong!");
    } else {
      setVersionData(adata.versions);
    }
    navigate(
      `/org/${data.params.orgId}/models/${data.params.modelId}/versions/logs`
    );
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
                        {metricsKeys.length !== 0 && versionData !== null ? (
                          <>
                            <table className="max-w-[1000px] w-full">
                              {metricsKeys.length !== 0 && (
                                <>
                                  <thead>
                                    <tr>
                                      <th className="text-slate-600 font-medium text-left border p-4">
                                        {" "}
                                      </th>
                                      <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                        Submitted commit {ver1}
                                      </th>
                                      {ver2 !== "" ? (
                                        <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                          {ver2}
                                        </th>
                                      ) : null}
                                    </tr>
                                  </thead>
                                  {metricsKeys.map((metric, i) => (
                                    <tr key={i}>
                                      <th className="text-slate-600 font-medium text-left border p-4">
                                        {metric}
                                      </th>
                                      <td className="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                                        {reviewMetrics[metric]
                                          ? reviewMetrics[metric].slice(0, 5)
                                          : "-"}
                                      </td>
                                      {ver2 !== "" && (
                                        <td className="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                                          {ver2Metrics[metric]
                                            ? ver2Metrics[metric].slice(0, 5)
                                            : "-"}
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
                        {paramsKeys.length !== 0 && versionData !== null ? (
                          <>
                            <table className=" max-w-[1000px] w-full">
                              {paramsKeys.length !== 0 && (
                                <>
                                  <thead>
                                    <tr>
                                      <th className="text-slate-600 font-medium text-left border p-4">
                                        {" "}
                                      </th>
                                      <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                        Submitted commit {ver1}
                                      </th>
                                      {ver2 !== "" ? (
                                        <th className="text-slate-600 font-medium text-left border p-4 w-1/5">
                                          {ver2}
                                        </th>
                                      ) : null}
                                    </tr>
                                  </thead>
                                  {paramsKeys.map(
                                    (param: any, index: number) => (
                                      <tr key={index}>
                                        <th className="text-slate-600 font-medium text-left border p-4">
                                          {param}
                                        </th>
                                        <td className="text-slate-600 font-medium text-left border p-4">
                                          {reviewParams[param]
                                            ? reviewParams[param].slice(0, 5)
                                            : "-"}
                                        </td>
                                        {ver2 !== "" && (
                                          <td className="text-slate-600 font-medium text-left border p-4">
                                            {ver2Params[param]
                                              ? ver2Params[param].slice(0, 5)
                                              : "-"}
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
                  <ul className="h-2/5 space-y-2 mt-8 overflow-auto">
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
                          <AvatarIcon>S</AvatarIcon>
                          <div>
                            <p className="px-4 font-medium">
                              Submitted review {reviewVersion}
                            </p>
                            <p className="px-4">from branch {reviewBranch}</p>
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
                            checked={version.version === ver2}
                            onChange={(e) => {
                              setVer2(version.version);
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
                          {`${datasetDetails[0].is_public ? "Yes" : "No"}`}
                        </span>
                      </div> */}
                    </div>
                    <div className="flex gap-x-4 mt-6">
                      <Form method="post">
                        <input name="_action" value="rejected" type="hidden" />
                        <input
                          name="fromBranch"
                          value={data.from_branch}
                          type="hidden"
                        />
                        <input
                          name="fromBranchVersion"
                          value={data.from_branch_version}
                          type="hidden"
                        />
                        <input name="isAccepted" value="false" type="hidden" />
                        <Button intent="secondary" fullWidth={false}>
                          Reject
                        </Button>
                      </Form>
                      <Form method="post">
                        <input name="_action" value="accepted" type="hidden" />
                        <input
                          name="fromBranch"
                          value={data.from_branch}
                          type="hidden"
                        />
                        <input
                          name="fromBranchVersion"
                          value={data.from_branch_version}
                          type="hidden"
                        />
                        <input name="isAccepted" value="true" type="hidden" />
                        <Button intent="primary" fullWidth={false}>
                          Approve
                        </Button>
                      </Form>
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
