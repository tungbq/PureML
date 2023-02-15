/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        brand: {
          50: "#c8fff4",
          100: "#70efde",
          200: "#03dac5",
          300: "#00c4b4",
          400: "#00b3a6",
          500: "#01a299",
          600: "#019592",
          700: "#018786",
          800: "#017374",
          900: "#005457",
        },
        slate: {
          0: "#ffffff",
          50: "#f8fafc",
          100: "#f1f5f9",
          150: "#eaeef3",
          200: "#e2e8f0",
          300: "#cbd5e1",
          400: "#94a3b8",
          500: "#64748b",
          600: "#475569",
          700: "#334155",
          800: "#1e293b",
          850: "#191f4d",
          900: "#0f172a",
        },
        blue: {
          750: "#0E4DDD",
        },
      },
      backgroundImage: {
        mobBG: "url('/imgs/landingPage/BGMobile.svg')",
        tabBG: "url('/imgs/landingPage/BGTab.svg')",
        desktopBG: "url('/imgs/landingPage/BGDesktop.svg')",
        largeBG: "url('/imgs/landingPage/BG4K.svg')",
      },
    },
    screens: {
      sm: "640px",
      md: "768px",
      lg: "1024px",
      xl: "1280px",
      "2xl": "2560px",
      "4xl": "3100px",
    },
    dropShadow: {
      "3xl": "0 0 30px rgba(203, 213, 225, 0.05)",
    },
  },
  plugins: [require("tailwindcss-radix")()],
};
