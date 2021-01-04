import axios from 'axios'

// create an axios instance
const service = axios.create({
  baseURL: 'https://www.bigfool.cn/api', // 'http://127.0.0.1:3003'
  timeout: 15000 // request timeout
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
