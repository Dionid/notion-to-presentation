/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["cmd/saas/**/*.templ", "libs/ntp/html.go"],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        mono: ["Courier Prime", "monospace"],
      },
    },
  },
  corePlugins: {
    preflight: true,
  },
};
