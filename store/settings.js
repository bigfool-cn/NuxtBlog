const state = () => ({
  shareTitle: '代码敲得再六又怎么样',
  isMobile: false,
  toolActiveTag: 'json',
  tags: {
    json: 'JSON格式器',
    xml: 'XML格式器',
    unicode: 'UNICODE转化器',
    url: 'URL解编码器',
    base64: 'BASE64解编码器',
    img_base64: '图片BASE64生成器',
    md5: 'MD5生成器',
    regex: '正则表达式'
  }
})

const mutations = {
  SET_TITLE (state,title) {
    state.shareTitle = title
  },
  SET_ISMOBILE (state,flag) {
    state.isMobile = flag
  },
  SET_TOOLACTIVETIONTAG(state,tag) {
     state.toolActiveTag = tag
  }
}

const actions = {
  setShareTitle ({ commit }, title) {
    commit('SET_TITLE', title)
  },
  setIsMobile ({ commit }, flag) {
    commit('SET_ISMOBILE', flag)
  },
  setToolActiveTag ({ commit }, tag) {
    commit('SET_TOOLACTIVETIONTAG', tag)
  }
}


export default {
  namespaced: true,
  state,
  mutations,
  actions
}