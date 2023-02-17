import Tabbar from "~/components/Tabbar";
import Dropdown from "~/components/ui/Dropdown";

export default function DatasetGraphs() {
  return (
    <main className="flex">
      <div className="w-full" id="main">
        <Tabbar intent="datasetTab" tab="graphs" />
        <div className="px-10 py-8">
          <section className="w-full">
            {/* <div
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
            )} */}
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
      </aside>
    </main>
  );
}
