import type { MetaFunction } from "@remix-run/node";
import { Link, Meta, useMatches } from "@remix-run/react";
import clsx from "clsx";
import { ArrowLeft, ChevronDown, ChevronUp } from "lucide-react";
import { useState } from "react";
import Tabbar from "~/components/Tabbar";
import Dropdown from "~/components/ui/Dropdown";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Review Model Commit | PureML",
  viewport: "width=device-width,initial-scale=1",
});

interface Props {
  tab: string;
  projectId?: string;
}

export default function Review() {
  const matches = useMatches();
  const path = matches[2].pathname;
  const pathname = decodeURI(path.slice(1));
  const orgId = pathname.split("/")[1];
  const modelId = pathname.split("/")[3];
  const [metrics, setMetrics] = useState(true);
  return (
    <div id="reviewCommit">
      <head>
        <Meta />
      </head>
      <Tabbar intent="primaryModelTab" tab="review" />
      <Link
        to={`/org/${orgId}/models/${modelId}/review`}
        className="flex font-medium text-sm text-slate-600 items-center px-12 pt-8"
      >
        <ArrowLeft /> View Commit
      </Link>
      <div className="flex px-4">
        <div className="w-full" id="main">
          <Tabbar intent="modelReviewTab" tab="metrics" />
          <div className="px-10 py-8">
            <section className="w-full">
              <div
                className="flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4"
                onClick={() => setMetrics(!metrics)}
              >
                <h1 className="text-slate-900 font-medium text-sm">Metrics</h1>
                {metrics ? (
                  <ChevronUp className="text-slate-400" />
                ) : (
                  <ChevronDown className="text-slate-400" />
                )}
              </div>
              {metrics && (
                <div className="py-6">
                  {/* {Object.keys(metricsData).length !== 0 ? (
                    <>
                      <table className=" max-w-[1000px] w-full">
                        {Object.keys(metricsData).map((metric, i: number) => (
                          <>
                            <tr>
                              <th className="text-slate-600 font-medium text-left border p-4">
                                {metric.charAt(0).toUpperCase() +
                                  metric.slice(1)}
                              </th>
                              <td className="text-slate-600 font-medium text-left border p-4">
                                {metricsData[metric].slice(0, 5)}
                              </td>
                              {v2 !== "" &&
                              Object.keys(metricsData2).length > 0 ? (
                                <td className="text-slate-600 font-medium text-left border p-4">
                                  {metricsData2[metric].value.slice(0, 5)}
                                </td>
                              ) : null}
                            </tr>
                          </>
                        ))}
                      </table>
                    </>
                  ) : (
                    <div>nothing</div>
                  )} */}
                  Metrics will be shown here
                </div>
              )}
            </section>
          </div>
        </div>
        <aside className="bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6">
          <Dropdown fullWidth={false} intent="branch">
            dev
          </Dropdown>
          {/* <div className="py-4">Status: {JSON.stringify(transition.state)}</div>
          <ul className="space-y-2">
            {versions.map((version: any) => (
              <li className="flex items-center space-x-2" key={version}>
                <Form method="post" onChange={versionChange}>
                  <input hidden name="version1" value={v1} />
                  <input hidden name="v" value={version === v1} />
                  <input
                    // name={version.version === v1 ? 'version1' : 'version2'}
                    name={"version2"}
                    value={version}
                    type="checkbox"
                    checked={version === v1 || version === v2}
                  />
                </Form>
                <p>{version}</p>
              </li>
            ))}
          </ul> */}
          List of versions
        </aside>
      </div>
    </div>
  );
}
