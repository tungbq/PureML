import { redirect } from "@remix-run/node";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Link from "~/components/ui/Link";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const email = form.get("email");
  return redirect("/auth/reset_password");
};
export default function Posts() {
  return (
    <div className="md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white">
      <div className="md:max-w-[450px] w-96 text-center">
        <h2 className="font-semibold text-2xl text-slate-800 mb-12">
          Forgot Password
        </h2>
        <form method="post" className="text-slate-400 flex flex-col text-left">
          <div className="pb-6">
            <label htmlFor="email" className="text-base pb-1">
              Email ID
            </label>
            <Input
              intent="primary"
              // onChange={(e) => setEmail(e.target.value)}
              type="email"
              name="email"
              placeholder="Enter email ID..."
              aria-label="emailid"
              data-testid="email-input3"
              required
            />
          </div>
          <Button aria-label="signin" intent="primary" icon="">
            Send Link
          </Button>
          <span className="text-sm text-zinc-400 pt-6">
            Reset password link will be sent to your mail ID given above. Click
            link to change or reset password.
          </span>
        </form>
        <div className="flex items-center text-slate-600 space-x-2 justify-center mt-6">
          <Link intent="secondary" hyperlink="/auth/signin">
            Go back
          </Link>
        </div>
      </div>
    </div>
  );
}
