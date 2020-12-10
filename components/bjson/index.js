import JsonComponent from './index.vue'

const blogJson = {
  install (Vue) {
    Vue.component('BlogJson', JsonComponent)
  }
}
export default blogJson
