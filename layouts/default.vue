<template>
  <div>
    <div class="container">
      <b-header />
      <Nuxt />
    </div>
  </div>
</template>
<script>
import BHeader from '~/layouts/components/bheader.vue'
export default {
  components: { BHeader },
  data() {
    return {
      maxClientWidth: 980
    }
  },
  mounted() {
    const flag = navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)
    if (flag) {
      this.$store.dispatch('setting/setIsMobile', true)
    }
    document.body.onresize = () => {
      if (document.body.clientWidth > this.maxClientWidth) {
        this.$store.dispatch('settings/setIsMobile', false)
      } else {
        this.$store.dispatch('settings/setIsMobile', true)
      }
    }
  }
}

</script>
<style>
  
</style>
