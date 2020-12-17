import axios from 'axios'
import Vue from 'vue'
// create an axios instance
const service = axios.create({
  baseURL: 'https://www.bigfool.cn/api', // url = base url + request url
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    if (process.client && /^\/admin\//.test(config.url)) {
      config.headers['Authorization'] = $nuxt.context.$cookies.get('_token_')
    }
    return config
  },
  error => {
    // do something with request error
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    return res
  },
  error => {
    return Promise.reject(error)
  }
)

export default service
