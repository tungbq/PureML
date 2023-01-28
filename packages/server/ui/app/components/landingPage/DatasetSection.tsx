export default function DatasetSection() {
  return (
    <div className="h-fit flex flex-col gap-y-6">
      <div className="flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850">
        Datasets
      </div>
      <div className="flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600">
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Create your own branching for experimentation
        </div>
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Track versions of dataset
        </div>
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Compare different versions of dataset
        </div>
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Review the submitted commit for versioning
        </div>
      </div>
      <div className="flex justify-end">
        <img
          src="/imgs/landingPage/Datasets.svg"
          alt="Datasets"
          className="w-64 sm:hidden"
        />
      </div>
    </div>
  );
}
