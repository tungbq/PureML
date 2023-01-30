import Tag from "../ui/Tag";

export default function RawDataSection() {
  return (
    <div className="h-fit flex flex-col gap-y-6">
      <div className="flex justify-start items-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850">
        Raw Data
      </div>
      <div className="flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600">
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Load from different sources
        </div>
        <div className="flex text-slate-600 items-center">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          <div>
            <div className="pr-2">Works with any orchestrator</div>
            <Tag intent="landingpg" />
          </div>
        </div>
      </div>
      <div className="flex justify-end">
        <img
          src="/imgs/landingPage/RawData.svg"
          alt="RawData"
          className="w-64 sm:hidden"
        />
      </div>
    </div>
  );
}
