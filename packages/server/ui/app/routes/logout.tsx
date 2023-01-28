import type { ActionArgs, MetaFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import { Form, Meta, useNavigate } from "@remix-run/react";
import toast from "react-hot-toast";
import Button from "~/components/ui/Button";
import { destroySession, getSession } from "~/session";

export const action = async ({ request }: ActionArgs) => {
  const session = await getSession(request.headers.get("Cookie"));
  toast.success("Signed Out Successfully!");
  return redirect("/", {
    headers: {
      "Set-Cookie": await destroySession(session),
    },
  });
};

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Logout | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function Logout() {
  const navigate = useNavigate();
  return (
    <div className="flex h-screen justify-center items-center bg-zinc-800 opacity-60">
      <head>
        <Meta />
      </head>
      <div className="bg-slate-0 p-4 rounded-lg">
        <div className="text-slate-800 font-medium">
          Are you sure you want to Sign out?
        </div>
        <div className="pt-12 grid justify-items-end w-full">
          <div className="flex justify-between w-1/2">
            <Button
              icon=""
              fullWidth={false}
              intent="secondary"
              onClick={() => {
                navigate("/");
              }}
            >
              No
            </Button>
            <Form method="post">
              <Button icon="" intent="danger" fullWidth={false} type="submit">
                Yes
              </Button>
            </Form>
          </div>
        </div>
      </div>
    </div>
  );
}
