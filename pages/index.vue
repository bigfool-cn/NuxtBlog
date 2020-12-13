<template>
  <div>
    <div ref="entry" class="articles">
      <div v-if="!articles || !articles.length" class="empty-data">该分类暂无文章...</div>
      <nuxt-link v-for="(item,key) in articles" :key="key" :to="'/article/' + item.article_id" class="article-to">
        <div class="li-article">
          <div class="article-title">
            {{ item.article_title }}
          </div>
          <div class="article-tools">
            <span class="article-tag">
              <span v-for="(tag,key) in item.tags" :key="key" :title="tag.tag_name">
                <img :src="tag.tag_icon">
              </span>
            </span>
            <span title="阅读量">{{ item.article_read }}</span>
            <span>{{ item.create_time | parseTime('{y}-{m}-{d}') }}</span>
          </div>
        </div>
      </nuxt-link>
    </div>
    <b-tag :tags="tags" @click="clickTag"/>
    <blog-eraser/>
  </div>
</template>

<script>
  import BlogEraser from '~/components/eraser/index.vue'
  import BTag from '~/components/btag/index.vue'
  import {getArticleList} from '~/apis/article'
  import {tagList} from '~/apis/tag'
  import {Message} from 'element-ui'
  export default {
    components: {
      BlogEraser,
      BTag
    },
    async asyncData(context) {
      const [articlesItem, tags] = await Promise.all([
        getArticleList().then((response) => {
          if (response.code !== 200) {
            return context.error({statusCode: 404, message: '页面未找到或无数据'})
          }
          return {
            articles: response.data.articles,
            total: response.data.total
          }
        }).catch((error) => {
          return context.error({statusCode: 404, message: '页面未找到或无数据'})
        }),
        tagList().then(response => {
          if (response.code !== 200) {
            return context.error({statusCode: 404, message: '页面未找到或无数据'})
          }
          return response.data
        }).catch((error) => {
          return context.error({statusCode: 404, message: '页面未找到或无数据'})
        })
      ])
      return {
        articles: articlesItem.articles,
        total: articlesItem.total,
        tags: tags
      }
    },
    data() {
      return {
        query: {
          page: 1,
          limit: 20,
          tag_id: -1
        },
        timer: undefined
      }
    },
    mounted() {
      window.addEventListener('scroll', this.handleScroll)
    },
    methods: {
      clickTag(tagId) {
        this.query.page = 1
        this.query.tag_id = tagId
        getArticleList(this.query).then((response) => {
          if (response.code !== 200) {
            return Message.error('加载数据失败')
          }
          this.articles = response.data.articles
          this.total = response.data.total
        })
      },
      handleScroll() {
        this.timer && clearTimeout(this.timer)
        this.timer = setTimeout(this.loadMoreData, 300)
      },
      loadMoreData() {
        return new Promise((resolve) => {
          const $el = document.documentElement
          const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
          const $entry = this.$refs.entry
          const clienHeight = $el.clientHeight
          const style = window.getComputedStyle ? window.getComputedStyle($entry, null) : null || $entry.currentStyle
          const containerHeight = ~~style.height.split('px')[0]
          // 滚动到一定高度，重新加载数据
          if (scrollTop + clienHeight > containerHeight - 50) {
            if (this.total / this.query.limit > this.query.page) {
              this.query.page = this.query.page + 1
              getArticleList(this.query).then((response) => {
                if (response.code !== 200) {
                  return Message.error('加载数据失败')
                }
                this.articles = this.articles.concat(response.data.articles)
                this.total = response.data.total
              })
            }
            resolve()
          }
        })
      }
    }
  }

</script>

<style lang="scss" scoped>
  .articles {
    margin: 10px 15px 0 25px;
    border-top: 1px solid var(--border-color);
    border-bottom: 1px solid var(--border-color);
    padding: 2px 0;
  }

  .empty-data {
    color: var(--color);
    text-align: center;
    padding: 15px;
    border-bottom: 1px solid var(--border-color);
  }

  .article-to {
    text-decoration: none;
  }

  .li-article {
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-top: 1px solid var(--border-color);
    height: 80px;
    padding: 5px 10px;
    margin-bottom: 1px;
  }

  .li-article:hover {
    .article-title {
      font-weight: bold;
      font-style: italic;
    }
  }

  .li-article:before {
    content: "";
    position: absolute;
    left: -22px;
    top: 30px;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: var(--bg);
  }

  .articles .article-to:last-child .li-article {
    border-bottom: 1px solid var(--border-color);
  }

  .article-title {
    color: var(--color);
    font-size: 16px;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden
  }

  .article-tools {
    display: flex;
    justify-content: flex-start;
    color: var(--color);
    font-size: 14px;
    line-height: 14px;
  }

  .article-tools > span {
    position: relative;
    &:nth-child(2) {
      padding: 0 20px;
    }
  }

  .article-tag {
    span {
      padding: 2px;
      img {
        width: 17px;
        height: 15px;
      }
    }
  }

</style>
