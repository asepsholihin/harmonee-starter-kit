 /** @type {import('tailwindcss').Config} */
 export default {
    content: ["./templates/**/*.{tmpl,html,js}"],
    theme: {
      extend: {
        fontFamily: {
          'sans': ['Roboto', 'Helvetica', 'Arial', 'sans-serif']
        }
      },
    },
    plugins: [],
  }