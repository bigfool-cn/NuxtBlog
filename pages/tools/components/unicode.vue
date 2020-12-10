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
      <svg-icon icon-class="direction" />
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
export default {
  name: 'ToolsUnicode',
  data () {
    return {
      input: '',
      output: ''
    }
  },
  methods: {
    changeInput () {
      try {
        this.output = eval("'" + this.input + "'")
      } catch (e) {
        this.$toast.error(e.toString())
      }
    },
    changeOutput () {
      try {
        this.input = escape(this.output).replace(/\%u/g, '\\u')
      } catch (e) {
        this.$toast.error(e.toString())
      }
    }
  }
}
</script>

<style scoped>
  
</style>
