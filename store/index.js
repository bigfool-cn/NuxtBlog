const actions = {

  async nuxtServerInit({ dispatch  }, { app }) {
    const token = app.$cookies.get('_token_')
    if (token) {
      await dispatch ('admin/setToken', token)
    }
  }
}

export default {
  actions
}
