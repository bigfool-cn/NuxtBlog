import Vue from 'vue'
import SvgIcon from '~/components/svgicon'

// 注册组件
Vue.component('svg-icon', SvgIcon)
//预请求svg组件(通过之前的svg-sprite-loader加载)
const req = require.context('@/assets/icons/svg', false, /\.svg$/)

const requireAll = requireContext => requireContext.keys().map(requireContext)

requireAll(req)