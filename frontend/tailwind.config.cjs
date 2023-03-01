/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: ["./src/**/*.{vue,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui"), require("flowbite")],
};
