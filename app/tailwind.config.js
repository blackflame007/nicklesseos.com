/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/**/*.{templ,js,html}',
    './layouts/**/*.{templ,js,html}',
    './components/**/*.{templ,js,html}',
    // Add any other directories where you use Tailwind CSS classes
  ],
  safelist: [
    {
      pattern: /alert-+/, // You can display all the colors that you need
    },
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
  plugins: [
    require("daisyui"),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/typography')
  ],
}

