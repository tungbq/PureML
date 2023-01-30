import type { MetaFunction } from "@remix-run/node";
import { Meta, useNavigate } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";
import AvatarIcon from "~/components/ui/Avatar";
import Button from "~/components/ui/Button";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Dataset Review | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function DatasetReview() {
  const navigate = useNavigate();
  return (
    <div id="datasetsReview">
      <head>
        <Meta />
      </head>
      <Tabbar intent="primaryDatasetTab" tab="review" />
      <div className="px-12 pt-8 w-2/3">
        <div className="pb-6">
          <div className="bg-slate-100 rounded-md flex justify-between p-4">
            <div className="flex items-center">
              <AvatarIcon children="A" />
              <div className="text-slate-600 pl-4">
                Dason J. submitted V1.2 of “Housing_prediction” from Dev
              </div>
            </div>
            <Button
              icon=""
              fullWidth={false}
              onClick={() => {
                navigate(`commit`);
              }}
            >
              View Commit
            </Button>
          </div>
        </div>
        <div className="pb-6">
          <div className="bg-slate-100 rounded-md flex justify-between p-4">
            <div className="flex items-center">
              <AvatarIcon children="B" />
              <div className="text-slate-600 pl-4">
                Dason J. submitted V1.2 of “Housing_prediction” from Dev
              </div>
            </div>
            <Button
              icon=""
              fullWidth={false}
              onClick={() => {
                navigate(`commit`);
              }}
            >
              View Commit
            </Button>
          </div>
        </div>
        <div className="pb-6">
          <div className="bg-slate-100 rounded-md flex justify-between p-4">
            <div className="flex items-center">
              <AvatarIcon children="C" />
              <div className="text-slate-600 pl-4">
                Dason J. submitted V1.2 of “Housing_prediction” from Dev
              </div>
            </div>
            <Button
              icon=""
              fullWidth={false}
              onClick={() => {
                navigate(`commit`);
              }}
            >
              View Commit
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
