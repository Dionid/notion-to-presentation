/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["cmd/saas/**/*.templ"],
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
