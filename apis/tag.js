import request from '~/utils/request'

export function tagList (params) {
  return request({
    url:'/tags',
    params
  })
}
