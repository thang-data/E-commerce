const createSitemapRoutes = async () => {
  let routes = [];
  const { $content } = require('@nuxt/content')
  const articles = await $content('articles', {deep: true}).fetch();
  for (const post of articles) {
    routes.push(`${post.path.replace("articles/", "")}`);
  }
  return routes;
}

export default {

  // Target: https://go.nuxtjs.dev/config-target
  target: "static",

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "nuxt-blog",
    htmlAttrs: {
      lang: "vn",
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "Nothing is impossible without trying" },
      { name: "format-detection", content: "telephone=no" },
      { hid: "og:type", name: "og:type", content: "website" },
      { hid: "og:title", name: "og:title", content: "Data-JS" },
      { hid: "og:description", name: "og:description", content: "Nothing is impossible without trying" },
      { hid: "og:url", name: "og:url", content: "https://nguyenducthang.name.vn/" },
      { name: "og:site_name", content: "Data-JS" },
      { name: "og:locale", content: "VN" },
      { name: "twitter:creator", content: "@nguyenducthang" },
      { name: "twitter:site", content: "@nguyenducthang" },
   ],
    link: [
      { rel: "alternate", type: "application/rss+xml", href: "https://nguyenducthang.name.vn/feed/" },
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" }

    ],
  },

  router: {
    trailingSlash: false,
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["@/assets/css/main.css", "@/assets/scss/main.scss"],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: ["@nuxt/postcss8", "@nuxt/image"],
  vite: {
    /* options for vite */
    // ssr: true // enable unstable server-side rendering for development (false by default)
    // experimentWarning: false // hide experimental warning message (disabled by default for tests)
    vue: {
      /* options for vite-plugin-vue2 */
    },
  },

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: ["@nuxt/content", "@nuxtjs/svg", "@nuxt/image", '@nuxtjs/redirect-module',
    '@nuxtjs/feed',
    '@nuxtjs/sitemap'],
  sitemap: {
    hostname: 'https://nguyenducthang.name.vn',
    gzip: true,
    routes: createSitemapRoutes,
    exclude: ["/blog"]
  },
  feed: [
    // A default feed configuration object
    {
      path: 'rss.xml', // The route to your feed.
      // eslint-disable-next-line require-await
      async create(feed) {
        const baseUrlArticles = 'https://nguyenducthang.name.vn';
        const baseLinkFeedArticles = '/feed/'

        feed.options = {
          title: 'Data-JS',
          description: "Nothing is impossible without trying",
          link: baseUrlArticles,
        }

        const { $content } = require('@nuxt/content');

        const articles = await $content('articles', {deep: true}).sortBy("date", 'desc').fetch()

        articles.forEach((article) => {
          const url = `${baseUrlArticles}${article.path.replace("articles/", "")}`

          feed.addItem({
            title: article.title,
            id: url,
            link: url,
            date: new Date(article.date),
            description: article.description,
            author: 'Nguyen Duc Thang'
          })
        })

      }, // The create function (see below)
      cacheTime: 1000 * 60 * 15, // How long should the feed be cached
      type: 'rss2', // Can be: rss2, atom1, json1
      data: []
    },
  ],
  redirect: [
    {
      from: '(?!^\/$|^\/[?].*$)(.*\/[?](.*)$|.*\/$)',
      to: (from, req) => {
        const base = req._parsedUrl.pathname.replace(/\/$/, '');
        const search = req._parsedUrl.search;
        return base + (search != null ? search : '');
      }
    }
  ],
  svg: {
    vueSvgLoader: {
      // vue-svg-loader options
    },
    svgSpriteLoader: {
      // svg-sprite-loader options
    },
    fileLoader: {
      // file-loader options
    },
  },

  content: {
    liveEdit: false,
    markdown: {
      prism: {
        theme: "prism-themes/themes/prism-dracula.css",
      },
    },
  },
  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    transpile: [
      "three"
    ],
    postcss: {
      plugins: {
        tailwindcss: {},
        autoprefixer: {},
      },
    }

  },
  hooks: {
    'content:file:beforeInsert': (document) => {
      const removeMd = require('remove-markdown');
      const stats = require('reading-time')(document.text);
      if (document.extension === '.md') {
        document.bodyPlainText = removeMd(document.text);
        document.readTime = stats;
      }
    },
  },
};
