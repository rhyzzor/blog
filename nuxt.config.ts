// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: "2024-11-01",
	devtools: { enabled: true },
  fonts: {
    families: [
      { name: 'JetBrains Mono', weight: 300, provider: 'google', global: true, preload: true  },
    ],
  },
  css: ["@/assets/css/main.css"],
	modules: ["@nuxt/content", "@nuxt/fonts", "@nuxt/image"],
});
