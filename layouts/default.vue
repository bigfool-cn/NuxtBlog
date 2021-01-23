<template>
  <div>
    <div class="container">
      <b-header/>
      <Nuxt keep-alive :keep-alive-props="{include: ['Index']}"/>
    </div>
  </div>
</template>
<script>
  import BHeader from '~/layouts/components/bheader.vue'

  export default {
    components: {BHeader},
    mounted() {
      console.log("%c ", "background:url('https://www.bigfool.cn/logo.png') no-repeat center;padding-left:80px;padding-bottom: 80px;");
      console.log("%c连心爱的人都留不住, 代码敲得再六又怎么样.","color:green")
    }
  }

</script>
<style>

</style>
