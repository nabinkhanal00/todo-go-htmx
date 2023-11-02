import { resolve } from "path"
import { defineConfig } from "vite"

export default defineConfig({
	build: {
		lib: {
			entry: [resolve(__dirname, "web/app.js")],
			formats: ["es"],
			name: "[name]",
			fileName: "[name]",
		},
		outDir: "web/static/js",
		emptyOutDir: false
	}
})
