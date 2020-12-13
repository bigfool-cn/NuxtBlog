<template>
  <div class="admin-container">
    <a-header/>
    <Nuxt/>
  </div>
</template>

<script>
  import AHeader from '~/layouts/components/aheader.vue'

  export default {
    name: 'Admin',
    components: {AHeader}
  }
</script>

<style>

</style>
