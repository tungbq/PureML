import type { MetaFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useActionData,
  useCatch,
  useLoaderData,
  useLocation,
  useMatches,
  useNavigate,
} from "@remix-run/react";
import type { ReactNode } from "react";
import { useEffect } from "react";
import CTASection from "./components/landingPage/CTASection";
import Footer from "./components/landingPage/Footer";
import HeroSection from "./components/landingPage/HeroSection";
import Navbar from "./components/landingPage/Navbar";
import PackageSection from "./components/landingPage/PackageSection";
import base, { fetchUserSettings } from "./routes/api/auth.server";
import { destroySession, getSession } from "./session";
import * as gtag from "~/analytics/gtags.client";
import { Analytics } from "@vercel/analytics/react";

// Tailwind CSS---------------------------------------------------------------
import styles from "./styles/app.css";
import style from "reactflow/dist/style.css";
import contributionStyle from "react-calendar-heatmap/dist/styles.css";
import toastStyle from "react-toastify/dist/ReactToastify.css";
import { Slide, ToastContainer } from "react-toastify";
import TrustedByCompanies from "./components/landingPage/TrustedByCompanies";
import VersionSection from "./components/landingPage/VersionSection";
import ReviewSection from "./components/landingPage/ReviewSection";
import TestingSection from "./components/landingPage/TestingSection";
import TestimonialsSection from "./components/landingPage/Testimonials";
import DeploySection from "./components/landingPage/DeploySection";

export function links() {
  return [
    { rel: "stylesheet", href: styles },
    { rel: "stylesheet", href: style },
    { rel: "stylesheet", href: toastStyle },
    { rel: "stylesheet", href: contributionStyle },
    { rel: "preconnect", href: "https://fonts.googleapis.com" },
    {
      rel: "preconnect",
      href: "https://fonts.gstatic.com",
      crossOrigin: "anonymous",
    },
    {
      rel: "stylesheet",
      href: "https://fonts.googleapis.com/css2?family=IBM+Plex+Mono&family=IBM+Plex+Sans:wght@400;500&family=Space+Grotesk:wght@400;500&display=swap",
    },
  ];
}
// ---------------------------------------------------------------------------

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "PureML",
  description: "Developer Platform for Production ML.",
  "og:image": "/PureMLMetaTag.svg",
  "twitter:image": "/PureMLMetaTag.svg",
  viewport: "width=device-width, height=device-height, initial-scale=1",
});

export async function loader({ request }: any) {
  const cookieHeader = request.headers.get("cookie");
  const cacheControlHeader = request.headers.get("cache-control");
  if (cookieHeader && cacheControlHeader) {
    const session = await getSession(request.headers.get("Cookie"));
    const accesstoken = session.get("accessToken");
    if (accesstoken) {
      const profile = await fetchUserSettings(accesstoken);
      if (!profile)
        return redirect("/", {
          headers: {
            "Set-Cookie": await destroySession(session),
          },
        });
      return { profile, gaTrackingId: process.env.GA_TRACKING_ID };
    }
  }
  return { profile: null, gaTrackingId: process.env.GA_TRACKING_ID };
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const email = form.get("email");
  console.log(email);
  const submit = base("tblxctpn3uWjLg9TP").create(
    [
      {
        fields: {
          Email: email,
        },
      },
    ],
    (err: string) => {
      if (err) {
        console.error(err);
      }
    }
  );
  console.log(submit);
  return redirect("/", {});
};

export default function App() {
  const location = useLocation();
  const navigate = useNavigate();
  const matches = useMatches();
  const pathname = matches[1].pathname;
  const profile = useLoaderData();
  const actData = useActionData();
  const { prof, gaTrackingId } = useLoaderData();
  useEffect(() => {
    if (profile.profile && pathname === "/") return navigate("/models");
  }, [profile]);

  useEffect(() => {
    if (gaTrackingId?.length && location) {
      gtag.pageview(location.pathname, gaTrackingId);
    }
  }, [location, gaTrackingId]);
  return (
    <Document gaTrackingId={gaTrackingId}>
      <ToastContainer
        theme="light"
        position="bottom-right"
        autoClose={5000}
        transition={Slide}
        hideProgressBar
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      {!prof && pathname === "/" ? (
        <>
          {!prof && (
            <div className="bg-slate-50 landingpg-font flex flex-col justify-center">
              <div className="flex flex-col justify-between h-screen 2xl:h-fit bgvideo bg-cover">
                <Navbar />
                <div className="flex flex-col gap-y-64 md:gap-y-48 pb-24 md:pb-24 2xl:pt-32">
                  <HeroSection />
                  <TrustedByCompanies />
                </div>
              </div>
              <div className="bg-slate-50 flex justify-center">
                <div className="md:max-w-screen-xl px-4 md:px-8">
                  <VersionSection />
                  <TestingSection />
                  <ReviewSection />
                  <PackageSection />
                  <DeploySection />
                </div>
              </div>
              {/* <div className="xl:flex xl:justify-center overflow-hidden pt-14 md:py-20">
                <div className="md:max-w-screen-xl">
                  <TestimonialsSection />
                  <JoinCommunitySection />
                </div>
              </div> */}
              <CTASection />
              <Footer />
            </div>
          )}
        </>
      ) : (
        ""
      )}
      <Outlet />
    </Document>
  );
}

function Document({
  children,
  gaTrackingId,
}: {
  children: ReactNode;
  gaTrackingId: string;
}) {
  return (
    <html lang="en">
      <head>
        <Meta />
        <meta property="og:image" content="/PureMLMetaTag.svg" />
        <meta property="twitter:image" content="/PureMLMetaTag.svg" />
        <Links />
      </head>
      <body>
        <script async defer src="https://buttons.github.io/buttons.js"></script>
        {!gaTrackingId ? null : (
          <>
            <script
              async
              src={`https://www.googletagmanager.com/gtag/js?id=${gaTrackingId}`}
            />
            <script
              async
              id="gtag-init"
              dangerouslySetInnerHTML={{
                __html: `
                window.dataLayer = window.dataLayer || [];
                function gtag(){dataLayer.push(arguments);}
                gtag('js', new Date());
                gtag('config', '${gaTrackingId}', {
                  page_path: window.location.pathname,
                });
              `,
              }}
            />
          </>
        )}
        <Scripts />
        <Analytics />
        <div>{children}</div>
        <ScrollRestoration />
        <LiveReload />
      </body>
    </html>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary({ error }: { error: Error }) {
  return (
    <Document gaTrackingId={""}>
      {/* <div className="p-12">
        <span className="text-3xl font-medium">Error</span>
        <p>{error.message}</p>
        <div className="text-xl pt-8 font-medium">The stack trace is:</div>
        <pre className="whitespace-pre-wrap">{error.stack}</pre> */}
      <div className="flex flex-col h-screen justify-center items-center">
        <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
        <div className="text-3xl text-slate-600 font-medium">
          Something went wrong :(
        </div>
        <img src="/error/FunctionalError.gif" alt="Error" width="500" />
      </div>
      {/* </div> */}
    </Document>
  );
}

export function CatchBoundary() {
  const caught = useCatch();

  return (
    <Document gaTrackingId={""}>
      {/* <div className="p-12">
        <span className="text-3xl font-medium">Status: {caught?.status}</span>
        <div className="text-xl pt-8 font-medium">Data: {caught?.data}</div> */}
      <div className="flex flex-col h-screen justify-center items-center">
        <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
        <div className="text-3xl text-slate-600 font-medium">
          Something went wrong :(
        </div>
        <img src="/error/FunctionalError.gif" alt="Error" width="500" />
      </div>
      {/* </div> */}
    </Document>
  );
}
