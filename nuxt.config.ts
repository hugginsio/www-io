// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  app: {
    head: {
      link: [
        {
          rel: "icon",
          type: "image/png",
          href: "/favicon-96x96.png",
          sizes: "96x96",
        },
        { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" },
        { rel: "shortcut icon", type: "image/x-icon", href: "/favicon.ico" },
      ],
    },
  },
  compatibilityDate: "2024-11-01",
  css: ["~/assets/css/main.css"],
  devtools: { enabled: true },
  googleFonts: {
    families: {
      Inter: [400, 700, 900],
      "JetBrains Mono": [400, 700],
    },
  },
  modules: ["@nuxt/ui", "@nuxtjs/google-fonts"],
  nitro: {
    prerender: {
      autoSubfolderIndex: false,
    },
    static: true,
  },
  ssr: false,
});
