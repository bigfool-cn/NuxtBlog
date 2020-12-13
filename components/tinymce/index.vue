<template>
  <div class="tinymce-box">
    <editor
      v-model="content"
      :init="init"
      :disabled="disabled"
      @onClick="onClick"
    />
  </div>
</template>

<script>
  import Editor from '@tinymce/tinymce-vue'

  let tinymce
  if (process.client) {
    tinymce = require('tinymce/tinymce')
    require('tinymce/themes/silver')
    require('tinymce/icons/default')
    // 编辑器插件plugins
    // 更多插件参考：https://www.tiny.cloud/docs/plugins/
    // 插入上传图片插件
    require('tinymce/plugins/image')
    // 插入视频插件
    require('tinymce/plugins/media')
    // 插入表格插件
    require('tinymce/plugins/table')
    // 列表插件
    require('tinymce/plugins/lists')
    // 字数统计插件
    require('tinymce/plugins/wordcount')
    // 自动保存
    require('tinymce/plugins/autosave')
    // 全屏
    require('tinymce/plugins/fullscreen')
    // 插入代码
    require('tinymce/plugins/codesample')
    // 预览
    require('tinymce/plugins/preview')
    // 查找替换
    require('tinymce/plugins/searchreplace')
    // 水平分割线
    require('tinymce/plugins/hr')
  }
  export default {
    name: 'Tinymce',
    components: {
      Editor
    },
    props: {
      value: {
        type: String,
        default: ''
      },
      height: {
        type: Number,
        default: 300
      },
      disabled: {
        type: Boolean,
        default: false
      },
      plugins: {
        type: [String, Array],
        default: 'lists image media table wordcount autosave fullscreen codesample preview searchreplace hr'
      },
      toolbar: {
        type: [String, Array],
        default: 'searchreplace | formatselect | bold italic forecolor backcolor removeformat codesample | image table lists bullist numlist alignleft aligncenter alignright alignjustify fullscreen | hr blockquote outdent indent outdent indent | media preview restoredraft'
      }
    },
    data() {
      return {
        init: {
          language_url: '/tinymce/langs/zh_CN.js',
          language: 'zh_CN',
          skin_url: '/tinymce/skins/ui/oxide',
          // skin_url: 'tinymce/skins/ui/oxide-dark',//暗色系
          content_css: '/tinymce/skins/content/default/content.css',
          height: this.height,
          plugins: this.plugins,
          toolbar: this.toolbar,
          branding: false,
          menubar: false,
          codesample_languages: [
            {text: 'C', value: 'c'},
            {text: 'C++', value: 'c++'},
            {text: 'SQL', value: 'sql'},
            {text: 'CSS', value: 'css'},
            {text: 'PHP', value: 'php'},
            {text: 'Golang', value: 'go'},
            {text: 'Java', value: 'java'},
            {text: 'Python', value: 'python'},
            {text: 'Bash', value: 'bash'},
            {text: 'YAML', value: 'yaml'},
            {text: 'JSON', value: 'json'},
            {text: 'HTML/XML', value: 'markup'},
            {text: 'JavaScript', value: 'javascript'}
          ],
          // 此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
          // 如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
          images_upload_handler: (blobInfo, success) => {
            const img = 'data:image/jpeg;base64,' + blobInfo.base64()
            success(img)
          }
        },
        content: this.value
      }
    },
    watch: {
      value(newValue) {
        this.content = newValue
      },
      content(newValue) {
        this.$emit('input', newValue)
      }
    },
    mounted() {
      tinymce.init({})
    },
    methods: {
      // 添加相关的事件，可用的事件参照文档=> https://github.com/tinymce/tinymce-vue => All available events
      // 需要什么事件可以自己增加
      onClick(e) {
        this.$emit('onclick', e, tinymce)
      },
      // 可以添加一些自己的自定义事件，如清空内容
      clear() {
        this.content = ''
      }
    }
  }
</script>
<style scoped>
</style>
