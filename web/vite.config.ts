import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import deno from "@deno/vite-plugin";
import react from "@vitejs/plugin-react-swc";

// https://vite.dev/config/
export default defineConfig({
    plugins: [deno(), react(), tailwindcss()],
    server: {
        proxy: {
            "/api": {
                target: "http://localhost:8080",
                changeOrigin: true,
            },
            "/auth": {
                target: "http://localhost:8080",
                changeOrigin: true,
            },
        },
    },
});
