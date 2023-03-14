import type { ActionArgs, MetaFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import {
  unstable_composeUploadHandlers,
  unstable_createFileUploadHandler,
  unstable_createMemoryUploadHandler,
  unstable_parseMultipartFormData,
} from "@remix-run/node";
import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
} from "@remix-run/react";
import { X } from "lucide-react";
import { useEffect } from "react";
import { toast } from "react-toastify";
import NavBar from "~/components/Navbar";
import Button from "~/components/ui/Button";
import Dropdown from "~/components/ui/Dropdown";
import Input from "~/components/ui/Input";
import { commitSession, getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";
import { fetchCreateOrg, fetchOrgDetails } from "./api/org.server";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Create Organization | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const profile = await fetchUserSettings(accesstoken);
  const orgId = session.get("orgId");
  const orgDetails = await fetchOrgDetails(orgId, session.get("accessToken"));
  return { profile, orgDetails };
}

export const action = async ({ request }: ActionArgs) => {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const uploadHandler = unstable_composeUploadHandlers(
    unstable_createFileUploadHandler({
      maxPartSize: 5_000_000,
      file: ({ filename }) => filename,
    }),
    unstable_createMemoryUploadHandler()
  );
  const form = await unstable_parseMultipartFormData(request, uploadHandler);
  const handle = form.get("orgname");
  const description = form.get("orgdesc");
  // const orgType = form.get("orgtype");
  // const orgLogo = form.get("orglogo");
  const data = await fetchCreateOrg(handle, description, accesstoken);
  if (data.message === "Organization created") {
    const orgID = data.data[0].uuid;
    session.unset("orgId");
    session.unset("orgName");
    session.set("orgId", orgID);
    session.set("orgName", handle);
    return json(data, {
      headers: { "Set-Cookie": await commitSession(session) },
    });
  } else {
    return data;
  }
};

export default function CreateOrg() {
  const prof = useLoaderData();
  const data = useActionData();
  const navigate = useNavigate();
  useEffect(() => {
    if (!data) return;
    if (data.message === "Organization created") {
      toast.success("Organization created successfully!");
      navigate("/models");
    } else toast.error("Something went wrong!");
  }, [data]);

  return (
    <div className="bg-slate-50 flex flex-col h-screen overflow-hidden">
      {prof ? (
        <NavBar
          intent="loggedIn"
          user={prof.profile[0].name.charAt(0).toUpperCase()}
          orgName={
            <a
              href={`/org/${prof.orgDetails[0].name}`}
              className="flex items-center justify-center"
            >
              {prof.orgDetails[0].name}
            </a>
          }
          orgAvatarName={prof.orgDetails[0].name.charAt(0)}
        />
      ) : (
        <NavBar
          intent="loggedOut"
          user=""
          orgName={
            <a href="/models" className="flex items-center justify-center pr-8">
              <img src="/PureMLLogoWText.svg" alt="Logo" className="w-20" />
            </a>
          }
          orgAvatarName=""
        />
      )}
      <div className="bg-slate-50 w-full h-full flex justify-center items-center">
        <div className="bg-slate-0 border border-slate-200 p-4 rounded-lg h-fit w-fit z-50">
          <div className="flex justify-between text-slate-800 font-medium pb-8">
            New Organization
            <X
              className="text-slate-400 w-4 cursor-pointer"
              onClick={() => {
                navigate("/models");
              }}
            />
          </div>
          <Form
            method="post"
            encType="multipart/form-data"
            className="text-sm text-slate-600"
          >
            <div className="pb-4">
              <label htmlFor="orgname">
                <div className="pb-2 font-medium">Organization Name</div>
                <Input
                  intent="primary"
                  type="text"
                  name="orgname"
                  aria-label="orgname"
                  data-testid="orgname"
                  required
                />
              </label>
            </div>
            <div className="pb-6">
              <label htmlFor="orgdesc">
                <div className="pb-2 font-medium">Organization Description</div>
                <Input
                  intent="primary"
                  type="text"
                  name="orgdesc"
                  aria-label="orgdesc"
                  data-testid="orgdesc"
                  required
                />
              </label>
            </div>
            <div className="flex justify-between gap-x-4">
              <div className="w-1/2">
                <label htmlFor="orglogo">
                  <div className="pb-2">Logo (optional)</div>
                  <input
                    type="file"
                    id="files"
                    name="orglogo"
                    accept="image/*"
                    disabled
                    className="cursor-not-allowed"
                  />
                </label>
              </div>
              <div className="w-1/2">
                <label htmlFor="orgtype">
                  <div className="pb-2">Organization Type</div>
                  {/* <Dropdown
                    intent="orgType"
                    name="orgtype"
                  >
                    Choose
                  </Dropdown> */}
                  <select disabled className="rounded cursor-not-allowed">
                    <option value="type">Select Type</option>
                  </select>
                </label>
              </div>
            </div>
            <div className="w-full pt-8">
              <Button intent="primary" type="submit">
                Create
              </Button>
            </div>
          </Form>
        </div>
      </div>
    </div>
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
