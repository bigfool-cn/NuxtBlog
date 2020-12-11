<template>
  <div>
    <div class="article">
      <div class="article-header">
        <div class="article-title">
          <h2>{{ article.article_title }}</h2>
        </div>
        <div class="article-tools">
          <span class="article-tag">
            <span v-for="(tag,key) in article.tags" :key="key" :title="tag.tag_name">
                <img :src="tag.tag_icon">
              </span>
          </span>
          <span title="阅读量">{{ article.article_read }}</span>
          <span>{{ article.create_time | parseTime }}</span>
        </div>
      </div>
      <div class="article-content" v-html="article.article_content">
      </div>
    </div>
    <blog-eraser />
  </div>
</template>

<script>
import Vue from 'vue'
import BlogEraser from '~/components/eraser/index.vue'
import { getArticle } from '~/apis/article'
export default {
  name: 'ArticleDetail',
  components: { BlogEraser },
  async asyncData({ context, params }) {
    const { article } = await getArticle(params.id).then(response => {
        if (response.code !== 200) {
          return context.error({ statusCode: 404, message: '页面未找到或无数据' })
        }
        return { article: response.data }
      }).catch((error) => {
        return context.error({ statusCode: 404, message: '页面未找到或无数据' })
      })
    return { article }
  },
  head() {
    return {
      link:[
        { rel: 'stylesheet', href: '/prism/prism.css' }
      ],
      script: [
        { src: '/prism/prism.js' },
        { src: '/prism/clipboard.min.js' }
      ]
    }
  },
  mounted() {
    process.browser && document.querySelectorAll("pre code").forEach(block => Prism.highlightElement(block));
  }
}
</script>

<style lang="scss" scoped>
  .article {
    margin-top: 15px;
    border-top: 1px solid var(--border-color);
    padding: 10px 25px;
  }
  .article-header {
    text-align: center;
    display: flex;
    flex-direction: column;
  }
  .article-title {
    color: var(--color);
    font-weight: bold;
  }
  .article-tools {
    display: flex;
    justify-content: center;
    padding-top: 12px;
    color: var(--color);
    font-size: 14px;
    line-height: 14px;
  }
  .article-tools {
     span {
      position: relative;
      &:nth-child(2) {
        padding: 0 20px;
      }
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
  .article-content {
    padding: 12px 0;
    color: var(--color);
  }
</style>
