export default function PackageSection() {
  return (
    <div className="h-fit flex flex-col gap-y-6">
      <div className="flex justify-start items-center text-3xl md:text-4xl lg:text-6xl text-slate-850">
        Packages
      </div>
      <div className="flex flex-col gap-y-2 lg:gap-y-6 text-base md:text-lg lg:text-2xl !leading-normal text-slate-600">
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Docker for shipping
        </div>
        <div className="flex text-slate-600">
          <img
            src="/imgs/landingPage/Bullet400.svg"
            alt=""
            className="pr-3 w-14 lg:w-20"
          />
          Streamlit & gradio for sharing with your team
        </div>
      </div>
      <div className="flex justify-end">
        <img
          src="/imgs/landingPage/Packages.svg"
          alt="Packages"
          className="w-64 sm:hidden"
        />
      </div>
    </div>
  );
}
