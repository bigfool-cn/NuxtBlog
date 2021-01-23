const { resolve } = require('path')
export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: '代码写得再六又怎么样...',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'keywords', name: "keywords", content: "博客 IT博客 技术博客 程序员博客 程序员技术 技术文章  在线工具 PHP Go MySQL Vue JavaScript Css Linux" },
      { hid: 'description', name: 'description', content: '代码敲得再六又怎么样, 连心爱的人都留不住.' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
    // 'element-ui/lib/theme-chalk/index.css'
    '~/assets/css/main.scss',
    '~/assets/css/themes.css',
    'codemirror/lib/codemirror.css',
    // merge css
    'codemirror/addon/merge/merge.css',
    // theme css
    'codemirror/theme/base16-dark.css'
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
     { src: '~/plugins/blog-filters' },
    { src: '~/plugins/code-mirror', ssr: false },
    { src: '~/plugins/blog-xml', ssr: false },
    { src: '~/plugins/blog-json', ssr: false },
    { src: '~/plugins/element-ui', ssr: false },
    { src: '~/plugins/drag', ssr: false },
    { src: '~/plugins/blog-router' },
    { src: '~/plugins/svg-icon' }
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    '@nuxtjs/color-mode'
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    'cookie-universal-nuxt'
  ],

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
    babel: {
      "plugins": [
        [
          "component",
          {
            "libraryName": "element-ui",
            "styleLibraryName": "theme-chalk"
          }
        ]
      ]
    },
    extend (config,ctx) {
      // 排除 nuxt 原配置的影响,Nuxt 默认有vue-loader,会处理svg,img等
      // 找到匹配.svg的规则,然后将存放svg文件的目录排除
      const svgRule = config.module.rules.find(rule => rule.test.test('.svg'))
      svgRule.exclude = [resolve(__dirname, 'assets/icons/svg')]
      //添加loader规则
      config.module.rules.push({
        test: /\.svg$/, //匹配.svg
        include: [resolve(__dirname, 'assets/icons/svg')], //将存放svg的目录加入到loader处理目录
        use: [{ loader: 'svg-sprite-loader',options: {symbolId: 'icon-[name]'}}]
      })
    }
  },

  colorMode: {
    classSuffix: "",
    preference: "dark",
    fallback: "dark",
  }
}
