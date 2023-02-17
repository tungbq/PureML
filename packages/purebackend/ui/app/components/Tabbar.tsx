import clsx from "clsx";
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import { Link, useMatches } from "@remix-run/react";

const tabStyles = cva("text-zinc-400 font-medium flex bg-slate-0 sticky z-10", {
  variants: {
    intent: {
      primaryModelTab: "pt-4 border-b-2 border-slate-100 top-44",
      primaryDatasetTab: "pt-4 border-b-2 border-slate-100 top-28",
      primarySettingTab: "pt-4 border-b-2 border-slate-100 top-28",
      modelTab: "pt-8 top-[16rem]",
      datasetTab: "pt-8 top-[11.7rem]",
      modelReviewTab: "pt-8 top-[16rem]",
      datasetReviewTab: "pt-8 top-[11.7rem]",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    intent: "primaryModelTab",
    fullWidth: true,
  },
});

interface Props extends VariantProps<typeof tabStyles> {
  // [children: string]: any;
  tab: string;
  projectId?: string;
}

function secondaryLinkCss(currentPage: boolean) {
  return clsx(
    currentPage ? "text-white" : "text-slate-600",
    "flex justify-center items-center"
  );
}

export default function TabBar({ intent, fullWidth, tab }: Props) {
  const matches = useMatches();
  const path = matches[2].pathname;
  const pathname = decodeURI(path.slice(1));
  const orgId = pathname.split("/")[1];
  const modelId = pathname.split("/")[3];
  const datasetId = pathname.split("/")[3];
  const primaryModelTabs = [
    {
      id: "modelCard",
      name: "Model Card",
      hyperlink: `/org/${orgId}/models/${modelId}`,
    },
    {
      id: "versions",
      name: "Versions",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/metrics`,
    },
    {
      id: "review",
      name: "Review",
      hyperlink: `/org/${orgId}/models/${modelId}/review`,
    },
  ];
  const primaryDatasetTabs = [
    {
      id: "datasetCard",
      name: "Dataset Card",
      hyperlink: `/org/${orgId}/datasets/${datasetId}`,
    },
    {
      id: "versions",
      name: "Versions",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/datalineage`,
    },
    {
      id: "review",
      name: "Review",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review`,
    },
  ];
  const primarySettingTabs = [
    {
      id: "profile",
      name: "Profile",
      hyperlink: `/settings`,
    },
    {
      id: "account",
      name: "Account",
      hyperlink: `/settings/account`,
    },
    {
      id: "members",
      name: "Members",
      hyperlink: `/settings/members`,
    },
    // {
    //   id: "billing",
    //   name: "Billing",
    //   hyperlink: `/settings/billing`,
    // },
    // {
    //   id: "apitoken",
    //   name: "API Token",
    //   hyperlink: `/settings/apitoken`,
    // },
  ];
  const modelTabs = [
    {
      id: "metrics",
      name: "Metrics",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/metrics`,
    },
    {
      id: "graphs",
      name: "Graphs",
      hyperlink: `/org/${orgId}/models/${modelId}/versions/graphs`,
    },
  ];
  const datasetTabs = [
    {
      id: "datalineage",
      name: "Data Lineage",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/datalineage`,
    },
    // {
    //   id: "graphs",
    //   name: "Graphs",
    //   hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/graphs`,
    // },
  ];
  const modelReviewTabs = [
    {
      id: "metrics",
      name: "Metrics",
      hyperlink: `/org/${orgId}/models/${modelId}/review/commit`,
    },
    // {
    //   id: "graphs",
    //   name: "Graphs",
    //   hyperlink: `/org/${orgId}/models/${modelId}/review/commit/graphs`,
    // },
  ];
  const datasetReviewTabs = [
    {
      id: "datalineage",
      name: "Data Lineage",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review/commit`,
    },
    // {
    //   id: "graphs",
    //   name: "Graphs",
    //   hyperlink: `/org/${orgId}/datasets/${datasetId}/graphs`,
    // },
  ];
  return (
    <div className={tabStyles({ intent, fullWidth })}>
      <div className="flex px-10">
        {intent === "primaryModelTab" ||
        intent === "primaryDatasetTab" ||
        intent === "primarySettingTab" ? (
          <>
            {intent === "primaryModelTab" ? (
              <>
                {Object.keys(primaryModelTabs).map((key: string) => (
                  <div
                    key={key}
                    className={`${
                      tab === primaryModelTabs[key as never].id
                        ? "text-blue-700 border-b-2 border-blue-700"
                        : "text-slate-600"
                    } p-4`}
                  >
                    <Link to={primaryModelTabs[key as any].hyperlink}>
                      <span>{primaryModelTabs[key as any].name}</span>
                    </Link>
                  </div>
                ))}
              </>
            ) : (
              <>
                {intent === "primaryDatasetTab" ? (
                  <>
                    {Object.keys(primaryDatasetTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === primaryDatasetTabs[key as never].id
                            ? "text-blue-700 border-b-2 border-blue-700"
                            : "text-slate-600"
                        } p-4`}
                      >
                        <Link to={primaryDatasetTabs[key as any].hyperlink}>
                          <span>{primaryDatasetTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                ) : (
                  <>
                    {Object.keys(primarySettingTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === primarySettingTabs[key as never].id
                            ? "text-blue-700 border-b-2 border-blue-700"
                            : "text-slate-600"
                        } p-4`}
                      >
                        <Link to={primarySettingTabs[key as any].hyperlink}>
                          <span>{primarySettingTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                )}
              </>
            )}
          </>
        ) : (
          <>
            {intent === "modelTab" || intent === "datasetTab" ? (
              <div className="flex">
                {intent === "modelTab" ? (
                  <>
                    {Object.keys(modelTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === modelTabs[key as never].id
                            ? "bg-blue-700 rounded text-white"
                            : ""
                        } px-4 py-2`}
                      >
                        <Link
                          to={modelTabs[key as any].hyperlink}
                          className={`${secondaryLinkCss(
                            tab === modelTabs[key as any].id
                          )}`}
                        >
                          <span>{modelTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                ) : (
                  <>
                    {Object.keys(datasetTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === datasetTabs[key as never].id
                            ? "bg-blue-700 rounded text-white"
                            : ""
                        } px-4 py-2`}
                      >
                        <Link
                          to={datasetTabs[key as any].hyperlink}
                          className={`${secondaryLinkCss(
                            tab === datasetTabs[key as any].id
                          )}`}
                        >
                          <span>{datasetTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                )}
              </div>
            ) : (
              <div className="flex">
                {intent === "modelReviewTab" ? (
                  <>
                    {Object.keys(modelReviewTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === modelReviewTabs[key as never].id
                            ? "bg-blue-700 rounded text-white"
                            : ""
                        } px-4 py-2`}
                      >
                        <Link
                          to={modelReviewTabs[key as any].hyperlink}
                          className={`${secondaryLinkCss(
                            tab === modelReviewTabs[key as any].id
                          )}`}
                        >
                          <span>{modelReviewTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                ) : (
                  <>
                    {Object.keys(datasetReviewTabs).map((key: string) => (
                      <div
                        key={key}
                        className={`${
                          tab === datasetReviewTabs[key as never].id
                            ? "bg-blue-700 rounded text-white"
                            : ""
                        } px-4 py-2`}
                      >
                        <Link
                          to={datasetReviewTabs[key as any].hyperlink}
                          className={`${secondaryLinkCss(
                            tab === datasetReviewTabs[key as any].id
                          )}`}
                        >
                          <span>{datasetReviewTabs[key as any].name}</span>
                        </Link>
                      </div>
                    ))}
                  </>
                )}
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
