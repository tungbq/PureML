import clsx from "clsx";
import { tv, type VariantProps } from "tailwind-variants";
import { Link, useMatches } from "@remix-run/react";

const tabStyles = tv({
  base: "text-zinc-400 font-medium flex bg-slate-50 sticky px-12",
  variants: {
    intent: {
      modelReviewTab: "pt-8",
      datasetReviewTab: "pt-8",
      modelReviewMetricsTab: "pt-8",
      datasetReviewLineageTab: "pt-8",
    },
    fullWidth: {
      true: "w-full",
      false: "w-fit",
    },
  },
  defaultVariants: {
    intent: "modelReviewTab",
    fullWidth: true,
  },
});

interface Props extends VariantProps<typeof tabStyles> {
  tab: string;
}

function primaryLinkCss(currentPage: boolean) {
  return clsx(
    currentPage ? "text-slate-600" : "text-slate-400",
    "flex justify-center items-center"
  );
}

export default function ReviewTabBar(props: Props) {
  const matches = useMatches();
  const path = matches[4].pathname;
  const pathname = decodeURI(path.slice(1));
  const orgId = pathname.split("/")[1];
  const modelId = pathname.split("/")[3];
  const datasetId = pathname.split("/")[3];
  const modelReviewId = pathname.split("/")[5];
  const datasetReviewId = pathname.split("/")[5];
  const modelReviewTabs = [
    {
      id: "newcommits",
      name: "New Commits",
      hyperlink: `/org/${orgId}/models/${modelId}/review`,
    },
    {
      id: "approved",
      name: "Approved",
      hyperlink: `/org/${orgId}/models/${modelId}/review/approved`,
    },
    {
      id: "rejected",
      name: "Rejected",
      hyperlink: `/org/${orgId}/models/${modelId}/review/rejected`,
    },
  ];
  const datasetReviewTabs = [
    {
      id: "newcommits",
      name: "New Commits",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review`,
    },
    {
      id: "approved",
      name: "Approved",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review/approved`,
    },
    {
      id: "rejected",
      name: "Rejected",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review/rejected`,
    },
  ];
  const modelReviewMetricsTabs = [
    {
      id: "metrics",
      name: "Metrics",
      hyperlink: `/org/${orgId}/models/${modelId}/review/${modelReviewId}/logs`,
    },
    // {
    //   id: "graphs",
    //   name: "Graphs",
    //   hyperlink: `/org/${orgId}/models/${modelId}/review/${reviewId}/graphs`,
    // },
  ];
  const datasetReviewLineageTabs = [
    {
      id: "datalineage",
      name: "Data Lineage",
      hyperlink: `/org/${orgId}/datasets/${datasetId}/review/${datasetReviewId}/datalineage`,
    },
    // {
    //   id: "graphs",
    //   name: "Graphs",
    //   hyperlink: `/org/${orgId}/datasets/${datasetId}/review/${reviewId}/graphs`,
    // },
  ];
  return (
    <div className={tabStyles(props)}>
      <div className="flex bg-slate-50 rounded-lg">
        {props.intent === "modelReviewTab" ||
        props.intent === "datasetReviewTab" ? (
          <>
            {props.intent === "modelReviewTab" ? (
              <>
                {Object.keys(modelReviewTabs).map((key: string) => (
                  <div
                    key={key}
                    className={`${
                      props.tab === modelReviewTabs[key as never].id
                        ? "text-slate-600 rounded-lg border border-brand-200"
                        : "text-slate-600"
                    } px-4 py-2`}
                  >
                    <Link
                      to={modelReviewTabs[key as any].hyperlink}
                      className={`${primaryLinkCss(
                        props.tab === modelReviewTabs[key as any].id
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
                      props.tab === datasetReviewTabs[key as never].id
                        ? "text-slate-600 rounded-lg border border-brand-200"
                        : "text-slate-600"
                    } px-4 py-2`}
                  >
                    <Link
                      to={datasetReviewTabs[key as any].hyperlink}
                      className={`${primaryLinkCss(
                        props.tab === datasetReviewTabs[key as any].id
                      )}`}
                    >
                      <span>{datasetReviewTabs[key as any].name}</span>
                    </Link>
                  </div>
                ))}
              </>
            )}
          </>
        ) : (
          <>
            {props.intent === "modelReviewMetricsTab" ? (
              <>
                {Object.keys(modelReviewMetricsTabs).map((key: string) => (
                  <div
                    key={key}
                    className={`${
                      props.tab === modelReviewMetricsTabs[key as never].id
                        ? "bg-slate-200 rounded text-slate-600"
                        : ""
                    } px-4 py-2`}
                  >
                    <Link
                      to={modelReviewMetricsTabs[key as any].hyperlink}
                      className={`${primaryLinkCss(
                        props.tab === modelReviewMetricsTabs[key as any].id
                      )}`}
                    >
                      <span>{modelReviewMetricsTabs[key as any].name}</span>
                    </Link>
                  </div>
                ))}
              </>
            ) : (
              <>
                {Object.keys(datasetReviewLineageTabs).map((key: string) => (
                  <div
                    key={key}
                    className={`${
                      props.tab === datasetReviewLineageTabs[key as never].id
                        ? "bg-slate-200 rounded text-slate-600"
                        : ""
                    } px-4 py-2`}
                  >
                    <Link
                      to={datasetReviewLineageTabs[key as any].hyperlink}
                      className={`${primaryLinkCss(
                        props.tab === datasetReviewLineageTabs[key as any].id
                      )}`}
                    >
                      <span>{datasetReviewLineageTabs[key as any].name}</span>
                    </Link>
                  </div>
                ))}
              </>
            )}
          </>
        )}
      </div>
    </div>
  );
}
