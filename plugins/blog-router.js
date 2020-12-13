export default (context) => {
  const app = context.app
  app.router.beforeEach((to, from, next) => {
    const path = to.path
    const token = context.$cookies.get('_token_')
    if (/^\/admin\/|^\/admin$/.test(path) && (!token || token === 'null' || token === 'undefined')) {
      context.error({statusCode: 403, message: '没有权限'})
    }
    next()
  })
}
