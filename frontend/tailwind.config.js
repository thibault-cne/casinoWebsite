/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{vue,js}","./node_modules/flowbite/**/*.js"],
    theme: {
        extend: {},
    },
    plugins: [require('flowbite/plugin')],
}
