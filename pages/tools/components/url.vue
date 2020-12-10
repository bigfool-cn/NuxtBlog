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
  name: 'ToolsUrl',
  data () {
    return {
      input: '',
      output: ''
    }
  },
  methods: {
    changeInput () {
      this.output = decodeURIComponent(this.input)
    },
    changeOutput () {
      this.input = encodeURIComponent(this.output)
    }
  }
}
</script>

<style scoped>
  
</style>
