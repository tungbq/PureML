import { redirect } from "@remix-run/node";
import { Form, useActionData, useLoaderData } from "@remix-run/react";
import { Copy } from "lucide-react";
import { Suspense, useEffect } from "react";
import { toast } from "react-toastify";
import Tabbar from "~/components/Tabbar";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Loader from "~/components/ui/Loading";
import { fetchOrgDetails, updateOrg } from "~/routes/api/org.server";
import { getSession } from "~/session";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const orgId = session.get("orgId");
  const orgSettings = await fetchOrgDetails(orgId, session.get("accessToken"));
  return orgSettings;
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const orgname = form.get("orgname");
  const orgdesc = form.get("orgdesc");
  const orgid = form.get("orgid");
  const session = await getSession(request.headers.get("Cookie"));
  const accessToken = session.get("accessToken");
  const data = await updateOrg(orgdesc, orgname, orgid, accessToken);
  return data;
};

export default function Organization() {
  const orgData = useLoaderData();
  const updateOrgData = useActionData();
  useEffect(() => {
    if (!updateOrgData) return;
    if (updateOrgData.message === "Organization updated")
      toast.success("Organization details Updated!");
    else toast.error("Something went wrong!");
  }, [updateOrgData]);

  function copyOrgId() {
    navigator.clipboard.writeText(orgData[0].uuid);
    toast.success("Copied to clipboard!");
  }
  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center w-full border-b-2 border-slate-100">
        <div className="w-full 2xl:max-w-screen-2xl">
          <Tabbar intent="primarySettingTab" tab="organization" />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <Form
            method="post"
            reloadDocument
            className="py-8 px-12 w-full h-[80%] overflow-auto"
          >
            <div className="pb-4">
              <label htmlFor="orgname" className="pb-1">
                Organization Name
                <Input
                  intent="valuePrimary"
                  type="text"
                  name="orgname"
                  fullWidth={false}
                  defaultValue={orgData[0].name || "Add your name"}
                  aria-label="orgname"
                  data-testid="orgname"
                  required
                />
              </label>
            </div>
            <div className="pb-4">
              <label htmlFor="orgdesc" className="text-sm pb-1">
                Organization Description
                <Input
                  intent="valuePrimary"
                  type="text"
                  name="orgdesc"
                  fullWidth={false}
                  defaultValue={orgData[0].description || "Enter email..."}
                  aria-label="orgdesc"
                  data-testid="orgdesc"
                  required
                />
              </label>
            </div>
            <div className="pb-4">
              <label htmlFor="orgid" className="pb-1">
                Organization ID
                <div className="input-icons">
                  <input
                    className="hidden"
                    name="orgid"
                    defaultValue={orgData[0].uuid}
                  />
                  <input
                    id="orgid"
                    type="text"
                    name="orgid"
                    defaultValue={orgData[0].uuid || "Organization ID"}
                    aria-label="orgid"
                    data-testid="orgid"
                    required
                    disabled
                    className="input-field rounded !w-[24rem]"
                  />
                  <i>
                    <Copy
                      className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                      onClick={() => copyOrgId()}
                    />
                  </i>
                </div>
              </label>
            </div>
            <div className="pb-8">
              <label htmlFor="jwttoken" className="pb-1">
                JWT Token
                <Input
                  intent="read"
                  type="text"
                  name="jwttoken"
                  fullWidth={false}
                  defaultValue={orgData[0].token || "Your JWT Token"}
                  aria-label="jwttoken"
                  data-testid="jwttoken"
                  required
                />
              </label>
            </div>
            <div className="w-fit">
              <Button fullWidth={false}>Save changes</Button>
            </div>
            {/* <div className="pt-12 font-medium text-slate-800">
              Delete Organization
            </div>
            <div className="pb-8 text-slate-600">
              Delete this organization permanently, this action is irreversible
              All its repositories (models, datasets) will be deleted.
            </div>
            <div className="w-fit">
              <Button intent="danger" fullWidth={false}>
                Delete Organization
              </Button>
            </div> */}
          </Form>
        </div>
      </div>
    </Suspense>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/FunctionalError.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/FunctionalError.gif" alt="Error" width="500" />
    </div>
  );
}
