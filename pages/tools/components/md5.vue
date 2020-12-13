<template>
  <div>
    <el-input type="textarea" v-model="slat" rows="3" placeholder="密钥..." @input="generateMd5"/>
    <el-input type="textarea" v-model="content" rows="15" placeholder="主体内容..." @input="generateMd5"/>
    <ul>
      <li>32位(小写): <span>{{ md5Val }}</span></li>
      <li>32位(大写): <span>{{ md5ValUp }}</span></li>
      <li>16位(小写): <span>{{ md5ValSm }}</span></li>
      <li>16位(大写): <span>{{ md5ValSmUp }}</span></li>
    </ul>
  </div>
</template>

<script>
  export default {
    name: 'ToolsMd5',
    data() {
      return {
        slat: '',
        content: '',
        md5Val: ''
      }
    },
    head() {
      return {
        script: [
          {src: '/js/md5.min.js'}
        ]
      }
    },
    computed: {
      md5ValUp() {
        return this.md5Val.toUpperCase()
      },
      md5ValSm() {
        return this.md5Val.substring(8, 24)
      },
      md5ValSmUp() {
        return this.md5Val.substring(8, 24).toUpperCase()
      }
    },
    methods: {
      generateMd5() {
        this.md5Val = md5(this.content, this.slat, false)
      }
    }
  }
</script>

<style scoped>
  .el-textarea {
    margin-bottom: 20px;
  }

  ul {
    list-style: none;
    margin: 0;
    padding: 20px 0;
    color: var(--color);
  }

  li {
    padding-bottom: 15px;
  }

  li > span {
    font-weight: 600;
  }
</style>
