<template>
  <div v-if="!this.$store.state.setting.isMobile" class="blog-tag">
    <span v-for="(tag,key) in tags" :key="key" :title="tag.tag_name" @click="handleClick(tag.tag_id)">
      <img :src="tag.tag_icon">
    </span>
  </div>
</template>

<script>
import { tagList } from '~/apis/tag'
export default {
  name: 'BTag',
  data () {
    return {
      tags: []
    }
  },
  created () {
    tagList().then((res) => {
      this.tags = res.data
    })
  },
  methods: {
    handleClick (tagId) {
      this.$emit('click', tagId)
    }
  }
}
</script>

<style scoped>
  .blog-tag {
    position: fixed;
    bottom: 150px;
    margin-left: 1010px;
    border-radius: 25px;
  }
  .blog-tag span {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 30px;
    height: 30px;
    margin-bottom: 10px;
    border-radius: 50%;
    background: #fff;
    cursor: pointer;
  }
  .blog-tag span img {
    width: 27px;
    height: 25px;
  }
</style>
