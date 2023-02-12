import Button from "~/components/ui/Button";

export default function HeroSection() {
  return (
    <div className="h-fit sm:flex gap-y-16 justify-between" data-aos="fade-up">
      <div className="flex flex-col gap-y-6 sm:w-1/2 justify-center md:pb-8">
        <div className="flex justify-center items-center !leading-snug text-3xl md:text-4xl lg:text-6xl text-slate-850">
          A bridge connecting Data Engineer and Data Scientist
        </div>
        <div className="py-18 text-lg md:text-2xl !leading-normal text-slate-600 justify-self-center items-center">
          We reduce friction between data engineer and data scientist to
          facilitate seamless collaboration and efficient model development.
        </div>
        <div className="flex items-center">
          <a href="https://calendly.com/pureml-inc/pureml">
            <Button
              intent="landingpg"
              icon=""
              fullWidth={false}
              className="!w-24"
            >
              Schedule Demo
            </Button>
          </a>
        </div>
      </div>
      <div className="flex flex-col justify-center items-center py-8 overflow-visible">
        <img
          src="/imgs/landingPage/HeroDataEngineer.svg"
          alt="Hero"
          className="absolute move-avatar w-32 -ml-72 -mt-20 md:w-36 md:-ml-96 lg:w-64 lg:-ml-[30rem] z-10"
        />
        <img src="/imgs/landingPage/HeroImg.svg" alt="Hero" />
        <img
          src="/imgs/landingPage/HeroDataScientist.svg"
          alt="Hero"
          className="absolute move-avatar w-32 mt-16 ml-48 md:w-36 md:ml-56 lg:w-56 lg:ml-72 z-10"
        />
      </div>
    </div>
  );
}
