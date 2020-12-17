<template>
  <div class="blog-xml">
    <div class="btn-block">
      <button @click="handleXml">
        格式化
      </button>
      <button @click="handleXml(true)">
        压缩
      </button>
    </div>
    <client-only>
      <codemirror v-model="code" :options="options"/>
    </client-only>
  </div>
</template>

<script>
  import {Message} from 'element-ui'

  export default {
    name: 'BlogXml',
    data() {
      return {
        code: '',
        options: {
          lineNumbers: true,
          matchBrackets: true,
          mode: 'application/xml',
          indentUnit: 4,
          indentWithTabs: true,
          autoCloseTags: true,
          moveOnDrag: false
        }
      }
    },
    methods: {
      handleXml(min) {
        try {
          const formattor = require('formattor')
          let options = {methods: 'xml'}
          if (min) {
            options = {methods: 'xmlmin'}
          }
          this.code = formattor(this.code, options)
        } catch (e) {
          Message.error('处理失败!' + e.toString())
        }
      }
    }
  }
</script>

<style lang="scss" scoped>
  .btn-block {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
  }

  .blog-xml {
    button {
      height: 35px;
      line-height: 32px;
      border-radius: 10px;
      color: var(--color);
      display: block;
      width: 49%;
      border: 1px dotted var(--border-color);
      outline: none;
      background: var(--bg);
      cursor: pointer;

      &:hover {
        background: var(--bg);
        border-color: var(--border-active-color);
      }
    }
  }

  ::v-deep .CodeMirror {
    font-size: 1.1rem;
    height: calc(100vh - 255px);
    border: 1px solid var(--border-color);
    margin: 0;
  }

  ::v-deep .CodeMirror-scroll {
    background: var(--bg);
    overflow: auto !important;
    margin: 0;
    color: var(--color);
  }

  ::v-deep .CodeMirror-vscrollbar {
    display: none !important;
  }

  ::v-deep .CodeMirror-scroll::-webkit-scrollbar {
    width: 7px;
    height: 7px;
    background-color: var(--bg);
  }

  /*定义滚动条轨道 内阴影+圆角*/
  ::v-deep .CodeMirror-scroll::-webkit-scrollbar-track {
    box-shadow: inset 0 0 6px var(--bg);;
    -webkit-box-shadow: inset 0 0 6px var(--bg);;
    border-radius: 10px;
    background-color: var(--bg);
  }

  /*定义滑块 内阴影+圆角*/
  ::v-deep .CodeMirror-scroll::-webkit-scrollbar-thumb {
    border-radius: 10px;
    box-shadow: inset 0 0 6px var(--bg);;
    -webkit-box-shadow: inset 0 0 6px var(--bg);;
    background-color: var(--bg);
  }

  ::v-deep .CodeMirror-gutters {
    background: var(--bg);
    border-right: 1px solid var(--border-color);
    color: var(--color);
  }

  ::v-deep .cm-s-default .cm-tag {
    color: #bd72b9;
  }
</style>
