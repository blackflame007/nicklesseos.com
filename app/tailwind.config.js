/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
    './layouts/**/*.templ',
    './components/**/*.templ',
  ],
  daisyui: {
    themes: [
      "night",
      "coffee",
      "synthwave",
      "halloween",
      "sunset"
    ],
  },
  plugins: [require("daisyui")],
}

