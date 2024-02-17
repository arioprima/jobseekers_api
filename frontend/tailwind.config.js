/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    container: {
      center: true,
      padding: "10px",
    },
    extend: {
      colors: {
        primary: "#334155",
        secondary: "#475569",
      }
    },
  },
  plugins: [],
}