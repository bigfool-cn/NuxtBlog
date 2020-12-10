<template>
  <div class="admin-container">
    <a-header />
    <Nuxt />
  </div>
</template>

<script>
import AHeader from '~/layouts/components/aheader.vue'
export default {
  name: 'Admin',
  components: { AHeader }
}
</script>

<style>
  .admin-container {
    margin: auto;
    min-height: 100vh;
    padding: 20px 0;
    background:var(--bg);
  }
  .content-container {
    padding: 20px 40px;
  }
  .serach-container {
    padding-bottom: 20px;
  }
  .table-btn {
    padding-bottom: 5px;
  }
</style>
