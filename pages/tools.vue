<template>
  <div>
    <div class="tools">
      <tools-tags :tag="activeTag" @change="changeTag"/>
      <div class="tools-comp">
        <nuxt-child/>
      </div>
    </div>
    <blog-back/>
    <blog-eraser/>
  </div>
</template>

<script>
  import ToolsTags from '~/pages/tools/components/tags.vue'
  import BlogEraser from '~/components/eraser/index.vue'
  import BlogBack from '~/components/back/index.vue'

  export default {
    name: 'Tools',
    components: {
      BlogEraser,
      ToolsTags,
      BlogBack
    },
    head() {
      return {
        title: '在线工具 - 代码写得再六又怎么样...',
        meta: [
          { hid: 'keywords', name: "keywords", content: "JSON格式器 XML格式器 UNICODE转化器 URL解编码器 BASE64解编码器 图片BASE64生成器 MD5生成器 正则表达式 Postwoman" },
          { hid: 'description', name: 'description', content: '程序员在线工具' }
        ],
      }
    },
    data() {
      return {
        activeTag: 'json'
      }
    },
    watch:{
      $route(route) {
        const tag = route.params.tag ? route.params.tag : this.activeTag 
        this.$store.dispatch('settings/setToolActiveTag', tag)
      }
    },
    mounted() {
      const tag = this.$route.params.tag ? this.$route.params.tag : this.activeTag 
      this.$store.dispatch('settings/setToolActiveTag', tag)
    },
    methods: {
      changeTag(tag) {
        this.activeTag = tag
      }
    }
  }

</script>

<style scoped>
  .tools {
    padding: 10px 40px;
  }

  .tools-comp {
    padding: 15px 0;
  }

</style>
