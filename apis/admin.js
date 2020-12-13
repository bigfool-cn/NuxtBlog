import request from '~/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function tagList(params) {
  return request({
    url: '/admin/tags',
    method: 'get',
    params
  })
}

export function addTag(data) {
  return request({
    url: '/admin/tags',
    method: 'post',
    data
  })
}

export function updateTag(tagId, data) {
  return request({
    url: '/admin/tags/' + tagId,
    method: 'put',
    data
  })
}

export function delTag(data) {
  return request({
    url: '/admin/tags',
    method: 'delete',
    data
  })
}

export function articleList(params) {
  return request({
    url: '/admin/articles',
    method: 'get',
    params
  })
}

export function getArticle(articleId) {
  return request({
    url: '/admin/articles/' + articleId,
    method: 'get'
  })
}

export function addArticle(data) {
  return request({
    url: '/admin/articles',
    method: 'post',
    data
  })
}

export function updateArticle(articleId, data) {
  return request({
    url: '/admin/articles/' + articleId,
    method: 'put',
    data
  })
}

export function delArticle(data) {
  return request({
    url: '/admin/articles',
    method: 'delete',
    data
  })
}
