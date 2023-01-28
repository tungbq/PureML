import { Form, useActionData, useLoaderData } from "@remix-run/react";
import { useState } from "react";
import { ChevronDown, ChevronUp } from "lucide-react";
import Tabbar from "~/components/Tabbar";
import { useSubmit, useTransition } from "@remix-run/react";

import {
  fetchModelMetrics,
  fetchModelVersions,
} from "~/routes/api/models.server";
import { getSession } from "~/session";
import Dropdown from "~/components/ui/Dropdown";

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const versions = await fetchModelVersions(
    params.orgId,
    params.modelId,
    session.get("accessToken")
  );
  const metrics = await fetchModelMetrics(
    params.orgId,
    params.modelId,
    versions.at(-1),
    session.get("accessToken")
  );
  // return [metrics, projectId, url.href];
  return {
    metrics1: metrics,
    versions: versions,
  };
}

export async function action({ params, request }: any) {
  const formData = await request.formData();
  let version1 = formData.get("version1");
  let version2 = formData.get("version2");
  let v = formData.get("v");
  // console.log('formData', Object.fromEntries(formData));
  // console.log("version=", version1);
  // console.log("v", v);
  const session = await getSession(request.headers.get("Cookie"));
  if (v === "true") {
    version1 = version2;
    version2 = null;
  }
  const metrics1 =
    version1 !== null
      ? await fetchModelMetrics(
          params.orgId,
          params.modelId,
          version1,
          session.get("accessToken")
        )
      : null;
  const metrics2 =
    version2 !== null
      ? await fetchModelMetrics(
          params.orgId,
          params.modelId,
          version2,
          session.get("accessToken")
        )
      : null;
  console.log(version1, version2);
  console.log(metrics1, metrics2);
  return {
    metrics1: metrics1,
    metrics2: metrics2,
    version1: version1,
    version2: version2,
  };
}

export default function ModelGraphs() {
  const data = useLoaderData();
  const adata = useActionData();
  const submit = useSubmit();
  const transition = useTransition();
  const [graphs, setGraphs] = useState(true);
  const versions = data.versions;
  let metricsData = JSON.parse(data.metrics1).metrics[1];
  let v1 = versions.at(-1);
  let v2 = "";
  let metricsData2: any;
  if (adata) {
    metricsData = adata.metrics1 !== null ? adata.metrics1 : data.metrics1;
    v1 = adata.version1 !== null ? adata.version1 : versions.at(-1);
    v2 = adata.version2 !== null ? adata.version2 : "";
    metricsData2 = adata.version2 !== null ? adata.metrics2 : [];
  }
  function versionChange(event: any) {
    submit(event.currentTarget, { replace: true });
  }
  return (
    <main className="flex">
      <div className="w-full" id="main">
        <Tabbar intent="modelTab" tab="graphs" />
        <div className="px-10 py-8">
          <section className="w-full">
            <div
              className="flex items-center justify-between px-4 w-full border-b-slate-300 border-b pb-4"
              onClick={() => setGraphs(!graphs)}
            >
              <h1 className="text-slate-900 font-medium text-sm">Graphs</h1>
              {graphs ? (
                <ChevronUp className="text-slate-400" />
              ) : (
                <ChevronDown className="text-slate-400" />
              )}
            </div>
            {graphs && (
              <div className="pt-2">
                <div className="py-6">
                  <div className="px-12 py-6 border-2 border-slate-200 rounded-lg">
                    <div className="text-slate-900 text-sm font-medium">
                      Confusion Matrix
                    </div>
                    <img
                      src="/imgs/ConfusionMatrix.svg"
                      alt="ConfusionMatrix"
                    />
                  </div>
                </div>
                <div className="py-6">
                  <div className="px-12 py-6 border-2 border-slate-200 rounded-lg">
                    <div className="text-slate-900 text-sm font-medium">
                      Classification Report
                    </div>
                    <img
                      src="/imgs/ClassificationReport.svg"
                      alt="ClassificationReport"
                    />
                  </div>
                </div>
              </div>
            )}
          </section>
        </div>
      </div>
      <aside className="bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6">
        <Dropdown fullWidth={false} intent="branch">
          dev
        </Dropdown>
        <div className="py-4">Status: {JSON.stringify(transition.state)}</div>
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
        </ul>
      </aside>
    </main>
  );
}
