import { definePerson } from "nuxt-schema-org/schema";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: "2024-11-01",
	devtools: { enabled: true },
	robots: {
		allow: "*",
	},
	site: {
		url: "https://blog.rhyzzor.com",
		name: "Rhyzzor's Blog",
	},
	schemaOrg: {
		identity: definePerson({
			name: "Rhyzzor",
			alternateName: "Ryan Vieira",
			description: "Software Engineer",
			image: "https://github.com/rhyzzor.png",
			url: "https://blog.rhyzzor.com",
			sameAs: [
				"https://twitter.com/rhyzzor",
				"https://github.com/rhyzzor",
				"https://linkedin.com/in/ryanvsouza",
				"https://rpg.rhyzzor.com",
			],
		}),
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
			author: "Rhyzzor",
			ogType: "website",
			description: "A blog to document my journey",
			themeColor: [
				{ content: "#FFFFFF", media: "(prefers-color-scheme: light)" },
			],
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
