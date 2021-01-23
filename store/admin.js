
const state = () => ({
  token: '',
})

const mutations = {
  SET_TOKEN (state,token) {
    state.token = token
  }
}

const actions = {
  inrLimit ({ commit }, limit = 1) {
    commit('INR_LIMIT', limit)
  },

  setToken ({ commit }, token) {
    commit('SET_TOKEN', token)
    
  },
  logout ({ commit }) {
    commit('SET_TOKEN', '')
  }
}


export default {
  namespaced: true,
  state,
  mutations,
  actions
}