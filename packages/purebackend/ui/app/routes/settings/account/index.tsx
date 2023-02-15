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
      <Tabbar intent="primarySettingTab" tab="account" />
      <form method="post" className="pt-8 px-12 w-2/3">
        <div className="pb-6">
          <label htmlFor="orgname" className="text-base pb-1">
            Email
            <Input
              intent="primary"
              // onChange={(e) => setEmail(e.target.value)}
              type="email"
              name="email"
              placeholder="Enter email..."
              aria-label="email"
              data-testid="email"
              required
            />
          </label>
        </div>
        <div className="pb-8">
          <label htmlFor="orgdesc" className="text-base pb-1">
            Organization domain name
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
        <Button icon="" fullWidth={false}>
          Save changes
        </Button>
        <div className="pt-12 text-slate-800">Delete Organization</div>
        <div className="text-base pb-8 text-slate-400">
          Delete this organization permanently, this action is irreversible All
          its repositories (models, datasets) will be deleted.
        </div>
        <Button intent="danger" icon="" fullWidth={false}>
          Delete Organization
        </Button>
      </form>
    </div>
  );
}
