import $axios from '~/utils/request'

export function login (data) {
  return $axios.$post(
    '/login',
    data
  )
}

export function tagList (params) {
  return $axios.$get(
    '/admin/tags',
    params
  )
}

export function addTag (data) {
  return $axios.$post(
    '/admin/tags',
    data
  )
}

export function updateTag (tagId, data) {
  return $axios.$put(
    '/admin/tags/' + tagId,
    data
  )
}

export function delTag (data) {
  return $axios.$delete(
    '/admin/tags',
    data 
  )
}

export function articleList (params) {
  return $axios.$get(
    '/admin/articles',
    params
  )
}

export function getArticle (articleId) {
  return $axios.$get(
    '/admin/articles/' + articleId
  )
}

export function addArticle (data) {
  return $axios.$post(
    '/admin/articles',
    data
  )
}

export function updateArticle (articleId, data) {
  return $axios.$put(
    '/admin/articles/' + articleId,
    data
  )
}
