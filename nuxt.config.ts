// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxt/content',
  ],
  routeRules: {
    '/': { prerender: true }
  },
  content: {
    documentDriven: true,
  },
  app: {
    head: {
      title: "Rhyzzor's Blog",
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      ],
    },
  },
})