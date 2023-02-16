import Tag from "../ui/Tag";

export default function TransformerSection() {
  return (
    <div className="h-fit flex flex-col gap-y-6">
      <div className="flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850">
        Transformers
      </div>
      <div className="flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600">
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Track data lineage from raw data to dataset
        </div>
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Save code snippets for the transformations
        </div>
        <div className="flex text-slate-600 items-center">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          <div>
            Automate data pipeline with changes in any transformation step.
            <Tag intent="landingpg" />
          </div>
        </div>
      </div>
      <div className="flex justify-end">
        <img
          src="/imgs/landingPage/Transformer.svg"
          alt="Transformer"
          className="w-[18rem] sm:hidden"
        />
      </div>
    </div>
  );
}
