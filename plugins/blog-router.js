export default ({ app }) => {
  app.router.beforeEach((to,from,next) => {
    const path = to.path
    const token = app.$cookies.get('_token_')
    if (/^\/admin\/|^\/admin$/.test(path) && (!token || token === 'null' || token === 'undefined')) {
      // app.router.push({ path: '/' })
    }
    next()  
  })
}
