<template>
  <div class="blog-tag">
    <span v-for="(tag,key) in tags" :key="key" :title="tag.tag_name" @click="handleClick(tag.tag_id)">
      <img :src="tag.tag_icon" :class="tag.tag_id === tagId ? 'active-tag' : ''">
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
    data() {
      return {
        tagId: -1
      }
    },
    methods: {
      handleClick(tagId) {
        this.tagId = tagId
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
    .active-tag {
      animation: turn 3s linear infinite;
    }
  }
  @keyframes turn{
    0%{-webkit-transform:rotate(0deg);}
    25%{-webkit-transform:rotate(90deg);}
    50%{-webkit-transform:rotate(180deg);}
    75%{-webkit-transform:rotate(270deg);}
    100%{-webkit-transform:rotate(360deg);}
  }
</style>
