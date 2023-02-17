import { redirect } from "@remix-run/node";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import Link from "~/components/ui/Link";

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const password = form.get("password");
  const cpassword = form.get("cpassword");
  return redirect("/auth/signin");
};
export default function Posts() {
  return (
    <div className="md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white">
      <div className="md:max-w-[450px] w-96 text-center">
        <h2 className="font-semibold text-2xl text-slate-800 mb-12">
          Reset Password
        </h2>
        <form method="post" className="text-slate-400 flex flex-col text-left">
          <div className="pb-6">
            <label htmlFor="password" className="text-base pb-1">
              New Password
              <Input
                intent="primary"
                required
                // onChange={(e) => setPassword(e.target.value)}
                type="password"
                name="password"
                placeholder="Enter password..."
                aria-label="password"
                data-testid="password-input"
              />
            </label>
          </div>
          <div className="pb-6">
            <label htmlFor="cpassword" className="text-base pb-1">
              Confirm Password
              <Input
                intent="primary"
                required
                // onChange={(e) => setConfirmPassword(e.target.value)}
                type="password"
                name="cpassword"
                placeholder="Enter password..."
                aria-label="password"
                data-testid="password-input"
              />
            </label>
          </div>
          <Button aria-label="signin" intent="primary" icon="">
            Sign In
          </Button>
        </form>
        <div className="flex items-center text-slate-600 space-x-2 justify-center mt-6">
          <Link intent="secondary" hyperlink="/forgot_password">
            Forgot Password?
          </Link>
          <p>|</p>
          <div className="flex items-center space-x-1">
            <span className="text-sm">Already have an account?</span>
            <Link intent="secondary" hyperlink="/signin">
              Sign In
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
