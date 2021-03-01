<template>
  <div>
    <div class="pencil">
      <span class="barrel" @click="toIndex">
        <div class="header">
          <span>一只小开发仔...</span>
        </div>
      </span>
      <span class="taper" @click="toLogin"/>
    </div>
    <client-only>
      <el-dialog v-drag-dialog :visible.sync="showDialog" title="进入后台" width="30%">
        <div class="admin-login">
          <el-input v-model="loginForm.username" placeholder="用户名..."/>
          <el-input v-model="loginForm.password" type="password" placeholder="密码..." @keyup.enter.native="login"/>
          <el-button type="button" @click="login">
            登 录
          </el-button>
        </div>
      </el-dialog>
    </client-only>
  </div>
</template>
<script>
  import {login} from '~/apis/admin'
  import {Message} from "element-ui"
  export default {
    name: 'BHeader',
    data() {
      return {
        showDialog: false,
        loginForm: {
          username: '',
          password: ''
        }
      }
    },
    methods: {
      toIndex() {
        this.$router.push({path: '/'})
      },
      toLogin() {
        const token = this.$store.state.admin.token
        /**
         if (token && token !== 'null' && token !== 'undefined') {
          this.$router.push({
            path: '/admin'
          })
        }
         */
        this.showDialog = true
      },
      login() {
        const limit = this.$cookies.get('_limit_') || 0
        // 年轻人不讲武德 耗子尾汁
        if (limit >= 5) {
          return Message.error('帐号已被限制')
        }

        if (!this.loginForm.username.trim()) {
          this.setLimit(limit + 1)
          return Message.error('请输入用户名')
        }

        if (!this.loginForm.password.trim()) {
          this.setLimit(limit + 1)
          return Message.error('请输入密码')
        }

        login(this.loginForm).then(async (response) => {
          if (response.code !== 200) {
            this.setLimit(limit + 1)
            return Message.error(response.message)
          }
          this.$cookies.set('_token_', response.data)
          await this.$store.dispatch('admin/setToken', response.data)
          Message.success(response.message)
          this.$router.push({
            path: '/admin'
          })
        })
      },
      setLimit(limit) {
        this.$cookies.set('_limit_',limit,{maxAge: 60 * 60 * 8})
      }
    }
  }

</script>
<style lang="scss" scoped>
  .pencil {
    margin: auto;
    display: flex;
    max-width: 900px;
    height: 30px;
    transform-origin: 150% center;

    .taper {
      width: 50px;
      position: relative;
      cursor: pointer;

      /* 三角形 */
      &:before {
        content: '';
        position: absolute;
        border-style: solid;
        border-width: calc(30px / 2) 0 calc(30px / 2) 50px;
        border-color: transparent;
        border-left-color: var(--pencil-bg-color);
      }

      &:after {
        content: '';
        position: absolute;
        border-style: solid;
        border-width: 15px 50px;
        border-color: transparent;
        border-right-color: #111;
        transform: scale(0.3) rotate(180deg);
      }
    }

    .barrel {
      border-radius: 0.2em 0 0 0.2em;
      width: 850px;
      background-color: var(--pencil-bg-color);
      border-top: 0.5em solid var(--pencil-border-top-color);
      border-bottom: 0.5em solid var(--pencil-border-bottom-color);
      line-height: 1em;
      text-align: center;
      text-transform: uppercase;
      letter-spacing: 0.1em;
      cursor: pointer;
    }
  }

  .header {
    font-weight: bold;
    font-style: italic;
    color: var(--color);
    font-size: 1.25em;
    line-height: 0.6em;
    text-align: center;
    padding-left: 50px;
  }

  .el-input {
    padding-bottom: 15px;
  }

</style>>
