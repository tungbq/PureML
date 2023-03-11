import { redirect } from "@remix-run/node";
import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
} from "@remix-run/react";
import { Eye, EyeOff } from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import { toast } from "react-toastify";
import Button from "~/components/ui/Button";
import Loader from "~/components/ui/Loading";
import {
  fetchResetPassword,
  fetchVerifyResetPassword,
} from "~/routes/api/auth.server";
import { getSession } from "~/session";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const verifiedReset = await fetchVerifyResetPassword(accesstoken);
  return verifiedReset;
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const newpassword = form.get("newpassword");
  const oldpassword = form.get("oldpassword");
  const url = new URL(request.url);
  const token = url.searchParams.get("token");
  const trimToken = token?.trim();
  const data = await fetchResetPassword(newpassword, oldpassword, trimToken);
  return data;
};

export default function ResetPassword() {
  useLoaderData();
  const data = useActionData();
  const [show, setShow] = useState(false);
  useEffect(() => {
    if (!data) return;
    if (data.message === "Password reset successfully") {
      toast.success("Password was reset successfully! Please sign in back.");
    } else if (data.message === "Token is required")
      toast.error("Invalid token!");
    else if (data.message === "Invalid credentials")
      toast.error("Old password did not match!");
    else toast.error("Something went wrong!");
  }, [data]);

  return (
    <Suspense fallback={<Loader />}>
      <div className="w-screen h-screen flex items-center justify-center bg-slate-50">
        <div className="bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white">
          <div className="w-fit text-center">
            <div className="flex justify-center items-center pb-16">
              <img src="/PureMLLogoWText.svg" alt="Logo" className="w-36" />
            </div>
            <Form
              method="post"
              className="text-slate-600 flex flex-col text-left"
            >
              <div className="flex flex-col gap-y-12">
                <div className="flex flex-col gap-y-6">
                  <label htmlFor="oldpassword" className="font-medium">
                    Enter Old password
                    <div className="input-icons">
                      <input
                        className="input-field rounded"
                        name="oldpassword"
                        aria-label="oldpassword"
                        data-testid="oldpassword-input1"
                        type={show ? "text" : "password"}
                        required
                      />
                      {show ? (
                        <i>
                          <Eye
                            className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() => setShow(!show)}
                          />
                        </i>
                      ) : (
                        <i>
                          <EyeOff
                            className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() => setShow(!show)}
                          />
                        </i>
                      )}
                    </div>
                  </label>
                  <label htmlFor="newpassword" className="font-medium">
                    Enter New Password
                    <div className="input-icons">
                      <input
                        className="input-field rounded"
                        name="newpassword"
                        aria-label="newpassword"
                        data-testid="newpassword-input"
                        type={show ? "text" : "password"}
                        required
                      />
                      {show ? (
                        <i>
                          <Eye
                            className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() => setShow(!show)}
                          />
                        </i>
                      ) : (
                        <i>
                          <EyeOff
                            className="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() => setShow(!show)}
                          />
                        </i>
                      )}
                    </div>
                  </label>
                </div>
                <Button intent="primary">Reset</Button>
              </div>
            </Form>
          </div>
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
