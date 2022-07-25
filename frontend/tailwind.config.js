/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Nanum Gothic',  ...defaultTheme.fontFamily.sans],
        'serif': ['Ibarra Real Nova',   ...defaultTheme.fontFamily.serif]
      },
    },
  },
  plugins: [],
}
