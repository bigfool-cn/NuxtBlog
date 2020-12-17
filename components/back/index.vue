<template>
  <div class="back-btn" @click="goBack"><svg-icon icon-class="back" /></div>
</template>
<script>
  export default {
    name: 'Back',
    methods: {
      goBack() {
        window.history.back(-1);
      }
    }
  }
</script>
<style lang="scss" scoped>
.back-btn {
  position: fixed;
  bottom: 180px;
  margin-left: 1010px;
  border-radius: 25px;
  cursor: pointer;

  .svg-icon {
    width: 27px;
    height: 25px;
  }
}
</style>