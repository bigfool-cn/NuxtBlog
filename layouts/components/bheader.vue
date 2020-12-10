<template>
  <div>
    <div class="pencil">
      <span class="barrel" @click="toIndex">
        <div class="header">
          <span>好记性不如烂笔头</span>
        </div>
      </span>
      <span class="taper" title="在线工具" @click="toTools" />
    </div>
  </div>
</template>
<script>
export default {
  name: 'BHeader',
  methods: {
    toIndex() {
      this.$router.push({ path: '/' })
    },
    toTools() {
      this.$router.push({ path: '/tools' })
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
}

.pencil .taper {
  width: 50px;
  position: relative;
  cursor: pointer;
}

/* 三角形 */
.pencil .taper::before {
  content: '';
  position: absolute;
  border-style: solid;
  border-width: calc(30px / 2) 0 calc(30px / 2) 50px;
  border-color: transparent;
}

.pencil .taper::before {
  border-left-color: var(--pencil-bg-color);
}

.pencil .taper::after {
  content: '';
  position: absolute;
  border-style: solid;
  border-width: 15px 50px;
  border-color: transparent;
  border-right-color: #111;
  transform: scale(0.3) rotate(180deg);
}

.pencil .barrel {
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

.header {
  font-weight: bold;
  font-style: italic;
  color: var(--color);
  font-size: 1.25em;
  line-height: 0.65em;
  padding-left: 2em;
}

</style>>
