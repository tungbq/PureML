import { redirect } from "@remix-run/node";
import { useLoaderData, useNavigate } from "@remix-run/react";
import { Suspense, useEffect } from "react";
import { toast } from "react-toastify";
import Loader from "~/components/ui/Loading";
import { getSession } from "~/session";
import { fetchVerifyEmail } from "./api/auth.server";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const url = new URL(request.url);
  const token = url.searchParams.get("token");
  const trimToken = token?.trim();
  if (accesstoken) {
    if (accesstoken === trimToken) {
      return null;
    } else return redirect("/", {});
  }
  const data = await fetchVerifyEmail(trimToken);
  return data;
}

export default function VerifyEmail() {
  const data = useLoaderData();
  const navigate = useNavigate();
  useEffect(() => {
    const timer = setTimeout(async () => {
      if (!data) return;
      if (data.message === "User verified") {
        toast.success("Successfully verified!");
        navigate("/auth/signin");
      } else if (data.message === "Token is required") {
        toast.error("Can't verify! Please re-verify.");
        navigate("/auth/signin");
      } else if (data.message === "User email is already verified") {
        toast.success("Email is already verified! Please sign in.");
        navigate("/auth/signin");
      } else {
        toast.error("Something went wrong!");
        navigate("/auth/signin");
      }
    }, 2000);
    return () => clearTimeout(timer);
  }, [data]);

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
          <div>Verifying your email, give us a moment :)</div>
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
