const path = require('path');

module.exports = {
  srcDir: 'src',
  server: {
    port: 3000,
    host: '0.0.0.0'
  },
  /*
  ** Headers of the page
  */
  head: {
    title: 'NorthernLife',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' },

      // OGP
      { hid: 'description', name: 'description', content: 'Northernlife' },
      { hid: 'og:site_name', property: 'og:site_name', content: 'Northernlife' },
      { hid: 'og:type', property: 'og:type', content: 'website' },
      { hid: 'og:url', property: 'og:url', content: 'https://northernlife.net' },
      { hid: 'og:title', property: 'og:title', content: 'Northernlife' },
      { hid: 'og:description', property: 'og:description', content: 'Northernlife' },
      { hid: 'og:image', property: 'og:image', content: 'https://example.com/img/ogp/common.jpg' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/northernlife.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    "@/plugins/axios"
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxt/typescript-build',
    '@nuxtjs/composition-api'
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    'bootstrap-vue/nuxt',
    '@nuxtjs/auth'
  ],
  /*
  ** Axios module configuration
  ** See https://axios.nuxtjs.org/options
  */
  axios: {
    proxy: true,
  },
  proxy: {
   '/api' : {target: (process.env.API_BASE_URL) ? `${process.env.API_BASE_URL}/api/v1` : 'http://localhost:9000/api/v1', pathRewrite: {'^/api': '/'}}
  },

  // Auth module to handle login
  auth: {
    redirect: {
      login: '/admin/login',
      logout: '/admin/login',
      callback: false,
      home: '/'
    },
    strategies: {
      local : {
        endpoints: {
          login: {url: '/api/auth/login', method: 'post', propertyName: 'token'},
          logout: {url: '/api/auth/logout', method: 'post'},
          refresh: {url: '/api/auth/refresh', method: 'get', propertyName: 'token'},
          user: false
        }
      }
    }
  },
  /*
  ** Build configuration
  */
  build: {
    extend(config, ctx) {
      // Added Line
      config.devtool = ctx.isClient ? 'eval-source-map' : 'inline-source-map'

      if (ctx.isDev && ctx.isClient) {
      }
    },
    transpile: ['@stylelib'],
    babel: {compact: true}
  },
  resolve: {
    extensions: ['.js', '.json', '.vue', '.ts'],
    root: path.resolve(__dirname),
    alias: {
      '@': path.resolve(__dirname),
      '~': path.resolve(__dirname),
    },
  }
}
