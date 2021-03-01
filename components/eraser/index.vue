<template>
  <div>
    <div class="eraser">
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
      <div class="eraser-header"  title="分享到QQ" >
        <svg-icon icon-class="share" @click="shareQQ"/>
      </div>
      <div class="eraser-body" title="返回顶部" @click="toTop">
        <svg-icon icon-class="top"/>
      </div>
      <div class="eraser-footer" title="在线工具" @click="toTools">
        <svg-icon icon-class="tools" class-name="tools-turn"/>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Eraser',
    methods: {
      shareQQ() {
        const url = window.location.href
        window.open(`http://connect.qq.com/widget/shareqq/index.html?url=${url}`)
      },
      toTop() {
        const timer = setInterval(() => {
          if (document.documentElement.scrollTop > 100) {
            document.documentElement.scrollTop -= 100
          } else {
            document.documentElement.scrollTop = 0
          }
          if (document.body.scrollTop > 100) {
            document.body.scrollTop -= 100
          } else {
            document.body.scrollTop = 0
          }
          if (document.body.scrollTop === 0 && document.documentElement.scrollTop === document.body.scrollTop) {
            clearInterval(timer)
          }
        },25)
      },
      toTools() {
        this.$router.push({path: '/tools'})
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


  }

  .theme-light:hover {
    .svg-icon {
      color: #243749;
     }
    background-color: #f3f5f4;
  }

  .theme-dark:hover {
    .svg-icon {
      color: #ebf4f1;
    }
    background-color: #091a28;
  }

  .theme-sepia:hover {
    .svg-icon {
      color: #433422;
    }
    background-color: #f1e7d0;
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

  .tools-turn {
    animation: turn 3s linear infinite;
  }

</style>
