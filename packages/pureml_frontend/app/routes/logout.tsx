import type { MetaFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { Suspense } from "react";
import Loader from "~/components/ui/Loading";
import { destroySession, getSession } from "~/session";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Logout | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  if (accesstoken) {
    return redirect("/", {
      headers: {
        "Set-Cookie": await destroySession(session),
      },
    });
  }
  return null;
}

export default function Logout() {
  useLoaderData();
  return (
    <Suspense fallback={<Loader />}>
      <div className="flex flex-col justify-center items-center text-slate-600 font-medium pt-80">
        <img src="/PureMLLogoWText.svg" alt="Logo" className="w-24 mb-6" />
        <div className="flex font-medium">
          <div className="mt-6 mr-5">
            <div className="loader-verify">
              <svg>
                <defs>
                  <filter id="goo">
                    <feGaussianBlur
                      in="SourceGraphic"
                      stdDeviation="2"
                      result="blur"
                    />
                    <feColorMatrix
                      in="blur"
                      mode="matrix"
                      values="1 0 0 0 0  0 1 0 0 0  0 0 1 0 0  0 0 0 5 -2"
                      result="gooey"
                    />
                    <feComposite
                      in="SourceGraphic"
                      in2="gooey"
                      operator="atop"
                    />
                  </filter>
                </defs>
              </svg>
            </div>
          </div>
          <div>Signing out, give us a moment :)</div>
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
