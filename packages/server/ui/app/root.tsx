import type { MetaFunction } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useCatch,
  useLoaderData,
  useMatches,
  useNavigate,
} from "@remix-run/react";
import type { ReactNode } from "react";
import { useEffect } from "react";
import RawDataSection from "./components/landingPage/RawDataSection";
import CTASection from "./components/landingPage/CTASection";
import DatasetSection from "./components/landingPage/DatasetSection";
import Footer from "./components/landingPage/Footer";
import HeroSection from "./components/landingPage/HeroSection";
import Navbar from "./components/landingPage/Navbar";
import ModelSection from "./components/landingPage/ModelSection";
import TransformerSection from "./components/landingPage/TransformerSection";
import { fetchUserSettings } from "./routes/api/auth.server";
import { getSession } from "./session";

// Tailwind CSS---------------------------------------------------------------
import styles from "./styles/app.css";
import style from "reactflow/dist/style.css";
import toast, { Toaster } from "react-hot-toast";
import PackageSection from "./components/landingPage/PackageSection";

export function links() {
  return [
    { rel: "stylesheet", href: styles },
    { rel: "stylesheet", href: style },
  ];
}
// ---------------------------------------------------------------------------

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const cookieHeader = request.headers.get("cookie");
  const cacheControlHeader = request.headers.get("cache-control");
  if (cookieHeader && cacheControlHeader) {
    const session = await getSession(request.headers.get("Cookie"));
    const accesstoken = session.get("accessToken");
    const profile = await fetchUserSettings(accesstoken);
    if (accesstoken) {
      if (profile.Data) {
        return profile;
      }
      return profile;
    }
  } else {
    request.headers.delete("cookie");
  }
  return null;
}

export default function App() {
  const navigate = useNavigate();
  const matches = useMatches();
  const pathname = matches[1].pathname;
  const prof = useLoaderData();
  useEffect(() => {
    if (pathname === "/" && prof) return navigate("/models");
  }, []);
  return (
    <Document>
      <Toaster position="bottom-right" reverseOrder={true} />
      {!prof && pathname === "/" ? (
        <>
          {!prof && (
            <div className="flex flex-col justify-center !font-outfit bg-mobBG md:bg-tabBG lg:bg-desktopBG 2xl:bg-largeBG bg-no-repeat bg-cover bg-center bg-fixed">
              <Navbar />
              <div className="flex justify-center bg-slate-50">
                <div className="w-full md:max-w-7xl px-6">
                  <HeroSection />
                </div>
              </div>
              <div className="flex justify-center">
                <div className="flex md:max-w-7xl p-8 md:py-16">
                  <div className="flex flex-col md:gap-y-24 lg:gap-y-48 xl:gap-y-80 sm:w-1/2">
                    <div className="flex flex-col md:gap-y-32 lg:gap-y-64 xl:gap-y-80 2xl:gap-y-96">
                      <RawDataSection />
                      <TransformerSection />
                      <DatasetSection />
                    </div>
                    <div className="flex flex-col md:gap-y-24 lg:gap-y-48 xl:gap-y-64 2xl:gap-y-72">
                      <ModelSection />
                      <PackageSection />
                    </div>
                  </div>
                  <img
                    src="/imgs/landingPage/LongPipeline.svg"
                    alt=""
                    className="hidden sm:block sm:w-3/5 xl:w-1/2"
                  />
                </div>
              </div>
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

function Document({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <head>
        <Meta />
        <Links />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;1,100;1,200;1,300;1,400;1,500;1,600;1,700&family=Outfit:wght@100;200;300;400;500;600;700;800;900&display=swap"
          rel="stylesheet"
        />
        <script async defer src="https://buttons.github.io/buttons.js"></script>
        <Scripts />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  );
}

export function ErrorBoundary({ error }: { error: Error }) {
  return (
    <Document>
      <div className="p-12">
        <span className="text-3xl font-medium">Error</span>
        <p>{error.message}</p>
        <div className="text-xl pt-8 font-medium">The stack trace is:</div>
        <pre className="whitespace-pre-wrap">{error.stack}</pre>
        <div className="flex justify-center items-center">
          <img src="/error/FunctionalError.gif" alt="Error" width="500" />
        </div>
      </div>
    </Document>
  );
}

export function CatchBoundary() {
  const caught = useCatch();

  return (
    <Document>
      <div className="p-12">
        <span className="text-3xl font-medium">Status: {caught?.status}</span>
        <div className="text-xl pt-8 font-medium">Data: {caught?.data}</div>
        <div className="flex justify-center items-center">
          <img src="/error/FunctionalError.gif" alt="Error" width="500" />
        </div>
      </div>
    </Document>
  );
}
