<template>
  <div v-if="!this.$store.state.settings.isMobile" class="blog-tag">
    <span v-for="(tag,key) in tags" :key="key" :title="tag.tag_name" @click="handleClick(tag.tag_id)">
      <img :src="tag.tag_icon">
    </span>
  </div>
</template>

<script>
export default {
  name: 'BTag',
  props: {
    tags: {
      type: Array,
      required: true
    }
  },
  methods: {
    handleClick (tagId) {
      this.$emit('click', tagId)
    }
  }
}
</script>

<style lang="scss" scoped>    
  .blog-tag {
    position: fixed;
    bottom: 180px;
    margin-left: 1010px;
    border-radius: 25px;
    span {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 30px;
      height: 30px;
      margin-bottom: 10px;
      border-radius: 50%;
      background-color: var(--bg-secondary);
      cursor: pointer;
      img {
        width: 27px;
        height: 25px;
      }
      &:hover {
        background-color: var(--border-active-color);
      }
    }
  }
</style>
