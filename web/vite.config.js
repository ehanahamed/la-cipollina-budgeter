import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
    server: {
		port: process?.env?.PORT ?? 8080,
		proxy: {
			"/api/": {
				target: process?.env?.API_URL ?? "http://localhost:3000",
				changeOrigin: true,
				rewrite: function (path) {
					return path.replace(/^\/api/, "");
				}
			},
		}
	},
	preview: {
		port: process?.env?.PORT ?? 8080,
		proxy: {
			"/api/": {
				target: process?.env?.API_URL ?? "http://localhost:3000",
				changeOrigin: true,
				rewrite: function (path) {
					return path.replace(/^\/api/, "");
				}
			},
		}
	}
});
