import { json, redirect } from "@remix-run/node";
import { Form, useActionData, useLoaderData } from "@remix-run/react";
import { Eye, EyeOff } from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import { toast } from "react-toastify";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Link from "~/components/ui/Link";
import Loader from "~/components/ui/Loading";
import { fetchSignIn } from "~/routes/api/auth.server";
import { fetchAllOrgs } from "~/routes/api/org.server";
import { commitSession, getSession } from "~/session";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  if (accesstoken) {
    return redirect("/models", {});
  }
  return null;
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const email = form.get("email");
  const password = form.get("password");
  const data = await fetchSignIn(email, password);
  const session = await getSession(request.headers.get("Cookie"));
  if (data.message === "User logged in") {
    const accessToken = data.data[0].accessToken;
    session.set("accessToken", accessToken);
    const org = await fetchAllOrgs(accessToken);
    session.set("orgId", org[0].org.uuid);
    session.set("orgName", org[0].org.name);
    return json(data, {
      headers: {
        "Set-Cookie": await commitSession(session),
      },
    });
  }
  return { data, ok: true };
};

export default function SignIn() {
  useLoaderData();
  const data = useActionData();
  const [show, setShow] = useState(false);
  useEffect(() => {
    if (!data) return;
    if (data.message === "User logged in") {
      toast.success("Successfully signed in");
    } else if (data.data.message === "User email is not verified")
      toast.error("Email not verified! Please verify to proceed.");
    else if (data.data.message === "User not found")
      toast.error("User not found");
    else if (data.data.message === "Invalid username")
      toast.error("Invalid username!");
    else if (data.data.message === "Invalid credentials")
      toast.error("Invalid credentials!");
    else toast.error("Something went wrong!");
  }, [data]);

  return (
    <Suspense fallback={<Loader />}>
      <Form method="post" className="text-slate-600 flex flex-col text-left">
        <div className="flex flex-col gap-y-12">
          <div className="flex flex-col gap-y-6">
            <label htmlFor="email" className="font-medium">
              Email
              <Input
                intent="primary"
                type="email"
                name="email"
                fullWidth={false}
                aria-label="emailid"
                data-testid="email-input1"
                required
              />
            </label>
            <label htmlFor="password" className="font-medium">
              Password
              <div className="input-icons">
                <input
                  className="input-field rounded"
                  name="password"
                  aria-label="password"
                  data-testid="password-input1"
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
          <Button intent="primary">Sign in</Button>
        </div>
      </Form>
      <div className="flex items-center text-slate-600 justify-center mt-6">
        <Link intent="secondary" hyperlink="/auth/forgot_password">
          Forgot Password?
        </Link>
        {/* <p className="px-2 text-slate-400">|</p>
        <div className="flex items-center space-x-1 font-medium">
          <Link intent="secondary" hyperlink="/auth/signup">
            Sign Up
          </Link>
        </div> */}
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
