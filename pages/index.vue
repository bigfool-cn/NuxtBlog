<template>
  <div>
    <div ref="entry" class="articles">
      <div v-if="!articles.length" class="empty-data">该分类暂无文章...</div>
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
            <span>{{ item.create_time | parseTime }}</span>
          </div>
        </div>
      </nuxt-link>
    </div>
    <blog-eraser />
  </div>
</template>

<script>
import { mapState } from 'vuex'
import BlogEraser from '~/components/eraser/index.vue'
export default {
  components: { BlogEraser },
  computed: {
    ...mapState({
      isMobile: state => state.settings.isMobile,
    })
  },
  data() {
    return {
      articles:[
        {
          article_id: 1,
          article_title: '哈哈哈哈',
          article_read: 10,
          create_time: '2020-12-09 15:15:15',
          tags: []
        }
      ]
    }
  }
}
</script>

<style scoped>
.articles {
    margin: 10px 15px 0 25px;
    border-top:1px solid var(--border-color);
    border-bottom:1px solid var(--border-color);
    padding: 2px 0;
  }
  .empty-data {
    color: var(--color);
    text-align: center;
    padding: 15px;
    border-bottom:1px solid var(--border-color);
  }
  .article-to {
    text-decoration: none;
  }
  .li-article {
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-top:1px solid var(--border-color);
    height: 80px;
    padding: 5px 10px;
    margin-bottom: 1px;
  }
  .li-article:hover {
    font-style: italic;
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
    border-bottom:1px solid var(--border-color);
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
    margin-right: 12px;
  }

  .article-tools > span:after {
    content: "/";
    font-weight: bold;
    position: absolute;
    margin-left: 3px;
  }
  .article-tools > span:last-child:after {
    content: "";
  }
  .article-tag > span {
    padding: 2px;
  }

  .article-tag > span  > img {
    width: 17px;
    height: 15px;
  }
</style>
