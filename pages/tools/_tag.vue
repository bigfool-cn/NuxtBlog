<template>
  <div>
    <client-only>
      <blog-json v-if="activeTag === 'json'"/>
      <blog-xml v-if="activeTag === 'xml'"/>
    </client-only>
    <tools-url v-if="activeTag === 'url'"/>
    <tools-base64 v-if="activeTag === 'base64'"/>
    <tools-unicode v-if="activeTag === 'unicode'"/>
    <tools-img-base64 v-if="activeTag === 'img_base64'"/>
    <tools-md5 v-if="activeTag === 'md5'"/>
    <tools-regex v-if="activeTag === 'regex'"/>
  </div>
</template>

<script>
  import ToolsUrl from './components/url.vue'
  import ToolsBase64 from './components/base64.vue'
  import ToolsUnicode from './components/unicode.vue'
  import ToolsImgBase64 from './components/imgBase64.vue'
  import ToolsMd5 from './components/md5.vue'
  import ToolsRegex from './components/regex.vue'

  export default {
    name: 'Tool',
    components: {ToolsUrl, ToolsBase64, ToolsUnicode, ToolsImgBase64, ToolsMd5, ToolsRegex},
    computed: {
      activeTag() {
        return this.$store.state.settings.toolActiveTag
      }
    },
    mounted() {
      if (!this.$store.state.settings.tags[this.$route.params.tag]) {
        return this.$router.push({path: '/tools'})
      }
      this.$store.dispatch('settings/setToolActiveTag', this.$route.params.tag)
    }
  }
</script>

<style scoped>

</style>
