const path = require('path');

module.exports = {
  mode: 'universal',
  srcDir: 'src',
  server: {
    port: 3000,
    host: '0.0.0.0'
  },
  /*
  ** Headers of the page
  */
  head: {
    title: 'northernlife',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
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
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxt/typescript-build',
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
  ],
  /*
  ** Axios module configuration
  ** See https://axios.nuxtjs.org/options
  */
  axios: {
    proxy: true,
  },
  proxy: {
   '/api/' : {target: (process.env.API_BASE_URL) ? `${process.env.API_BASE_URL}/api/v1/` : "http://192.168.3.5:9000/api/v1/", pathRewrite: {'^/api/': ''}}
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
    transpile: ['@stylelib']
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
