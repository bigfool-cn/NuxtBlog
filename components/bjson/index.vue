<template>
  <div class="blog-json">
    <vue-json-editor v-model="json" :mode="mode" :modes="modes" @has-error="handleError"/>
  </div>
</template>

<script>
  import VueJsonEditor from 'vue-json-editor'
  import {Message} from 'element-ui'

  export default {
    name: 'BlogJson',
    components: {
      VueJsonEditor
    },
    data() {
      return {
        slider: 450,
        mode: 'code',
        modes: ['code'],
        json: ''
      }
    },
    watch: {
      slider(value) {
        const ele = window.document.getElementsByClassName('jsoneditor-vue')[0]
        ele.style.height = value + 'px'
      }
    },
    mounted() {
      const formatBtn = window.document.getElementsByClassName('jsoneditor-format')[0]
      formatBtn.textContent = '格式化'

      const compactBtn = window.document.getElementsByClassName('jsoneditor-compact')[0]
      compactBtn.textContent = '压缩'

      const modesBtn = window.document.getElementsByClassName('jsoneditor-modes')[0]
      modesBtn.style.display = 'none'

      const poweredBy = window.document.getElementsByClassName('jsoneditor-poweredBy')[0]
      poweredBy.style.display = 'none'
    },
    methods: {
      handleError(error) {
        Message.error(error)
      }
    }
  }
</script>

<style lang="scss" scoped>
  ::v-deep .jsoneditor-format,
  ::v-deep .jsoneditor-compact {
    outline: none;
  }

  ::v-deep .jsoneditor-menu {
    background: var(--bg-secondary);
    border-color: var(--border-color);
  }

  ::v-deep .jsoneditor-menu button {
    width: 55px;
    line-height: 26px;
    background: var(--bg-secondary);
    color: var(--color);
  }

  ::v-deep .ace-jsoneditor .ace_gutter, ::v-deep .ace-jsoneditor .ace_variable {
    background: var(--bg-secondary);
    color: var(--color);
  }

  ::v-deep .ace-jsoneditor .ace_gutter {
    border-right: 1px solid var(--border-color);
  }

  ::v-deep .ace_gutter-cell {
    background: var(--bg);
  }

  ::v-deep .jsoneditor {
    border: 1px solid var(--border-color);
    box-sizing: border-box;
  }

  ::v-deep .jsoneditor-vue {
    height: calc(100vh - 210px);
  }

  ::v-deep .ace_content {
    background: var(--bg);
    color: var(--color);
  }

  ::v-deep .ace_active-line {
    background: var(--bg-secondary) !important;
  }
</style>
