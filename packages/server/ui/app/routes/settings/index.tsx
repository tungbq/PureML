import { redirect } from "@remix-run/node";
import Tabbar from "~/components/Tabbar";
import Button from "~/components/ui/Button";
import Dropdown from "~/components/ui/Dropdown";
import Input from "~/components/ui/Input";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const orgname = form.get("orgname");
  const orgdesc = form.get("desc");
  return redirect("/settings");
};
export default function Index() {
  return (
    <div>
      <Tabbar intent="primarySettingTab" tab="profile" />
      <form method="post" className="p-12 w-2/3">
        <div className="pb-6">
          <label htmlFor="orgname" className="text-base pb-1">
            Org Name
            <Input
              intent="primary"
              // onChange={(e) => setEmail(e.target.value)}
              type="text"
              name="orgname"
              placeholder="Enter Org Name..."
              aria-label="org-name"
              data-testid="org-name"
              required
            />
          </label>
        </div>
        <div className="pb-6">
          <label htmlFor="orgdesc" className="text-base pb-1">
            Org Description
            <Input
              intent="primary"
              // onChange={(e) => setEmail(e.target.value)}
              type="text"
              name="orgdesc"
              placeholder="Enter Org Description..."
              aria-label="org-desc"
              data-testid="org-desc"
              required
            />
          </label>
        </div>
        <div className="pb-8">
          <label htmlFor="orgtype" className="text-base pb-1">
            Org Type
            <Dropdown intent="orgType" fullWidth={false}>
              Choose
            </Dropdown>
          </label>
        </div>
        <Button icon="" fullWidth={false}>
          Save changes
        </Button>
      </form>
    </div>
  );
}
