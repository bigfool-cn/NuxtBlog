<template>
  <div>
    <div v-if="!this.$store.state.settings.isMobile" class="eraser">
      <div class="theme-item">
        <span class="theme-btn" title="明亮" @click="$colorMode.preference = 'light'">
          <svg-icon icon-class="light" />
        </span>
        <span class="theme-btn" title="暗黑" @click="$colorMode.preference = 'dark'">
          <svg-icon icon-class="dark" />
        </span>
        <span class="theme-btn" title="深褐" @click="$colorMode.preference = 'sepia'">
          <svg-icon icon-class="sepia" />
        </span>
      </div>
      <div class="eraser-header">
        <svg-icon icon-class="share" title="分享到QQ" @click="shareQQ" />
      </div>
      <div class="eraser-body" title="返回顶部" @click="toTop">
        <svg-icon icon-class="top" />
      </div>
      <div class="eraser-footer" title="前往后台" @click="toLogin" />
    </div>
    <client-only>
    <el-dialog v-drag-dialog :visible.sync="showDialog" title="进入后台" width="30%">
      <div class="admin-login">
        <el-input v-model="loginForm.username" placeholder="用户名..." />
        <el-input v-model="loginForm.password" type="password" placeholder="密码..." />
        <el-button type="button" @click="login">
          登 录
        </el-button>
      </div>
    </el-dialog>
    </client-only>
  </div>
</template>

<script>
  import { login } from '~/apis/admin'
  import { Message } from 'element-ui'
  export default {
    name: 'Eraser',
    data() {
      return {
        showDialog: false,
        loginForm: {
          username: '',
          password: ''
        }
      }
    },
    mounted() {
      document.body.onresize = () => {
        if (document.body.clientWidth < 1170) {
          this.$store.dispatch('settings/setIsMobile', true)
        } else {
          this.$store.dispatch('settings/setIsMobile', false)
        }
      }
    },
    methods: {
      shareQQ() {

        // window.open('https://connect.qq.com/widget/shareqq/index.html?url=https://www.bigfool.cn&title=打工人&summary=打工魂&desc=打工人上人&pics=http://img.doutula.com/production/uploads/image/2018/06/05/20180605165104_EnKBLo.jpg&sharesource=qzone')
      },
      toLogin() {
        console.log(this.$store.state)
        const token = this.$store.state.admin.token
        if (token && token !== 'null' && token !== 'undefined') {
          this.$router.push({
            path: '/admin'
          })
        }
        this.showDialog = true
      },
      toTop() {
        document.documentElement.scrollTop = 0
        document.body.scrollTop = 0
      },
     login() {
        // 前台防君子
        if (this.$store.state.admin.limit > 5) {
          return Message.error('帐号已被限制')
        }

        if (!this.loginForm.username.trim()) {
          this.$store.dispatch('admin/inrLimit')
          return Message.error('请输入用户名')
        }

        if (!this.loginForm.password.trim()) {
          this.$store.dispatch('admin/inrLimit')
          return Message.error('请输入密码')
        }

        login(this.loginForm).then(async (response) => {
          if (response.code !== 200) {
            this.$store.dispatch('admin/inrLimit')
            return Message.error(response.message)
          }
          this.$cookies.set('_token_',response.data)
          await this.$store.dispatch('admin/setToken', response.data)
          Message.success(response.message)
          this.$router.push({
            path: '/admin'
          })
        })
      }
    }
  }

</script>

<style lang="scss" scoped>
  .theme-item {
    position: absolute;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    top: -32px;
    width: 80px;
  }

  .theme-btn {
    width: 25px;
    height: 25px;
    line-height: 25px;
    border-radius: 50%;
    text-align: center;
    cursor: pointer;
    background-color: var(--bg-secondary);
    &:hover {
      background-color: var(--border-active-color);
    }
  }

  .eraser {
    position: fixed;
    width: 80px;
    height: 120px;
    bottom: 20px;
    margin-left: 985px;
    background: var(--bg-secondary);
    border-radius: 25px;
    cursor: pointer;
  }

  .eraser-header,
  .eraser-footer {
    position: relative;
    text-align: center;
    height: 30px;
    padding-top: 4px;
  }

  .eraser-header {
    .svg-icon {
      width: 20px;
      height: 20px;
    }
  }

  .eraser-body {
    height: 60px;
    background: var(--bg);
    border: 1px solid var(--bg-secondary);
    text-align: center;
    padding-top: 14px;

    .svg-icon {
      width: 25px;
      height: 25px;
    }
  }

  .el-input {
    margin-bottom: 15px;
  }

</style>
