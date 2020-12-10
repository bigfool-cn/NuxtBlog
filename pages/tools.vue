<template>
  <div>
    <div class="tools">
      <tools-tags :tag="activeTag" @change="changeTag" />
      <div class="tools-comp">
        <nuxt-child />
      </div>
    </div>
    <blog-eraser />
  </div>
</template>

<script>
  import ToolsTags from '~/pages/tools/components/tags.vue'
  import BlogEraser from '~/components/eraser/index.vue'
  export default {
    name: 'Tools',
    components: {
      BlogEraser,
      ToolsTags
    },
    data() {
      return {
        activeTag: 'json'
      }
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
