import Button from "../ui/Button";

export default function CTASection() {
  return (
    <div className="p-8 pb-32 h-fit flex flex-col justify-center items-center">
      <div className="lg:w-[60rem] flex flex-col gap-y-6">
        <div className="flex justify-center items-center md:pt-28 text-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850">
          PureML empowers everyone work together, and build knowledge.
        </div>
        <div className="py-18 !leading-normal text-lg md:text-2xl text-slate-600 justify-self-center items-center text-center font-medium">
          No more jumping between tools, struggling with versions, compare
          changes or sharing via cloud providers.
        </div>
        <div className="flex justify-center items-center pt-10">
          <Button
            intent="landingpg"
            icon=""
            fullWidth={false}
            className="!w-24"
          >
            Schedule Demo
          </Button>
        </div>
      </div>
    </div>
  );
}
