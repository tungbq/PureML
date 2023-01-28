import Tabbar from "~/components/Tabbar";
import { Link, Meta, useLoaderData, useMatches } from "@remix-run/react";
import { getSession } from "~/session";
import { fetchDatasetVersions } from "~/routes/api/datasets.server";
import Pipeline from "../Pipeline";
import { useEffect, useState } from "react";
import type { MetaFunction } from "@remix-run/node";
import { ChevronDown, ChevronUp } from "lucide-react";
import clsx from "clsx";
export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Dataset Versions | PureML",
  viewport: "width=device-width,initial-scale=1",
});

function tabCss(currentPage: boolean) {
  return clsx(
    currentPage ? "text-white" : "text-slate-600",
    "flex justify-center items-center"
  );
}

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const versions = await fetchDatasetVersions(
    params.orgId,
    params.datasetId,
    session.get("accessToken")
  );
  // console.log('ver=', versions);
  // console.log('lin=', versions[0].lineage.lineage);
  return {
    versions: versions,
  };
}

interface Props {
  tab: string;
}

export default function DatasetVersions({ tab }: Props) {
  const data = useLoaderData();
  const matches = useMatches();
  const path = matches[2].pathname;
  const pathname = decodeURI(path.slice(1));
  const orgId = pathname.split("/")[1];
  const datasetId = pathname.split("/")[3];
  const [lineage, setLineage] = useState(true);
  const [graphs, setGraphs] = useState(true);
  const versionData = data.versions;
  // console.log('v=', versionData);
  const [ver1, setVer1] = useState("");
  const [ver2, setVer2] = useState("");
  const [node, setNode] = useState(null);
  const [edge, setEdge] = useState(null);
  const [node2, setNode2] = useState(null);
  const [edge2, setEdge2] = useState(null);
  useEffect(() => {
    // console.log("v0");
    if (!versionData[0]) return;
    // console.log("l0");
    if (!versionData[0].lineage.lineage) return;
    // console.log(versionData.at(-1).version);
    setVer1(versionData.at(-1).version);
    setVer2("");
    // console.log("passed 1");
  }, [versionData]);

  useEffect(() => {
    if (!versionData) return;
    if (ver2 === "") {
      setNode2(null);
      setEdge2(null);
      return;
    }
    // console.log("slice=", versionData.at(Number(ver2.slice(1)) - 1));
    const n = JSON.parse(
      versionData.at(Number(ver2.slice(1)) - 1).lineage.lineage
    ).nodes;
    n.forEach((e: any) => {
      e.data = { label: e.text };
    });
    setNode2(null);
    setTimeout(() => {
      setNode2(n);
    }, 10);
    const ed = JSON.parse(
      versionData.at(Number(ver2.slice(1)) - 1).lineage.lineage
    ).edges;
    ed.forEach((e: any) => {
      e.source = e.from;
      e.target = e.to;
    });
    setEdge2(null);
    setTimeout(() => {
      setEdge2(ed);
    }, 10);
    // console.log("passed 2");
  }, [ver2, versionData]);

  useEffect(() => {
    if (!versionData) return;
    // console.log(versionData.at(Number(ver1.slice(1)) - 1).lineage.lineage);
    const n = JSON.parse(
      versionData.at(Number(ver1.slice(1)) - 1).lineage.lineage
    ).nodes;
    n.forEach((e: any) => {
      e.data = { label: e.text };
    });
    setNode(null);
    setTimeout(() => {
      setNode(n);
    }, 10);
    const ed = JSON.parse(
      versionData.at(Number(ver1.slice(1)) - 1).lineage.lineage
    ).edges;
    ed.forEach((e: any) => {
      e.source = e.from;
      e.target = e.to;
    });
    setEdge(null);
    setTimeout(() => {
      setEdge(ed);
    }, 10);
  }, [ver1, versionData]);

  return (
    <main className="flex px-2">
      <div className="w-full" id="main">
        <Tabbar intent="datasetTab" tab="datalineage" />
        <div className="px-10 py-8">
          <section className="w-full">
            <div
              className="flex items-center justify-between w-full border-b-slate-300 border-b"
              onClick={() => setLineage(!lineage)}
            >
              <h1 className="text-slate-900 font-medium text-sm rounded-lg">
                Data Lineage
              </h1>
              {lineage ? (
                <ChevronUp className="text-slate-400" />
              ) : (
                <ChevronDown className="text-slate-400" />
              )}
            </div>
            {lineage && (
              <div className="flex">
                <div className="w-full h-screen max-h-[600px]">
                  {node && <Pipeline pnode={node} pedge={edge} />}
                </div>
                <div className="w-full h-screen max-h-[600px]">
                  {node2 && <Pipeline pnode={node2} pedge={edge2} />}
                </div>
              </div>
            )}
            {!lineage && <div>Add Data lineage</div>}
          </section>
          {/* <div id="graphs" className="pt-16">
            <section className="w-full pt-8">
              <div
                className="flex items-center justify-between w-full border-b-slate-300 border-b"
                onClick={() => setGraphs(!graphs)}
              >
                <h1 className="text-slate-900 font-medium text-sm rounded-lg">
                  Graphs
                </h1>
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
          </div> */}
        </div>
      </div>
      <aside className="sticky top-40 bg-slate-50 w-1/3 max-w-[400px] max-h-[700px] m-8 px-4 py-6">
        <p className="mb-8">Main</p>
        {versionData ? (
          <ul className="space-y-2">
            {versionData.map((version: any) => (
              <li className="flex items-center space-x-2" key={version.version}>
                <input
                  name={"version2"}
                  value={version.version}
                  type="checkbox"
                  checked={version.version === ver1 || version.version === ver2}
                  onChange={(e) => {
                    if (e.target.checked) {
                      setVer2(version.version);
                    } else if (ver1 === version.version && ver2 === "") {
                      new Error("You can't uncheck the present version");
                    } else if (ver1 === version.version) {
                      setVer1(ver2);
                      setVer2("");
                    } else {
                      setVer2("");
                    }
                  }}
                />
                <p>{version.version}</p>
              </li>
            ))}
          </ul>
        ) : (
          "No side"
        )}
      </aside>
    </main>
  );
}
