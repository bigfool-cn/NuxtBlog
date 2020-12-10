import $axios from '~/utils/request'

export function getArticleList (params) {
  return $axios.$get(
    '/articles',
    params
  )
}

export function getArticle (articleId) {
  return $axios.$get(
    '/articles/' + articleId
  )
}
