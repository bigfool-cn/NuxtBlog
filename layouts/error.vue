<template>
  <div>
    <div class="error">
      <svg-icon v-if="error.statusCode === 404" icon-class="404" />
      <svg-icon v-if="error.statusCode === 403" icon-class="403" />
      <svg-icon v-if="error.statusCode === 500" icon-class="500" />
    </div>
  </div>
</template>
<script>
  export default {
    props: ['error']
  }

</script>
<style lang="scss" scoped>
.error {
  padding-top: 100px;
  text-align: center;
  .svg-icon {
    color: var(--border-active-color);
    width: 300px;
    height: 300px;
  }
}
</style>
