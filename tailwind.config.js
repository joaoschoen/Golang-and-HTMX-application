const plugin = require("tailwindcss/plugin");

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./public/**/*.html"],
  purge: ["./public/**/*.html"],
  theme: {
    extend: {
      
    },
    colors: {
      Denim: "#1571B3",
      WhiteIce: "#DCF3F9",
      Malibu: "#77C4F5",
    },
  },
  plugins: [
    plugin(function ({ addBase, theme }) {
      addBase({
        h1: { fontSize: theme("fontSize.2xl"), textColor: "WhiteIce" },
        h2: { fontSize: theme("fontSize.xl") },
        h3: { fontSize: theme("fontSize.lg") },
      });
    }),
  ],
};
