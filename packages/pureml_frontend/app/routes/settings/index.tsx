import { Suspense, useEffect } from "react";
import Tabbar from "~/components/Tabbar";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Loader from "~/components/ui/Loading";
import { getSession } from "~/session";
import { updateProfile, fetchUserSettings } from "../api/auth.server";
import { Form, useActionData, useLoaderData } from "@remix-run/react";
import { toast } from "react-toastify";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const userProfile = await fetchUserSettings(session.get("accessToken"));
  return userProfile;
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const name = form.get("name");
  const bio = form.get("bio");
  const session = await getSession(request.headers.get("Cookie"));
  const accessToken = session.get("accessToken");
  const data = await updateProfile(name, bio, accessToken);
  return data;
};

export default function ProfileSetting() {
  const profileData = useLoaderData();
  const updateProfileData = useActionData();
  useEffect(() => {
    if (!updateProfileData) return;
    if (updateProfileData.message === "User profile updated")
      toast.success("Profile Updated!");
    else toast.error("Something went wrong!");
  }, [updateProfileData]);

  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center w-full border-b-2 border-slate-100">
        <div className="w-full 2xl:max-w-screen-2xl">
          <Tabbar intent="primarySettingTab" tab="profile" />
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
              <label htmlFor="name" className="text-sm pb-1">
                Name
                <Input
                  intent="valuePrimary"
                  type="text"
                  name="name"
                  fullWidth={false}
                  defaultValue={profileData[0].name || "Add your name"}
                  aria-label="name"
                  data-testid="name"
                  required
                />
              </label>
            </div>
            <div className="pb-4">
              <label htmlFor="bio" className="text-sm pb-1">
                Bio
                <Input
                  intent="valuePrimary"
                  type="text"
                  name="bio"
                  fullWidth={false}
                  defaultValue={profileData[0].bio || "Enter your bio..."}
                  aria-label="bio"
                  data-testid="bio"
                  required
                />
              </label>
            </div>
            <div className="pb-4">
              <label htmlFor="email" className="text-sm pb-1">
                Email
                <Input
                  intent="read"
                  type="text"
                  name="email"
                  fullWidth={false}
                  defaultValue={profileData[0].email}
                  aria-label="email"
                  data-testid="email"
                  required
                />
              </label>
            </div>
            <div className="pb-8">
              <label htmlFor="username" className="text-sm pb-1">
                Username
                <Input
                  intent="read"
                  type="text"
                  name="username"
                  fullWidth={false}
                  defaultValue={profileData[0].handle || "Your username"}
                  aria-label="username"
                  data-testid="username"
                  required
                />
              </label>
            </div>
            <div className="w-fit">
              <Button fullWidth={false}>Save changes</Button>
            </div>
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
