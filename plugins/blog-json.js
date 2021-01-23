import Vue from 'vue'
import JsonComponent from '~/components/bjson/index.vue'

const BlogJson = {
  install (Vue) {
    Vue.component('BlogJson', JsonComponent)
  }
}
Vue.use(BlogJson)
