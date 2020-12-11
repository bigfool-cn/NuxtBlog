import request from '~/utils/request'

export function getArticleList (params) {
  return request({
    url:'/articles',
    method:'get',
    params
  })
}

export function getArticle (articleId) {
  return request({
    url:'/articles/' + articleId,
    method:'get'
  })
}
