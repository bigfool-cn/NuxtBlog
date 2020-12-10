<template>
  <div>
    <input id="file" type="file" name="file" style="display: none" @change="changeFile">
    <div class="upload-box" @dragover="fileDragover" @drop="fileDrop" @click="uploadFile">
      <span>将文件拖拽至此或点击上传</span>
    </div>
    <svg-icon icon-class="copy" title="复制" class="copy" @click="copyContent" />
    <el-input type="textarea" id="base64" v-model="base64" rows="25" readonly="readonly" placeholder="输出base64..." />
  </div>
</template>

<script>
export default {
  name: 'ToolsImgBase64',
  data () {
    return {
      base64: '',
      maxSize: 1024 * 1024 * 10
    }
  },
  methods: {
    uploadFile () {
      document.querySelector('#file').click()
    },
    changeFile (event) {
      const file = event.target.files.item(0)

      if (!file) {
        return false
      }
      if (file.size > this.maxSize) {
        this.$toast.error('文件不能超多10MB')
        return false
      }

      this.setFileUrl(file)
      const obj = document.querySelector('.upload-box span')
      obj.innerHTML = file.name
      event.target.value = ''
    },
    fileDragover (event) {
      event.preventDefault()
    },
    fileDrop (event) {
      event.preventDefault()
      const file = event.dataTransfer.files[0]
      if (file.size > this.maxSize) {
        this.$toast.error('文件不能超多10MB')
        return
      }
      this.setFileUrl(file)
    },
    setFileUrl (file) {
      const fr = new FileReader()
      fr.onload = (e) => {
        this.base64 = e.target.result
        const obj = document.querySelector('.upload-box span')
        obj.innerHTML = file.name
      }
      fr.readAsDataURL(file)
    },
    copyContent () {
      const obj = document.querySelector('#base64')
      obj.select()
      document.execCommand('copy')
      this.$toast.success('复制成功!')
    }
  }
}
</script>

<style scoped>
  .upload-box {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    background: var(--bg);
    height: 120px;
    border: 1px dotted var(--border-color);
    margin-bottom: 15px;
    cursor: pointer;
    border-radius: 10px;
  }
  .upload-box span {
    color: var(--color);;
  }
  .copy {
    float: right;
    margin-bottom: 5px;
    cursor: pointer;
  }
</style>
