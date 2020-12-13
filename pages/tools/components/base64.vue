<template>
  <div>
    <el-input
      v-model="input"
      type="textarea"
      rows="10"
      class="input"
      spellcheck="false"
      placeholder="待转化内容..."
      @input="changeInput"
    />
    <div class="direction">
      <svg-icon icon-class="direction"/>
    </div>
    <el-input
      v-model="output"
      type="textarea"
      rows="10"
      class="output"
      spellcheck="false"
      placeholder="转化内容..."
      @input="changeOutput"
    />
  </div>
</template>

<script>
  import {Message} from 'element-ui'
  export default {
    name: 'ToolsBase64',
    data() {
      return {
        input: '',
        output: ''
      }
    },
    methods: {
      changeInput() {
        try {
          this.output = window.atob(this.input)
        } catch (e) {
          Message.error('解码失败!' + e.toString())
          return
        }
      },
      changeOutput() {
        try {
          this.input = window.btoa(this.output)
        } catch (e) {
          Message.error('编码码失败!' + e.toString())
          return
        }
      }
    }
  }
</script>

<style scoped>

</style>
