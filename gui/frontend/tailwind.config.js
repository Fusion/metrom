/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./dist/index.html","../components/*.go","./src/**/*.{vue,html,js,ts,jsx,tsx,svelte}",],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    logs: false,
    themes: [
      "light",
      "dark"
    ],
  },
}
