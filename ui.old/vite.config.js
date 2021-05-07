"use strict";
exports.__esModule = true;
var vite_1 = require("vite");
var preset_vite_1 = require("@preact/preset-vite");
// https://vitejs.dev/config/
exports["default"] = vite_1.defineConfig({
    plugins: [preset_vite_1["default"]()]
});
