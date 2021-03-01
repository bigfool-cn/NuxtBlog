<template>
  <div>
    <el-input type="textarea" id="regex-context" v-model="regexContent" rows="10" placeholder="待匹配内容..." spellcheck="false"/>
    <el-input type="textarea" id="regex" v-model="regex" rows="2" placeholder="正则表达式..." spellcheck="false"/>
    <el-button type="button" @click="handleMatch">匹 配</el-button>
    <el-input type="textarea" id="regex-result" v-model="regexResult" rows="10" placeholder="匹配结果..."/>
  </div>
</template>

<script>
  import {Message} from 'element-ui'
  export default {
    name: 'ToolsRegex',
    data() {
      return {
        regexContent: '',
        regex: '',
        regexResult: ''
      }
    },
    methods: {
      handleMatch() {
        if (this.regexContent === null || this.regexContent.length < 1) {
          Message.error('请输入带匹配内容!')
          return
        }
        if (this.regex === null || this.regex.length < 1) {
          Message.error('请输入正则表达式!')
          return
        }
        try {
          const reg = new RegExp(eval(this.regex))
          const res = this.regexContent.match(reg)
          if (res === null || res.length < 1) {
            this.regexResult = '没有匹配'
          } else {
            let resStr = ''
            if (this.regex.includes('g')) {
              resStr = '共找到 ' + res.length + ' 处匹配：\r\n'
              for (let i = 0; i < res.length; ++i) {
                resStr = resStr + res[i] + '\r\n'
              }
            } else {
              resStr = '匹配位置：' + res.index + '\r\n匹配结果：' + res[0]
            }
            this.regexResult = resStr
          }
        } catch (e) {
          Message.error('匹配失败，请稍后重试!')
          return
        }
      }
    }
  }
</script>

<style scoped>
  #regex {
    margin: 10px 0;
  }

  .el-textarea {
    margin-bottom: 20px;
  }

  button {
    height: 35px;
    line-height: 10px;
    width: 100%;
    border-radius: 10px;
    background: var(--bg);
    border: 1px dotted var(--border-color);
    color: var(--color);
    outline: none;
    cursor: pointer;
    margin-bottom: 15px;
  }

  button:hover {
    background: var(--bg);
    border-color: var(--border-active-color);
  }
</style>
