import { useState } from "react";
import Tabs from "react-simply-tabs";

export default function ToolsSection() {
  const [activeTabIndex, setActiveTabIndex] = useState(0);
  const AllToolsTabs = [
    {
      logo: "/imgs/landingPage/mltools/Airflow.svg",
      category: "Data",
      name: "Airflow",
      // licence: "Apache-2.0 License",
      goto: "https://airflow.apache.org/",
      desc: "Airflow is an open-source data orchestration tool that allows users to schedule, monitor, and manage complex data pipelines.",
    },
    {
      logo: "/imgs/landingPage/mltools/ApacheZepplin.svg",
      category: "Data",
      name: "Apache Zeppelin",
      // licence: "Apache-2.0 License",
      goto: "https://zeppelin.apache.org/",
      desc: "Apache Zeppelin is a web-based data science and visualization platform that enables interactive data exploration, analysis, and collaboration with multiple programming languages.",
    },
    {
      logo: "/imgs/landingPage/mltools/Argo.svg",
      category: "Orchestration",
      name: "Argo",
      // licence: "Apache-2.0 License",
      goto: "https://argoproj.github.io/",
      desc: "Argo is a container-native workflow engine for orchestrating parallel jobs on Kubernetes.",
    },
    {
      logo: "/imgs/landingPage/mltools/BentoML.svg",
      category: "Model Serving",
      name: "BentoML",
      // licence: "Apache-2.0 License",
      goto: "https://www.bentoml.ai/",
      desc: "BentoML is a platform for serving, deploying, and managing machine learning models.",
    },
    {
      logo: "/imgs/landingPage/mltools/Cronitor.svg",
      category: "Monitoring",
      name: "Cronitor",
      // licence: "MIT License",
      goto: "https://cronitor.io/",
      desc: "Monitoring product that provides real-time insights into the health and performance of your website or application, helping you to identify and fix issues quickly to ensure a seamless user experience.",
    },
    {
      logo: "/imgs/landingPage/mltools/Cortex.svg",
      category: "Model Serving",
      name: "Cortex",
      // licence: "Apache-2.0 License",
      goto: "https://www.cortex.dev/",
      desc: "Platform that allows developers to deploy machine learning models at scale in production environments, with features such as model versioning, automated scaling, and real-time monitoring.",
    },
    {
      logo: "/imgs/landingPage/mltools/Dask.svg",
      category: "Orchestration",
      name: "Dask",
      // licence: "BSD-3-Clause License",
      goto: "https://dask.org/",
      desc: "Distributed computing framework that allows users to parallelize and scale their data analysis tasks across multiple CPUs or GPUs, clusters, or cloud services using a flexible and user-friendly API.",
    },
    {
      logo: "/imgs/landingPage/mltools/DataRobot.svg",
      category: "AutoML",
      name: "DataRobot",
      // licence: "Apache-2.0 License",
      goto: "https://www.datarobot.com/platform/automated-machine-learning/",
      desc: "Platform for people who need to automate, ensure, and accelerate predictive analytics, assisting data scientists and analysts in developing and deploying correct predictive models.",
    },
    {
      logo: "/imgs/landingPage/mltools/DVC.svg",
      category: "Data",
      name: "DVC",
      // licence: "Apache-2.0 License",
      goto: "https://dvc.org/",
      desc: "DVC (Data Version Control) is a tool for managing data science projects and provides a version control system for data products, enabling reproducibility and collaboration.",
    },
    {
      logo: "/imgs/landingPage/mltools/ELi5.svg",
      category: "Monitoring",
      name: "ELI5",
      // licence: "MIT License",
      goto: "https://eli5.readthedocs.io/en/latest/#",
      desc: "Tool that helps keep track of a system or application's performance and alerts the user when there are any issues or abnormalities detected, using simple language that even a child could understand.",
    },
    {
      logo: "/imgs/landingPage/mltools/Feast.svg",
      category: "Model",
      name: "Feast",
      // licence: "Apache-2.0 License",
      goto: "https://feast.dev/",
      desc: "Feast is an open-source feature store that simplifies the management, discovery, and deployment of machine learning features for building and deploying ML models.",
    },
    {
      logo: "/imgs/landingPage/mltools/Flyte.svg",
      category: "Orchestration",
      name: "Flyte",
      // licence: "Apache-2.0 License",
      goto: "https://flyte.org/",
      desc: "Platform that helps organizations manage and execute complex workflows at scale, with a focus on machine learning and data processing pipelines.",
    },
    {
      logo: "/imgs/landingPage/mltools/GoogleAutoML.svg",
      category: "AutoML",
      name: "Google AutoML Cloud",
      // licence: "Apache-2.0 License",
      goto: "https://cloud.google.com/automl",
      desc: "The architecture of neural networks is used by Cloud AutoML.",
    },
    {
      logo: "/imgs/landingPage/mltools/Healthchecks.svg",
      category: "Monitoring",
      name: "Healthchecks.io",
      // licence: "BSD-3-Clause License",
      goto: "https://healthchecks.io/",
      desc: "A healthchecks monitoring product is a tool that regularly checks the status of a system or application to ensure that it is functioning properly.",
    },
    {
      logo: "/imgs/landingPage/mltools/Horovod.svg",
      category: "Model",
      name: "Horovod",
      // licence: "Apache-2.0 License",
      goto: "https://horovod.ai/",
      desc: "Horovod is a distributed deep learning framework that allows for parallel training of large neural network models across multiple GPUs or nodes.",
    },
    {
      logo: "/imgs/landingPage/mltools/KFServing.svg",
      category: "Model Serving",
      name: "KFServing",
      // licence: "Apache-2.0 License",
      goto: "https://www.kubeflow.org/docs/components/kfserving/",
      desc: "Product that enables the deployment and management of machine learning models in production environments using Kubernetes.",
    },
    {
      logo: "/imgs/landingPage/mltools/Lime.svg",
      category: "Model",
      name: "LIME",
      // licence: "BSD-2-Clause Licence",
      goto: "https://github.com/marcotcr/lime",
      desc: "Lime is an open-source framework for explaining the predictions of machine learning models, allowing users to understand how the model arrived at its results.",
    },
    {
      logo: "/imgs/landingPage/mltools/Luigi.svg",
      category: "Orchestration",
      name: "Luigi",
      // licence: "Apache-2.0 License",
      goto: "https://luigi.readthedocs.io/en/stable/",
      desc: "Luigi is an open-source workflow management system that helps automate complex data pipelines and tasks, allowing for efficient and scalable data processing.",
    },
    {
      logo: "/imgs/landingPage/mltools/MLflow.svg",
      category: "Model",
      name: "MLflow",
      // licence: "Apache-2.0 License",
      goto: "https://mlflow.org/",
      desc: "MLflow is a model management product that provides a comprehensive platform for tracking, managing, and deploying machine learning models across different frameworks and environments.",
    },
    {
      logo: "/imgs/landingPage/mltools/MLlib.svg",
      category: "Orchestration",
      name: "MLlib",
      // licence: "Apache-2.0 License",
      goto: "http://spark.apache.org/mllib/",
      desc: "MLlib is a machine learning library in Apache Spark that provides distributed algorithms for data processing and analysis, making it easier to build scalable machine learning pipelines.",
    },
    {
      logo: "/imgs/landingPage/mltools/Neptune.svg",
      category: "Model",
      name: "Neptune",
      // licence: "Apache-2.0 License",
      goto: "https://neptune.ai/",
      desc: "The Neptune Model Product is a machine learning experiment tracking and management tool.",
    },
    {
      logo: "/imgs/landingPage/mltools/Prefect.svg",
      category: "Orchestration",
      name: "Prefect",
      // licence: "Apache-2.0 License",
      goto: "https://www.prefect.io/",
      desc: "Prefect Orchestration Product is a workflow automation and management platform.",
    },
    {
      logo: "/imgs/landingPage/mltools/PyCaret.svg",
      category: "AutoML",
      name: "PyCaret",
      // licence: "Apache-2.0 License",
      goto: "https://pycaret.org/",
      desc: "Well-liked open-source and low-code Python machine learning library for automating machine learning models.",
    },
    {
      logo: "/imgs/landingPage/mltools/Seldoncore.svg",
      category: "Model Serving",
      name: "Seldon Core",
      // licence: "Apache-2.0 License",
      goto: "https://www.seldon.io/tech/products/core/",
      desc: "Seldon Core Model Serving Product is a scalable platform for deploying and managing machine learning models in production.",
    },
    {
      logo: "/imgs/landingPage/mltools/Shap.svg",
      category: "Model",
      name: "SHAP",
      // licence: "MIT License",
      goto: "https://shap.readthedocs.io/en/latest/index.html#",
      desc: "SHAP Model Product is an open-source library for explaining the output of machine learning models.",
    },
    {
      logo: "/imgs/landingPage/mltools/Splunk.svg",
      category: "AutoML",
      name: "Splunk",
      // licence: "Apache-2.0 License",
      goto: "https://www.splunk.com/",
      desc: "Platform aids in searching, analyzing, and visualizing the information obtained from the many websites, sensors, devices, and other applications that make up your company’s IT infrastructure.",
    },
    {
      logo: "/imgs/landingPage/mltools/Streamlit.svg",
      category: "Model Serving",
      name: "Streamlit",
      // licence: "Apache-2.0 License",
      goto: "https://streamlit.io/",
      desc: "Streamlit Model Serving Product is an open-source library for building interactive web-based applications for machine learning models.",
    },
    {
      logo: "/imgs/landingPage/mltools/Superset.svg",
      category: "Data",
      name: "Superset",
      // licence: "Apache-2.0 License",
      goto: "https://superset.apache.org/",
      desc: "Superset Data Product is an open-source business intelligence and data visualization platform.",
    },
  ];
  return (
    <div className="py-8 md:py-16">
      <Tabs
        activeTabIndex={activeTabIndex}
        onRequestChange={setActiveTabIndex}
        controls={[
          <button type="button" className="text-base text-slate-400" key="tab1">
            ALL TOOLS
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab2">
            DATA
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab3">
            MODEL
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab4">
            MODEL SERVING
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab5">
            ORCHESTRATION
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab6">
            MONITORING
          </button>,
          <button type="button" className="text-base text-slate-400" key="tab7">
            AUTO ML
          </button>,
        ]}
        controlsWrapperProps={{
          style: {
            paddingBottom: "10px",
            paddingTop: "10px",
            borderTop: "1px solid #CBD5E1",
            borderBottom: "1px solid #CBD5E1",
            display: "flex",
            columnGap: "2rem",
            overflow: "auto",
          },
        }}
        activeControlProps={{
          className: "active",
          style: { color: "#1E293B" },
        }}
      >
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <a
                href={AllToolsTabs[key as any].goto}
                key={key}
                className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
              >
                <div className="flex gap-x-4">
                  <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                    <img
                      src={AllToolsTabs[key as any].logo}
                      alt="Logo"
                      className="w-8 h-8"
                    />
                  </div>
                  <div>
                    <div className="text-base text-slate-400">
                      {AllToolsTabs[key as any].category}
                    </div>
                    <div className="text-2xl font-medium">
                      {AllToolsTabs[key as any].name}
                    </div>
                  </div>
                </div>
                <div className="flex flex-col gap-y-2">
                  {/* <div className="text-lg">
                    {AllToolsTabs[key as any].licence}
                  </div> */}
                  <span className="text-xl text-justify">
                    {AllToolsTabs[key as any].desc}
                  </span>
                </div>
              </a>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "Data" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "Model" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "Model Serving" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "Orchestration" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "Monitoring" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
        <div className="pt-16">
          <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-16 min-w-72">
            {Object.keys(AllToolsTabs).map((key: string) => (
              <>
                {AllToolsTabs[key as any].category === "AutoML" && (
                  <a
                    href={AllToolsTabs[key as any].goto}
                    key={key}
                    className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                  >
                    <div className="flex gap-x-4">
                      <div className="bg-slate-100 w-14 h-14 flex justify-center items-center rounded-lg border border-slate-200">
                        <img
                          src={AllToolsTabs[key as any].logo}
                          alt="Logo"
                          className="w-8 h-8"
                        />
                      </div>
                      <div>
                        <div className="text-base text-slate-400">
                          {AllToolsTabs[key as any].category}
                        </div>
                        <div className="text-2xl font-medium">
                          {AllToolsTabs[key as any].name}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col gap-y-2">
                      {/* <div className="text-lg">
                        {AllToolsTabs[key as any].licence}
                      </div> */}
                      <span className="text-xl text-justify">
                        {AllToolsTabs[key as any].desc}
                      </span>
                    </div>
                  </a>
                )}
              </>
            ))}
          </div>
        </div>
      </Tabs>
    </div>
  );
}
