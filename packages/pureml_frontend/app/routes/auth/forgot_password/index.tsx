import { Form, useActionData } from "@remix-run/react";
import { Suspense, useEffect } from "react";
import { toast } from "react-toastify";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Loader from "~/components/ui/Loading";
import { fetchForgotPassword } from "~/routes/api/auth.server";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const email = form.get("email");
  const data = await fetchForgotPassword(email);
  return data;
};

export default function ForgotPassword() {
  const data = useActionData();
  useEffect(() => {
    if (!data) return;
    if (data.message === "Reset password link sent to your mail") {
      toast.success("Reset password link shared! Check your inbox.");
    } else if (data.message === "User with given email not found")
      toast.error("Email does not exist! Please sign up.");
    else toast.error("Something went wrong!");
  }, [data]);

  return (
    <Suspense fallback={<Loader />}>
      <Form method="post" className="text-slate-600 flex flex-col text-left">
        <div className="flex flex-col gap-y-12">
          <label htmlFor="email" className="font-medium">
            Email
            <Input
              intent="primary"
              type="email"
              name="email"
              fullWidth={false}
              aria-label="email"
              data-testid="email"
              required
            />
          </label>
          <Button intent="primary">Send Link</Button>
        </div>
        <span className="w-[24rem] text-slate-600 pt-6">
          Reset password link will be sent to your mail ID given above. Click
          link to change or reset password.
        </span>
      </Form>
    </Suspense>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}
