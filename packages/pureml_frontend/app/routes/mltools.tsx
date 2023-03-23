import type { MetaFunction } from "@remix-run/node";
import CTASection from "~/components/landingPage/CTASection";
import Footer from "~/components/landingPage/Footer";
import Navbar from "~/components/landingPage/Navbar";
import ToolsSection from "~/components/landingPage/ToolsSection";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "ML Tools | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function SwitchOrg() {
  return (
    <div className="bg-slate-50 landingpg-font flex flex-col justify-center">
      <div className="mltoolsbg bg-cover">
        <Navbar />
        <div className="flex flex-col justify-center md:items-center md:text-center gap-y-4 px-4 md:px-0 h-80 md:h-96 lg:h-[32rem]">
          <h1 className="font-bold text-4xl md:text-5xl lg:text-6xl">
            MLOps Tools
          </h1>
          <h1 className="text-slate-600 text-lg md:text-2xl lg:text-3xl md:w-4/5 lg:w-3/5 xl:w-2/5 2xl:w-1/5">
            A selection of the finest MLOps tools to help you construct the
            ideal machine learning stack.
          </h1>
        </div>
      </div>
      <div className="bg-slate-50 flex justify-center">
        <div className="w-full md:max-w-screen-xl px-4 md:px-8">
          <ToolsSection />
        </div>
      </div>
      <CTASection />
      <Footer />
    </div>
  );
}
