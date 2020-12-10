import $axios from '~/utils/request'

export function tagList (params) {
  return $axios.$get(
    '/tags',
    params
  )
}
