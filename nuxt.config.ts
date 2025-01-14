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
      meta: [
        {
          name: "theme-color",
          content: "#FAFAFA",
          media: "(prefers-color-scheme: light)",
        },
        {
          name: "theme-color",
          content: "#000000",
          media: "(prefers-color-scheme: dark)",
        },
      ],
      title: "Kyle Huggins",
    },
  },
  compatibilityDate: "2024-11-01",
  css: ["~/assets/css/main.css"],
  devtools: { enabled: true },
  googleFonts: {
    families: {
      "JetBrains Mono": [400, 700],
      Inter: [400, 600, 700],
    },
  },
  modules: ["@nuxtjs/google-fonts", "@nuxtjs/tailwindcss"],
  nitro: {
    prerender: {
      autoSubfolderIndex: false,
    },
    static: true,
  },
  ssr: false,
});
