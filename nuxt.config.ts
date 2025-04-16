// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: "2024-11-01",
	devtools: { enabled: true },
	site: {
		url: "https://blog.rhyzzor.com",
		name: "Rhyzzor's Blog",
	},
	schemaOrg: {
		identity: "Person",
	},
	app: {
		head: {
			templateParams: {
				separator: " | ",
			},
		},
	},
	seo: {
		meta: {
			description: "A blog to document my journey",
			themeColor: [{ content: "white", color: "#ffffff" }],
		},
	},
	content: {
		renderer: {
			anchorLinks: false,
		},
		build: {
			markdown: {
				highlight: {
					langs: ["js", "ts", "go", "python", "bash", "cpp", "json"],
					theme: "nord",
				},
			},
		},
	},
	fonts: {
		families: [
			{
				name: "JetBrains Mono",
				weight: 300,
				provider: "google",
				global: true,
				preload: true,
			},
		],
	},
	image: {
		format: ["png"],
	},
	css: ["@/assets/css/main.css"],
	modules: ["@nuxtjs/seo", "@nuxt/content", "@nuxt/fonts", "@nuxt/image"],
});
