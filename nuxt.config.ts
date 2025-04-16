// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: "2024-11-01",
    devtools: { enabled: true },
  app:  {
    head: {
      title: "Rhyzzor's Blog",
      htmlAttrs: {
        lang: 'en'
      },
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      ]
    }
  },
  content: {
    renderer: {
      anchorLinks: false,
    },
    build: {
      markdown: {
        highlight: {
          langs: ['js', 'ts', 'go', 'python', 'bash', 'cpp', "json"],
          theme: "nord"
        }
      }
    }
  },
  fonts: {
    families: [
      { name: 'JetBrains Mono', weight: 300, provider: 'google', global: true, preload: true  },
    ],
  },
  css: ["@/assets/css/main.css"],
    modules: ['@nuxtjs/seo', "@nuxt/content", "@nuxt/fonts", "@nuxt/image"],
});