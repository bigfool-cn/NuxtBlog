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
export default{
  name: 'ToolsBase64',
  data () {
    return {
      input: '',
      output: ''
    }
  },
  methods: {
    changeInput () {
      const obj = document.querySelector('.input')
      try {
        obj.style.borderColor = '#856043'
        this.output = window.atob(this.input)
      } catch (e) {
        obj.style.borderColor = 'red'
      }
    },
    changeOutput () {
      const obj = document.querySelector('.input')
      try {
        obj.style.borderColor = '#856043'
        this.input = window.btoa(this.output)
      } catch (e) {
        console.log(e)
        obj.style.borderColor = 'red'
      }
    }
  }
}
</script>

<style scoped>

</style>
