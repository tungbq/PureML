import { useState } from "react";
import Carousel from "react-simply-carousel";

export default function TestimonialsSection() {
  const [activeSlide, setActiveSlide] = useState(0);

  return (
    <div className="h-fit flex flex-col gap-y-6">
      <Carousel
        containerProps={{
          style: {
            justifyContent: "space-between",
            userSelect: "none",
          },
        }}
        itemsListProps={{
          style: {
            paddingBottom: "48px",
          },
        }}
        responsiveProps={[
          { minWidth: 1024, itemsToShow: 4 },
          { minWidth: 640, maxWidth: 1024, itemsToShow: 3 },
          { maxWidth: 640, itemsToShow: 2 },
        ]}
        swipeTreshold={60}
        activeSlideIndex={activeSlide}
        onRequestChange={setActiveSlide}
        dotsNav={{
          show: true,
          itemBtnProps: {
            style: {
              height: 12,
              width: 12,
              borderRadius: "50%",
              border: 0,
              background: "#e2e8f0",
              marginRight: "16px",
            },
          },
          activeItemBtnProps: {
            style: {
              height: 12,
              width: 12,
              borderRadius: "50%",
              border: 0,
              background: "#475569",
              marginRight: "16px",
            },
          },
        }}
        itemsToShow={3}
        speed={2000}
        easing="linear"
      >
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/1Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/2Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/3Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/4Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/5Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/6Test.svg"
            alt="Testimonial"
          />
        </div>
        <div className="bg-slate-50 flex flex-col gap-y-6 px-6 w-[23rem] md:w-96 xl:w-[30rem]">
          <img
            src="/imgs/landingPage/testimonials/7Test.svg"
            alt="Testimonial"
          />
        </div>
      </Carousel>
    </div>
  );
}
