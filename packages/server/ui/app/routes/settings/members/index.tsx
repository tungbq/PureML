import { redirect } from "@remix-run/node";
import { Link } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const name = form.get("name");
  const desc = form.get("desc");
  return redirect("/org/oegId/settings");
};
export default function Index() {
  return (
    <div>
      <Tabbar intent="primarySettingTab" tab="members" />
      <div className="pt-8 px-12 w-2/3">Members List will be shown here</div>
    </div>
  );
}
