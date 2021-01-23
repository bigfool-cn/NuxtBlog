<template>
  <div class="admin-header">
    <img src="/logo.png">
    <ul>
      <nuxt-link v-for="(menu,key) in menus" :key="key" :to="menu.path"
                 :class="currentPath === menu.path ? 'menu active-menu': 'menu'">
        <li>{{ menu.name }}</li>
      </nuxt-link>
    </ul>
    <div class="theme-item">
      <span class="theme-btn theme-light" title="明亮" @click="$colorMode.preference = 'light'">
        <svg-icon icon-class="light"/>
      </span>
      <span class="theme-btn theme-dark" title="暗黑" @click="$colorMode.preference = 'dark'">
        <svg-icon icon-class="dark"/>
      </span>
      <span class="theme-btn theme-sepia" title="深褐" @click="$colorMode.preference = 'sepia'">
        <svg-icon icon-class="sepia"/>
      </span>
    </div>
    <svg-icon icon-class="logout" class="logout" title="退出" @click="logout"/>
  </div>
</template>

<script>
  export default {
    name: 'AdminHeader',
    data() {
      return {
        menus: [
          {
            path: '/admin',
            name: '首页'
          },
          {
            path: '/admin/tag',
            name: '标签'
          },
          {
            path: '/admin/article',
            name: '文章'
          }
        ]
      }
    },
    computed: {
      currentPath() {
        return this.$route.path
      }
    },
    methods: {
      async logout() {
        await this.$store.dispatch('admin/logout')
        this.$cookies.remove('_token_')
        this.$router.replace({path: '/'})
      }
    }
  }
</script>

<style lang="scss" scoped>
  .admin-header {
    position: relative;
    display: flex;
    align-items: center;
    padding: 0 40px 15px 40px;
    border-bottom: 1px solid var(--border-color);
  }

  .menu {
    color: var(--color);
    text-decoration: none;
  }

  .active-menu,
  .menu:hover {
    color: var(--color-active);
  }

  ul {
    margin-left: 20px;
    padding: 0;
    list-style: none;
  }

  li {
    display: inline-block;
    padding: 0 25px;
  }

  .theme-item {
    position: absolute;
    right: 90px;
  }

  .theme-light,
  .theme-dark,
  .theme-sepia {
    border-radius: 50%;
    text-align: center;
    cursor: pointer;

    .svg-icon {
      width: 20px;
      height: 20px;
    }
  }

  .logout {
    position: absolute;
    right: 40px;
    color: var(--color-active);
    width: 30px !important;
    height: 30px !important;
    cursor: pointer;
  }
</style>
