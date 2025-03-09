import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import deno from "@deno/vite-plugin";
import react from "@vitejs/plugin-react-swc";

// https://vite.dev/config/
export default defineConfig({
    plugins: [deno(), react(), tailwindcss()],
    server: {
        proxy: {
            "/api": "http://localhost:8080",
            "/auth": "http://localhost:8080",
        },
    },
});
