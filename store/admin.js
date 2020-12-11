

const state = () => ({
  token: '',
  limit: 0 
})

const mutations = {
  INR_LIMIT (state) {
    state.limit += 1
  },
  SET_TOKEN (state,token) {
    state.token = token
  }
}

const actions = {
  inrLimit ({ commit }) {
    commit('INR_LIMIT')
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