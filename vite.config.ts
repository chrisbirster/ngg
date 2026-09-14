import stylex from "@stylexjs/unplugin/vite";
import solid from "@solidjs/vite-plugin";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [stylex({ dev: true, unstable_moduleResolution: { type: "commonJS", rootDir: import.meta.dirname } }), solid()],
  build: { target: "es2023", outDir: "internal/web/dist", emptyOutDir: true },
  server: { proxy: { "/api": "http://localhost:8080" } }
});
