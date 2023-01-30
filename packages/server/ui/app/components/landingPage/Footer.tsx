export default function Footer() {
  return (
    <div className="p-8 md:p-24 lg:p-14 bg-slate-850 xl:h-fit flex md:justify-center">
      <div className="md:flex justify-center text-lg text-slate-400 font-medium gap-y-2 md:gap-y-0">
        <span className="px-4">
          © 2022,{" "}
          <a href="https://pureml.com" className="text-brand-200">
            PureML Inc
          </a>
        </span>
        <div className="py-4 md:py-0 px-4 flex flex-col md:flex md:flex-row">
          <a href="https://docs.pureml.com">Docs</a>
          <span className="px-2 hidden md:block"> | </span>
          <a href="https://discord.com/invite/xNUHt9yguJ">Join Discord</a>
        </div>
        <div className="flex justify-between px-4">
          <a href="https://twitter.com/getPureML">
            <img
              src="/imgs/landingPage/Twitter.svg"
              alt="Twitter"
              width="36"
              height="36"
              className="px-2"
            />
          </a>
          <a href="mailto:contact@pureml.com">
            <img
              src="/imgs/landingPage/Mail.svg"
              alt="Mail"
              width="36"
              height="36"
              className="px-2"
            />
          </a>
          <a href="https://www.linkedin.com/company/pureml-inc/">
            <img
              src="/imgs/landingPage/Linkedin.svg"
              alt="Linkedin"
              width="36"
              height="36"
              className="px-2"
            />
          </a>
        </div>
      </div>
    </div>
  );
}
