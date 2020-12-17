<template>
  <div>
    <ul>
      <li v-for="(item,key) in tags" :key="key" @click="clickTag(key)">
        <span :class="activeTag === key ? 'active' : ''">{{ item }}</span>
      </li>
      <li @click="openUrl">
        <span>Postwoman</span>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    name: 'ToolsTags',
    data() {
      return {
        tags: {}
      }
    },
    computed: {
      activeTag() {
        return this.$store.state.settings.toolActiveTag
      }
    },
    created() {
      this.tags = this.$store.state.settings.tags
    },
    methods: {
      clickTag(tag) {
        this.$store.dispatch('settings/setToolActiveTag', tag)
        this.$router.push({path: '/tools/' + tag})
      },
      openUrl() {
        window.open('https://hoppscotch.io/cn')
      }
    }
  }
</script>

<style scoped>
  ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  ul li {
    display: inline-block;
    color: var(--color);
    padding: 5px;
    margin: 10px 30px 5px 0;
    background: var(--bg-tag-color);
    height: 30px;
    line-height: 18px;
    position: relative;
    cursor: pointer;
  }

  .active,
  ul li:hover {
    color: var(--tag-active-color);
  }

  ul li:after {
    content: "";
    position: absolute;
    top: -1px;
    right: -32px;
    border-radius: 2px;
    border: 16px solid transparent;
    border-left: 16px solid var(--bg-tag-color);
  }
</style>
