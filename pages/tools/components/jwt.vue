<template>
  <div>
    <div class="label-title">Token:</div>
    <el-input type="textarea" v-model="token" rows="4" resize="none" placeholder="token..." spellcheck="false"/>
    <div class="label-title">Header:</div>
    <el-input type="textarea" v-model="header" rows="8" resize="none" readonly spellcheck="false"/>
    <div class="label-title">Playload:</div>
    <el-input type="textarea" v-model="playload" rows="12" resize="none" readonly spellcheck="false"/>
  </div>
</template>

<script>
  import {Message} from 'element-ui'
  export default {
    name: "ToolJwt",
    data() {
      return {
        token: ''
      }
    },
    computed: {
      header() {
        const tk = this.token.replace('Bearer','').trim();
        const tokenArr = tk.split('.')
        if (tokenArr && tokenArr.length === 3) {
          try {
            return JSON.stringify(JSON.parse( window.atob(tokenArr[0])),undefined,4)
          } catch (e) {
            Message.error('解析失败，请检查token格式')
          }
        } else {
          return ''
        }
      },
      playload() {
        const tk = this.token.replace('Bearer','').trim();
        const tokenArr = tk.split('.')
        if (tokenArr && tokenArr.length === 3) {
          try {
            return JSON.stringify(JSON.parse(window.atob(tokenArr[1])),undefined,4)
          } catch (e) {
            Message.error('解析失败，请检查token格式')
          }
        } else {
          return ''
        }
      }
    }
  }
</script>

<style scoped>
 .label-title {
   margin: 10px 0 3px 0;
   color: var(--color);
 }
</style>
