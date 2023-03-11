import { Form, useActionData } from "@remix-run/react";
import { Eye, EyeOff } from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import { toast } from "react-toastify";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Link from "~/components/ui/Link";
import Loader from "~/components/ui/Loading";
import { fetchSignUp } from "~/routes/api/auth.server";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const name = form.get("name");
  const username = form.get("username");
  const email = form.get("email");
  const password = form.get("password");
  const bio = form.get("bio");
  const data = await fetchSignUp(name, username, email, password, bio);
  return data;
};

export default function SignUp(err: string) {
  const data = useActionData();
  const [show, setShow] = useState(false);
  useEffect(() => {
    if (!data) return;
    if (data.message === "User created. Please verify your email address") {
      toast.success("Email verification sent! Please check your inbox.");
    } else if (data.message === "User with email already exists")
      toast.error("User with email already exists");
    else toast.error("Something went wrong!");
  }, [data]);

  return (
    <Suspense fallback={<Loader />}>
      <Form method="post" className="text-slate-600 flex flex-col text-left">
        <div className="flex flex-col gap-y-12">
          <div className="flex flex-col gap-y-6">
            <label htmlFor="name" className="font-medium">
              Name
              <Input
                intent="primary"
                type="text"
                name="name"
                fullWidth={false}
                aria-label="name"
                data-testid="name-input"
                required
              />
            </label>
            <label htmlFor="username" className="font-medium">
              Username
              <Input
                intent="primary"
                type="text"
                name="username"
                fullWidth={false}
                aria-label="username"
                data-testid="username-input"
                required
              />
            </label>
            <label htmlFor="email" className="font-medium">
              Email
              <Input
                intent="primary"
                required
                type="email"
                name="email"
                fullWidth={false}
                aria-label="emalid"
                data-testid="email-input2"
              />
            </label>
            <label htmlFor="password" className="font-medium">
              Password
              <div className="input-icons">
                <input
                  className="input-field rounded"
                  name="password"
                  type={show ? "text" : "password"}
                  required
                  aria-label="password"
                  data-testid="password-input1"
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
          <div className="w-[24rem] flex flex-col gap-y-6">
            <span className="text-slate-600">
              By clicking the Sign in button, you agree to PureML’s{" "}
              <a
                href="/TermsAndCondition.pdf"
                target="_blank"
                className="text-blue-250"
              >
                Terms of Service
              </a>{" "}
              and{" "}
              <a
                href="/PrivacyPolicy.pdf"
                target="_blank"
                className="text-blue-250"
              >
                Privacy Policy
              </a>
              .
            </span>
            <Button intent="primary">Sign Up</Button>
          </div>
        </div>
      </Form>
      <div className="flex items-center text-slate-600 justify-center mt-6">
        <Link intent="secondary" hyperlink="/auth/forgot_password">
          Forgot Password?
        </Link>
        <p className="px-2 text-slate-400">|</p>
        <div className="flex items-center space-x-1 font-medium">
          {/* <span>Already a user?</span> */}
          <Link intent="secondary" hyperlink="/auth/signin">
            Sign In
          </Link>
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
