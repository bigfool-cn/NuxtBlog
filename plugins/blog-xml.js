import Vue from 'vue'
import XmlComponent from '~/components/bxml/index.vue'

const BlogXml = {
  install (Vue) {
    Vue.component('BlogXml', XmlComponent)
  }
}

Vue.use(BlogXml)
